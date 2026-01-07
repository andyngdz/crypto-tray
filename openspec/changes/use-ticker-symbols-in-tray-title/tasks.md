# Tasks

## 1. Add symbolMap to tray manager
- [x] 1.1 Add `symbolMap map[string]string` field to `Manager` struct in `tray/types.go`
- [x] 1.2 Initialize `symbolMap` in `New()` function

## 2. Populate symbolMap from price data
- [x] 2.1 Update `UpdatePrices()` to populate `symbolMap` from `d.CoinID → d.Symbol`

## 3. Add helper to convert coinIDs to display symbols
- [x] 3.1 Add `getDisplaySymbols()` method that maps coinIDs to tickers via symbolMap
- [x] 3.2 Fallback to uppercase coinID if not in map (for initial load before first price fetch)

## 4. Update display functions to use ticker symbols
- [x] 4.1 Update `onReady()` to use `getDisplaySymbols()`
- [x] 4.2 Update `SetSymbols()` to use `getDisplaySymbols()`
- [x] 4.3 Update `SetLoading()` to use `getDisplaySymbols()`
- [x] 4.4 Update `SetError()` to use `getDisplaySymbols()`
- [x] 4.5 Update `updateSlots()` to use `getDisplaySymbols()`

## 5. Validation
- [x] 5.1 Build and verify loading state shows "ETH ... | BTC ..."
- [x] 5.2 Verify error state shows "ETH $??? | BTC $???"
- [x] 5.3 Verify price display still works correctly
