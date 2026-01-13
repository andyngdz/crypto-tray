package services

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"

	"crypto-tray/config"
	"crypto-tray/providers"
)

// ProviderRegistry provides access to price data providers
type ProviderRegistry interface {
	List() []providers.Provider
	GetSymbols(ctx context.Context, providerID string) ([]providers.SymbolInfo, error)
	GetDefaultCoinIDs(providerID string) []string
	FetchPrices(ctx context.Context, providerID string, coinIDs []string) ([]*providers.PriceData, error)
}

// AppService is the main service for CryptoTray application
type AppService struct {
	app           *application.App
	configManager *config.Manager
	registry      ProviderRegistry

	onSymbolsChanged         func(symbols []string)
	onNumberFormatChanged    func(format string)
	onDisplayCurrencyChanged func(currency string)
	onRefreshPrices          func()
}

// NewAppService creates a new AppService
func NewAppService(configManager *config.Manager, registry ProviderRegistry, app *application.App) *AppService {
	return &AppService{
		app:           app,
		configManager: configManager,
		registry:      registry,
	}
}

// ServiceName returns the name of the service (used by Wails v3)
func (s *AppService) ServiceName() string {
	return "AppService"
}

// ServiceStartup is called when the application starts
func (s *AppService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	s.app = application.Get()

	return nil
}

// ServiceShutdown is called when the application shuts down
func (s *AppService) ServiceShutdown() error {
	return nil
}

// SetOnSymbolsChanged sets the callback for symbol changes
func (s *AppService) SetOnSymbolsChanged(callback func(symbols []string)) {
	s.onSymbolsChanged = callback
}

// SetOnNumberFormatChanged sets the callback for number format changes
func (s *AppService) SetOnNumberFormatChanged(callback func(format string)) {
	s.onNumberFormatChanged = callback
}

// SetOnDisplayCurrencyChanged sets the callback for display currency changes
func (s *AppService) SetOnDisplayCurrencyChanged(callback func(currency string)) {
	s.onDisplayCurrencyChanged = callback
}

// SetOnRefreshPrices sets the callback for manual price refresh
func (s *AppService) SetOnRefreshPrices(callback func()) {
	s.onRefreshPrices = callback
}

// ShowWindow shows the settings window
func (s *AppService) ShowWindow() {
	windows := s.app.Window.GetAll()
	if len(windows) > 0 {
		windows[0].Show()
	}
}

// HideWindow hides the settings window
func (s *AppService) HideWindow() {
	windows := s.app.Window.GetAll()
	if len(windows) > 0 {
		windows[0].Hide()
	}
}

// QuitApp quits the application
func (s *AppService) QuitApp() {
	s.app.Quit()
}

// GetApp returns the Wails application instance
func (s *AppService) GetApp() *application.App {
	return s.app
}
