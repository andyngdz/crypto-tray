//go:build !bindings

package main

import (
	"embed"
	"log"
	"os"

	"github.com/wailsapp/wails/v3/pkg/application"

	"crypto-tray/exchange"
	"crypto-tray/movement"
	"crypto-tray/price"
	"crypto-tray/services"
	"crypto-tray/tray"
	"crypto-tray/window"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed frontend/src/assets/images/logo.png
var windowIcon []byte

func main() {
	deps, err := InitApp()
	if err != nil {
		log.Fatal("Failed to initialize:", err)
	}

	cfg := deps.ConfigManager.Get()

	// Create services and UI components
	appService := services.NewAppService(deps.ConfigManager, deps.Registry, application.Get())
	app := NewApp(assets, appService)
	settingsWindow := window.NewSettings(app, windowIcon)

	// Create tray manager with callbacks
	var priceService *price.Service
	trayManager := tray.New(
		cfg.Symbols,
		cfg.NumberFormat,
		func() { settingsWindow.Show() },
		func() { priceService.RefreshNow() },
		func() {
			app.Quit()
			os.Exit(0)
		},
	)
	trayManager.Setup(app, settingsWindow)

	// Create remaining services
	exchangeService := exchange.NewService(deps.ConfigManager, app)
	converter := exchangeService.GetConverter()
	movementTracker := movement.NewTracker()
	priceService = price.NewService(
		deps.Registry,
		deps.ConfigManager,
		trayManager,
		converter,
		app,
		movementTracker,
	)

	// Wire up callbacks and shutdown handler
	SetupCallbacks(appService, trayManager, priceService)
	app.OnShutdown(func() {
		priceService.Stop()
		exchangeService.Stop()
	})

	// Start services in background
	go StartServices(cfg, deps.Registry, exchangeService, priceService, trayManager, movementTracker, converter)

	if err := app.Run(); err != nil {
		log.Fatal("Application error:", err)
	}
}
