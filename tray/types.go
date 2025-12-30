package tray

import "github.com/getlantern/systray"

// Manager handles system tray operations
type Manager struct {
	onOpenSettings func()
	onRefreshNow   func()
	onQuit         func()
	priceItems     map[string]*systray.MenuItem
	symbols        []string
}
