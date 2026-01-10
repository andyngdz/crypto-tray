## ADDED Requirements

### Requirement: Fiat Currency Selector
The settings UI SHALL provide a dropdown to select the display currency for crypto prices.

#### Scenario: Display dropdown in settings
- **WHEN** user opens the settings window
- **THEN** a "Display Currency" dropdown appears below the "Currencies" dropdown
- **AND** the dropdown shows the currently configured currency

#### Scenario: Select currency
- **WHEN** user selects a different currency from the dropdown
- **THEN** the configuration is updated with the new currency
- **AND** prices in the tray update to show the selected currency

#### Scenario: Default currency
- **WHEN** no currency has been configured
- **THEN** the dropdown shows USD as the default selection

### Requirement: Supported Fiat Currencies
The system SHALL support a predefined list of common fiat currencies for selection.

#### Scenario: Available currencies
- **WHEN** user opens the currency dropdown
- **THEN** the following currencies are available: USD, EUR, GBP, JPY, CHF, CAD, AUD, NZD, CNY, HKD, SGD, SEK, NOK, DKK, KRW, TWD, INR, THB, VND, BRL, MXN, ZAR
- **AND** each currency displays with its symbol (e.g., "USD ($)")
