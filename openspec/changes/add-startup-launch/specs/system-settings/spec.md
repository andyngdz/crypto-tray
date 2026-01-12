# System Settings

## ADDED Requirements

### Requirement: Auto-Start Configuration
The system SHALL provide a user-configurable option to automatically launch the application when the operating system starts.

#### Scenario: Default enabled for new installations
- **WHEN** the application is installed and run for the first time
- **THEN** auto-start SHALL be enabled by default
- **AND** the application SHALL register itself with the operating system's startup mechanism

#### Scenario: User disables auto-start
- **WHEN** the user unchecks the "Start on system startup" checkbox in Settings
- **THEN** the application SHALL immediately unregister from the operating system's startup mechanism
- **AND** the preference SHALL be persisted to the configuration file

#### Scenario: User enables auto-start
- **WHEN** the user checks the "Start on system startup" checkbox in Settings
- **THEN** the application SHALL immediately register with the operating system's startup mechanism
- **AND** the preference SHALL be persisted to the configuration file

#### Scenario: Cross-platform support
- **WHEN** the application is running on Windows
- **THEN** auto-start SHALL be managed via the Windows Startup folder shortcut
- **WHEN** the application is running on macOS
- **THEN** auto-start SHALL be managed via LaunchAgent plist
- **WHEN** the application is running on Linux
- **THEN** auto-start SHALL be managed via XDG autostart (.desktop file)

### Requirement: System Settings UI Section
The Settings window SHALL include a "System" section for operating system integration settings.

#### Scenario: System section displayed
- **WHEN** the user opens the Settings window
- **THEN** a "System" section SHALL be displayed after the "Formatting" section
- **AND** the section SHALL contain the "Start on system startup" checkbox

#### Scenario: Checkbox reflects current state
- **WHEN** the Settings window is opened
- **THEN** the "Start on system startup" checkbox SHALL reflect the current auto-start configuration value
