# Change: Add Multi-Currency Support

## Why
Currently the application only supports tracking a single cryptocurrency at a time. Users want to monitor multiple currencies (e.g., BTC, ETH, SOL) simultaneously without switching between them.

## What Changes
- **Config**: Change `symbol: string` to `symbols: string[]` to store multiple currencies
- **Provider Interface**: Add batch fetch method `FetchPrices(symbols []string)` with fallback for providers that don't support batch
- **CoinGecko Provider**: Implement batch fetch (API already supports multiple coins in one request)
- **Price Fetcher**: Update to fetch and return multiple prices
- **Tray Display**: Show each currency as a separate menu item (primary currency in tray title)
- **Settings UI**: Add multi-select dropdown for currency selection using HeroUI

## Impact
- Affected specs: None (first spec for this capability)
- Affected code:
  - `config/config.go` - Config struct, defaults, validation
  - `providers/provider.go` - Provider interface, PriceData
  - `providers/coingecko.go` - Batch fetch implementation
  - `price/fetcher.go` - Multi-price fetching logic
  - `tray/tray.go` - Multiple menu items
  - `app.go` - New binding for available symbols
  - `frontend/src/features/settings/` - New symbols selector component
