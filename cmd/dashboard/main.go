package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/ethereum/go-ethereum/common"

	atk "asbg-token/bindings/attendance_tracker"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/members"
	"asbg-token/internal/secret"
	"asbg-token/internal/store"
)

var (
	rpcURL         string
	privateKey     string
	trackerAddress string
	dynStore       *store.DynamoStore
)

func init() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("init: load aws config: %v", err)
	}
	sm := secret.NewManager(cfg)
	dynStore = store.NewDynamoStore(cfg)

	rpcURL, _ = sm.Get(ctx, "asbg/RPC_URL")
	privateKey, _ = sm.Get(ctx, "asbg/PRIVATE_KEY")
	trackerAddress, _ = sm.Get(ctx, "asbg/ATTENDANCE_TRACKER_ADDRESS")
}

func jsonResp(status int, body any) (events.APIGatewayProxyResponse, error) {
	b, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		Body:       string(b),
	}, nil
}

func errResp(status int, msg string) (events.APIGatewayProxyResponse, error) {
	return jsonResp(status, map[string]string{"error": msg})
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod + " " + req.Path {
	case "GET /members":
		return handleMembers(ctx)
	case "POST /mint":
		return handleMint(ctx, req.Body)
	}
	return errResp(404, "not found")
}

// ── GET /members ──────────────────────────────────────────────────────────────

type memberResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	Attendance string `json:"attendance"`
}

func handleMembers(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return errResp(500, "aws config: "+err.Error())
	}
	bucket := os.Getenv("DATA_BUCKET")
	if bucket == "" {
		bucket = "asbg-data"
	}
	mems, err := members.LoadFromS3(ctx, awsCfg, bucket, "members.csv")
	if err != nil {
		return errResp(500, "load members: "+err.Error())
	}

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		return errResp(500, "blockchain client: "+err.Error())
	}
	trackerAddr := common.HexToAddress(trackerAddress)
	tracker, err := atk.NewAttendanceTracker(trackerAddr, client.Eth())
	if err != nil {
		return errResp(500, "bind tracker: "+err.Error())
	}

	var resp []memberResponse
	for _, m := range mems {
		bal, err := tracker.Attendance(nil, big.NewInt(int64(m.ID)))
		attStr := "0"
		if err == nil {
			attStr = bal.String()
		}
		resp = append(resp, memberResponse{
			ID:         m.ID,
			Name:       m.Name,
			Role:       m.Role,
			Attendance: attStr,
		})
	}
	return jsonResp(200, resp)
}

// ── POST /mint ────────────────────────────────────────────────────────────────

type mintRequest struct {
	EventID string `json:"event_id"`
	Present []uint `json:"present"` // member IDs (CSV row indices)
}

type mintResult struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	TxHash string `json:"tx_hash,omitempty"`
	Error  string `json:"error,omitempty"`
}

func handleMint(ctx context.Context, body string) (events.APIGatewayProxyResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	var req mintRequest
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return errResp(400, "invalid json: "+err.Error())
	}
	if req.EventID == "" {
		return errResp(400, "event_id required")
	}
	if len(req.Present) == 0 {
		return jsonResp(200, map[string]any{"success": []any{}, "failed": []any{}})
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return errResp(500, "aws config: "+err.Error())
	}
	bucket := os.Getenv("DATA_BUCKET")
	if bucket == "" {
		bucket = "asbg-data"
	}
	mems, err := members.LoadFromS3(ctx, awsCfg, bucket, "members.csv")
	if err != nil {
		return errResp(500, "load members: "+err.Error())
	}
	nameByID := make(map[uint]string, len(mems))
	for _, m := range mems {
		nameByID[m.ID] = m.Name
	}

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		return errResp(500, "blockchain client: "+err.Error())
	}
	trackerAddr := common.HexToAddress(trackerAddress)
	tracker, err := atk.NewAttendanceTracker(trackerAddr, client.Eth())
	if err != nil {
		return errResp(500, "bind tracker: "+err.Error())
	}

	// batch mint in one tx
	ids := make([]*big.Int, len(req.Present))
	for i, id := range req.Present {
		ids[i] = big.NewInt(int64(id))
	}

	var txHash string
	err = blockchain.WithRetry(3, func() error {
		auth, err := client.NewTransactOpts(ctx)
		if err != nil {
			return err
		}
		tx, err := tracker.MintAttendance(auth, ids, req.EventID)
		if err != nil {
			return fmt.Errorf("mintAttendance: %w", err)
		}
		receipt, err := client.WaitMined(ctx, tx)
		if err != nil {
			return fmt.Errorf("WaitMined: %w", err)
		}
		if receipt.Status == 0 {
			return fmt.Errorf("tx reverted")
		}
		txHash = tx.Hash().Hex()
		return nil
	})

	if err != nil {
		return jsonResp(200, map[string]any{
			"success": []any{},
			"failed": []mintResult{{Error: err.Error()}},
		})
	}

	// log each member
	var successes []mintResult
	for _, id := range req.Present {
		name := nameByID[id]
		successes = append(successes, mintResult{ID: id, Name: name, TxHash: txHash})
		_ = dynStore.PutTx(ctx, store.TxRecord{
			TxHash:  txHash,
			Action:  "mint_attendance",
			Member:  fmt.Sprintf("memberId:%d", id),
			Name:    name,
			EventID: req.EventID,
		})
	}

	return jsonResp(200, map[string]any{"success": successes, "failed": []any{}})
}

func main() {
	lambda.Start(Handler)
}
