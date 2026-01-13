//go:build !bindings

package main

import (
	"embed"

	"github.com/wailsapp/wails/v3/pkg/application"

	"crypto-tray/services"
)

// NewApp creates the Wails v3 application
func NewApp(assets embed.FS, appService *services.AppService) *application.App {
	app := application.New(application.Options{
		Name:        "Crypto Tray",
		Description: "Cryptocurrency price tracker in your system tray",
		Services: []application.Service{
			application.NewService(appService),
		},
		Assets: application.AssetOptions{
			Handler: application.BundledAssetFileServer(assets),
		},
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
		},
	})

	// Disable application menu bar (tray app doesn't need one)
	app.Menu.SetApplicationMenu(app.NewMenu())

	return app
}
