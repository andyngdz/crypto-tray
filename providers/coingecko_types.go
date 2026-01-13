package providers

import "time"

// CoinGecko API constants
const (
	coingeckoBaseURL      = "https://api.coingecko.com/api/v3"
	coingeckoCurrency     = "usd"
	coingeckoAPIKeyHeader = "x-cg-demo-api-key"
	coingeckoTimeout      = 10 * time.Second
	coingeckoCacheTTL     = 24 * time.Hour
)

// CoinGeckoPriceQuery represents query params for /simple/price endpoint
type CoinGeckoPriceQuery struct {
	IDs              string `url:"ids"`
	VsCurrencies     string `url:"vs_currencies"`
	Include24hChange string `url:"include_24hr_change"`
}

// CoinGeckoMarketsQuery represents query params for /coins/markets endpoint
type CoinGeckoMarketsQuery struct {
	VsCurrency string `url:"vs_currency"`
	Order      string `url:"order"`
	PerPage    string `url:"per_page"`
	Page       string `url:"page"`
}

// CoinGeckoMarketCoin represents a coin from the CoinGecko /coins/markets API
type CoinGeckoMarketCoin struct {
	ID     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

// CoinGeckoPriceResponse represents the price API response
type CoinGeckoPriceResponse map[string]struct {
	USD         float64 `json:"usd"`
	USDChange24 float64 `json:"usd_24h_change"`
}
