package price

import (
	"context"

	"crypto-tray/movement"
	"crypto-tray/providers"
)

// Callback is called when price data is fetched
type Callback func(data []*providers.PriceData, err error)

// TrayUpdater handles tray display updates
type TrayUpdater interface {
	SetError(msg string)
	UpdatePrices(data []*providers.PriceData, movements map[string]movement.Direction)
}

// PriceConverter handles currency conversion for price data
type PriceConverter interface {
	ConvertPrices(data []*providers.PriceData)
}

// ContextProvider provides Wails context for event emission
type ContextProvider interface {
	GetContext() context.Context
}

// MovementTracker tracks price movements between fetches
type MovementTracker interface {
	Track(data []*providers.PriceData) map[string]movement.Direction
}
