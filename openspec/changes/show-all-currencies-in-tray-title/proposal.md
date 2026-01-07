# Proposal: Show All Currencies in Tray Title

## Summary
Display all selected currencies in the system tray title instead of only showing the first/primary currency. Users who select multiple currencies (e.g., ETH, BTC, USDT) should see all prices in the tray title.

## Motivation
Currently, when a user selects 3 currencies, only the first one appears in the tray title (e.g., "ETH $3,252"). Users expect to see all their selected currencies at a glance without opening the dropdown menu.

## Scope
- **In scope**: Update tray title to show all currencies with prices separated by " | "
- **Out of scope**: Changing the dropdown menu display, adding user preference for title format

## Target Format
```
ETH $3,252 | BTC $97,500 | USDT $1.00
```

## Affected Files
- `tray/tray.go` - Update title formatting in `UpdatePrices()`, `onReady()`, `SetSymbols()`, `SetLoading()`, `SetError()`

## Dependencies
- Modifies existing behavior from `add-multi-currency-support` change
