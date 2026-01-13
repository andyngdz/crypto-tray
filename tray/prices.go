package tray

import (
	"fmt"
	"strings"

	"crypto-tray/movement"
	"crypto-tray/providers"
	"crypto-tray/services"
)

// UpdatePrices updates the tray display with multiple price data and movement indicators
func (t *Manager) UpdatePrices(data []*providers.PriceData, movements map[string]movement.Direction) {
	if len(data) == 0 {
		return
	}

	// Build a map for quick lookup by coinID and update symbolMap
	priceMap := make(map[string]*providers.PriceData)
	for dataIdx := range data {
		priceData := data[dataIdx]
		priceMap[priceData.CoinID] = priceData
		t.symbolMap[priceData.CoinID] = priceData.Symbol
	}

	// Update each slot with its coinID's price and movement indicator
	for symbolIdx := range t.symbols {
		if symbolIdx >= len(t.priceItems) {
			break
		}

		coinID := t.symbols[symbolIdx]
		priceData, ok := priceMap[coinID]
		if !ok {
			continue
		}

		price := priceData.ConvertedPrice
		if price == 0 {
			price = priceData.Price
		}

		dir := movement.Neutral
		if movementDir, ok := movements[coinID]; ok {
			dir = movementDir
		}

		priceText := services.FormatPriceWithCurrency(price, t.numberFormat, priceData.Currency)
		t.updatePriceItem(t.priceItems[symbolIdx], dir, priceData.Symbol, priceText)
	}

	// Update tray title with all currencies and movement indicators
	t.updateTrayTitle(priceMap, movements)
}

// updateTrayTitle updates the systray label with price information
func (t *Manager) updateTrayTitle(priceMap map[string]*providers.PriceData, movements map[string]movement.Direction) {
	var titleParts []string

	for symbolIdx := range t.symbols {
		coinID := t.symbols[symbolIdx]
		priceData, ok := priceMap[coinID]
		if !ok {
			continue
		}

		price := priceData.ConvertedPrice
		if price == 0 {
			price = priceData.Price
		}

		indicator := movement.IndicatorNeutral
		if dir, ok := movements[coinID]; ok {
			indicator = dir.Indicator()
		}

		priceText := services.FormatPriceWithCurrency(price, t.numberFormat, priceData.Currency)
		titleParts = append(titleParts, fmt.Sprintf("%s %s %s", indicator, priceData.Symbol, priceText))
	}

	if len(titleParts) > 0 {
		title := strings.Join(titleParts, " | ")
		t.systray.SetLabel(title)
	}
}

// SetError updates tray to show error state
func (t *Manager) SetError(msg string) {
	displaySymbols := t.getDisplaySymbols()

	if len(t.symbols) > 0 {
		label := services.FormatTrayTitle(displaySymbols, "$???")
		t.systray.SetLabel(label)
	}

	for symbolIdx := range t.symbols {
		if symbolIdx >= len(t.priceItems) {
			break
		}
		t.priceItems[symbolIdx].SetLabel(fmt.Sprintf("%s Error", displaySymbols[symbolIdx]))
	}
}

// SetLoading shows loading state
func (t *Manager) SetLoading() {
	if len(t.symbols) > 0 {
		label := services.FormatTrayTitle(t.getDisplaySymbols(), "...")
		t.systray.SetLabel(label)
	}
}
