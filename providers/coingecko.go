package providers

import (
	"context"
	"strings"
	"sync"
	"time"

	"crypto-tray/services"
)

// CoinGecko API constants
const (
	coingeckoBaseURL      = "https://api.coingecko.com/api/v3"
	coingeckoCurrency     = "usd"
	coingeckoAPIKeyHeader = "x-cg-demo-api-key"
	coingeckoTimeout      = 10 * time.Second
	coingeckoCacheTTL     = 24 * time.Hour
)

// coingeckoMarketCoin represents a coin from the CoinGecko /coins/markets API
type coingeckoMarketCoin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// coingeckoPriceResponse represents the price API response
type coingeckoPriceResponse map[string]struct {
	USD         float64 `json:"usd"`
	USDChange24 float64 `json:"usd_24h_change"`
}

// CoinGecko implements the Provider interface for the CoinGecko API
type CoinGecko struct {
	httpClient *services.HTTPClient

	// Symbol cache
	cacheMu    sync.RWMutex
	symbols    []SymbolInfo
	coinIDMap  map[string]SymbolInfo // keyed by coinID for reverse lookup
	cacheTime  time.Time
}

// NewCoinGecko creates a new CoinGecko provider
func NewCoinGecko() *CoinGecko {
	return &CoinGecko{
		httpClient: services.NewHTTPClient(services.HTTPClientConfig{
			BaseURL:      coingeckoBaseURL,
			Timeout:      coingeckoTimeout,
			APIKeyHeader: coingeckoAPIKeyHeader,
		}),
	}
}

func (c *CoinGecko) ID() string           { return "coingecko" }
func (c *CoinGecko) Name() string         { return "CoinGecko" }
func (c *CoinGecko) RequiresAPIKey() bool { return false }
func (c *CoinGecko) DefaultCoinID() string { return "bitcoin" }

func (c *CoinGecko) SetAPIKey(key string) {
	c.httpClient.SetAPIKey(key)
}

// FetchPrices retrieves prices for multiple cryptocurrencies by coinID
func (c *CoinGecko) FetchPrices(ctx context.Context, coinIDs []string) ([]*PriceData, error) {
	if len(coinIDs) == 0 {
		return []*PriceData{}, nil
	}

	query := map[string]string{
		"ids":                  strings.Join(coinIDs, ","),
		"vs_currencies":        coingeckoCurrency,
		"include_24hr_change": "true",
	}

	var result coingeckoPriceResponse
	if err := c.httpClient.GetWithQuery(ctx, "/simple/price", query, &result); err != nil {
		return nil, err
	}

	prices := make([]*PriceData, 0, len(coinIDs))
	for coinIdx := range coinIDs {
		coinID := coinIDs[coinIdx]
		data, ok := result[coinID]
		if !ok {
			continue
		}
		prices = append(prices, &PriceData{
			CoinID:    coinID,
			Symbol:    c.coinIDToSymbol(coinID),
			Price:     data.USD,
			Change24h: data.USDChange24,
		})
	}

	return prices, nil
}

// FetchSymbols fetches top cryptocurrencies by market cap from CoinGecko API
func (c *CoinGecko) FetchSymbols(ctx context.Context) ([]SymbolInfo, error) {
	// Check cache first
	c.cacheMu.RLock()
	if len(c.symbols) > 0 && time.Since(c.cacheTime) < coingeckoCacheTTL {
		symbols := c.symbols
		c.cacheMu.RUnlock()
		return symbols, nil
	}
	c.cacheMu.RUnlock()

	// Fetch top 250 coins by market cap
	query := map[string]string{
		"vs_currency": coingeckoCurrency,
		"order":       "market_cap_desc",
		"per_page":    "250",
		"page":        "1",
	}

	var coins []coingeckoMarketCoin
	if err := c.httpClient.GetWithQuery(ctx, "/coins/markets", query, &coins); err != nil {
		return []SymbolInfo{}, nil
	}

	// Map to SymbolInfo and build lookup map
	symbols := make([]SymbolInfo, 0, len(coins))
	coinIDMap := make(map[string]SymbolInfo, len(coins))

	for coinIdx := range coins {
		coin := coins[coinIdx]
		info := SymbolInfo{
			CoinID: coin.ID,
			Symbol: strings.ToUpper(coin.Symbol),
			Name:   coin.Name,
		}
		symbols = append(symbols, info)
		coinIDMap[coin.ID] = info
	}

	// Update cache
	c.cacheMu.Lock()
	c.symbols = symbols
	c.coinIDMap = coinIDMap
	c.cacheTime = time.Now()
	c.cacheMu.Unlock()

	return symbols, nil
}

// coinIDToSymbol maps a coinID to its display symbol
func (c *CoinGecko) coinIDToSymbol(coinID string) string {
	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	if info, ok := c.coinIDMap[coinID]; ok {
		return info.Symbol
	}
	return strings.ToUpper(coinID)
}
