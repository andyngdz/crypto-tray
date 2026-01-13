package httpclient

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/imroc/req/v3"
)

// Config holds configuration for creating an HTTP client
type Config struct {
	BaseURL      string
	Timeout      time.Duration
	APIKeyHeader string
	APIKey       string
}

// Client wraps req.Client with provider-specific configuration
type Client struct {
	client       *req.Client
	apiKeyHeader string
	apiKey       string
}

// New creates a new HTTP client with the given configuration
func New(cfg Config) *Client {
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

	return &Client{
		client:       client,
		apiKeyHeader: cfg.APIKeyHeader,
		apiKey:       cfg.APIKey,
	}
}

// SetAPIKey updates the API key for subsequent requests
func (h *Client) SetAPIKey(key string) {
	h.apiKey = key
}

// Get performs a GET request and unmarshals the JSON response into result
func (h *Client) Get(ctx context.Context, path string, result interface{}) error {
	r := h.client.R().SetContext(ctx).SetSuccessResult(result)

	if h.apiKey != "" && h.apiKeyHeader != "" {
		r.SetHeader(h.apiKeyHeader, h.apiKey)
	}

	_, err := r.Get(path)

	return err
}

// GetWithQuery performs a GET request with query parameters.
// The params argument should be a struct with `url` tags for field mapping.
func (h *Client) GetWithQuery(ctx context.Context, path string, params any, result any) error {
	queryValues, err := query.Values(params)
	if err != nil {
		return fmt.Errorf("failed to encode query params: %w", err)
	}

	r := h.client.R().SetContext(ctx).SetSuccessResult(result)

	if h.apiKey != "" && h.apiKeyHeader != "" {
		r.SetHeader(h.apiKeyHeader, h.apiKey)
	}

	r.SetQueryParamsFromValues(queryValues)

	_, err = r.Get(path)

	return err
}
