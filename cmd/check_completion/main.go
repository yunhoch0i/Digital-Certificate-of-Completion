package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"

	tracker "asbg-token/bindings/attendance_tracker"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/members"
	"asbg-token/internal/metadata"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	rpcURL := mustEnv("RPC_URL")
	privateKey := mustEnv("PRIVATE_KEY")
	trackerAddrHex := mustEnv("ATTENDANCE_TRACKER_ADDRESS")
	totalSessionsStr := mustEnv("TOTAL_SESSIONS")

	var totalSessions int
	if _, err := fmt.Sscan(totalSessionsStr, &totalSessions); err != nil || totalSessions <= 0 {
		log.Fatalf("TOTAL_SESSIONS must be a positive integer")
	}
	threshold := metadata.CalcThreshold(totalSessions)

	mems, err := members.LoadMembers("members.csv")
	if err != nil {
		log.Fatalf("load members: %v", err)
	}

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		log.Fatalf("blockchain client: %v", err)
	}
	trackerAddr := common.HexToAddress(trackerAddrHex)
	atk, err := tracker.NewAttendanceTracker(trackerAddr, client.Eth())
	if err != nil {
		log.Fatalf("bind tracker: %v", err)
	}

	type result struct {
		members.Member
		Attendance int
		Completed  bool
	}

	var results []result
	for _, m := range mems {
		bal, err := atk.Attendance(nil, big.NewInt(int64(m.ID)))
		att := 0
		if err != nil {
			log.Printf("attendance(%d=%s): %v", m.ID, m.Name, err)
		} else {
			att = int(bal.Int64())
		}
		results = append(results, result{
			Member:     m,
			Attendance: att,
			Completed:  att >= threshold,
		})
	}

	fmt.Printf("ASBG 수료 확인 결과 (기준: %d회 / %d회, 90%%)\n", threshold, totalSessions)
	fmt.Println("──────────────────────────────────────────────")
	fmt.Printf("%-12s %-14s %4s %4s %4s\n", "이름", "직책", "출석", "기준", "수료")
	pass, fail := 0, 0
	for _, r := range results {
		mark := "X"
		if r.Completed {
			mark = "O"
			pass++
		} else {
			fail++
		}
		fmt.Printf("%-12s %-14s %4d %4d %4s\n", r.Name, r.Role, r.Attendance, threshold, mark)
	}
	fmt.Println("──────────────────────────────────────────────")
	fmt.Printf("수료: %d명 / 미수료: %d명\n", pass, fail)

	outFile, err := os.Create("completion_result.csv")
	if err != nil {
		log.Fatalf("create completion_result.csv: %v", err)
	}
	defer outFile.Close()
	w := csv.NewWriter(outFile)
	_ = w.Write([]string{"id", "name", "role", "attendance", "threshold", "completed"})
	for _, r := range results {
		_ = w.Write([]string{
			strconv.Itoa(int(r.ID)), r.Name, r.Role,
			strconv.Itoa(r.Attendance),
			strconv.Itoa(threshold),
			strconv.FormatBool(r.Completed),
		})
	}
	w.Flush()
	fmt.Println("결과 저장: completion_result.csv")
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("환경변수 %q 가 설정되지 않았습니다", key)
	}
	return v
}
