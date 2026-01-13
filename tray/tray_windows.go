//go:build windows

package tray

import (
	_ "embed"
	"fmt"

	"crypto-tray/movement"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed up.ico
var iconUp []byte

//go:embed down.ico
var iconDown []byte

//go:embed neutral.ico
var iconNeutral []byte

// iconForDirection returns icon bytes for a given movement direction
func iconForDirection(dir movement.Direction) []byte {
	switch dir {
	case movement.Up:
		return iconUp
	case movement.Down:
		return iconDown
	default:
		return iconNeutral
	}
}

// updatePriceItem updates a menu item with price and movement indicator
// On Windows: uses SetBitmap() with ICO files because Win32 menus don't render emoji well
func (t *Manager) updatePriceItem(item *application.MenuItem, dir movement.Direction, symbol, priceText string) {
	item.SetBitmap(iconForDirection(dir))
	item.SetLabel(fmt.Sprintf("%s %s", symbol, priceText))
}
