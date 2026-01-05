package tray

import "github.com/getlantern/systray"

// maxPriceSlots is the maximum number of currency slots in the tray menu
const maxPriceSlots = 10

// Manager handles system tray operations
type Manager struct {
	onOpenSettings func()
	onRefreshNow   func()
	onQuit         func()
	priceSlots     []*systray.MenuItem // Pre-allocated menu item slots
	symbols        []string            // Currently active symbols (maps to slots by index)
}
