# Change: Update Default Configuration

## Why
Binance has better rate limits than CoinGecko, making it a more reliable default provider. Additionally, the default currency selection should include the most popular cryptocurrencies (BTC, ETH, SOL) rather than just Bitcoin.

## What Changes
- Change default provider from CoinGecko to Binance
- Change default symbols from `["bitcoin"]` to `["BTCUSDT", "ETHUSDT", "SOLUSDT"]` (Binance format)
- Update provider interface to return multiple default symbols instead of a single one
- Each provider defines its own default symbols in their respective format
- Remove redundant `DefaultSymbol` constant from config

## Impact
- Affected specs: configuration
- Affected code:
  - `providers/provider.go` - Interface change
  - `providers/binance.go` - New defaults
  - `providers/coingecko.go` - New defaults
  - `config/types.go` - Default provider and symbols
  - `app.go` - Provider switch handling
  - `frontend/src/features/settings/constants/defaults.ts` - Frontend defaults
