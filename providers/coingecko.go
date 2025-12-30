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

func (c *CoinGecko) ID() string           { return "coingecko" }
func (c *CoinGecko) Name() string         { return "CoinGecko" }
func (c *CoinGecko) RequiresAPIKey() bool { return false }
func (c *CoinGecko) SetAPIKey(key string) { c.apiKey = key }

// FetchPrices retrieves prices for multiple cryptocurrencies in a single API call
func (c *CoinGecko) FetchPrices(ctx context.Context, symbols []string) ([]*PriceData, error) {
	if len(symbols) == 0 {
		return []*PriceData{}, nil
	}

	coinIDs := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		coinIDs = append(coinIDs, symbolToCoinID(symbol))
	}

	ids := strings.Join(coinIDs, ",")
	url := fmt.Sprintf(
		"%s/simple/price?ids=%s&vs_currencies=%s&include_24hr_change=true",
		coingeckoBaseURL,
		ids,
		coingeckoCurrency,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

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

	prices := make([]*PriceData, 0, len(symbols))
	for _, symbol := range symbols {
		coinID := symbolToCoinID(symbol)
		data, ok := result[coinID]
		if !ok {
			continue
		}

		prices = append(prices, &PriceData{
			Symbol:    symbol,
			Price:     data.USD,
			Change24h: data.USDChange24,
		})
	}

	return prices, nil
}

// GetSupportedSymbols returns list of supported cryptocurrencies
func (c *CoinGecko) GetSupportedSymbols() []SymbolInfo {
	return []SymbolInfo{
		{ID: "BTC", Name: "Bitcoin"},
		{ID: "ETH", Name: "Ethereum"},
		{ID: "SOL", Name: "Solana"},
		{ID: "ADA", Name: "Cardano"},
		{ID: "DOT", Name: "Polkadot"},
		{ID: "LINK", Name: "Chainlink"},
		{ID: "AVAX", Name: "Avalanche"},
		{ID: "MATIC", Name: "Polygon"},
		{ID: "ATOM", Name: "Cosmos"},
		{ID: "XRP", Name: "Ripple"},
		{ID: "USDT", Name: "Tether"},
		{ID: "USDC", Name: "USD Coin"},
		{ID: "BNB", Name: "BNB"},
		{ID: "DOGE", Name: "Dogecoin"},
	}
}

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
