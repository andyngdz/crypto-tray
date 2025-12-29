# Tasks: Add Multi-Currency Support

## 1. Backend - Config
- [ ] 1.1 Update `Config` struct: change `Symbol string` to `Symbols []string`
- [ ] 1.2 Update `DefaultConfig()` to use `Symbols: []string{"BTC"}`
- [ ] 1.3 Add migration logic in `Load()` to convert old `symbol` to `symbols`
- [ ] 1.4 Update validation to ensure at least one symbol

## 2. Backend - Provider Interface
- [ ] 2.1 Add `FetchPrices(ctx, symbols []string) ([]*PriceData, error)` to Provider interface
- [ ] 2.2 Add `SupportsBatchFetch() bool` to Provider interface
- [ ] 2.3 Add `GetSupportedSymbols() []SymbolInfo` to Provider interface
- [ ] 2.4 Create `SymbolInfo` struct with `ID` and `Name` fields

## 3. Backend - CoinGecko Provider
- [ ] 3.1 Add supported symbols list with ID mapping (BTC, ETH, SOL, etc.)
- [ ] 3.2 Implement `FetchPrices()` using batch API call
- [ ] 3.3 Implement `SupportsBatchFetch()` returning true
- [ ] 3.4 Implement `GetSupportedSymbols()` returning available coins

## 4. Backend - Price Fetcher
- [ ] 4.1 Update callback type to `func(data []*PriceData, err error)`
- [ ] 4.2 Update `fetchPrice()` to call `FetchPrices()` with all symbols
- [ ] 4.3 Handle partial failures (some symbols succeed, some fail)

## 5. Backend - Tray Manager
- [ ] 5.1 Change `priceItem *systray.MenuItem` to `priceItems map[string]*systray.MenuItem`
- [ ] 5.2 Update `New()` to accept initial symbols list
- [ ] 5.3 Create menu items dynamically in `onReady()` based on symbols
- [ ] 5.4 Add `UpdatePrices(data []*PriceData)` method
- [ ] 5.5 Update tray title with first/primary currency
- [ ] 5.6 Add `SetSymbols(symbols []string)` to update tracked symbols

## 6. Backend - App Bindings
- [ ] 6.1 Add `GetAvailableSymbols()` binding returning `[]SymbolInfo`
- [ ] 6.2 Update `SaveConfig()` to handle symbols array
- [ ] 6.3 Update main.go to wire new multi-currency flow

## 7. Frontend - Types
- [ ] 7.1 Update `Config` interface: `symbol: string` to `symbols: string[]`
- [ ] 7.2 Add `SymbolInfo` interface with `id` and `name`
- [ ] 7.3 Update type guards for new structure

## 8. Frontend - Services
- [ ] 8.1 Add `fetchAvailableSymbols()` in configService.ts
- [ ] 8.2 Generate Wails bindings for new `GetAvailableSymbols` function

## 9. Frontend - State
- [ ] 9.1 Create `useSettingsSymbols.ts` hook for symbol selection state
- [ ] 9.2 Update `useConfig.ts` to include symbols in query

## 10. Frontend - UI
- [ ] 10.1 Create `SettingsSymbols.tsx` component with multi-select dropdown
- [ ] 10.2 Add `SettingsSymbols` to `SettingsView.tsx`
- [ ] 10.3 Export new component from feature index

## 11. Testing & Validation
- [ ] 11.1 Test config migration from old format
- [ ] 11.2 Test batch fetch with multiple symbols
- [ ] 11.3 Test tray display with multiple currencies
- [ ] 11.4 Test UI multi-select functionality
- [ ] 11.5 Build and run full application
