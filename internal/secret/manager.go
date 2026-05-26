package secret

import (
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// Manager fetches SecureString parameters from AWS SSM Parameter Store.
// Parameter names are stored as /asbg/<KEY> (e.g. /asbg/PRIVATE_KEY).
// Values are cached in-process after the first fetch (Lambda warm reuse).
type Manager struct {
	client *ssm.Client
	mu     sync.RWMutex
	cache  map[string]string
}

func NewManager(cfg aws.Config) *Manager {
	return &Manager{
		client: ssm.NewFromConfig(cfg),
		cache:  make(map[string]string),
	}
}

// Get fetches the parameter at /<key> (e.g. key="asbg/PRIVATE_KEY" → "/asbg/PRIVATE_KEY").
func (m *Manager) Get(ctx context.Context, key string) (string, error) {
	m.mu.RLock()
	if v, ok := m.cache[key]; ok {
		m.mu.RUnlock()
		return v, nil
	}
	m.mu.RUnlock()

	path := "/" + key
	out, err := m.client.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(true), // required for SecureString
	})
	if err != nil {
		return "", fmt.Errorf("secret.Get %q: %w", path, err)
	}
	val := aws.ToString(out.Parameter.Value)

	m.mu.Lock()
	m.cache[key] = val
	m.mu.Unlock()

	return val, nil
}
