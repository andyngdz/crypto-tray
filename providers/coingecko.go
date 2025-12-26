package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// CoinGecko API constants
const (
	coingeckoBaseURL      = "https://api.coingecko.com/api/v3"
	coingeckoCurrency     = "usd"
	coingeckoAPIKeyHeader = "x-cg-demo-api-key"
	coingeckoTimeout      = 10 * time.Second
)

// CoinGecko implements the Provider interface for the CoinGecko API
type CoinGecko struct {
	apiKey     string
	httpClient *http.Client
}

// NewCoinGecko creates a new CoinGecko provider
func NewCoinGecko() *CoinGecko {
	return &CoinGecko{
		httpClient: &http.Client{Timeout: coingeckoTimeout},
	}
}

func (c *CoinGecko) ID() string   { return "coingecko" }
func (c *CoinGecko) Name() string { return "CoinGecko" }

func (c *CoinGecko) FetchPrice(ctx context.Context, symbol string) (*PriceData, error) {
	coinID := symbolToCoinID(symbol)

	url := fmt.Sprintf(
		"%s/simple/price?ids=%s&vs_currencies=%s&include_24hr_change=true",
		coingeckoBaseURL,
		coinID,
		coingeckoCurrency,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Add API key header if available (for higher rate limits)
	if c.apiKey != "" {
		req.Header.Set(coingeckoAPIKeyHeader, c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result map[string]struct {
		USD         float64 `json:"usd"`
		USDChange24 float64 `json:"usd_24h_change"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	data, ok := result[coinID]
	if !ok {
		return nil, fmt.Errorf("no data for %s", symbol)
	}

	return &PriceData{
		Symbol:    symbol,
		Price:     data.USD,
		Change24h: data.USDChange24,
	}, nil
}

func (c *CoinGecko) RequiresAPIKey() bool { return false }
func (c *CoinGecko) SetAPIKey(key string) { c.apiKey = key }

// symbolToCoinID maps common symbols to CoinGecko coin IDs
func symbolToCoinID(symbol string) string {
	mapping := map[string]string{
		"BTC":  "bitcoin",
		"ETH":  "ethereum",
		"USDT": "tether",
		"BNB":  "binancecoin",
		"SOL":  "solana",
		"XRP":  "ripple",
		"USDC": "usd-coin",
		"ADA":  "cardano",
		"DOGE": "dogecoin",
		"AVAX": "avalanche-2",
	}
	if id, ok := mapping[strings.ToUpper(symbol)]; ok {
		return id
	}
	return strings.ToLower(symbol)
}
