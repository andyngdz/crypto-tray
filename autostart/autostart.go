package autostart

import (
	"os"

	"github.com/emersion/go-autostart"
)

const (
	appName     = "CryptoTray"
	displayName = "Crypto Tray"
)

// getApp returns an autostart.App configured for this application
func getApp() *autostart.App {
	exe, _ := os.Executable()
	return &autostart.App{
		Name:        appName,
		DisplayName: displayName,
		Exec:        []string{exe},
	}
}

// IsEnabled checks if auto-start is currently enabled at OS level
func IsEnabled() bool {
	return getApp().IsEnabled()
}

// SetEnabled enables or disables auto-start at OS level
func SetEnabled(enabled bool) error {
	app := getApp()
	if enabled {
		return app.Enable()
	}
	return app.Disable()
}
