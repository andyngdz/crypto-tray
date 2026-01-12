## 1. Backend Provider Interface
- [x] 1.1 Update `providers/provider.go` - rename `DefaultCoinID() string` to `DefaultCoinIDs() []string`

## 2. Provider Implementations
- [x] 2.1 Update `providers/binance.go` - implement `DefaultCoinIDs()` returning `["BTCUSDT", "ETHUSDT", "SOLUSDT"]`
- [x] 2.2 Update `providers/coingecko.go` - implement `DefaultCoinIDs()` returning `["bitcoin", "ethereum", "solana"]`

## 3. Configuration Defaults
- [x] 3.1 Update `config/types.go` - change `DefaultProviderID` to `"binance"`
- [x] 3.2 Update `config/types.go` - remove `DefaultSymbol` constant
- [x] 3.3 Update `config/types.go` - update `defaultConfig()` to use Binance default symbols

## 4. App Logic
- [x] 4.1 Update `app.go` - change `newProvider.DefaultCoinID()` to `newProvider.DefaultCoinIDs()`

## 5. Frontend
- [x] 5.1 Update `frontend/src/features/settings/constants/defaults.ts` - change `providerId` to `'binance'`

## 6. Testing
- [x] 6.1 Verified on Windows - defaults work correctly (Binance provider, BTC/ETH/SOL selected)
