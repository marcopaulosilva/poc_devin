package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient interface {
	Get(ctx context.Context, url string) ([]byte, error)
	Do(req *http.Request) (*http.Response, error)
}

type HTTPClientImpl struct {
	client *http.Client
}

func NewHTTPClient(timeout time.Duration) HTTPClient {
	return &HTTPClientImpl{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HTTPClientImpl) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func (c *HTTPClientImpl) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}

func ParseJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
