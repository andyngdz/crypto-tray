//go:build !bindings

package main

import (
	"context"

	"crypto-tray/config"
	"crypto-tray/exchange"
	"crypto-tray/movement"
	"crypto-tray/price"
	"crypto-tray/providers"
	"crypto-tray/tray"
)

// StartServices initializes and starts all background services
func StartServices(
	cfg config.Config,
	registry *providers.Registry,
	exchangeService *exchange.Service,
	priceService *price.Service,
	trayManager *tray.Manager,
	movementTracker *movement.Tracker,
	converter *exchange.Converter,
) {
	// Start exchange service first to get rates
	exchangeService.Start()

	// Preload symbols and fetch initial prices
	provider, ok := registry.Get(cfg.ProviderID)
	if !ok {
		priceService.Start()
		return
	}

	ctx := context.Background()
	provider.FetchSymbols(ctx)

	data, err := provider.FetchPrices(ctx, cfg.Symbols)
	if err == nil && len(data) > 0 {
		converter.ConvertPrices(data)
		movements := movementTracker.Track(data)
		trayManager.UpdatePrices(data, movements)
	}

	priceService.Start()
}
