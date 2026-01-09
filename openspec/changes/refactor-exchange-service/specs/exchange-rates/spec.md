## MODIFIED Requirements

### Requirement: Exchange Rate Fetching

The system SHALL provide an exchange.Service that encapsulates fetching, conversion, and event emission.

#### Scenario: Service initialization

- **WHEN** main.go initializes exchange service
- **THEN** it only calls `NewService(configManager, contextProvider)`
- **AND** no callback logic is defined in main.go

#### Scenario: Service lifecycle

- **WHEN** app starts
- **THEN** main.go calls `exchangeService.Start()`
- **WHEN** app shuts down
- **THEN** main.go calls `exchangeService.Stop()`

#### Scenario: Converter access

- **WHEN** price fetcher needs to convert prices
- **THEN** it calls `exchangeService.GetConverter()`
