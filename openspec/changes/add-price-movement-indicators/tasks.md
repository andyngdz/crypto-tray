## 1. Movement Package

- [x] 1.1 Create `movement/types.go` with Direction enum (Neutral, Up, Down)
- [x] 1.2 Add indicator constants (IndicatorUp, IndicatorDown, IndicatorNeutral)
- [x] 1.3 Add Indicator() method to Direction type
- [x] 1.4 Create `movement/tracker.go` with Tracker struct and NewTracker()
- [x] 1.5 Implement Track() method that compares prices and returns directions
- [x] 1.6 Add thread-safety with sync.RWMutex

## 2. Price Service Integration

- [x] 2.1 Add MovementTracker interface to `price/types.go`
- [x] 2.2 Add movement tracker field to Service struct in `price/service.go`
- [x] 2.3 Update NewService() to accept movement tracker parameter
- [x] 2.4 Call tracker.Track() in price callback and pass to tray

## 3. Tray Updates

- [x] 3.1 Update TrayUpdater interface to include movement data
- [x] 3.2 Update UpdatePrices() method to accept movements parameter
- [x] 3.3 Prepend indicator emoji to price display text

## 4. Main Wiring

- [x] 4.1 Create movement tracker in `main.go`
- [x] 4.2 Pass tracker to price service constructor
- [x] 4.3 Update initial price fetch to use movement tracking

## 5. Verification

- [x] 5.1 Build succeeds with `wails build -tags webkit2_41`
- [ ] 5.2 Run `wails dev -tags webkit2_41` - app starts successfully
- [ ] 5.3 Verify first fetch shows ⚪ (neutral) indicators
- [ ] 5.4 Wait for refresh, verify 🟢/🔴 indicators appear for changed prices
- [ ] 5.5 Restart app, verify indicators reset to ⚪
