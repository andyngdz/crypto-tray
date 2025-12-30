package tray

import (
	_ "embed"
	"fmt"

	"crypto-tray/providers"
	"crypto-tray/services"

	"github.com/getlantern/systray"
)

//go:embed icon.png
var iconData []byte

// New creates a new tray manager with the initial symbols
func New(symbols []string, onOpenSettings, onRefreshNow, onQuit func()) *Manager {
	if len(symbols) == 0 {
		symbols = []string{"---"}
	}
	return &Manager{
		onOpenSettings: onOpenSettings,
		onRefreshNow:   onRefreshNow,
		onQuit:         onQuit,
		priceItems:     make(map[string]*systray.MenuItem),
		symbols:        symbols,
	}
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
	systray.SetTitle(fmt.Sprintf("%s $--,---", t.symbols[0]))
	systray.SetTooltip("Crypto Tray - Loading...")

	for _, symbol := range t.symbols {
		item := systray.AddMenuItem(fmt.Sprintf("%s $--,---", symbol), "Current price")
		item.Disable()
		t.priceItems[symbol] = item
	}

	systray.AddSeparator()

	settingsItem := systray.AddMenuItem("Open Settings", "Configure the application")
	refreshItem := systray.AddMenuItem("Refresh Now", "Fetch latest price")

	systray.AddSeparator()

	quitItem := systray.AddMenuItem("Quit", "Exit the application")

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

// UpdatePrices updates the tray display with multiple price data
func (t *Manager) UpdatePrices(data []*providers.PriceData) {
	if len(data) == 0 {
		return
	}

	primarySymbol := t.symbols[0]
	var primaryData *providers.PriceData
	for _, item := range data {
		displayText := fmt.Sprintf("%s %s", item.Symbol, services.FormatPrice(item.Price))
		if menuItem, ok := t.priceItems[item.Symbol]; ok {
			menuItem.SetTitle(displayText)
		}

		if item.Symbol == primarySymbol {
			primaryData = item
		}
	}

	if primaryData != nil {
		displayText := fmt.Sprintf("%s %s", primaryData.Symbol, services.FormatPrice(primaryData.Price))
		systray.SetTitle(displayText)
		systray.SetTooltip(fmt.Sprintf("Crypto Tray - %s", displayText))
	}
}

// SetError updates tray to show error state
func (t *Manager) SetError(msg string) {
	primarySymbol := t.symbols[0]
	systray.SetTitle(fmt.Sprintf("%s $???", primarySymbol))
	systray.SetTooltip("Error: " + msg)
	for symbol, menuItem := range t.priceItems {
		menuItem.SetTitle(fmt.Sprintf("%s Error", symbol))
	}
}

// SetLoading shows loading state
func (t *Manager) SetLoading() {
	primarySymbol := t.symbols[0]
	systray.SetTitle(fmt.Sprintf("%s ...", primarySymbol))
	systray.SetTooltip("Crypto Tray - Loading...")
}

// SetSymbols updates the tracked currencies
func (t *Manager) SetSymbols(symbols []string) {
	if len(symbols) == 0 {
		return
	}

	t.symbols = symbols
	primarySymbol := symbols[0]
	systray.SetTitle(fmt.Sprintf("%s $--,---", primarySymbol))

	// Note: systray doesn't support removing menu items at runtime
	// The menu will be rebuilt on next app restart
	t.priceItems = make(map[string]*systray.MenuItem)
	for _, symbol := range symbols {
		item := systray.AddMenuItem(fmt.Sprintf("%s $--,---", symbol), "Current price")
		item.Disable()
		t.priceItems[symbol] = item
	}
}
