package providers

// PriceData represents cryptocurrency price information
type PriceData struct {
	CoinID         string  `json:"coinId"`
	Symbol         string  `json:"symbol"`
	Price          float64 `json:"price"`
	Change24h      float64 `json:"change_24h"`
	ConvertedPrice float64 `json:"convertedPrice"`
	Currency       string  `json:"currency"`
}

// SymbolInfo represents a supported cryptocurrency
type SymbolInfo struct {
	CoinID string `json:"coinId"` // Provider-specific ID for API calls (e.g., "bitcoin")
	Symbol string `json:"symbol"` // User-facing ticker in uppercase (e.g., "BTC")
	Name   string `json:"name"`   // Full display name (e.g., "Bitcoin")
}
