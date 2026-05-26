package secret

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Manager struct {
	client *secretsmanager.Client
	mu     sync.RWMutex
	cache  map[string]string
}

func NewManager(cfg aws.Config) *Manager {
	return &Manager{
		client: secretsmanager.NewFromConfig(cfg),
		cache:  make(map[string]string),
	}
}

func (m *Manager) Get(ctx context.Context, key string) (string, error) {
	m.mu.RLock()
	if v, ok := m.cache[key]; ok {
		m.mu.RUnlock()
		return v, nil
	}
	m.mu.RUnlock()

	out, err := m.client.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("secret.Get %q: %w", key, err)
	}
	val := aws.ToString(out.SecretString)

	m.mu.Lock()
	m.cache[key] = val
	m.mu.Unlock()

	return val, nil
}
