package price

import (
	"log"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"crypto-tray/config"
	"crypto-tray/providers"
)

// Service orchestrates price fetching, conversion, tray updates, and event emission
type Service struct {
	fetcher         *Fetcher
	tray            TrayUpdater
	converter       PriceConverter
	contextProvider ContextProvider
	movementTracker MovementTracker
}

// NewService creates a new price service
func NewService(
	registry *providers.Registry,
	configManager *config.Manager,
	tray TrayUpdater,
	converter PriceConverter,
	contextProvider ContextProvider,
	movementTracker MovementTracker,
) *Service {
	fetcher := newFetcher(registry, configManager)

	return &Service{
		fetcher:         fetcher,
		tray:            tray,
		converter:       converter,
		contextProvider: contextProvider,
		movementTracker: movementTracker,
	}
}

// Start begins the price fetching loop
func (s *Service) Start() {
	s.fetcher.Start(func(data []*providers.PriceData, err error) {
		if err != nil {
			log.Printf("Error fetching price: %v", err)
			s.tray.SetError(err.Error())
			return
		}

		if len(data) > 0 {
			s.converter.ConvertPrices(data)
			movements := s.movementTracker.Track(data)
			s.tray.UpdatePrices(data, movements)
			runtime.EventsEmit(s.contextProvider.GetContext(), "price:update", data)
		}
	})
}

// Stop stops the price fetching loop
func (s *Service) Stop() {
	s.fetcher.Stop()
}

// RefreshNow triggers an immediate price fetch
func (s *Service) RefreshNow() {
	s.fetcher.RefreshNow()
}
