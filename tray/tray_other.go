//go:build !windows

package tray

import (
	"fmt"

	"crypto-tray/movement"

	"fyne.io/systray"
)

// updatePriceSlot updates a menu item with price and movement indicator
// On Linux/macOS: uses emoji in text (SetIcon is no-op on Linux anyway)
func (t *Manager) updatePriceSlot(slot *systray.MenuItem, dir movement.Direction, symbol, priceText string) {
	slot.SetTitle(fmt.Sprintf("%s %s %s", dir.Indicator(), symbol, priceText))
}
