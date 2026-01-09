package exchange

import (
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"crypto-tray/config"
)

// Service orchestrates exchange rate fetching, conversion, and event emission
type Service struct {
	fetcher         *Fetcher
	converter       *Converter
	contextProvider ContextProvider
}

// NewService creates a new exchange service
func NewService(configManager *config.Manager, contextProvider ContextProvider) *Service {
	fetcher := newFetcher(configManager)
	converter := NewConverter(fetcher, configManager)

	return &Service{
		fetcher:         fetcher,
		converter:       converter,
		contextProvider: contextProvider,
	}
}

// Start begins the exchange rate fetching loop
func (s *Service) Start() {
	s.fetcher.Start(func(rates *ExchangeRates, err error) {
		if err != nil {
			log.Printf("Error fetching exchange rates: %v", err)
			return
		}

		runtime.EventsEmit(s.contextProvider.GetContext(), "exchange:update", rates)
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
