package providers

import "context"

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

// Registry holds all available providers
type Registry struct {
	providers map[string]Provider
}

// NewRegistry creates a new provider registry
func NewRegistry() *Registry {
	return &Registry{
		providers: make(map[string]Provider),
	}
}

// Register adds a provider to the registry
func (r *Registry) Register(p Provider) {
	r.providers[p.ID()] = p
}

// Get retrieves a provider by ID
func (r *Registry) Get(id string) (Provider, bool) {
	p, ok := r.providers[id]
	return p, ok
}

// List returns all registered providers
func (r *Registry) List() []Provider {
	result := make([]Provider, 0, len(r.providers))
	for _, p := range r.providers {
		result = append(result, p)
	}
	return result
}
