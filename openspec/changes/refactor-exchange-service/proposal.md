# Change: Refactor Exchange to Service Layer

## Why

main.go is polluted with callback logic. Business logic (error handling, event emission) is mixed with initialization code.

## What Changes

- Create `exchange.Service` that wraps fetcher, converter, and event emission
- Simplify main.go to just call `exchangeService.Start()` and `exchangeService.Stop()`
- Move callback logic from main.go into the service layer

## Impact

- New file: `exchange/service.go`
- Modified: `exchange/fetcher.go`, `exchange/types.go`, `main.go`
