# Change: Add Price Movement Indicators

## Why

Users want visual feedback on whether crypto prices have increased or decreased since the last fetch. This helps them quickly identify market trends at a glance without needing to remember previous values.

## What Changes

- Add emoji indicators to show price movement direction in the tray menu
- Track price changes between fetches (in-memory, resets on restart)
- Display: 🟢 (up), 🔴 (down), ⚪ (neutral/first fetch)

## Impact

- Affected specs: price-movement (new capability)
- Affected code:
  - `tray/tray.go` - Display emoji indicators with prices
  - `price/service.go` - Integrate movement tracking
  - `price/types.go` - Add MovementTracker interface
  - `main.go` - Wire up movement tracker
- New files:
  - `movement/types.go` - Direction enum and indicator constants
  - `movement/tracker.go` - Price movement tracking service
