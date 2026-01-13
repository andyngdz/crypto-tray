//go:build !bindings

package main

import (
	"crypto-tray/price"
	"crypto-tray/services"
	"crypto-tray/tray"
)

// SetupCallbacks wires configuration change callbacks between services
func SetupCallbacks(
	appService *services.AppService,
	trayManager *tray.Manager,
	priceService *price.Service,
) {
	appService.SetOnSymbolsChanged(func(symbols []string) {
		trayManager.SetSymbols(symbols)
		priceService.RefreshNow()
	})

	appService.SetOnNumberFormatChanged(func(format string) {
		trayManager.SetNumberFormat(format)
		priceService.RefreshNow()
	})

	appService.SetOnDisplayCurrencyChanged(func(currency string) {
		priceService.RefreshNow()
	})

	appService.SetOnRefreshPrices(func() {
		priceService.RefreshNow()
	})
}
