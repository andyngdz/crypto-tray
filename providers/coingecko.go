package providers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
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

// CoinGecko implements the Provider interface for the CoinGecko API
type CoinGecko struct {
	apiKey     string
	httpClient *http.Client

	// Symbol cache
	cacheMu    sync.RWMutex
	symbols    []SymbolInfo
	symbolMap  map[string]SymbolInfo // symbol -> SymbolInfo lookup
	cacheTime  time.Time
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
		coinIDs = append(coinIDs, c.symbolToCoinID(symbol))
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
	url := fmt.Sprintf("%s/coins/list", coingeckoBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	if c.apiKey != "" {
		req.Header.Set(coingeckoAPIKeyHeader, c.apiKey)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("FetchSymbols: API request failed: %v", err)
		return []SymbolInfo{}, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("FetchSymbols: API returned status %d", resp.StatusCode)
		return []SymbolInfo{}, nil
	}

	var coins []coingeckoCoin

	if err := json.NewDecoder(resp.Body).Decode(&coins); err != nil {
		log.Printf("FetchSymbols: failed to decode response: %v", err)
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
