package providers

import "time"

// Binance API constants
const (
	binanceBaseURL    = "https://api.binance.com"
	binanceTimeout    = 10 * time.Second
	binanceCacheTTL   = 24 * time.Hour
	binanceQuoteAsset = "USDT"
)

// BinanceTickerQuery represents query params for /api/v3/ticker/24hr endpoint
type BinanceTickerQuery struct {
	Symbols string `url:"symbols"`
}

// BinanceExchangeInfoQuery represents query params for /api/v3/exchangeInfo endpoint
type BinanceExchangeInfoQuery struct {
	Permissions string `url:"permissions"`
}

// BinanceTicker24hr represents a single ticker from /api/v3/ticker/24hr
type BinanceTicker24hr struct {
	Symbol             string `json:"symbol"`
	LastPrice          string `json:"lastPrice"`
	PriceChangePercent string `json:"priceChangePercent"`
}

// BinanceExchangeInfo represents the response from /api/v3/exchangeInfo
type BinanceExchangeInfo struct {
	Symbols []BinanceSymbol `json:"symbols"`
}

// BinanceSymbol represents a trading pair from exchangeInfo
type BinanceSymbol struct {
	Symbol     string `json:"symbol"`
	BaseAsset  string `json:"baseAsset"`
	QuoteAsset string `json:"quoteAsset"`
	Status     string `json:"status"`
}
