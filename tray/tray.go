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
		priceSlots:     make([]*systray.MenuItem, 0, maxPriceSlots),
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

	// Pre-allocate menu item slots
	for i := 0; i < maxPriceSlots; i++ {
		item := systray.AddMenuItem("", "Current price")
		item.Disable()
		item.Hide()
		t.priceSlots = append(t.priceSlots, item)
	}

	// Initialize with current symbols
	t.updateSlots()

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

// updateSlots syncs the pre-allocated slots with current symbols
func (t *Manager) updateSlots() {
	for i, slot := range t.priceSlots {
		if i < len(t.symbols) {
			symbol := t.symbols[i]
			slot.SetTitle(fmt.Sprintf("%s $--,---", symbol))
			slot.Show()
		} else {
			slot.Hide()
		}
	}
}

// UpdatePrices updates the tray display with multiple price data
func (t *Manager) UpdatePrices(data []*providers.PriceData) {
	if len(data) == 0 {
		return
	}

	// Build a map for quick lookup by coinID
	priceMap := make(map[string]*providers.PriceData)
	for _, d := range data {
		priceMap[d.CoinID] = d
	}

	// Update each slot with its coinID's price
	for i, coinID := range t.symbols {
		if i >= len(t.priceSlots) {
			break
		}
		if d, ok := priceMap[coinID]; ok {
			displayText := fmt.Sprintf("%s %s", d.Symbol, services.FormatPrice(d.Price))
			t.priceSlots[i].SetTitle(displayText)
		}
	}

	// Update tray title with primary symbol
	if len(t.symbols) > 0 {
		if d, ok := priceMap[t.symbols[0]]; ok {
			displayText := fmt.Sprintf("%s %s", d.Symbol, services.FormatPrice(d.Price))
			systray.SetTitle(displayText)
			systray.SetTooltip(fmt.Sprintf("Crypto Tray - %s", displayText))
		}
	}
}

// SetError updates tray to show error state
func (t *Manager) SetError(msg string) {
	if len(t.symbols) > 0 {
		systray.SetTitle(fmt.Sprintf("%s $???", t.symbols[0]))
	}
	systray.SetTooltip("Error: " + msg)
	for i, symbol := range t.symbols {
		if i >= len(t.priceSlots) {
			break
		}
		t.priceSlots[i].SetTitle(fmt.Sprintf("%s Error", symbol))
	}
}

// SetLoading shows loading state
func (t *Manager) SetLoading() {
	if len(t.symbols) > 0 {
		systray.SetTitle(fmt.Sprintf("%s ...", t.symbols[0]))
	}
	systray.SetTooltip("Crypto Tray - Loading...")
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
	systray.SetTitle(fmt.Sprintf("%s $--,---", symbols[0]))
	t.updateSlots()
}
