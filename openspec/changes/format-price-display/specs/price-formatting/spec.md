# Price Formatting Specification

## Overview
Defines how cryptocurrency prices are formatted for display in the system tray.

## ADDED Requirements

### Requirement: Thousand separator formatting
Prices SHALL be displayed with comma thousand separators for improved readability.

#### Scenario: Price in thousands
- Given a price of 97000
- When displayed in the system tray
- Then it shows as "$97,000"

#### Scenario: Price in tens of thousands
- Given a price of 123456
- When displayed in the system tray
- Then it shows as "$123,456"

#### Scenario: Price under one thousand
- Given a price of 500
- When displayed in the system tray
- Then it shows as "$500"

### Requirement: Whole number display
Prices SHALL be displayed as whole numbers without decimal places.

#### Scenario: Price with decimals
- Given a price of 97123.45
- When displayed in the system tray
- Then it shows as "$97,123"
