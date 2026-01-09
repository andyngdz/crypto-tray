package main

import (
	"context"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"crypto-tray/config"
	"crypto-tray/providers"
)

// ProviderInfo represents provider metadata for the frontend
type ProviderInfo struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	RequiresAPIKey bool   `json:"requiresApiKey"`
}

// App struct
type App struct {
	ctx                      context.Context
	ctxMu                    sync.RWMutex
	configManager            *config.Manager
	registry                 *providers.Registry
	onSymbolsChanged         func(symbols []string)
	onNumberFormatChanged    func(format string)
	onDisplayCurrencyChanged func(currency string)
	onRefreshPrices          func()
}

// NewApp creates a new App application struct
func NewApp(configManager *config.Manager, registry *providers.Registry) *App {
	return &App{
		configManager: configManager,
		registry:      registry,
	}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctxMu.Lock()
	a.ctx = ctx
	a.ctxMu.Unlock()
}

// GetContext returns the current Wails context (thread-safe)
func (a *App) GetContext() context.Context {
	a.ctxMu.RLock()
	defer a.ctxMu.RUnlock()
	return a.ctx
}

// ShowWindow shows the settings window (for tray callback)
func (a *App) ShowWindow() {
	runtime.WindowShow(a.GetContext())
}

// QuitApp quits the application (for tray callback)
func (a *App) QuitApp() {
	runtime.Quit(a.GetContext())
}

// GetConfig returns the current configuration
func (a *App) GetConfig() config.Config {
	return a.configManager.Get()
}

// SaveConfig saves updated configuration
func (a *App) SaveConfig(cfg config.Config) error {
	oldCfg := a.configManager.Get()

	// Reset symbols to provider's default when provider changes
	if oldCfg.ProviderID != cfg.ProviderID {
		newProvider, ok := a.registry.Get(cfg.ProviderID)
		if ok {
			cfg.Symbols = []string{newProvider.DefaultCoinID()}
		}
	}

	if err := a.configManager.Update(cfg); err != nil {
		return err
	}
	// Notify if symbols changed
	if !equalSymbols(oldCfg.Symbols, cfg.Symbols) {
		a.onSymbolsChanged(cfg.Symbols)
	}
	// Notify if number format changed
	if oldCfg.NumberFormat != cfg.NumberFormat {
		a.onNumberFormatChanged(cfg.NumberFormat)
	}
	// Notify if display currency changed
	if oldCfg.DisplayCurrency != cfg.DisplayCurrency {
		a.onDisplayCurrencyChanged(cfg.DisplayCurrency)
	}
	return nil
}

// setOnSymbolsChanged sets callback for when symbols change (internal use only)
func (a *App) setOnSymbolsChanged(callback func(symbols []string)) {
	a.onSymbolsChanged = callback
}

// setOnNumberFormatChanged sets callback for when number format changes (internal use only)
func (a *App) setOnNumberFormatChanged(callback func(format string)) {
	a.onNumberFormatChanged = callback
}

// setOnDisplayCurrencyChanged sets callback for when display currency changes (internal use only)
func (a *App) setOnDisplayCurrencyChanged(callback func(currency string)) {
	a.onDisplayCurrencyChanged = callback
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

// GetAvailableProviders returns list of available API providers
func (a *App) GetAvailableProviders() []ProviderInfo {
	providerList := a.registry.List()
	result := make([]ProviderInfo, len(providerList))
	for i, p := range providerList {
		result[i] = ProviderInfo{
			ID:             p.ID(),
			Name:           p.Name(),
			RequiresAPIKey: p.RequiresAPIKey(),
		}
	}
	return result
}

// GetAvailableSymbols returns list of supported cryptocurrency symbols
func (a *App) GetAvailableSymbols() []providers.SymbolInfo {
	cfg := a.configManager.Get()
	provider, ok := a.registry.Get(cfg.ProviderID)
	if !ok {
		return []providers.SymbolInfo{}
	}

	symbols, _ := provider.FetchSymbols(a.GetContext())
	return symbols
}

// HideWindow hides the settings window
func (a *App) HideWindow() {
	runtime.WindowHide(a.ctx)
}

// FetchPrices fetches current prices for the specified symbols
func (a *App) FetchPrices(symbols []string) ([]*providers.PriceData, error) {
	cfg := a.configManager.Get()
	provider, ok := a.registry.Get(cfg.ProviderID)
	if !ok {
		return nil, nil
	}
	return provider.FetchPrices(a.GetContext(), symbols)
}

// RefreshPrices triggers an immediate price refresh
func (a *App) RefreshPrices() {
	a.onRefreshPrices()
}

// setOnRefreshPrices sets callback for manual price refresh (internal use only)
func (a *App) setOnRefreshPrices(callback func()) {
	a.onRefreshPrices = callback
}
