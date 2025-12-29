package tray

import (
	_ "embed"
	"fmt"

	"crypto-tray/services"

	"github.com/getlantern/systray"
)

//go:embed icon.png
var iconData []byte

// Manager handles system tray operations
type Manager struct {
	onOpenSettings func()
	onRefreshNow   func()
	onQuit         func()
	priceItem      *systray.MenuItem
	currentSymbol  string
}

// New creates a new tray manager with the initial symbol
func New(initialSymbol string, onOpenSettings, onRefreshNow, onQuit func()) *Manager {
	symbol := initialSymbol
	if symbol == "" {
		symbol = "---"
	}
	return &Manager{
		onOpenSettings: onOpenSettings,
		onRefreshNow:   onRefreshNow,
		onQuit:         onQuit,
		currentSymbol:  symbol,
	}
}

// SetSymbol updates the current symbol used in display
func (t *Manager) SetSymbol(symbol string) {
	t.currentSymbol = symbol
}

// Setup registers the tray without blocking - use this with Wails
func (t *Manager) Setup() {
	systray.Register(t.onReady, t.onExit)
}

// Run starts the system tray - this blocks! Use Setup() instead with Wails
func (t *Manager) Run() {
	systray.Run(t.onReady, t.onExit)
}

func (t *Manager) onReady() {
	systray.SetIcon(iconData)
	systray.SetTitle(fmt.Sprintf("%s $--,---", t.currentSymbol))
	systray.SetTooltip("Crypto Tray - Loading...")

	// Create menu items
	t.priceItem = systray.AddMenuItem(fmt.Sprintf("%s $--,---", t.currentSymbol), "Current price")
	t.priceItem.Disable()

	systray.AddSeparator()

	settingsItem := systray.AddMenuItem("Open Settings", "Configure the application")
	refreshItem := systray.AddMenuItem("Refresh Now", "Fetch latest price")

	systray.AddSeparator()

	quitItem := systray.AddMenuItem("Quit", "Exit the application")

	// Handle menu clicks
	go func() {
		for {
			select {
			case <-settingsItem.ClickedCh:
				t.onOpenSettings()
			case <-refreshItem.ClickedCh:
				t.onRefreshNow()
			case <-quitItem.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func (t *Manager) onExit() {
	t.onQuit()
}

// UpdatePrice updates the tray display with new price data
func (t *Manager) UpdatePrice(symbol string, price float64) {
	displayText := fmt.Sprintf("%s %s", symbol, services.FormatPrice(price))
	systray.SetTitle(displayText)
	systray.SetTooltip(fmt.Sprintf("Crypto Tray - %s", displayText))
	t.priceItem.SetTitle(displayText)
}

// SetError updates tray to show error state
func (t *Manager) SetError(msg string) {
	systray.SetTitle(fmt.Sprintf("%s $???", t.currentSymbol))
	systray.SetTooltip("Error: " + msg)
	t.priceItem.SetTitle("Error fetching price")
}

// SetLoading shows loading state
func (t *Manager) SetLoading() {
	systray.SetTitle(fmt.Sprintf("%s ...", t.currentSymbol))
	systray.SetTooltip("Crypto Tray - Loading...")
}
