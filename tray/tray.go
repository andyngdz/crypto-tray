package tray

import (
	"fmt"
	"strings"

	"crypto-tray/movement"
	"crypto-tray/providers"
	"crypto-tray/services"

	"github.com/getlantern/systray"
)

// New creates a new tray manager with the initial symbols and number format
func New(symbols []string, numberFormat string, onOpenSettings, onRefreshNow, onQuit func()) *Manager {
	if len(symbols) == 0 {
		symbols = []string{"---"}
	}
	return &Manager{
		onOpenSettings: onOpenSettings,
		onRefreshNow:   onRefreshNow,
		onQuit:         onQuit,
		priceSlots:     make([]*systray.MenuItem, 0, maxPriceSlots),
		symbols:        symbols,
		symbolMap:      make(map[string]string),
		numberFormat:   numberFormat,
	}
}

// SetNumberFormat updates the number format for price display
func (t *Manager) SetNumberFormat(format string) {
	t.numberFormat = format
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
	systray.SetTitle(services.FormatTrayTitle(t.getDisplaySymbols(), "$--,---"))
	systray.SetTooltip("Crypto Tray - Loading...")

	// Pre-allocate menu item slots
	for range maxPriceSlots {
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
	displaySymbols := t.getDisplaySymbols()
	for slotIdx := range t.priceSlots {
		slot := t.priceSlots[slotIdx]
		if slotIdx < len(t.symbols) {
			slot.SetTitle(fmt.Sprintf("%s $--,---", displaySymbols[slotIdx]))
			slot.Show()
		} else {
			slot.Hide()
		}
	}
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

// UpdatePrices updates the tray display with multiple price data and movement indicators
func (t *Manager) UpdatePrices(data []*providers.PriceData, movements map[string]movement.Direction) {
	if len(data) == 0 {
		return
	}

	// Build a map for quick lookup by coinID and update symbolMap
	priceMap := make(map[string]*providers.PriceData)
	for dataIdx := range data {
		d := data[dataIdx]
		priceMap[d.CoinID] = d
		t.symbolMap[d.CoinID] = d.Symbol
	}

	// Update each slot with its coinID's price and movement indicator
	for symbolIdx := range t.symbols {
		if symbolIdx >= len(t.priceSlots) {
			break
		}
		coinID := t.symbols[symbolIdx]
		if d, ok := priceMap[coinID]; ok {
			price := d.ConvertedPrice
			if price == 0 {
				price = d.Price
			}
			indicator := movement.IndicatorNeutral
			if dir, ok := movements[coinID]; ok {
				indicator = dir.Indicator()
			}
			displayText := fmt.Sprintf("%s %s %s", indicator, d.Symbol, services.FormatPriceWithCurrency(price, t.numberFormat, d.Currency))
			t.priceSlots[symbolIdx].SetTitle(displayText)
		}
	}

	// Update tray title with all currencies and movement indicators
	var titleParts []string
	for symbolIdx := range t.symbols {
		coinID := t.symbols[symbolIdx]
		if d, ok := priceMap[coinID]; ok {
			price := d.ConvertedPrice
			if price == 0 {
				price = d.Price
			}
			indicator := movement.IndicatorNeutral
			if dir, ok := movements[coinID]; ok {
				indicator = dir.Indicator()
			}
			titleParts = append(titleParts, fmt.Sprintf("%s %s %s", indicator, d.Symbol, services.FormatPriceWithCurrency(price, t.numberFormat, d.Currency)))
		}
	}
	if len(titleParts) > 0 {
		title := strings.Join(titleParts, " | ")
		systray.SetTitle(title)
		systray.SetTooltip("Crypto Tray - " + title)
	}
}

// SetError updates tray to show error state
func (t *Manager) SetError(msg string) {
	displaySymbols := t.getDisplaySymbols()
	if len(t.symbols) > 0 {
		systray.SetTitle(services.FormatTrayTitle(displaySymbols, "$???"))
	}
	systray.SetTooltip("Error: " + msg)
	for symbolIdx := range t.symbols {
		if symbolIdx >= len(t.priceSlots) {
			break
		}
		t.priceSlots[symbolIdx].SetTitle(fmt.Sprintf("%s Error", displaySymbols[symbolIdx]))
	}
}

// SetLoading shows loading state
func (t *Manager) SetLoading() {
	if len(t.symbols) > 0 {
		systray.SetTitle(services.FormatTrayTitle(t.getDisplaySymbols(), "..."))
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
	systray.SetTitle(services.FormatTrayTitle(t.getDisplaySymbols(), "$--,---"))
	t.updateSlots()
}
