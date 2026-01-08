package tray

import "github.com/getlantern/systray"

// maxPriceSlots is the maximum number of currency slots in the tray menu
const maxPriceSlots = 10

// Manager handles system tray operations
type Manager struct {
	onOpenSettings  func()
	onRefreshNow    func()
	onQuit          func()
	priceSlots      []*systray.MenuItem // Pre-allocated menu item slots
	symbols         []string            // Currently active coinIDs (maps to slots by index)
	symbolMap       map[string]string   // coinID -> ticker symbol (e.g., "ethereum" -> "ETH")
	numberFormat    string              // Number format: "us", "european", or "asian"
	displayCurrency string              // Display currency code (e.g., "usd", "eur")
}
