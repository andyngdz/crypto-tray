package price

import (
	"crypto-tray/config"
	"crypto-tray/providers"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Service orchestrates price fetching, conversion, tray updates, and event emission
type Service struct {
	fetcher         *Fetcher
	tray            TrayUpdater
	converter       PriceConverter
	app             *application.App
	movementTracker MovementTracker
}

// NewService creates a new price service
func NewService(
	registry *providers.Registry,
	configManager *config.Manager,
	tray TrayUpdater,
	converter PriceConverter,
	app *application.App,
	movementTracker MovementTracker,
) *Service {
	fetcher := newFetcher(registry, configManager)

	return &Service{
		fetcher:         fetcher,
		tray:            tray,
		converter:       converter,
		app:             app,
		movementTracker: movementTracker,
	}
}

// Start begins the price fetching loop
func (s *Service) Start() {
	s.fetcher.Start(func(data []*providers.PriceData, err error) {
		if err != nil {
			s.tray.SetError(err.Error())
			return
		}

		if len(data) > 0 {
			s.converter.ConvertPrices(data)
			movements := s.movementTracker.Track(data)
			s.tray.UpdatePrices(data, movements)
			s.app.Event.Emit("price:update", data)
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
