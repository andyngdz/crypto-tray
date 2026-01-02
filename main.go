//go:build !bindings

package main

import (
	"context"
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"crypto-tray/price"
	"crypto-tray/providers"
	"crypto-tray/tray"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Initialize all dependencies
	deps, err := InitApp()
	if err != nil {
		log.Fatal("Failed to initialize:", err)
	}

	// Get initial config for tray display
	cfg := deps.ConfigManager.Get()

	// Create tray manager
	var fetcher *price.Fetcher
	trayManager := tray.New(
		cfg.Symbols,
		deps.App.ShowWindow,
		func() { // onRefreshNow
			if fetcher != nil {
				fetcher.RefreshNow()
			}
		},
		func() { // onQuit
			deps.App.QuitApp()
			os.Exit(0)
		},
	)

	// Create price fetcher
	fetcher = price.NewFetcher(deps.Registry, deps.ConfigManager, func(data []*providers.PriceData, err error) {
		if err != nil {
			log.Printf("Error fetching price: %v", err)
			trayManager.SetError(err.Error())
			return
		}
		if len(data) > 0 {
			trayManager.UpdatePrices(data)
			runtime.EventsEmit(deps.App.GetContext(), "price:update", data)
		}
	})

	// Connect symbol changes to tray
	deps.App.setOnSymbolsChanged(func(symbols []string) {
		trayManager.SetSymbols(symbols)
		fetcher.RefreshNow()
	})

	// Connect manual refresh from frontend
	deps.App.setOnRefreshPrices(func() {
		fetcher.RefreshNow()
	})

	// Setup systray before Wails starts (shares GTK context)
	trayManager.Setup()

	// Run Wails in main goroutine (GTK requires main thread)
	err = wails.Run(&options.App{
		Title:     "Crypto Tray Settings",
		Width:     500,
		Height:    400,
		MinWidth:  400,
		MinHeight: 300,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		StartHidden:      true,
		OnStartup: func(ctx context.Context) {
			deps.App.startup(ctx)
			// Preload symbol cache for price lookups
			if provider, ok := deps.Registry.Get(cfg.ProviderID); ok {
				provider.FetchSymbols(ctx)
			}
			fetcher.Start()
		},
		OnShutdown: func(ctx context.Context) {
			fetcher.Stop()
		},
		Bind: []interface{}{
			deps.App,
		},
	})
	if err != nil {
		log.Fatal("Wails error:", err)
	}
}
