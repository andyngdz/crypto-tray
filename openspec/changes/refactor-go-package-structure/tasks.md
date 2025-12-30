# Tasks: Refactor Go Package Structure

## 1. Providers Package
- [x] 1.1 Create `providers/types.go` with `PriceData` and `SymbolInfo` structs
- [x] 1.2 Create `providers/registry.go` with `Registry` struct and methods
- [x] 1.3 Update `providers/provider.go` to contain only `Provider` interface
- [x] 1.4 Remove old content from `provider.go`, verify imports

## 2. Config Package
- [x] 2.1 Create `config/types.go` with constants and `Config` struct
- [x] 2.2 Create `config/validation.go` with `Validate()` method
- [x] 2.3 Create `config/manager.go` with `Manager` struct and methods
- [x] 2.4 Delete `config/config.go`

## 3. Price Package
- [x] 3.1 Create `price/types.go` with `Callback` type
- [x] 3.2 Update `price/fetcher.go` to remove `Callback` type definition

## 4. Tray Package
- [x] 4.1 Create `tray/types.go` with `Manager` struct definition
- [x] 4.2 Update `tray/tray.go` to keep icon embed and methods only

## 5. Validation
- [x] 5.1 Run `go vet ./...` to verify no errors
- [x] 5.2 Run `go build .` to verify compilation
- [x] 5.3 Run `wails generate module` to verify bindings still work
