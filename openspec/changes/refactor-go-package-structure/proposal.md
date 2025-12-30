# Change: Refactor Go Package Structure

## Why
The current Go packages mix types, interfaces, and logic in single files. Splitting by responsibility improves readability, maintainability, and follows Go conventions for medium-sized projects.

## What Changes
- Split `providers/provider.go` into `types.go`, `provider.go`, `registry.go`
- Split `config/config.go` into `types.go`, `validation.go`, `manager.go`
- Split `price/fetcher.go` into `types.go`, `fetcher.go`
- Split `tray/tray.go` into `types.go`, `tray.go`
- Keep `services/formatter.go` unchanged (only 13 lines)
- Keep root `main` package unchanged (already well organized)

## Impact
- Affected specs: None (internal refactoring only)
- Affected code: `providers/`, `config/`, `price/`, `tray/` packages
- No breaking changes to public APIs
- No behavior changes
