package providers

import (
	"context"
	"encoding/json"
	"strconv"
	"sync"
	"time"

	"crypto-tray/internal/httpclient"
)

// Binance implements the Provider interface for the Binance API
type Binance struct {
	httpClient *httpclient.Client

	// Symbol cache
	cacheMu   sync.RWMutex
	symbols   []SymbolInfo
	coinIDMap map[string]SymbolInfo
	cacheTime time.Time
}

// NewBinance creates a new Binance provider
func NewBinance() *Binance {
	return &Binance{
		httpClient: httpclient.New(httpclient.Config{
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

	query := BinanceTickerQuery{
		Symbols: string(symbolsJSON),
	}

	var tickers []BinanceTicker24hr
	if err := b.httpClient.GetWithQuery(ctx, "/api/v3/ticker/24hr", query, &tickers); err != nil {
		return nil, err
	}

	// Build response map for quick lookup
	tickerMap := make(map[string]BinanceTicker24hr)
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
