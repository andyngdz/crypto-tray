package tray

import (
	"github.com/wailsapp/wails/v3/pkg/application"

	"crypto-tray/movement"
	"crypto-tray/providers"
)

// maxPriceSlots is the maximum number of currency slots in the tray menu
const maxPriceSlots = 10

// TrayUpdater interface for price service integration
type TrayUpdater interface {
	SetError(msg string)
	UpdatePrices(data []*providers.PriceData, movements map[string]movement.Direction)
}

// Manager handles system tray operations using Wails v3 native systray
type Manager struct {
	app            *application.App
	systray        *application.SystemTray
	window         *application.WebviewWindow
	menu           *application.Menu
	priceItems     []*application.MenuItem
	onOpenSettings func()
	onRefreshNow   func()
	onQuit         func()
	symbols        []string
	symbolMap      map[string]string
	numberFormat   string
}
