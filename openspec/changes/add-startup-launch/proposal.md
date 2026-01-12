# Change: Add Start on System Startup Feature

## Why
Users want the app to automatically start when their operating system boots, so they can see cryptocurrency prices in the system tray without manually launching the application each time they log in.

## What Changes
- Add a new "System" section in the Settings window
- Add a "Start on system startup" checkbox within the System section
- Implement cross-platform auto-start functionality (Windows, macOS, Linux)
- Auto-start is enabled by default for new installations
- When toggled, immediately register/unregister with the operating system's startup mechanism

## Impact
- Affected specs: `system-settings` (new capability)
- Affected code:
  - Backend: `config/types.go`, `app.go`, new `autostart/` package
  - Frontend: Settings UI components and state hooks
- New dependency: `github.com/emersion/go-autostart` (Go package for cross-platform auto-start)
