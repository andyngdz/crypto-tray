## MODIFIED Requirements

### Requirement: Platform-Aware Movement Indicators
The system SHALL display price movement indicators in tray menu items using platform-appropriate methods to ensure consistent colored display across operating systems.

#### Scenario: Windows uses icon images with transparency
- **GIVEN** the application is running on Windows
- **WHEN** a price is updated with an Up, Down, or Neutral movement
- **THEN** the system SHALL call `MenuItem.SetIcon()` with an ICO file containing a colored circle with transparent background
- **AND** the icon SHALL render with proper alpha transparency (no dark background)
- **AND** the menu item text SHALL NOT contain emoji characters

#### Scenario: Linux uses emoji text
- **GIVEN** the application is running on Linux
- **WHEN** a price is updated with an Up, Down, or Neutral movement
- **THEN** the system SHALL display emoji characters in the menu item text
- **AND** `MenuItem.SetIcon()` SHALL NOT be called (it is a no-op on Linux)

#### Scenario: macOS uses emoji text
- **GIVEN** the application is running on macOS
- **WHEN** a price is updated with an Up, Down, or Neutral movement
- **THEN** the system SHALL display emoji characters in the menu item text

### Requirement: Systray Library Dependency
The system SHALL use the `fyne.io/systray` library for cross-platform system tray functionality.

#### Scenario: Library provides transparent icon support on Windows
- **GIVEN** the application uses `fyne.io/systray`
- **WHEN** `MenuItem.SetIcon()` is called with an ICO file containing alpha transparency
- **THEN** the icon SHALL be rendered with proper transparency using 32-bit DIB bitmap
- **AND** the background SHALL NOT appear dark or black
