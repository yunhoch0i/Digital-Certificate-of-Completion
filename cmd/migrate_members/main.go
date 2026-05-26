package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"

	"asbg-token/internal/members"
)

func main() {
	_ = godotenv.Load(".env")

	localPath := "members.csv"
	if len(os.Args) > 1 {
		localPath = os.Args[1]
	}
	bucket := os.Getenv("DATA_BUCKET")
	if bucket == "" {
		bucket = "asbg-data"
	}

	// validate before upload
	mems, err := members.LoadMembers(localPath)
	if err != nil {
		log.Fatalf("validate members.csv: %v", err)
	}
	fmt.Printf("검증 완료: %d명\n", len(mems))

	ctx := context.Background()
	cfg, err := awscfg.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("aws config: %v", err)
	}

	f, err := os.Open(localPath)
	if err != nil {
		log.Fatalf("open %s: %v", localPath, err)
	}
	defer f.Close()

	client := s3.NewFromConfig(cfg)
	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String("members.csv"),
		Body:        f,
		ContentType: aws.String("text/csv"),
	})
	if err != nil {
		log.Fatalf("upload to S3: %v", err)
	}
	fmt.Printf("업로드 완료: s3://%s/members.csv\n", bucket)
}
