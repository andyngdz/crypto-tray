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

// coingeckoCoin represents a coin from the CoinGecko /coins/list API
type coingeckoCoin struct {
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
	cacheMu   sync.RWMutex
	symbols   []SymbolInfo
	symbolMap map[string]SymbolInfo
	cacheTime time.Time
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

func (c *CoinGecko) SetAPIKey(key string) {
	c.httpClient.SetAPIKey(key)
}

// FetchPrices retrieves prices for multiple cryptocurrencies
func (c *CoinGecko) FetchPrices(ctx context.Context, symbols []string) ([]*PriceData, error) {
	if len(symbols) == 0 {
		return []*PriceData{}, nil
	}

	coinIDs := make([]string, 0, len(symbols))
	for _, symbol := range symbols {
		coinIDs = append(coinIDs, c.symbolToCoinID(symbol))
	}

	query := map[string]string{
		"ids":               strings.Join(coinIDs, ","),
		"vs_currencies":     coingeckoCurrency,
		"include_24hr_change": "true",
	}

	var result coingeckoPriceResponse
	if err := c.httpClient.GetWithQuery(ctx, "/simple/price", query, &result); err != nil {
		return nil, err
	}

	prices := make([]*PriceData, 0, len(symbols))
	for _, symbol := range symbols {
		coinID := c.symbolToCoinID(symbol)
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

// FetchSymbols fetches the list of supported cryptocurrencies from CoinGecko API
func (c *CoinGecko) FetchSymbols(ctx context.Context) ([]SymbolInfo, error) {
	// Check cache first
	c.cacheMu.RLock()
	if len(c.symbols) > 0 && time.Since(c.cacheTime) < coingeckoCacheTTL {
		symbols := c.symbols
		c.cacheMu.RUnlock()
		return symbols, nil
	}
	c.cacheMu.RUnlock()

	// Fetch from API
	var coins []coingeckoCoin
	if err := c.httpClient.Get(ctx, "/coins/list", &coins); err != nil {
		return []SymbolInfo{}, nil
	}

	// Map to SymbolInfo and build lookup map
	symbols := make([]SymbolInfo, 0, len(coins))
	symbolMap := make(map[string]SymbolInfo, len(coins))

	for _, coin := range coins {
		info := SymbolInfo{
			CoinID: coin.ID,
			Symbol: strings.ToUpper(coin.Symbol),
			Name:   coin.Name,
		}
		symbols = append(symbols, info)
		symbolMap[info.Symbol] = info
	}

	// Update cache
	c.cacheMu.Lock()
	c.symbols = symbols
	c.symbolMap = symbolMap
	c.cacheTime = time.Now()
	c.cacheMu.Unlock()

	return symbols, nil
}

// symbolToCoinID maps a symbol to its CoinGecko coin ID using cached data
func (c *CoinGecko) symbolToCoinID(symbol string) string {
	upperSymbol := strings.ToUpper(symbol)

	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	if info, ok := c.symbolMap[upperSymbol]; ok {
		return info.CoinID
	}
	return strings.ToLower(symbol)
}
