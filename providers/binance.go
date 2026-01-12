package providers

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"sync"
	"time"

	"crypto-tray/services"
)

// Binance API constants
const (
	binanceBaseURL    = "https://api.binance.com"
	binanceTimeout    = 10 * time.Second
	binanceCacheTTL   = 24 * time.Hour
	binanceQuoteAsset = "USDT"
)

// binanceTicker24hr represents a single ticker from /api/v3/ticker/24hr
type binanceTicker24hr struct {
	Symbol             string `json:"symbol"`
	LastPrice          string `json:"lastPrice"`
	PriceChangePercent string `json:"priceChangePercent"`
}

// binanceExchangeInfo represents the response from /api/v3/exchangeInfo
type binanceExchangeInfo struct {
	Symbols []binanceSymbol `json:"symbols"`
}

// binanceSymbol represents a trading pair from exchangeInfo
type binanceSymbol struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	Status     string `json:"status"`
}

// Binance implements the Provider interface for the Binance API
type Binance struct {
	httpClient *services.HTTPClient

	// Symbol cache
	cacheMu   sync.RWMutex
	symbols   []SymbolInfo
	coinIDMap map[string]SymbolInfo
	cacheTime time.Time
}

// NewBinance creates a new Binance provider
func NewBinance() *Binance {
	return &Binance{
		httpClient: services.NewHTTPClient(services.HTTPClientConfig{
			BaseURL: binanceBaseURL,
			Timeout: binanceTimeout,
		}),
	}
}

func (b *Binance) ID() string               { return "binance" }
func (b *Binance) Name() string             { return "Binance" }
func (b *Binance) RequiresAPIKey() bool     { return false }
func (b *Binance) SetAPIKey(key string)     {}
func (b *Binance) DefaultCoinIDs() []string { return []string{"BTCUSDT", "ETHUSDT", "SOLUSDT"} }

// FetchPrices retrieves prices for multiple cryptocurrencies by coinID (trading pair)
func (b *Binance) FetchPrices(ctx context.Context, coinIDs []string) ([]*PriceData, error) {
	if len(coinIDs) == 0 {
		return []*PriceData{}, nil
	}

	// Build JSON array format for symbols query parameter
	symbolsJSON, err := json.Marshal(coinIDs)
	if err != nil {
		return nil, err
	}

	query := map[string]string{
		"symbols": string(symbolsJSON),
	}

	var tickers []binanceTicker24hr
	if err := b.httpClient.GetWithQuery(ctx, "/api/v3/ticker/24hr", query, &tickers); err != nil {
		return nil, err
	}

	// Build response map for quick lookup
	tickerMap := make(map[string]binanceTicker24hr)
	for tickerIdx := range tickers {
		t := tickers[tickerIdx]
		tickerMap[t.Symbol] = t
	}

	prices := make([]*PriceData, 0, len(coinIDs))
	for coinIdx := range coinIDs {
		coinID := coinIDs[coinIdx]
		ticker, ok := tickerMap[coinID]
		if !ok {
			continue
		}

		price, err := strconv.ParseFloat(ticker.LastPrice, 64)
		if err != nil {
			continue
		}

		change, err := strconv.ParseFloat(ticker.PriceChangePercent, 64)
		if err != nil {
			change = 0
		}

		prices = append(prices, &PriceData{
			CoinID:    coinID,
			Symbol:    b.coinIDToSymbol(coinID),
			Price:     price,
			Change24h: change,
		})
	}

	return prices, nil
}

// FetchSymbols fetches USDT trading pairs from Binance API
func (b *Binance) FetchSymbols(ctx context.Context) ([]SymbolInfo, error) {
	// Check cache first
	b.cacheMu.RLock()
	if len(b.symbols) > 0 && time.Since(b.cacheTime) < binanceCacheTTL {
		symbols := b.symbols
		b.cacheMu.RUnlock()
		return symbols, nil
	}
	b.cacheMu.RUnlock()

	// Fetch exchange info with SPOT permission filter
	query := map[string]string{
		"permissions": "SPOT",
	}

	var info binanceExchangeInfo
	if err := b.httpClient.GetWithQuery(ctx, "/api/v3/exchangeInfo", query, &info); err != nil {
		return []SymbolInfo{}, nil
	}

	// Filter to USDT pairs only and build symbol list
	symbols := make([]SymbolInfo, 0)
	coinIDMap := make(map[string]SymbolInfo)
	seen := make(map[string]bool)

	for symbolIdx := range info.Symbols {
		s := info.Symbols[symbolIdx]
		// Only include USDT pairs that are actively trading
		if s.QuoteAsset != binanceQuoteAsset || s.Status != "TRADING" {
			continue
		}

		// Skip if we've already seen this base asset
		if seen[s.BaseAsset] {
			continue
		}
		seen[s.BaseAsset] = true

		symbolInfo := SymbolInfo{
			CoinID: s.Symbol,
			Symbol: s.BaseAsset,
			Name:   s.BaseAsset,
		}
		symbols = append(symbols, symbolInfo)
		coinIDMap[s.Symbol] = symbolInfo
	}

	// Update cache
	b.cacheMu.Lock()
	b.symbols = symbols
	b.coinIDMap = coinIDMap
	b.cacheTime = time.Now()
	b.cacheMu.Unlock()

	return symbols, nil
}

// coinIDToSymbol maps a coinID (e.g., "BTCUSDT") to its display symbol (e.g., "BTC")
func (b *Binance) coinIDToSymbol(coinID string) string {
	b.cacheMu.RLock()
	defer b.cacheMu.RUnlock()

	if info, ok := b.coinIDMap[coinID]; ok {
		return info.Symbol
	}
	// Fallback: strip quote asset suffix
	return strings.TrimSuffix(coinID, binanceQuoteAsset)
}
