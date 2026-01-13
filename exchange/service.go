package exchange

import (
	"crypto-tray/config"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Service orchestrates exchange rate fetching, conversion, and event emission
type Service struct {
	fetcher   *Fetcher
	converter *Converter
	app       *application.App
}

// NewService creates a new exchange service
func NewService(configManager *config.Manager, app *application.App) *Service {
	fetcher := newFetcher(configManager)
	converter := NewConverter(fetcher, configManager)

	return &Service{
		fetcher:   fetcher,
		converter: converter,
		app:       app,
	}
}

// Start begins the exchange rate fetching loop
func (s *Service) Start() {
	s.fetcher.Start(func(rates *ExchangeRates, err error) {
		if err != nil {
			return
		}

		s.app.Event.Emit("exchange:update", rates.Rates)
	})
}

// Stop stops the exchange rate fetching loop
func (s *Service) Stop() {
	s.fetcher.Stop()
}

// GetConverter returns the currency converter
func (s *Service) GetConverter() *Converter {
	return s.converter
}
