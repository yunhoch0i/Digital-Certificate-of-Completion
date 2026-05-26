package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	lambdasdk "github.com/aws/aws-sdk-go-v2/service/lambda"
	lambdatypes "github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/ethereum/go-ethereum/common"

	atk "asbg-token/bindings/attendance_tracker"
	cnft "asbg-token/bindings/certificate_nft"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/members"
	"asbg-token/internal/metadata"
	"asbg-token/internal/secret"
	"asbg-token/internal/store"
)

var (
	rpcURL         string
	privateKey     string
	trackerAddress string
	certNFTAddress string
	dynStore       *store.DynamoStore
)

func mustGet(ctx context.Context, sm *secret.Manager, key string) string {
	v, err := sm.Get(ctx, key)
	if err != nil {
		log.Fatalf("init: SSM %q 조회 실패: %v", key, err)
	}
	if v == "" {
		log.Fatalf("init: SSM %q 값이 비어 있습니다", key)
	}
	return v
}

func init() {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("init: load aws config: %v", err)
	}
	sm := secret.NewManager(cfg)
	dynStore = store.NewDynamoStore(cfg)

	rpcURL = mustGet(ctx, sm, "asbg/RPC_URL")
	privateKey = mustGet(ctx, sm, "asbg/PRIVATE_KEY")
	trackerAddress = mustGet(ctx, sm, "asbg/ATTENDANCE_TRACKER_ADDRESS")
	certNFTAddress = mustGet(ctx, sm, "asbg/CERTIFICATE_NFT_ADDRESS")
}

func jsonResp(status int, body any) (events.APIGatewayV2HTTPResponse, error) {
	b, _ := json.Marshal(body)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		Body:       string(b),
	}, nil
}

func errResp(status int, msg string) (events.APIGatewayV2HTTPResponse, error) {
	return jsonResp(status, map[string]string{"error": msg})
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch req.RequestContext.HTTP.Method + " " + req.RawPath {
	case "GET /members":
		return handleMembers(ctx)
	case "GET /certificates":
		return handleCertificates(ctx)
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

func handleMembers(ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {
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

// ── GET /certificates ─────────────────────────────────────────────────────────

type certResponse struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Role         string  `json:"role"`
	Attendance   int64   `json:"attendance"`
	TotalSessions int    `json:"total_sessions"`
	Threshold    int     `json:"threshold"`
	Certified    bool    `json:"certified"`
	TokenID      *uint64 `json:"token_id,omitempty"`
	TokenURI     string  `json:"token_uri,omitempty"`
}

func handleCertificates(ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {
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

	totalSessions, _ := strconv.Atoi(os.Getenv("TOTAL_SESSIONS"))
	if totalSessions == 0 {
		totalSessions = 10
	}
	threshold := metadata.CalcThreshold(totalSessions)

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		return errResp(500, "blockchain client: "+err.Error())
	}

	trackerAddr := common.HexToAddress(trackerAddress)
	tracker, err := atk.NewAttendanceTracker(trackerAddr, client.Eth())
	if err != nil {
		return errResp(500, "bind tracker: "+err.Error())
	}

	certAddr := common.HexToAddress(certNFTAddress)
	certContract, err := cnft.NewCertificateNFT(certAddr, client.Eth())
	if err != nil {
		return errResp(500, "bind cert nft: "+err.Error())
	}

	resp := make([]certResponse, 0, len(mems))
	for _, m := range mems {
		bal, _ := tracker.Attendance(nil, big.NewInt(int64(m.ID)))
		att := int64(0)
		if bal != nil {
			att = bal.Int64()
		}

		certified, _ := certContract.HasCertificate(nil, big.NewInt(int64(m.ID)))

		cr := certResponse{
			ID:           m.ID,
			Name:         m.Name,
			Role:         m.Role,
			Attendance:   att,
			TotalSessions: totalSessions,
			Threshold:    threshold,
			Certified:    certified,
		}

		if certified {
			tokenId, err := certContract.MemberTokenId(nil, big.NewInt(int64(m.ID)))
			if err == nil && tokenId != nil {
				tid := tokenId.Uint64()
				cr.TokenID = &tid
				uri, err := certContract.TokenURI(nil, tokenId)
				if err == nil {
					cr.TokenURI = uri
				}
			}
		}

		resp = append(resp, cr)
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

func handleMint(ctx context.Context, body string) (events.APIGatewayV2HTTPResponse, error) {
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

	// Mint 성공 → Job Lambda를 비동기로 트리거해 수료 기준 달성자에게 인증서 자동 발급
	invokeCertJob()

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

// invokeCertJob triggers the Job Lambda asynchronously (fire-and-forget).
// Returns immediately — certificate issuance happens in the background.
func invokeCertJob() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Printf("invokeCertJob: aws config: %v", err)
		return
	}
	name := os.Getenv("JOB_FUNCTION_NAME")
	if name == "" {
		name = "asbg-job"
	}
	client := lambdasdk.NewFromConfig(cfg)
	_, err = client.Invoke(ctx, &lambdasdk.InvokeInput{
		FunctionName:   aws.String(name),
		InvocationType: lambdatypes.InvocationTypeEvent, // async — returns instantly
	})
	if err != nil {
		log.Printf("invokeCertJob: %v", err)
	} else {
		log.Printf("certificate job triggered (async)")
	}
}

func main() {
	lambda.Start(Handler)
}
