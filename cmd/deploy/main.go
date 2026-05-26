package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	tracker "asbg-token/bindings/attendance_tracker"
	cert "asbg-token/bindings/certificate_nft"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/metadata"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	rpcURL := mustEnv("RPC_URL")
	privateKey := mustEnv("PRIVATE_KEY")
	totalSessionsStr := mustEnv("TOTAL_SESSIONS")

	var totalSessions int
	if _, err := fmt.Sscan(totalSessionsStr, &totalSessions); err != nil || totalSessions <= 0 {
		log.Fatalf("TOTAL_SESSIONS must be a positive integer, got: %q", totalSessionsStr)
	}

	threshold := metadata.CalcThreshold(totalSessions)
	fmt.Printf("TOTAL_SESSIONS       : %d\n", totalSessions)
	fmt.Printf("COMPLETION_THRESHOLD : %d  (ceil(%d × 0.9))\n\n", threshold, totalSessions)

	fmt.Print("배포를 진행하시겠습니까? [y/N] ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		fmt.Println("배포 취소.")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		log.Fatalf("blockchain client: %v", err)
	}

	// ── AttendanceTracker: 이미 배포된 경우 스킵 ──────────────────────────────
	var trackerAddr common.Address
	if existing := os.Getenv("ATTENDANCE_TRACKER_ADDRESS"); existing != "" && common.IsHexAddress(existing) {
		trackerAddr = common.HexToAddress(existing)
		fmt.Printf("AttendanceTracker : 기존 주소 사용 — %s\n\n", trackerAddr.Hex())
	} else {
		auth, err := client.NewTransactOpts(ctx)
		if err != nil {
			log.Fatalf("transact opts: %v", err)
		}
		fmt.Println("AttendanceTracker 배포 중...")
		addr, tx, _, err := tracker.DeployAttendanceTracker(auth, client.Eth())
		if err != nil {
			log.Fatalf("deploy AttendanceTracker: %v", err)
		}
		if _, err := client.WaitMined(ctx, tx); err != nil {
			log.Fatalf("wait AttendanceTracker: %v", err)
		}
		trackerAddr = addr
		fmt.Printf("AttendanceTracker : %s\n  %s/address/%s\n\n", trackerAddr.Hex(), explorerURL(), trackerAddr.Hex())
	}

	// ── CertificateNFT 배포 ───────────────────────────────────────────────────
	auth2, err := client.NewTransactOpts(ctx)
	if err != nil {
		log.Fatalf("transact opts (cert): %v", err)
	}
	fmt.Println("CertificateNFT 배포 중...")
	certAddr, certTx, _, err := cert.DeployCertificateNFT(auth2, client.Eth(), trackerAddr, big.NewInt(int64(threshold)))
	if err != nil {
		log.Fatalf("deploy CertificateNFT: %v", err)
	}
	if _, err := client.WaitMined(ctx, certTx); err != nil {
		log.Fatalf("wait CertificateNFT: %v", err)
	}
	fmt.Printf("CertificateNFT    : %s\n  %s/address/%s\n\n", certAddr.Hex(), explorerURL(), certAddr.Hex())

	fmt.Println("─────────────────────────────────────────────────")
	fmt.Println(".env에 아래 주소를 기입하세요:")
	fmt.Printf("ATTENDANCE_TRACKER_ADDRESS=%s\n", trackerAddr.Hex())
	fmt.Printf("CERTIFICATE_NFT_ADDRESS=%s\n", certAddr.Hex())
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("환경변수 %q 가 설정되지 않았습니다", key)
	}
	return v
}

func explorerURL() string {
	if u := os.Getenv("EXPLORER_URL"); u != "" {
		return u
	}
	return "https://sepolia.etherscan.io"
}
