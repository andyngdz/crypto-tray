# Change: Use Dynamic Currency Symbols

## Why

Currency symbols currently display as text codes (e.g., "EUR77,702") instead of proper Unicode symbols (e.g., "€77,702"). The hardcoded symbol map is incomplete and unmaintainable.

## What Changes

- Replace hardcoded `currencySymbols` map with dynamic lookup using `golang.org/x/text/currency`
- Use `currency.ParseISO()` to get proper Unicode symbols for any ISO 4217 currency code
- Graceful fallback to uppercase code for unknown currencies

## Impact

- Affected specs: exchange-rates (modified requirement)
- Affected code: `services/formatter.go`
- No new dependencies (golang.org/x/text already in use)
