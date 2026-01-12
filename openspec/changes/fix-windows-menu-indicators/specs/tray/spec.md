## ADDED Requirements

### Requirement: Platform-Aware Movement Indicators
The system SHALL display price movement indicators in tray menu items using platform-appropriate methods to ensure consistent colored display across operating systems.

#### Scenario: Windows uses icon images
- **GIVEN** the application is running on Windows
- **WHEN** a price is updated with an Up, Down, or Neutral movement
- **THEN** the system SHALL call `MenuItem.SetIcon()` with a generated PNG colored circle icon (green/red/grey)
- **AND** the menu item text SHALL NOT contain emoji characters

#### Scenario: Linux uses emoji text
- **GIVEN** the application is running on Linux
- **WHEN** a price is updated with an Up, Down, or Neutral movement
- **THEN** the system SHALL display emoji characters (🟢/🔴/⚪) in the menu item text
- **AND** `MenuItem.SetIcon()` SHALL NOT be called (it is a no-op on Linux)

#### Scenario: macOS prefers emoji text with icon fallback
- **GIVEN** the application is running on macOS
- **WHEN** a price is updated with an Up, Down, or Neutral movement
- **THEN** the system SHALL attempt to use emoji characters (🟢/🔴/⚪) in the menu item text
- **AND** if emoji rendering is detected as broken, the system SHALL fall back to `MenuItem.SetIcon()` with generated PNG icons

### Requirement: Runtime Icon Generation
The system SHALL generate colored circle icon images at runtime without requiring external image assets.

#### Scenario: Generate green circle icon
- **GIVEN** the movement direction is Up
- **WHEN** the icon generation function is called
- **THEN** it SHALL create a 16x16 pixel PNG image with a solid green filled circle
- **AND** the image SHALL be encoded as PNG bytes
- **AND** the bytes SHALL be cached for reuse

#### Scenario: Generate red circle icon
- **GIVEN** the movement direction is Down
- **WHEN** the icon generation function is called
- **THEN** it SHALL create a 16x16 pixel PNG image with a solid red filled circle
- **AND** the image SHALL be encoded as PNG bytes

#### Scenario: Generate grey circle icon
- **GIVEN** the movement direction is Neutral
- **WHEN** the icon generation function is called
- **THEN** it SHALL create a 16x16 pixel PNG image with a solid grey filled circle
- **AND** the image SHALL be encoded as PNG bytes

### Requirement: Icon Color Standards
The system SHALL use standardized colors for movement indicators to ensure visual clarity and accessibility.

#### Scenario: Up indicator color
- **GIVEN** a price increase indicator
- **WHEN** generating the icon
- **THEN** the color SHALL be green (#00C800 or RGB 0,200,0)

#### Scenario: Down indicator color
- **GIVEN** a price decrease indicator
- **WHEN** generating the icon
- **THEN** the color SHALL be red (#C80000 or RGB 200,0,0)

#### Scenario: Neutral indicator color
- **GIVEN** a neutral/no change indicator
- **WHEN** generating the icon
- **THEN** the color SHALL be grey (#C8C8C8 or RGB 200,200,200)
