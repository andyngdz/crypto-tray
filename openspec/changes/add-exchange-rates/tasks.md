## 1. Exchange Package

- [x] 1.1 Create `exchange/types.go` with ExchangeRates struct and Callback type
- [x] 1.2 Create `exchange/fetcher.go` with NewFetcher, Start, Stop, RefreshNow, GetRates
- [x] 1.3 Implement fetchFromURL with primary/fallback URL logic
- [x] 1.4 Implement in-memory caching (only update on success)

## 2. Config Updates

- [x] 2.1 Add `DisplayCurrency string` field to `config/types.go`
- [x] 2.2 Add `DefaultDisplayCurrency = "usd"` constant
- [x] 2.3 Update defaultConfig() to include DisplayCurrency

## 3. Price Data Updates

- [x] 3.1 Add `ConvertedPrice float64` to PriceData in `providers/types.go`
- [x] 3.2 Add `Currency string` to PriceData in `providers/types.go`

## 4. Converter Service

- [x] 4.1 Create `exchange/converter.go` with Converter struct (moved from services/ to avoid import cycle)
- [x] 4.2 Implement NewConverter(fetcher, configManager)
- [x] 4.3 Implement ConvertPrices(data []*PriceData) method

## 5. Formatter Updates

- [x] 5.1 Add currency symbol map to `services/formatter.go`
- [x] 5.2 Add FormatPriceWithCurrency function

## 6. App Integration

- [x] 6.1 Add `onDisplayCurrencyChanged` callback to App struct in `app.go`
- [x] 6.2 Update SaveConfig to trigger callback when DisplayCurrency changes
- [x] 6.3 Add setOnDisplayCurrencyChanged setter method

## 7. Main Wiring

- [x] 7.1 Create exchange fetcher in `main.go`
- [x] 7.2 Create converter with fetcher + config manager
- [x] 7.3 Update price callback to call converter.ConvertPrices
- [x] 7.4 Connect onDisplayCurrencyChanged to trigger refresh
- [x] 7.5 Start/stop exchange fetcher in OnStartup/OnShutdown

## 8. Tray Updates

- [x] 8.1 Update `tray/tray.go` to use ConvertedPrice when available
- [x] 8.2 Pass currency to formatter for correct symbol display
- [x] 8.3 Add displayCurrency field and SetDisplayCurrency method

## 9. Verification

- [x] 9.1 Build succeeds with `wails build -tags webkit2_41`
- [x] 9.2 Run `wails dev -tags webkit2_41` - app starts successfully
- [x] 9.3 Manually set DisplayCurrency in config.json, restart, verify conversion
