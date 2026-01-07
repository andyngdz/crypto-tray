## MODIFIED Requirements

### Requirement: Multi-Currency Tray Display
The system SHALL display each tracked currency using its ticker symbol (not coinID) in all tray title states.

#### Scenario: Loading state shows ticker symbols
- **WHEN** user has configured ethereum, bitcoin, and tether (coinIDs)
- **AND** prices are being fetched
- **THEN** the system tray title shows "ETH ... | BTC ... | USDT ..."
- **AND** NOT "ethereum ... | bitcoin ... | tether ..."

#### Scenario: Error state shows ticker symbols
- **WHEN** user has configured ethereum, bitcoin, and tether (coinIDs)
- **AND** price fetch fails
- **THEN** the system tray title shows "ETH $??? | BTC $??? | USDT $???"
- **AND** NOT "ethereum $??? | bitcoin $??? | tether $???"

#### Scenario: Placeholder state shows ticker symbols
- **WHEN** user has configured ethereum, bitcoin, and tether (coinIDs)
- **AND** prices have not been fetched yet
- **THEN** the system tray title shows "ETH $--,--- | BTC $--,--- | USDT $--,---"
- **AND** NOT "ethereum $--,--- | bitcoin $--,--- | tether $--,---"

#### Scenario: Fallback for unknown coinIDs
- **WHEN** a coinID has no known ticker mapping
- **THEN** the system displays the uppercase coinID as fallback
- **AND** example: "solana" becomes "SOLANA" until price data provides "SOL"
