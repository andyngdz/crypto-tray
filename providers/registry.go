package providers

import (
	"context"
	"maps"
	"slices"
)

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
	return slices.Collect(maps.Values(r.providers))
}

// GetSymbols fetches symbols list from provider by ID
func (r *Registry) GetSymbols(ctx context.Context, providerID string) ([]SymbolInfo, error) {
	provider, ok := r.Get(providerID)
	if !ok {
		return []SymbolInfo{}, nil
	}

	return provider.FetchSymbols(ctx)
}

// GetDefaultCoinIDs returns default coin IDs from provider by ID
func (r *Registry) GetDefaultCoinIDs(providerID string) []string {
	provider, ok := r.Get(providerID)
	if !ok {
		return []string{}
	}

	return provider.DefaultCoinIDs()
}

// FetchPrices fetches prices from provider by ID
func (r *Registry) FetchPrices(ctx context.Context, providerID string, coinIDs []string) ([]*PriceData, error) {
	provider, ok := r.Get(providerID)
	if !ok {
		return nil, nil
	}

	return provider.FetchPrices(ctx, coinIDs)
}
