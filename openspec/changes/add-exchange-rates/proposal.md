# Change: Add Exchange Rate Fetching for Currency Conversion

## Why

Users need to view crypto prices in their local currency (EUR, GBP, JPY, etc.) rather than just USD/USDT. Currently, prices are always displayed in the provider's quote currency (typically USDT).

## What Changes

- Fetch exchange rates from fawazahmed0/exchange-api
- Add `DisplayCurrency` config field (default: "usd")
- Convert crypto prices to user's selected display currency before showing in tray/frontend
- Cache exchange rates and refresh on the same interval as price data
- Graceful fallback to USDT if exchange rate fetch fails

## Impact

- Affected specs: exchange-rates (new capability)
- Affected code:
  - `config/types.go` - Add DisplayCurrency field
  - `providers/types.go` - Add ConvertedPrice/Currency to PriceData
  - `main.go` - Wire exchange fetcher and converter
  - `app.go` - Add callback for currency change
  - `tray/tray.go` - Display converted prices
  - `services/formatter.go` - Currency symbol mapping
- New files:
  - `exchange/types.go` - Exchange rate data structures
  - `exchange/fetcher.go` - API client with fallback
  - `exchange/converter.go` - Price conversion service (in exchange/ to avoid import cycle with providers)
