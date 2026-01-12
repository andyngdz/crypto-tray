## ADDED Requirements

### Requirement: Platform-Specific Systray Icons

The system SHALL use platform-appropriate icon formats for the systray display based on the operating system.

#### Scenario: Windows uses ICO format

- **WHEN** running on Windows
- **THEN** the system embeds and uses an `.ico` file for the systray icon
- **AND** the systray icon displays correctly without errors

#### Scenario: macOS uses PNG format

- **WHEN** running on macOS
- **THEN** the system embeds and uses a `.png` file for the systray icon
- **AND** the systray icon displays correctly

#### Scenario: Linux uses PNG format

- **WHEN** running on Linux
- **THEN** the system embeds and uses a `.png` file for the systray icon
- **AND** the systray icon displays correctly

#### Scenario: No icon errors on Windows

- **WHEN** the application starts on Windows
- **THEN** no "Unable to set icon" errors appear in the logs
- **AND** the icon appears in the Windows system tray
