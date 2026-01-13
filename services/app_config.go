package services

import (
	"context"

	"crypto-tray/config"
	"crypto-tray/providers"
)

// GetConfig returns the current configuration
func (s *AppService) GetConfig() config.Config {
	return s.configManager.Get()
}

// SaveConfig saves an updated configuration
func (s *AppService) SaveConfig(cfg config.Config) error {
	oldCfg := s.configManager.Get()

	// Reset symbols to the provider's default when provider changes
	if oldCfg.ProviderID != cfg.ProviderID {
		cfg.Symbols = s.registry.GetDefaultCoinIDs(cfg.ProviderID)
	}

	if err := s.configManager.Update(cfg); err != nil {
		return err
	}

	// Notify if symbols changed
	if !equalSymbols(oldCfg.Symbols, cfg.Symbols) && s.onSymbolsChanged != nil {
		s.onSymbolsChanged(cfg.Symbols)
	}

	// Notify if number format changed
	if oldCfg.NumberFormat != cfg.NumberFormat && s.onNumberFormatChanged != nil {
		s.onNumberFormatChanged(cfg.NumberFormat)
	}

	// Notify if display currency changed
	if oldCfg.DisplayCurrency != cfg.DisplayCurrency && s.onDisplayCurrencyChanged != nil {
		s.onDisplayCurrencyChanged(cfg.DisplayCurrency)
	}

	return nil
}

// GetAvailableProviders returns a list of available API providers
func (s *AppService) GetAvailableProviders() []providers.ProviderInfo {
	providerList := s.registry.List()
	result := make([]providers.ProviderInfo, len(providerList))

	for providerIdx := range providerList {
		p := providerList[providerIdx]
		result[providerIdx] = providers.ProviderInfo{
			ID:             p.ID(),
			Name:           p.Name(),
			RequiresAPIKey: p.RequiresAPIKey(),
		}
	}

	return result
}

// GetAvailableSymbols returns a list of supported cryptocurrency symbols
func (s *AppService) GetAvailableSymbols() []providers.SymbolInfo {
	cfg := s.configManager.Get()
	symbols, _ := s.registry.GetSymbols(context.Background(), cfg.ProviderID)

	return symbols
}

// FetchPrices fetches the current prices for the specified symbols
func (s *AppService) FetchPrices(symbols []string) ([]*providers.PriceData, error) {
	cfg := s.configManager.Get()

	return s.registry.FetchPrices(context.Background(), cfg.ProviderID, symbols)
}

// RefreshPrices triggers an immediate price refresh
func (s *AppService) RefreshPrices() {
	if s.onRefreshPrices != nil {
		s.onRefreshPrices()
	}
}

// equalSymbols compares two symbol slices for equality
func equalSymbols(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for symbolIdx := range a {
		if a[symbolIdx] != b[symbolIdx] {
			return false
		}
	}

	return true
}
