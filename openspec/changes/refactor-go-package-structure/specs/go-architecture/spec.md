## ADDED Requirements

### Requirement: Go Package File Organization
Each Go package SHALL organize code into separate files by responsibility:
- `types.go` - Data structures and type definitions
- `<entity>.go` - Interface definitions (e.g., `provider.go` for `Provider` interface)
- `<component>.go` - Implementation logic (e.g., `manager.go`, `registry.go`, `fetcher.go`)

#### Scenario: Providers package structure
- **WHEN** viewing the `providers/` package
- **THEN** it contains `types.go` (PriceData, SymbolInfo), `provider.go` (Provider interface), `registry.go` (Registry), and implementation files

#### Scenario: Config package structure
- **WHEN** viewing the `config/` package
- **THEN** it contains `types.go` (Config struct, constants), `validation.go` (Validate method), and `manager.go` (Manager)

#### Scenario: Price package structure
- **WHEN** viewing the `price/` package
- **THEN** it contains `types.go` (Callback type) and `fetcher.go` (Fetcher implementation)

#### Scenario: Tray package structure
- **WHEN** viewing the `tray/` package
- **THEN** it contains `types.go` (Manager struct) and `tray.go` (icon embed, methods)
