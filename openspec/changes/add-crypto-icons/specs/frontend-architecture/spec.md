## MODIFIED Requirements

### Requirement: Currency Selection Component
The currency selector SHALL use an autocomplete search with tags pattern for multi-selection, displaying cryptocurrency icons alongside text.

#### Scenario: Display selected currencies as tags
- **WHEN** the user has selected one or more currencies
- **THEN** each selected currency is displayed as a removable tag showing the coin icon and symbol (e.g., icon + "BTC")

#### Scenario: Search and filter currencies
- **WHEN** the user types in the search input
- **THEN** the dropdown filters available currencies matching the input against both symbol ID and name (case-insensitive)

#### Scenario: Add currency via autocomplete
- **WHEN** the user selects a currency from the filtered dropdown
- **THEN** the currency is added to the selection and appears as a new tag with icon
- **AND** the search input is cleared

#### Scenario: Remove currency via tag
- **WHEN** the user clicks the remove button on a tag
- **THEN** that currency is removed from the selection

#### Scenario: Exclude selected currencies from dropdown
- **WHEN** a currency is already selected
- **THEN** it SHALL NOT appear in the dropdown options

#### Scenario: Display icons in dropdown options
- **WHEN** the dropdown is open showing available currencies
- **THEN** each option SHALL display the coin icon, symbol, and name (e.g., icon + "BTC - Bitcoin")

#### Scenario: Fallback icon for unsupported coins
- **WHEN** a cryptocurrency does not have an icon in the icon package
- **THEN** the system SHALL display the project logo as a fallback icon
