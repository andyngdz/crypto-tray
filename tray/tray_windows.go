//go:build windows

package tray

import (
	_ "embed"
	"fmt"

	"crypto-tray/movement"

	"github.com/getlantern/systray"
)

//go:embed up.ico
var iconUp []byte

//go:embed down.ico
var iconDown []byte

//go:embed neutral.ico
var iconNeutral []byte

// iconForDirection returns the icon bytes for a given movement direction
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

// updatePriceSlot updates a menu item with price and movement indicator
// On Windows: uses SetIcon() with ICO files because emoji renders as grey in Win32 menus
func (t *Manager) updatePriceSlot(slot *systray.MenuItem, dir movement.Direction, symbol, priceText string) {
	slot.SetIcon(iconForDirection(dir))
	slot.SetTitle(fmt.Sprintf("%s %s", symbol, priceText))
}
