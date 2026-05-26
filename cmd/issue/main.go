package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	atk "asbg-token/bindings/attendance_tracker"
	cert "asbg-token/bindings/certificate_nft"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/ipfs"
	"asbg-token/internal/members"
	"asbg-token/internal/metadata"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	rpcURL      := mustEnv("RPC_URL")
	privateKey  := mustEnv("PRIVATE_KEY")
	trackerAddr := mustEnv("ATTENDANCE_TRACKER_ADDRESS")
	certAddr    := mustEnv("CERTIFICATE_NFT_ADDRESS")
	pinataJWT   := mustEnv("PINATA_JWT")

	totalSessions, err := strconv.Atoi(mustEnv("TOTAL_SESSIONS"))
	if err != nil || totalSessions <= 0 {
		log.Fatalf("TOTAL_SESSIONS invalid")
	}
	threshold := metadata.CalcThreshold(totalSessions)

	mems, err := members.LoadMembers("members.csv")
	if err != nil {
		log.Fatalf("load members.csv: %v", err)
	}

	ctx := context.Background()
	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		log.Fatalf("blockchain client: %v", err)
	}

	tracker, err := atk.NewAttendanceTracker(common.HexToAddress(trackerAddr), client.Eth())
	if err != nil {
		log.Fatalf("bind tracker: %v", err)
	}
	nft, err := cert.NewCertificateNFT(common.HexToAddress(certAddr), client.Eth())
	if err != nil {
		log.Fatalf("bind nft: %v", err)
	}

	// ── 멤버 목록 출력 ─────────────────────────────────────────────────────
	fmt.Printf("\n수료 기준: %d / %d회 (90%%)\n", threshold, totalSessions)
	fmt.Println("──────────────────────────────────────────────────────────")
	fmt.Printf("  %-4s  %-14s  %-12s  %-6s  %s\n", "ID", "이름", "직책", "출석", "상태")
	fmt.Println("──────────────────────────────────────────────────────────")

	type memberStatus struct {
		m          members.Member
		attendance int64
		hasCert    bool
		eligible   bool
	}
	statuses := make([]memberStatus, 0, len(mems))

	for _, m := range mems {
		bal, _ := tracker.Attendance(nil, big.NewInt(int64(m.ID)))
		att := int64(0)
		if bal != nil {
			att = bal.Int64()
		}
		has, _ := nft.HasCertificate(nil, big.NewInt(int64(m.ID)))
		eligible := att >= int64(threshold)

		status := "미달"
		if has {
			status = "✅ 발급완료"
		} else if eligible {
			status = "🏆 발급가능"
		}
		mark := "  "
		if eligible && !has {
			mark = ">>"
		}
		fmt.Printf("%s %-4d  %-14s  %-12s  %2d/%d    %s\n",
			mark, m.ID, m.Name, m.Role, att, totalSessions, status)

		statuses = append(statuses, memberStatus{m, att, has, eligible})
	}
	fmt.Println("──────────────────────────────────────────────────────────")

	// ── 멤버 선택 ──────────────────────────────────────────────────────────
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n발급할 멤버 ID 입력 (취소: Enter): ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	if input == "" {
		fmt.Println("취소.")
		return
	}
	id, err := strconv.Atoi(input)
	if err != nil || id < 0 || id >= len(statuses) {
		log.Fatalf("잘못된 ID: %q", input)
	}

	st := statuses[id]
	m := st.m

	if st.hasCert {
		fmt.Printf("\n%s 는 이미 수료증이 발급되어 있습니다.\n", m.Name)
		return
	}

	// ── 출석 부족 시 테스트 민팅 제안 ─────────────────────────────────────
	if !st.eligible {
		need := int64(threshold) - st.attendance
		fmt.Printf("\n%s 의 현재 출석(%d)이 수료 기준(%d)에 %d회 부족합니다.\n",
			m.Name, st.attendance, threshold, need)
		fmt.Printf("데모 목적으로 테스트 출석 %d회를 즉시 민팅하시겠습니까? [y/N] ", need)
		scanner.Scan()
		if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
			fmt.Println("취소.")
			return
		}

		fmt.Printf("\n테스트 출석 민팅 중 (%d회)...\n", need)
		for i := int64(1); i <= need; i++ {
			eventID := fmt.Sprintf("test-session-%02d", i)
			issueCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
			err := blockchain.WithRetry(3, func() error {
				auth, err := client.NewTransactOpts(issueCtx)
				if err != nil {
					return err
				}
				tx, err := tracker.MintAttendance(auth, []*big.Int{big.NewInt(int64(m.ID))}, eventID)
				if err != nil {
					return fmt.Errorf("mintAttendance: %w", err)
				}
				receipt, err := client.WaitMined(issueCtx, tx)
				if err != nil {
					return err
				}
				if receipt.Status == 0 {
					return fmt.Errorf("tx reverted")
				}
				return nil
			})
			cancel()
			if err != nil {
				log.Fatalf("테스트 민팅 실패 (round %d): %v", i, err)
			}
			fmt.Printf("  %d/%d 완료\n", i, need)
		}
		fmt.Println("테스트 출석 민팅 완료!")
	}

	// ── 수료증 발급 ────────────────────────────────────────────────────────
	fmt.Printf("\n%s (%s) 수료증 발급을 진행합니다.\n", m.Name, m.Role)
	fmt.Print("계속하시겠습니까? [y/N] ")
	scanner.Scan()
	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		fmt.Println("취소.")
		return
	}

	issueCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// 최신 출석 조회
	bal, _ := tracker.Attendance(nil, big.NewInt(int64(m.ID)))
	att := int64(0)
	if bal != nil {
		att = bal.Int64()
	}

	issuedDate := time.Now().Format("2006-01-02")
	cloudfrontURL := os.Getenv("CLOUDFRONT_URL")
	var credURL string
	if cloudfrontURL != "" {
		credURL = fmt.Sprintf("%s/certificate.html?member=%d", cloudfrontURL, m.ID)
	} else {
		credURL = fmt.Sprintf("https://sepolia.etherscan.io/token/%s?a=%d", certAddr, m.ID)
	}

	meta := metadata.Build(metadata.CertParams{
		MemberName:      m.Name,
		Role:            m.Role,
		IssuedDate:      issuedDate,
		CredentialURL:   credURL,
		MemberAddress:   fmt.Sprintf("memberId:%d", m.ID),
		Attendance:      int(att),
		TotalSessions:   totalSessions,
		ContractAddress: certAddr,
	})

	fmt.Print("IPFS 메타데이터 업로드 중... ")
	cid, err := ipfs.UploadJSON(issueCtx, pinataJWT, meta)
	if err != nil {
		log.Fatalf("\nIPFS 업로드 실패: %v", err)
	}
	fmt.Printf("완료\n  CID: %s\n", cid)

	fmt.Print("수료증 민팅 중... ")
	var tokenID uint64
	var txHash string

	err = blockchain.WithRetry(3, func() error {
		auth, err := client.NewTransactOpts(issueCtx)
		if err != nil {
			return err
		}
		tx, err := nft.IssueCertificate(auth, big.NewInt(int64(m.ID)), cid, m.Name)
		if err != nil {
			return fmt.Errorf("issueCertificate: %w", err)
		}
		receipt, err := client.WaitMined(issueCtx, tx)
		if err != nil {
			return err
		}
		if receipt.Status == 0 {
			return fmt.Errorf("tx reverted")
		}
		txHash = tx.Hash().Hex()
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
		log.Fatalf("\n수료증 발급 실패: %v", err)
	}

	fmt.Printf("완료!\n")
	fmt.Println("\n══════════════════════════════════════════════════════════")
	fmt.Printf("  수료증 발급 완료\n")
	fmt.Println("══════════════════════════════════════════════════════════")
	fmt.Printf("  이름     : %s (%s)\n", m.Name, m.Role)
	fmt.Printf("  출석     : %d / %d\n", att, totalSessions)
	fmt.Printf("  발급일   : %s\n", issuedDate)
	fmt.Printf("  Token ID : #%d\n", tokenID)
	fmt.Printf("  Tx Hash  : %s\n", txHash)
	fmt.Println("──────────────────────────────────────────────────────────")
	fmt.Printf("  Etherscan: https://sepolia.etherscan.io/token/%s?a=%d\n", certAddr, tokenID)
	fmt.Printf("  OpenSea  : https://testnets.opensea.io/assets/sepolia/%s/%d\n", certAddr, tokenID)
	fmt.Println("══════════════════════════════════════════════════════════")
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("환경변수 %q 가 설정되지 않았습니다", key)
	}
	return v
}
