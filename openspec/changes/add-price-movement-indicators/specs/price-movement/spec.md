## ADDED Requirements

### Requirement: Price Movement Tracking
The system SHALL track price changes between consecutive fetches to determine movement direction.

#### Scenario: First price fetch
- **WHEN** the app starts and fetches prices for the first time
- **THEN** all prices have Neutral direction
- **AND** prices are stored as "last known" for future comparison

#### Scenario: Price increased
- **WHEN** a price fetch returns a higher price than the last fetch
- **THEN** the movement direction is Up
- **AND** the new price becomes the "last known" price

#### Scenario: Price decreased
- **WHEN** a price fetch returns a lower price than the last fetch
- **THEN** the movement direction is Down
- **AND** the new price becomes the "last known" price

#### Scenario: Price unchanged
- **WHEN** a price fetch returns the same price as the last fetch
- **THEN** the movement direction remains unchanged from previous
- **AND** the price is still stored as "last known"

#### Scenario: App restart
- **WHEN** the app is restarted
- **THEN** all tracked prices are cleared
- **AND** the next fetch is treated as first fetch (Neutral)

### Requirement: Movement Indicator Display
The system SHALL display emoji indicators in the tray menu to show price movement direction.

#### Scenario: Display up indicator
- **WHEN** a price has Up movement direction
- **THEN** the tray menu shows "🟢" before the price
- **AND** the display format is "🟢 BTC $97,000"

#### Scenario: Display down indicator
- **WHEN** a price has Down movement direction
- **THEN** the tray menu shows "🔴" before the price
- **AND** the display format is "🔴 BTC $97,000"

#### Scenario: Display neutral indicator
- **WHEN** a price has Neutral movement direction
- **THEN** the tray menu shows "⚪" before the price
- **AND** the display format is "⚪ BTC $97,000"

### Requirement: Indicator Constants
The system SHALL define movement indicators as named constants for maintainability.

#### Scenario: Constant definitions
- **GIVEN** the movement package
- **THEN** IndicatorUp is defined as "🟢"
- **AND** IndicatorDown is defined as "🔴"
- **AND** IndicatorNeutral is defined as "⚪"

### Requirement: Thread Safety
The system SHALL ensure thread-safe access to movement tracking data.

#### Scenario: Concurrent access
- **WHEN** multiple goroutines access the movement tracker
- **THEN** data races are prevented via mutex synchronization
- **AND** no corrupted state occurs
