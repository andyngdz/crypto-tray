//go:build bindings

package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	deps, err := InitApp()
	if err != nil {
		log.Fatal("Failed to initialize:", err)
	}

	err = wails.Run(&options.App{
		Title:  "Crypto Tray Settings",
		Width:  500,
		Height: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: deps.App.startup,
		Bind: []interface{}{
			deps.App,
		},
	})
	if err != nil {
		log.Fatal("Wails error:", err)
	}
}
