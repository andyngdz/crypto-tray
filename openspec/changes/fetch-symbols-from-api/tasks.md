# Tasks

## 1. Refactor SymbolInfo Struct

- [x] 1.1 Update `SymbolInfo` in `providers/types.go`:
  - Rename `ID` → `Symbol` (user-facing ticker, e.g., "BTC")
  - Add `CoinID` field (provider-specific ID, e.g., "bitcoin")
- [x] 1.2 Update all usages of `SymbolInfo.ID` to `SymbolInfo.Symbol` across codebase
- [x] 1.3 Update frontend types to match new struct

## 2. Backend Implementation

- [x] 2.1 Add `FetchSymbols(ctx) ([]SymbolInfo, error)` method to `Provider` interface
- [x] 2.2 Implement CoinGecko `/coins/list` API call in `CoinGecko.FetchSymbols()`
- [x] 2.3 Add in-memory cache for symbols with TTL (24 hours)
- [x] 2.4 Refactor `symbolToCoinID()` to use `SymbolInfo.CoinID` from cached data
- [x] 2.5 Update `App.GetAvailableSymbols()` to call the new async method

## 3. Frontend Integration

- [x] 3.1 Update `fetchAvailableSymbols` to handle loading state during initial fetch
- [x] 3.2 Add error handling for symbol fetch failures

## 4. Testing

- [x] 4.1 Build and verify the application runs correctly
- [x] 4.2 Test symbol autocomplete with API-fetched symbols
