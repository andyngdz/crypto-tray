## 1. Provider Interface Update

- [x] 1.1 Add `DefaultCoinID() string` method to Provider interface in `providers/provider.go`
- [x] 1.2 Implement `DefaultCoinID()` in CoinGecko provider returning `"bitcoin"`

## 2. Binance Provider Implementation

- [x] 2.1 Create `providers/binance.go` with constants (base URL, timeout, cache TTL, quote asset)
- [x] 2.2 Add Binance API response types (`binanceTicker24hr`, `binanceExchangeInfo`, `binanceSymbol`)
- [x] 2.3 Add hardcoded coin name map for top cryptocurrencies
- [x] 2.4 Implement `Binance` struct with HTTPClient and symbol cache
- [x] 2.5 Implement `ID()`, `Name()`, `RequiresAPIKey()`, `SetAPIKey()`, `DefaultCoinID()` methods
- [x] 2.6 Implement `FetchSymbols()` - fetch from exchangeInfo, filter USDT pairs, cache 24h
- [x] 2.7 Implement `FetchPrices()` - fetch from ticker/24hr with JSON array symbols format
- [x] 2.8 Implement `coinIDToSymbol()` helper for mapping BTCUSDT -> BTC

## 3. Provider Registration

- [x] 3.1 Register Binance provider in `init.go`

## 4. Provider Switch Handling

- [x] 4.1 Update `app.go` SaveConfig to reset symbols to provider's DefaultCoinID when provider changes

## 5. Frontend Fix

- [x] 5.1 Add `onSuccess` callback to `saveMutation` in `useConfig.ts` to invalidate symbols and config queries when provider changes

## 6. Validation

- [x] 6.1 Run `wails build -tags webkit2_41` to verify compilation
- [x] 6.2 Verify Binance API responds correctly
- [x] 6.3 Manual testing: Run app and test provider selection in Settings
- [x] 6.4 Manual testing: Verify prices display with 24h change
- [x] 6.5 Manual testing: Verify provider switch resets to BTC default
