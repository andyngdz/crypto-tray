## ADDED Requirements

### Requirement: Number Format Setting
The system SHALL provide a user setting to select number formatting style for price display.

#### Scenario: Display format options
- **WHEN** user opens the settings window
- **THEN** a "Formatting" section is displayed with a "Number Format" dropdown
- **AND** dropdown contains options: US (1,234.56), European (1.234,56), Asian (1,234.56)

#### Scenario: Default format
- **WHEN** no format has been configured
- **THEN** US format is used as default

#### Scenario: Format persistence
- **WHEN** user selects a number format
- **AND** closes the settings window
- **AND** reopens the settings window
- **THEN** the previously selected format is displayed

#### Scenario: Price display uses selected format
- **WHEN** user has selected European format
- **THEN** prices in system tray display as "1.234,56" style
