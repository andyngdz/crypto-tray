# Tasks

## 1. Update tray title formatting
- [x] 1.1 Add `FormatTrayTitle` helper to `services/formatter.go`
- [x] 1.2 Update `UpdatePrices()` to concatenate all currency prices with " | " separator
- [x] 1.3 Update `onReady()` initial title to show all symbols with placeholders
- [x] 1.4 Update `SetSymbols()` to show all symbols in title
- [x] 1.5 Update `SetLoading()` to show loading state for all symbols
- [x] 1.6 Update `SetError()` to show error state for all symbols

## 2. Refactoring
- [x] 2.1 Extract repeated title-building logic to `services.FormatTrayTitle()`
- [x] 2.2 Simplify `onReady()`, `SetSymbols()`, `SetLoading()`, `SetError()` with new helper
- [x] 2.3 Keep `UpdatePrices()` inline (uses price data, not simple suffix)

## 3. Validation
- [x] 3.1 Build and test with single currency
- [x] 3.2 Build and test with multiple currencies (3+)
- [x] 3.3 Verify tooltip also shows all currencies
