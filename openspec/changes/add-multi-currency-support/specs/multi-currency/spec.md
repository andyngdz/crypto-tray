## ADDED Requirements

### Requirement: Multi-Currency Configuration
The system SHALL allow users to configure multiple cryptocurrencies to track simultaneously.

#### Scenario: User selects multiple currencies
- **WHEN** user opens settings
- **AND** selects BTC, ETH, and SOL from the currency dropdown
- **THEN** the configuration saves all three symbols
- **AND** the system tracks all selected currencies

#### Scenario: At least one currency required
- **WHEN** user attempts to deselect all currencies
- **THEN** the system prevents saving with zero currencies
- **AND** displays validation feedback

#### Scenario: Config migration from single symbol
- **WHEN** application loads with old config format containing `symbol: "BTC"`
- **THEN** the system migrates to new format `symbols: ["BTC"]`
- **AND** preserves the user's previous selection

### Requirement: Batch Price Fetching
The system SHALL fetch prices for all configured currencies efficiently using batch requests when supported by the provider.

#### Scenario: Provider supports batch fetch
- **WHEN** price refresh is triggered
- **AND** provider supports batch fetching
- **THEN** the system fetches all currency prices in a single API call

#### Scenario: Provider does not support batch fetch
- **WHEN** price refresh is triggered
- **AND** provider does not support batch fetching
- **THEN** the system fetches each currency price individually
- **AND** aggregates results before updating display

#### Scenario: Partial fetch failure
- **WHEN** price refresh is triggered
- **AND** some currencies fail to fetch
- **THEN** the system displays successful prices
- **AND** shows error state for failed currencies

### Requirement: Multi-Currency Tray Display
The system SHALL display each tracked currency as a separate menu item in the system tray dropdown.

#### Scenario: Display multiple currencies in menu
- **WHEN** user has configured BTC, ETH, and SOL
- **AND** prices are fetched successfully
- **THEN** the tray menu shows three separate price items
- **AND** each item displays the symbol and formatted price

#### Scenario: Primary currency in tray title
- **WHEN** user has configured multiple currencies
- **THEN** the system tray title shows the first/primary currency
- **AND** all currencies are visible in the dropdown menu

#### Scenario: Dynamic menu updates
- **WHEN** user changes currency selection in settings
- **THEN** the tray menu updates to reflect new selection
- **AND** removed currencies disappear from menu
- **AND** added currencies appear in menu

### Requirement: Available Symbols List
The system SHALL provide a list of supported cryptocurrency symbols for user selection.

#### Scenario: Fetch available symbols
- **WHEN** user opens settings
- **THEN** the system displays available symbols from the current provider
- **AND** shows symbol ID and display name for each

#### Scenario: Provider-specific symbols
- **WHEN** user switches providers
- **THEN** the available symbols list updates for the new provider
- **AND** previously selected symbols remain if supported by new provider
