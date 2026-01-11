## ADDED Requirements

### Requirement: macOS Systray Compatibility

The build system SHALL use a locally-vendored copy of `github.com/getlantern/systray` with a renamed `AppDelegate` class to avoid symbol conflicts with Wails on macOS.

#### Scenario: macOS build succeeds

- **WHEN** running `wails build -platform darwin/arm64`
- **THEN** the build completes without duplicate symbol errors
- **AND** the systray functionality works correctly (icon, title, menu items)

#### Scenario: Other platforms unaffected

- **WHEN** running `wails build` for Linux or Windows
- **THEN** the build uses the vendored systray without issues
- **AND** no behavioral changes occur
