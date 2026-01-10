# Window Branding

## ADDED Requirements

### Requirement: Settings window displays correct branding

The settings window MUST display the application name "CryptoTray" as the window title and MUST show the application icon in the window titlebar and taskbar.

#### Scenario: User opens settings window

**Given** the application is running
**When** the user clicks "Open Settings" from the system tray
**Then** the window title displays "CryptoTray"
**And** the window icon shows the purple/pink gradient logo in the taskbar
