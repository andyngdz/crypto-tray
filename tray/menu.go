package tray

import (
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// createMenu builds the systray context menu
func (m *Manager) createMenu() {
	m.menu = m.app.NewMenu()

	// Pre-allocate price menu items
	displaySymbols := m.getDisplaySymbols()

	for slotIdx := range maxPriceSlots {
		var item *application.MenuItem

		if slotIdx < len(m.symbols) {
			item = m.menu.Add(fmt.Sprintf("%s $--,---", displaySymbols[slotIdx]))
		} else {
			item = m.menu.Add("")
			item.SetHidden(true)
		}

		m.priceItems = append(m.priceItems, item)
	}

	m.menu.AddSeparator()

	m.menu.Add("Open Settings").OnClick(func(ctx *application.Context) {
		m.onOpenSettings()
	})

	m.menu.Add("Refresh Now").OnClick(func(ctx *application.Context) {
		m.onRefreshNow()
	})

	m.menu.AddSeparator()

	m.menu.Add("Quit").OnClick(func(ctx *application.Context) {
		m.onQuit()
	})

	m.systray.SetMenu(m.menu)
}

// updateSlots syncs the pre-allocated slots with current symbols
func (m *Manager) updateSlots() {
	displaySymbols := m.getDisplaySymbols()
	for slotIdx := range m.priceItems {
		item := m.priceItems[slotIdx]
		if slotIdx < len(m.symbols) {
			item.SetLabel(fmt.Sprintf("%s $--,---", displaySymbols[slotIdx]))
			item.SetHidden(false)
		} else {
			item.SetHidden(true)
		}
	}
}
