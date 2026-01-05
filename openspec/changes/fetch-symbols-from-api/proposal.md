# Change: Fetch Symbols from API

## Why

Currently, supported cryptocurrency symbols are hard-coded in `providers/coingecko.go:104-121` (14 coins). This limits users to a fixed set of cryptocurrencies and requires code changes to add new coins. CoinGecko's API provides a `/coins/list` endpoint that returns all 15,000+ supported coins.

## What Changes

- Replace hard-coded `GetSupportedSymbols()` with an API call to CoinGecko's `/coins/list` endpoint
- Cache the fetched symbols list to avoid repeated API calls
- Update the `Provider` interface to support async symbol fetching
- Refactor `SymbolInfo` struct to use provider-agnostic fields:
  - `CoinID` - provider-specific identifier for API calls (e.g., "bitcoin")
  - `Symbol` - user-facing ticker (e.g., "BTC")
  - `Name` - full display name (e.g., "Bitcoin")
- This abstraction allows future providers to map their own ID schemes to the common interface

## Impact

- Affected code: `providers/types.go`, `providers/coingecko.go`, `providers/provider.go`, `app.go`
- Affected specs: `symbol-management` (new capability)
- **BREAKING**: `SymbolInfo.ID` renamed to `SymbolInfo.Symbol`, new `CoinID` field added
