## MODIFIED Requirements

### Requirement: Currency Selection Component
The currency selector SHALL use an autocomplete search with tags pattern for multi-selection.

#### Scenario: Display selected currencies as tags
- **WHEN** the user has selected one or more currencies
- **THEN** each selected currency is displayed as a removable tag showing the symbol ID (e.g., "BTC")

#### Scenario: Search and filter currencies
- **WHEN** the user types in the search input
- **THEN** the dropdown filters available currencies matching the input against both symbol ID and name (case-insensitive)

#### Scenario: Add currency via autocomplete
- **WHEN** the user selects a currency from the filtered dropdown
- **THEN** the currency is added to the selection and appears as a new tag
- **AND** the search input is cleared

#### Scenario: Remove currency via tag
- **WHEN** the user clicks the remove button on a tag
- **THEN** that currency is removed from the selection

#### Scenario: Exclude selected currencies from dropdown
- **WHEN** a currency is already selected
- **THEN** it SHALL NOT appear in the dropdown options
