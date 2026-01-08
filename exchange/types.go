package exchange

import "context"

// ExchangeRates holds rates from base currency to other currencies
type ExchangeRates struct {
	Date  string             `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

// Callback is called when exchange rates are fetched
type Callback func(rates *ExchangeRates, err error)

// ContextProvider provides Wails context for event emission
type ContextProvider interface {
	GetContext() context.Context
}
