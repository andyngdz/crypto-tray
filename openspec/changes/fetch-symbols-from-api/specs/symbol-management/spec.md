## ADDED Requirements

### Requirement: Provider-Agnostic Symbol Interface

The system SHALL use a standardized `SymbolInfo` structure with provider-agnostic fields to support multiple price providers.

#### Scenario: SymbolInfo structure

- **WHEN** representing a cryptocurrency symbol
- **THEN** the structure SHALL contain:
  - `CoinID` - provider-specific identifier for API calls (e.g., "bitcoin")
  - `Symbol` - user-facing ticker in uppercase (e.g., "BTC")
  - `Name` - full display name (e.g., "Bitcoin")

#### Scenario: Provider maps to common interface

- **WHEN** a provider fetches its coin list
- **THEN** it SHALL map provider-specific response fields to the common `SymbolInfo` interface
- **AND** normalize the `Symbol` field to uppercase

### Requirement: Dynamic Symbol Fetching

The system SHALL fetch the list of supported cryptocurrency symbols from the provider's API instead of using a hard-coded list.

#### Scenario: Symbols fetched from CoinGecko API

- **WHEN** the application requests available symbols
- **THEN** it SHALL call the CoinGecko `/coins/list` endpoint
- **AND** return symbol info mapped to the common interface

#### Scenario: API unavailable

- **WHEN** the symbol fetch API call fails
- **THEN** the system SHALL return an empty list
- **AND** log the error for debugging

### Requirement: Symbol Cache

The system SHALL cache fetched symbols to minimize API calls.

#### Scenario: Cache hit

- **WHEN** symbols are requested within the cache TTL (24 hours)
- **THEN** the system SHALL return cached symbols without making an API call

#### Scenario: Cache miss or expired

- **WHEN** symbols are requested and cache is empty or expired
- **THEN** the system SHALL fetch fresh symbols from the API
- **AND** update the cache with the new data

### Requirement: Symbol-to-CoinID Mapping

The system SHALL use the cached symbol data to map user-facing symbols to provider-specific coin IDs.

#### Scenario: Price lookup uses CoinID from cache

- **WHEN** fetching prices for a symbol (e.g., "BTC")
- **THEN** the system SHALL lookup the `CoinID` field from cached symbol data (e.g., "bitcoin")
- **AND** NOT rely on a hard-coded mapping table
