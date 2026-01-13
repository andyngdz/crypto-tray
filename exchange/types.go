package exchange

import "time"

// API constants for exchange rate fetching
const (
	primaryURL   = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies"
	fallbackURL  = "https://latest.currency-api.pages.dev/v1/currencies"
	baseCurrency = "usdt"
	timeout      = 10 * time.Second
)

// ExchangeRates holds rates from base currency to other currencies
type ExchangeRates struct {
	Date  string             `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// RatesHandler is called when exchange rates are fetched
type RatesHandler func(rates *ExchangeRates, err error)

// APIResponse represents the raw API response
type APIResponse struct {
	Date string             `json:"date"`
	USDT map[string]float64 `json:"usdt"`
}
