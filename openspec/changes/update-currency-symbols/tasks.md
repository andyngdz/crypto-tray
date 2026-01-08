## 1. Update Formatter

- [x] 1.1 Remove hardcoded `currencySymbols` map from `services/formatter.go`
- [x] 1.2 Add import for `golang.org/x/text/currency`
- [x] 1.3 Update `GetCurrencySymbol` to use `currency.ParseISO()`
- [x] 1.4 FormatPriceWithCurrency already uses GetCurrencySymbol

## 2. Verification

- [x] 2.1 Build with `wails build -tags webkit2_41`
- [x] 2.2 Test EUR displays as "€"
- [x] 2.3 Test GBP displays as "£"
- [x] 2.4 Test JPY displays as "¥"
