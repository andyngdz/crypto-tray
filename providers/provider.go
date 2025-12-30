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

	// GetSupportedSymbols returns list of supported cryptocurrencies
	GetSupportedSymbols() []SymbolInfo

	// RequiresAPIKey returns true if the provider needs an API key
	RequiresAPIKey() bool

	// SetAPIKey configures the provider's API key
	SetAPIKey(key string)
}
