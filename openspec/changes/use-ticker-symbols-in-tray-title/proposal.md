# Proposal: Use Ticker Symbols in Tray Title

## Summary
Fix the tray title to show ticker symbols (ETH, BTC, USDT) instead of coinIDs (ethereum, bitcoin, tether) during loading, error, and placeholder states.

## Motivation
The spec defines that loading state should show "ETH ... | BTC ... | USDT ..." but the current implementation shows "ethereum ... | bitcoin ... | tether ..." because `t.symbols` stores coinIDs, not ticker symbols.

## Scope
- **In scope**: Fix loading/error/placeholder states to show ticker symbols
- **Out of scope**: The price display already works correctly (uses `d.Symbol` from price data)

## Root Cause
- `tray.Manager.symbols` stores coinIDs (e.g., "ethereum")
- `FormatTrayTitle()` uses these directly without mapping to tickers
- `UpdatePrices()` works correctly because it uses `d.Symbol` from price data

## Solution
Add a `symbolMap` (coinID→ticker) to the tray manager, populated when prices are fetched.

## Affected Files
- `tray/types.go` - Add `symbolMap` field
- `tray/tray.go` - Initialize map, populate from prices, use for display
