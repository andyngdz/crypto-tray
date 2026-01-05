package services

import (
	"context"
	"fmt"
	"time"

	"github.com/imroc/req/v3"
)

// HTTPClientConfig holds configuration for creating an HTTP client
type HTTPClientConfig struct {
	BaseURL      string
	Timeout      time.Duration
	APIKeyHeader string
	APIKey       string
}

// HTTPClient wraps req.Client with provider-specific configuration
type HTTPClient struct {
	client       *req.Client
	apiKeyHeader string
	apiKey       string
}

// NewHTTPClient creates a new HTTP client with the given configuration
func NewHTTPClient(cfg HTTPClientConfig) *HTTPClient {
	client := req.C().
		SetBaseURL(cfg.BaseURL).
		SetTimeout(cfg.Timeout).
		OnAfterResponse(func(c *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = fmt.Errorf("API returned status %d: %s", resp.StatusCode, resp.Status)
			}
			return nil
		})

	return &HTTPClient{
		client:       client,
		apiKeyHeader: cfg.APIKeyHeader,
		apiKey:       cfg.APIKey,
	}
}

// SetAPIKey updates the API key for subsequent requests
func (h *HTTPClient) SetAPIKey(key string) {
	h.apiKey = key
}

// Get performs a GET request and unmarshals the JSON response into result
func (h *HTTPClient) Get(ctx context.Context, path string, result interface{}) error {
	r := h.client.R().SetContext(ctx).SetSuccessResult(result)

	if h.apiKey != "" && h.apiKeyHeader != "" {
		r.SetHeader(h.apiKeyHeader, h.apiKey)
	}

	_, err := r.Get(path)
	return err
}

// GetWithQuery performs a GET request with query parameters
func (h *HTTPClient) GetWithQuery(ctx context.Context, path string, query map[string]string, result interface{}) error {
	r := h.client.R().SetContext(ctx).SetSuccessResult(result)

	if h.apiKey != "" && h.apiKeyHeader != "" {
		r.SetHeader(h.apiKeyHeader, h.apiKey)
	}

	for k, v := range query {
		r.SetQueryParam(k, v)
	}

	_, err := r.Get(path)
	return err
}
