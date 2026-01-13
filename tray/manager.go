package tray

import (
	"runtime"
	"strings"

	"github.com/wailsapp/wails/v3/pkg/application"

	"crypto-tray/services"
)

// New creates a new tray manager for Wails v3
func New(symbols []string, numberFormat string, onOpenSettings, onRefreshNow, onQuit func()) *Manager {
	if len(symbols) == 0 {
		symbols = []string{"---"}
	}

	return &Manager{
		onOpenSettings: onOpenSettings,
		onRefreshNow:   onRefreshNow,
		onQuit:         onQuit,
		priceItems:     make([]*application.MenuItem, 0, maxPriceSlots),
		symbols:        symbols,
		symbolMap:      make(map[string]string),
		numberFormat:   numberFormat,
	}
}

// Setup initializes the systray with the Wails v3 application
func (t *Manager) Setup(app *application.App, window *application.WebviewWindow) {
	t.app = app
	t.window = window

	// Create system tray via the manager
	t.systray = app.SystemTray.New()

	// Set icon based on platform
	if runtime.GOOS == "darwin" {
		t.systray.SetTemplateIcon(iconData)
	} else {
		t.systray.SetIcon(iconData)
	}

	// Set empty click handler to prevent "openMenu not implemented on Linux" errors
	// Users access settings via right-click context menu
	t.systray.OnClick(func() {})

	// Set initial label
	initialLabel := services.FormatTrayTitle(t.getDisplaySymbols(), "$--,---")
	t.systray.SetLabel(initialLabel)

	// Create menu
	t.createMenu()
}

// SetNumberFormat updates the number format for price display
func (t *Manager) SetNumberFormat(format string) {
	t.numberFormat = format
}

// SetSymbols updates the tracked currencies
func (t *Manager) SetSymbols(symbols []string) {
	if len(symbols) == 0 {
		return
	}

	// Limit to max slots
	if len(symbols) > maxPriceSlots {
		symbols = symbols[:maxPriceSlots]
	}

	t.symbols = symbols
	label := services.FormatTrayTitle(t.getDisplaySymbols(), "$--,---")
	t.systray.SetLabel(label)
	t.updateSlots()
}

// getDisplaySymbols converts coinIDs to ticker symbols for display
func (t *Manager) getDisplaySymbols() []string {
	result := make([]string, len(t.symbols))

	for symbolIdx := range t.symbols {
		coinID := t.symbols[symbolIdx]
		if symbol, ok := t.symbolMap[coinID]; ok {
			result[symbolIdx] = symbol
		} else {
			result[symbolIdx] = strings.ToUpper(coinID)
		}
	}

	return result
}
