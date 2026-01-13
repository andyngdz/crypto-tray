//go:build !windows

package tray

import (
	"fmt"

	"crypto-tray/movement"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// updatePriceItem updates a menu item with price and movement indicator
// On Linux/macOS: uses emoji in text
func (t *Manager) updatePriceItem(item *application.MenuItem, dir movement.Direction, symbol, priceText string) {
	item.SetLabel(fmt.Sprintf("%s %s %s", dir.Indicator(), symbol, priceText))
}
