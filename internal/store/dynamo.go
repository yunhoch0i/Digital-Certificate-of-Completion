package store

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const tableName = "asbg-tx-history"

type TxRecord struct {
	TxHash    string
	Timestamp string
	Action    string
	Member    string
	Name      string
	EventID   string
}

type DynamoStore struct {
	client *dynamodb.Client
}

func NewDynamoStore(cfg aws.Config) *DynamoStore {
	return &DynamoStore{client: dynamodb.NewFromConfig(cfg)}
}

func (s *DynamoStore) PutTx(ctx context.Context, r TxRecord) error {
	if r.Timestamp == "" {
		r.Timestamp = time.Now().UTC().Format(time.RFC3339)
	}
	_, err := s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"tx_hash":   &types.AttributeValueMemberS{Value: r.TxHash},
			"timestamp": &types.AttributeValueMemberS{Value: r.Timestamp},
			"action":    &types.AttributeValueMemberS{Value: r.Action},
			"member":    &types.AttributeValueMemberS{Value: r.Member},
			"name":      &types.AttributeValueMemberS{Value: r.Name},
			"event_id":  &types.AttributeValueMemberS{Value: r.EventID},
		},
	})
	if err != nil {
		return fmt.Errorf("PutTx: %w", err)
	}
	return nil
}
