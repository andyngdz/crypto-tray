package providers

// PriceData represents cryptocurrency price information
type PriceData struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	Change24h float64 `json:"change_24h"`
}

// SymbolInfo represents a supported cryptocurrency
type SymbolInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
