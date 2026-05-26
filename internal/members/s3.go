package members

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func LoadFromS3(ctx context.Context, cfg aws.Config, bucket, key string) ([]Member, error) {
	client := s3.NewFromConfig(cfg)
	out, err := client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("LoadFromS3: get s3://%s/%s: %w", bucket, key, err)
	}
	defer out.Body.Close()
	return LoadMembersFromReader(out.Body)
}
