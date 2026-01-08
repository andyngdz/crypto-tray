package tray

import (
	_ "embed"
	"fmt"
	"strings"

	"crypto-tray/providers"
	"crypto-tray/services"

	"github.com/getlantern/systray"
)

//go:embed icon.png
var iconData []byte

// New creates a new tray manager with the initial symbols, number format, and display currency
func New(symbols []string, numberFormat string, displayCurrency string, onOpenSettings, onRefreshNow, onQuit func()) *Manager {
	if len(symbols) == 0 {
		symbols = []string{"---"}
	}
	return &Manager{
		onOpenSettings:  onOpenSettings,
		onRefreshNow:    onRefreshNow,
		onQuit:          onQuit,
		priceSlots:      make([]*systray.MenuItem, 0, maxPriceSlots),
		symbols:         symbols,
		symbolMap:       make(map[string]string),
		numberFormat:    numberFormat,
		displayCurrency: displayCurrency,
	}
}

// SetNumberFormat updates the number format for price display
func (t *Manager) SetNumberFormat(format string) {
	t.numberFormat = format
}

// SetDisplayCurrency updates the display currency for price formatting
func (t *Manager) SetDisplayCurrency(currency string) {
	t.displayCurrency = currency
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
	for i, slot := range t.priceSlots {
		if i < len(t.symbols) {
			slot.SetTitle(fmt.Sprintf("%s $--,---", displaySymbols[i]))
			slot.Show()
		} else {
			slot.Hide()
		}
	}
}

// getDisplaySymbols converts coinIDs to ticker symbols for display
func (t *Manager) getDisplaySymbols() []string {
	result := make([]string, len(t.symbols))
	for i, coinID := range t.symbols {
		if symbol, ok := t.symbolMap[coinID]; ok {
			result[i] = symbol
		} else {
			result[i] = strings.ToUpper(coinID)
		}
	}
	return result
}

// UpdatePrices updates the tray display with multiple price data
func (t *Manager) UpdatePrices(data []*providers.PriceData) {
	if len(data) == 0 {
		return
	}

	// Build a map for quick lookup by coinID and update symbolMap
	priceMap := make(map[string]*providers.PriceData)
	for _, d := range data {
		priceMap[d.CoinID] = d
		t.symbolMap[d.CoinID] = d.Symbol
	}

	// Update each slot with its coinID's price
	for i, coinID := range t.symbols {
		if i >= len(t.priceSlots) {
			break
		}
		if d, ok := priceMap[coinID]; ok {
			price := d.ConvertedPrice
			if price == 0 {
				price = d.Price
			}
			currency := d.Currency
			if currency == "" {
				currency = t.displayCurrency
			}
			displayText := fmt.Sprintf("%s %s", d.Symbol, services.FormatPriceWithCurrency(price, t.numberFormat, currency))
			t.priceSlots[i].SetTitle(displayText)
		}
	}

	// Update tray title with all currencies
	var titleParts []string
	for _, coinID := range t.symbols {
		if d, ok := priceMap[coinID]; ok {
			price := d.ConvertedPrice
			if price == 0 {
				price = d.Price
			}
			currency := d.Currency
			if currency == "" {
				currency = t.displayCurrency
			}
			titleParts = append(titleParts, fmt.Sprintf("%s %s", d.Symbol, services.FormatPriceWithCurrency(price, t.numberFormat, currency)))
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
	for i := range t.symbols {
		if i >= len(t.priceSlots) {
			break
		}
		t.priceSlots[i].SetTitle(fmt.Sprintf("%s Error", displaySymbols[i]))
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
