## MODIFIED Requirements

### Requirement: Currency Symbol Display
The system SHALL display the appropriate currency symbol for the configured display currency using dynamic lookup from golang.org/x/text/currency.

#### Scenario: Known ISO currency
- **WHEN** display currency is a valid ISO 4217 code (USD, EUR, GBP, JPY)
- **THEN** the system shows the correct Unicode symbol ($, €, £, ¥)

#### Scenario: Unknown currency
- **WHEN** display currency is not a valid ISO code
- **THEN** the system shows the uppercase currency code as fallback
