## ADDED Requirements

### Requirement: Exchange Rate Fetching
The system SHALL fetch exchange rates from fawazahmed0/exchange-api to convert crypto prices to the user's display currency.

#### Scenario: Successful rate fetch
- **WHEN** the refresh interval triggers
- **THEN** the system fetches rates from the primary API URL
- **AND** caches the rates in memory

#### Scenario: Primary URL fails
- **WHEN** the primary API URL is unavailable
- **THEN** the system retries with the fallback URL
- **AND** logs a warning

#### Scenario: Both URLs fail
- **WHEN** both primary and fallback URLs fail
- **THEN** the system continues using cached rates
- **AND** logs a warning (not shown to user)

#### Scenario: No cached rates on startup
- **WHEN** the system starts and exchange rate fetch fails
- **AND** no cached rates exist
- **THEN** prices display in USDT (rate = 1.0)

### Requirement: Price Conversion
The system SHALL convert crypto prices from USDT to the configured display currency before displaying.

#### Scenario: Successful conversion
- **WHEN** price data is fetched
- **AND** exchange rates are available for display currency
- **THEN** prices are multiplied by the exchange rate
- **AND** the currency code is attached to price data

#### Scenario: Missing currency rate
- **WHEN** the configured display currency is not in the rates
- **THEN** the system falls back to USD
- **AND** logs a warning

### Requirement: Display Currency Configuration
The system SHALL store the user's preferred display currency in configuration.

#### Scenario: Default value
- **WHEN** no display currency is configured
- **THEN** the system uses "usd" as default

#### Scenario: Currency change triggers refresh
- **WHEN** user changes display currency setting
- **THEN** the system triggers a price refresh
- **AND** prices update with new currency

### Requirement: Currency Symbol Display
The system SHALL display the appropriate currency symbol for the configured display currency.

#### Scenario: Known currency
- **WHEN** display currency is USD, EUR, GBP, or other known currency
- **THEN** the system shows the correct symbol ($, EUR, GBP)

#### Scenario: Unknown currency
- **WHEN** display currency symbol is not mapped
- **THEN** the system shows the uppercase currency code
