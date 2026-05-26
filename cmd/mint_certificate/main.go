package main

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	cert "asbg-token/bindings/certificate_nft"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/ipfs"
	"asbg-token/internal/metadata"
)

type completionRow struct {
	ID         int
	Name       string
	Role       string
	Attendance int
	Threshold  int
	Completed  bool
}

type txLog struct {
	Timestamp string `json:"timestamp"`
	Action    string `json:"action"`
	MemberID  int    `json:"member_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	TokenID   uint64 `json:"token_id"`
	TxHash    string `json:"tx_hash"`
	Block     uint64 `json:"block"`
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	rpcURL := mustEnv("RPC_URL")
	privateKey := mustEnv("PRIVATE_KEY")
	certAddrHex := mustEnv("CERTIFICATE_NFT_ADDRESS")
	pinataJWT := mustEnv("PINATA_JWT")
	totalSessionsStr := mustEnv("TOTAL_SESSIONS")

	var totalSessions int
	if _, err := fmt.Sscan(totalSessionsStr, &totalSessions); err != nil {
		log.Fatalf("TOTAL_SESSIONS invalid: %v", err)
	}

	rows, err := loadCompletionResult("completion_result.csv")
	if err != nil {
		log.Fatalf("load completion_result.csv: %v", err)
	}
	var qualified []completionRow
	for _, r := range rows {
		if r.Completed {
			qualified = append(qualified, r)
		}
	}
	if len(qualified) == 0 {
		fmt.Println("수료 대상자가 없습니다.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		log.Fatalf("blockchain client: %v", err)
	}
	certAddr := common.HexToAddress(certAddrHex)
	nft, err := cert.NewCertificateNFT(certAddr, client.Eth())
	if err != nil {
		log.Fatalf("bind CertificateNFT: %v", err)
	}

	issuedDate := time.Now().Format("2006-01-02")

	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatalf("create logs dir: %v", err)
	}
	logFile, err := os.OpenFile("logs/tx_history.jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("open log file: %v", err)
	}
	defer logFile.Close()

	var guides []string
	issued, skipped, failed := 0, 0, 0

	fmt.Printf("수료 대상 %d명 인증서 발행 시작\n", len(qualified))
	fmt.Println("─────────────────────────────────────────────────")

	for _, row := range qualified {
		memberID := big.NewInt(int64(row.ID))

		has, err := nft.HasCertificate(nil, memberID)
		if err != nil {
			log.Printf("[SKIP] %s: hasCertificate error: %v", row.Name, err)
			skipped++
			continue
		}
		if has {
			fmt.Printf("[SKIP] %s — 이미 발행됨\n", row.Name)
			skipped++
			continue
		}

		credURL := fmt.Sprintf("https://sepolia.etherscan.io/token/%s?a=", certAddrHex)
		meta := metadata.Build(metadata.CertParams{
			MemberName:      row.Name,
			Role:            row.Role,
			IssuedDate:      issuedDate,
			CredentialURL:   credURL,
			MemberAddress:   fmt.Sprintf("memberId:%d", row.ID),
			Attendance:      row.Attendance,
			TotalSessions:   totalSessions,
			ContractAddress: certAddrHex,
		})

		cid, err := ipfs.UploadJSON(ctx, pinataJWT, meta)
		if err != nil {
			log.Printf("[FAIL] %s: upload IPFS: %v", row.Name, err)
			failed++
			continue
		}

		var tokenID uint64
		var txHash string
		var blockNum uint64

		err = blockchain.WithRetry(3, func() error {
			auth, err := client.NewTransactOpts(ctx)
			if err != nil {
				return err
			}
			tx, err := nft.IssueCertificate(auth, memberID, cid, row.Name)
			if err != nil {
				return fmt.Errorf("issueCertificate: %w", err)
			}
			receipt, err := client.WaitMined(ctx, tx)
			if err != nil {
				return fmt.Errorf("WaitMined: %w", err)
			}
			if receipt.Status == 0 {
				return fmt.Errorf("tx reverted")
			}
			txHash = tx.Hash().Hex()
			blockNum = receipt.BlockNumber.Uint64()

			nftAbi, _ := cert.CertificateNFTMetaData.GetAbi()
			for _, lg := range receipt.Logs {
				ev, err := nftAbi.EventByID(lg.Topics[0])
				if err != nil || ev.Name != "CertificateIssued" {
					continue
				}
				if parsed, err := nft.ParseCertificateIssued(*lg); err == nil {
					tokenID = parsed.TokenId.Uint64()
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[FAIL] %s: %v", row.Name, err)
			failed++
			continue
		}

		credURLFull := fmt.Sprintf("%s%d", credURL, tokenID)
		fmt.Printf("[OK] %s (%s) → tokenId=%d  tx=%s...\n", row.Name, row.Role, tokenID, txHash[:18])

		entry := txLog{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Action:    "mint_certificate",
			MemberID:  row.ID,
			Name:      row.Name,
			Role:      row.Role,
			TokenID:   tokenID,
			TxHash:    txHash,
			Block:     blockNum,
		}
		line, _ := json.Marshal(entry)
		fmt.Fprintln(logFile, string(line))

		guides = append(guides, fmt.Sprintf(
			"=== %s (%s) ===\n자격증 이름 : ASBG 수료 인증서\n발급 기관   : AWS Student Builders Group\n발급일      : %s\n만료일      : (입력 안 함 — 영구)\n자격증 URL  : %s\n",
			row.Name, row.Role, issuedDate, credURLFull,
		))
		issued++
	}

	fmt.Println("─────────────────────────────────────────────────")
	fmt.Printf("발행 완료: %d명 / 스킵: %d명 / 실패: %d명\n", issued, skipped, failed)

	if len(guides) > 0 {
		guideFile := fmt.Sprintf("linkedin_guide_%s.txt", time.Now().Format("2006"))
		if err := os.WriteFile(guideFile, []byte(strings.Join(guides, "\n")), 0644); err != nil {
			log.Printf("write linkedin guide: %v", err)
		} else {
			fmt.Printf("LinkedIn 가이드 저장: %s\n", guideFile)
		}
	}
}

func loadCompletionResult(path string) ([]completionRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return nil, err
	}
	var rows []completionRow
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		id, _ := strconv.Atoi(rec[0])
		att, _ := strconv.Atoi(rec[3])
		thr, _ := strconv.Atoi(rec[4])
		rows = append(rows, completionRow{
			ID: id, Name: rec[1], Role: rec[2],
			Attendance: att, Threshold: thr,
			Completed: rec[5] == "true",
		})
	}
	return rows, nil
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("환경변수 %q 가 설정되지 않았습니다", key)
	}
	return v
}
