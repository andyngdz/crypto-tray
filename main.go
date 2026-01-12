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
	"github.com/wailsapp/wails/v2/pkg/options/linux"

	"crypto-tray/exchange"
	"crypto-tray/movement"
	"crypto-tray/price"
	"crypto-tray/tray"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed frontend/src/assets/images/logo.png
var windowIcon []byte

func main() {
	// Initialize all dependencies
	deps, err := InitApp()
	if err != nil {
		log.Fatal("Failed to initialize:", err)
	}

	// Get initial config for tray display
	cfg := deps.ConfigManager.Get()

	// Create tray manager
	var priceService *price.Service
	trayManager := tray.New(
		cfg.Symbols,
		cfg.NumberFormat,
		deps.App.ShowWindow,
		func() { // onRefreshNow
			priceService.RefreshNow()
		},
		func() { // onQuit
			deps.App.QuitApp()
			os.Exit(0)
		},
	)

	// Create exchange service
	exchangeService := exchange.NewService(deps.ConfigManager, deps.App)
	converter := exchangeService.GetConverter()

	// Create movement tracker
	movementTracker := movement.NewTracker()

	// Create price service
	priceService = price.NewService(
		deps.Registry,
		deps.ConfigManager,
		trayManager,
		exchangeService.GetConverter(),
		deps.App,
		movementTracker,
	)

	// Connect symbol changes to tray
	deps.App.setOnSymbolsChanged(func(symbols []string) {
		trayManager.SetSymbols(symbols)
		priceService.RefreshNow()
	})

	// Connect number format changes to tray
	deps.App.setOnNumberFormatChanged(func(format string) {
		trayManager.SetNumberFormat(format)
		priceService.RefreshNow() // Refresh to update display with new format
	})

	// Connect display currency changes
	deps.App.setOnDisplayCurrencyChanged(func(currency string) {
		priceService.RefreshNow() // Refresh to update display with new currency
	})

	// Connect manual refresh from frontend
	deps.App.setOnRefreshPrices(func() {
		priceService.RefreshNow()
	})

	// Setup systray before Wails starts (shares GTK context)
	trayManager.Setup()

	// Run Wails in main goroutine (GTK requires main thread)
	err = wails.Run(&options.App{
		Title:     "Crypto Tray",
		Width:     640,
		Height:    800,
		MinWidth:  640,
		MinHeight: 600,
		Linux: &linux.Options{
			Icon:        windowIcon,
			ProgramName: "Crypto Tray",
		},
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Frameless:        true,
		CSSDragProperty:  "--wails-draggable",
		CSSDragValue:     "drag",
		StartHidden:      true,
		OnStartup: func(ctx context.Context) {
			deps.App.startup(ctx)

			// Start exchange service first to get rates
			exchangeService.Start()

			if provider, ok := deps.Registry.Get(cfg.ProviderID); ok {
				// Preload symbol cache
				provider.FetchSymbols(ctx)

				// Fetch initial prices synchronously
				if data, err := provider.FetchPrices(ctx, cfg.Symbols); err == nil && len(data) > 0 {
					converter.ConvertPrices(data)
					movements := movementTracker.Track(data)
					trayManager.UpdatePrices(data, movements)
				}
			}

			priceService.Start()
		},
		OnShutdown: func(ctx context.Context) {
			priceService.Stop()
			exchangeService.Stop()
		},
		Bind: []interface{}{
			deps.App,
		},
	})

	if err != nil {
		log.Fatal("Wails error:", err)
	}
}
