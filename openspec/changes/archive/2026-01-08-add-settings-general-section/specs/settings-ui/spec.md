## ADDED Requirements

### Requirement: Settings Sections
The settings modal SHALL organize settings into visual sections with titled headers.

#### Scenario: General section displays
- **WHEN** the settings modal is opened
- **THEN** a "General" section header is visible
- **AND** the API Provider, Currencies, and Refresh Interval settings are grouped under it

#### Scenario: Conditional settings within section
- **WHEN** the selected provider requires an API key
- **THEN** the API Key field is displayed within the General section
