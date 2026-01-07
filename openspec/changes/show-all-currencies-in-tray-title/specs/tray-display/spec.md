## MODIFIED Requirements

### Requirement: Multi-Currency Tray Display
The system SHALL display each tracked currency as a separate menu item in the system tray dropdown.

#### Scenario: Display multiple currencies in menu
- **WHEN** user has configured BTC, ETH, and SOL
- **AND** prices are fetched successfully
- **THEN** the tray menu shows three separate price items
- **AND** each item displays the symbol and formatted price

#### Scenario: All currencies shown in tray title
- **WHEN** user has configured ETH, BTC, and USDT
- **AND** prices are fetched successfully (ETH=$3,252, BTC=$97,500, USDT=$1)
- **THEN** the system tray title shows "ETH $3,252 | BTC $97,500 | USDT $1"
- **AND** each currency price is separated by " | "

#### Scenario: Single currency in tray title
- **WHEN** user has configured only BTC
- **AND** price is fetched successfully (BTC=$97,500)
- **THEN** the system tray title shows "BTC $97,500"
- **AND** no separator is displayed

#### Scenario: Loading state shows all currencies
- **WHEN** user has configured ETH, BTC, and USDT
- **AND** prices are being fetched
- **THEN** the system tray title shows "ETH ... | BTC ... | USDT ..."

#### Scenario: Error state shows all currencies
- **WHEN** user has configured ETH, BTC, and USDT
- **AND** price fetch fails
- **THEN** the system tray title shows "ETH $??? | BTC $??? | USDT $???"

#### Scenario: Dynamic menu updates
- **WHEN** user changes currency selection in settings
- **THEN** the tray menu updates to reflect new selection
- **AND** removed currencies disappear from menu
- **AND** added currencies appear in menu
