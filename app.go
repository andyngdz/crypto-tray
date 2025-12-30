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
	ctx           context.Context
	ctxMu         sync.RWMutex
	configManager *config.Manager
	registry      *providers.Registry
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
	ctx := a.GetContext()
	if ctx != nil {
		runtime.WindowShow(ctx)
	}
}

// QuitApp quits the application (for tray callback)
func (a *App) QuitApp() {
	ctx := a.GetContext()
	if ctx != nil {
		runtime.Quit(ctx)
	}
}

// GetConfig returns the current configuration
func (a *App) GetConfig() config.Config {
	return a.configManager.Get()
}

// SaveConfig saves updated configuration
func (a *App) SaveConfig(cfg config.Config) error {
	return a.configManager.Update(cfg)
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
	return provider.GetSupportedSymbols()
}

// HideWindow hides the settings window
func (a *App) HideWindow() {
	runtime.WindowHide(a.ctx)
}
