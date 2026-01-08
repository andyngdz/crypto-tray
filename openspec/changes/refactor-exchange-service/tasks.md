## 1. Types Update

- [x] 1.1 Add `ContextProvider` interface to `exchange/types.go`

## 2. Service Layer

- [x] 2.1 Create `exchange/service.go` with Service struct
- [x] 2.2 Implement `NewService(configManager, contextProvider)`
- [x] 2.3 Implement `Start()`, `Stop()`, `GetConverter()` methods

## 3. Fetcher Update

- [x] 3.1 Update `exchange/fetcher.go` to accept callback in Start() instead of constructor

## 4. Main Integration

- [x] 4.1 Update `main.go` to use `exchange.NewService()`
- [x] 4.2 Remove inline callback logic from main.go

## 5. Verification

- [x] 5.1 Build with `wails build -tags webkit2_41`
- [x] 5.2 Run and verify exchange rates update correctly
