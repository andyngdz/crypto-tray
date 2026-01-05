## ADDED Requirements

### Requirement: MUI Theme Configuration
The application SHALL use Material-UI as the UI component library with a dark theme by default.

#### Scenario: Dark theme applied
- **WHEN** the application loads
- **THEN** MUI components render with dark theme styling

### Requirement: Multi-Select Currency Autocomplete
The currency selector SHALL use MUI Autocomplete with `multiple` prop to allow searching and selecting multiple currencies in a single component.

#### Scenario: Search and select currencies
- **WHEN** user types in the currency search field
- **THEN** matching currencies are filtered and displayed
- **AND** user can select multiple currencies that appear as chips/tags

#### Scenario: Remove selected currency
- **WHEN** user clicks the remove icon on a currency chip
- **THEN** the currency is removed from the selection

### Requirement: Settings Form Components
All settings form components SHALL use MUI equivalents for consistent styling.

#### Scenario: Provider selection uses MUI Select
- **WHEN** user opens the provider dropdown
- **THEN** it renders as a MUI Select with MenuItem options

#### Scenario: Refresh interval uses MUI Select
- **WHEN** user opens the refresh interval dropdown
- **THEN** it renders as a MUI Select with MenuItem options

#### Scenario: API key input uses MUI TextField
- **WHEN** user focuses the API key field
- **THEN** it renders as a MUI TextField with password type
