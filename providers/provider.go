package providers

import "context"

// Provider defines the interface for price data providers
type Provider interface {
	// ID returns the provider's unique identifier
	ID() string

	// Name returns the provider's display name
	Name() string

	// FetchPrices retrieves prices for multiple cryptocurrencies
	FetchPrices(ctx context.Context, symbols []string) ([]*PriceData, error)

	// FetchSymbols fetches list of supported cryptocurrencies from the API
	FetchSymbols(ctx context.Context) ([]SymbolInfo, error)

	// RequiresAPIKey returns true if the provider needs an API key
	RequiresAPIKey() bool

	// SetAPIKey configures the provider's API key
	SetAPIKey(key string)
}
