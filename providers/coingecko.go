package providers

import (
	"context"
	"strings"
	"sync"
	"time"

	"crypto-tray/internal/httpclient"
)

// CoinGecko implements the Provider interface for the CoinGecko API
type CoinGecko struct {
	httpClient *httpclient.Client

	// Symbol cache
	cacheMu   sync.RWMutex
	symbols   []SymbolInfo
	coinIDMap map[string]SymbolInfo // keyed by coinID for reverse lookup
	cacheTime time.Time
}

// NewCoinGecko creates a new CoinGecko provider
func NewCoinGecko() *CoinGecko {
	return &CoinGecko{
		httpClient: httpclient.New(httpclient.Config{
			BaseURL:      coingeckoBaseURL,
			Timeout:      coingeckoTimeout,
			APIKeyHeader: coingeckoAPIKeyHeader,
		}),
	}
}

func (c *CoinGecko) ID() string               { return "coingecko" }
func (c *CoinGecko) Name() string             { return "CoinGecko" }
func (c *CoinGecko) RequiresAPIKey() bool     { return false }
func (c *CoinGecko) DefaultCoinIDs() []string { return []string{"bitcoin", "ethereum", "solana"} }

func (c *CoinGecko) SetAPIKey(key string) {
	c.httpClient.SetAPIKey(key)
}

// FetchPrices retrieves prices for multiple cryptocurrencies by coinID
func (c *CoinGecko) FetchPrices(ctx context.Context, coinIDs []string) ([]*PriceData, error) {
	if len(coinIDs) == 0 {
		return []*PriceData{}, nil
	}

	query := CoinGeckoPriceQuery{
		IDs:              strings.Join(coinIDs, ","),
		VsCurrencies:     coingeckoCurrency,
		Include24hChange: "true",
	}

	var result CoinGeckoPriceResponse
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
	query := CoinGeckoMarketsQuery{
		VsCurrency: coingeckoCurrency,
		Order:      "market_cap_desc",
		PerPage:    "250",
		Page:       "1",
	}

	var coins []CoinGeckoMarketCoin
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
