package ipfs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const pinataURL = "https://api.pinata.cloud/pinning/pinJSONToIPFS"

type pinataResponse struct {
	IpfsHash string `json:"IpfsHash"`
}

func UploadJSON(ctx context.Context, jwt string, data any) (string, error) {
	body, err := json.Marshal(map[string]any{
		"pinataContent": data,
	})
	if err != nil {
		return "", fmt.Errorf("UploadJSON: marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, pinataURL, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("UploadJSON: build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+jwt)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("UploadJSON: http post: %w", err)
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("UploadJSON: pinata status %d: %s", resp.StatusCode, string(raw))
	}

	var result pinataResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return "", fmt.Errorf("UploadJSON: decode response: %w", err)
	}
	if result.IpfsHash == "" {
		return "", fmt.Errorf("UploadJSON: empty CID in response")
	}
	return result.IpfsHash, nil
}
