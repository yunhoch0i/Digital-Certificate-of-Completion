package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/ethereum/go-ethereum/common"

	atk "asbg-token/bindings/attendance_tracker"
	cert "asbg-token/bindings/certificate_nft"
	"asbg-token/internal/blockchain"
	"asbg-token/internal/ipfs"
	"asbg-token/internal/members"
	"asbg-token/internal/metadata"
	"asbg-token/internal/secret"
	"asbg-token/internal/store"
)

var (
	rpcURL         string
	privateKey     string
	trackerAddress string
	certAddress    string
	pinataJWT      string
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
	cfg, err := awscfg.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("init: %v", err)
	}
	sm := secret.NewManager(cfg)
	dynStore = store.NewDynamoStore(cfg)

	rpcURL = mustGet(ctx, sm, "asbg/RPC_URL")
	privateKey = mustGet(ctx, sm, "asbg/PRIVATE_KEY")
	trackerAddress = mustGet(ctx, sm, "asbg/ATTENDANCE_TRACKER_ADDRESS")
	certAddress = mustGet(ctx, sm, "asbg/CERTIFICATE_NFT_ADDRESS")
	pinataJWT = mustGet(ctx, sm, "asbg/PINATA_JWT")
}

func Handler(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Minute)
	defer cancel()

	totalSessions, err := strconv.Atoi(os.Getenv("TOTAL_SESSIONS"))
	if err != nil || totalSessions <= 0 {
		return fmt.Errorf("TOTAL_SESSIONS env invalid")
	}
	threshold := metadata.CalcThreshold(totalSessions)

	cfg, err := awscfg.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("aws config: %w", err)
	}
	bucket := os.Getenv("DATA_BUCKET")
	if bucket == "" {
		bucket = "asbg-data"
	}

	mems, err := members.LoadFromS3(ctx, cfg, bucket, "members.csv")
	if err != nil {
		return fmt.Errorf("load members: %w", err)
	}

	client, err := blockchain.NewClient(rpcURL, privateKey)
	if err != nil {
		return fmt.Errorf("blockchain client: %w", err)
	}
	trackerAddr := common.HexToAddress(trackerAddress)
	tracker, err := atk.NewAttendanceTracker(trackerAddr, client.Eth())
	if err != nil {
		return fmt.Errorf("bind tracker: %w", err)
	}
	certAddr := common.HexToAddress(certAddress)
	nft, err := cert.NewCertificateNFT(certAddr, client.Eth())
	if err != nil {
		return fmt.Errorf("bind nft: %w", err)
	}

	issuedDate := time.Now().Format("2006-01-02")
	var guides []string
	issued, skipped, failed := 0, 0, 0

	for _, m := range mems {
		memberID := big.NewInt(int64(m.ID))

		bal, err := tracker.Attendance(nil, memberID)
		if err != nil || bal.Int64() < int64(threshold) {
			continue
		}
		has, err := nft.HasCertificate(nil, memberID)
		if err != nil || has {
			skipped++
			continue
		}

		cloudfrontURL := os.Getenv("CLOUDFRONT_URL")
		var credURL string
		if cloudfrontURL != "" {
			credURL = fmt.Sprintf("%s/certificate.html?member=%d", cloudfrontURL, m.ID)
		} else {
			credURL = fmt.Sprintf("https://sepolia.etherscan.io/token/%s?a=%d", certAddress, m.ID)
		}
		meta := metadata.Build(metadata.CertParams{
			MemberName:      m.Name,
			Role:            m.Role,
			IssuedDate:      issuedDate,
			CredentialURL:   credURL,
			MemberAddress:   fmt.Sprintf("memberId:%d", m.ID),
			Attendance:      int(bal.Int64()),
			TotalSessions:   totalSessions,
			ContractAddress: certAddress,
		})

		cid, err := ipfs.UploadJSON(ctx, pinataJWT, meta)
		if err != nil {
			log.Printf("[FAIL] %s: upload IPFS: %v", m.Name, err)
			failed++
			continue
		}

		var tokenID uint64
		var txHash string

		err = blockchain.WithRetry(3, func() error {
			auth, err := client.NewTransactOpts(ctx)
			if err != nil {
				return err
			}
			tx, err := nft.IssueCertificate(auth, memberID, cid, m.Name)
			if err != nil {
				return fmt.Errorf("issueCertificate: %w", err)
			}
			receipt, err := client.WaitMined(ctx, tx)
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
			log.Printf("[FAIL] %s: %v", m.Name, err)
			failed++
			continue
		}

		_ = dynStore.PutTx(ctx, store.TxRecord{
			TxHash:  txHash,
			Action:  "mint_certificate",
			Member:  fmt.Sprintf("memberId:%d", m.ID),
			Name:    m.Name,
			EventID: fmt.Sprintf("yearly-%s", issuedDate),
		})

		credURLFull := fmt.Sprintf("%s%d", credURL, tokenID)
		guides = append(guides, fmt.Sprintf(
			"=== %s (%s) ===\n자격증 이름 : ASBG 수료 인증서\n발급 기관   : AWS Student Builders Group\n발급일      : %s\n만료일      : (입력 안 함 — 영구)\n자격증 URL  : %s\n",
			m.Name, m.Role, issuedDate, credURLFull,
		))
		issued++
	}

	log.Printf("발행: %d / 스킵: %d / 실패: %d", issued, skipped, failed)

	if len(guides) > 0 {
		year := time.Now().Format("2006")
		content := strings.Join(guides, "\n")
		s3Client := s3.NewFromConfig(cfg)
		key := fmt.Sprintf("results/linkedin_guide_%s.txt", year)
		_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(key),
			Body:        strings.NewReader(content),
			ContentType: aws.String("text/plain; charset=utf-8"),
			ACL:         s3types.ObjectCannedACLPrivate,
		})
		if err != nil {
			log.Printf("upload linkedin guide: %v", err)
		} else {
			log.Printf("LinkedIn 가이드: s3://%s/%s", bucket, key)
		}
	}
	return nil
}

func main() {
	lambda.Start(Handler)
}
