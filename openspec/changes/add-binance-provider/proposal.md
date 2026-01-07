# Change: Add Binance as a Cryptocurrency Data Provider

## Why

Currently the app only supports CoinGecko as a data provider. Users want the option to use Binance for cryptocurrency price data, as it offers real-time trading prices and may be preferred in certain regions or use cases.

## What Changes

- Add new `Binance` provider implementing the `Provider` interface
- Add `DefaultCoinID()` method to Provider interface for provider-specific defaults
- Update CoinGecko to implement `DefaultCoinID()`
- Reset symbols to BTC default when switching providers (coin IDs are provider-specific)
- Register Binance provider in application initialization

## Impact

- Affected specs: `price-providers` (new capability)
- Affected code:
  - `providers/binance.go` (new file)
  - `providers/provider.go` (interface update)
  - `providers/coingecko.go` (add DefaultCoinID)
  - `init.go` (register provider)
  - `app.go` (handle provider switch)
