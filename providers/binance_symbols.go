package providers

import (
	"context"
	"strings"
	"time"
)

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
	query := BinanceExchangeInfoQuery{
		Permissions: "SPOT",
	}

	var info BinanceExchangeInfo
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
