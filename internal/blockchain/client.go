package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const sepoliaChainID = 11155111

type Client struct {
	eth        *ethclient.Client
	privateKey *ecdsa.PrivateKey
}

func NewClient(rpcURL, privateKey string) (*Client, error) {
	eth, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("NewClient: dial rpc: %w", err)
	}
	pk, err := crypto.HexToECDSA(strings.TrimPrefix(privateKey, "0x"))
	if err != nil {
		return nil, fmt.Errorf("NewClient: parse private key: %w", err)
	}
	return &Client{eth: eth, privateKey: pk}, nil
}

func (c *Client) Eth() *ethclient.Client {
	return c.eth
}

func (c *Client) NewTransactOpts(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, big.NewInt(sepoliaChainID))
	if err != nil {
		return nil, fmt.Errorf("NewTransactOpts: %w", err)
	}
	auth.Context = ctx
	return auth, nil
}

func (c *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.eth, tx)
	if err != nil {
		return nil, fmt.Errorf("WaitMined: %w", err)
	}
	return receipt, nil
}

func WithRetry(attempts int, fn func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		if err = fn(); err == nil {
			return nil
		}
		time.Sleep(time.Duration(1<<uint(i)) * time.Second)
	}
	return fmt.Errorf("all %d attempts failed: %w", attempts, err)
}
