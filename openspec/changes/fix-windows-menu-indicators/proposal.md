# Change: Fix Windows Menu Movement Indicators

## Why

On Linux, the system tray menu displays colored emoji indicators (🟢/🔴/⚪) next to cryptocurrency prices to show price movement. On Windows, these emojis render as grey/monochrome due to limited Unicode emoji support in native Win32 menus. This creates an inconsistent user experience across platforms.

## What Changes

- Generate colored circle icon images (PNG) at runtime using Go's `image` package
- Use `MenuItem.SetIcon()` on Windows to display colored indicators as actual icons
- Keep emoji-in-text approach on Linux (where it works, and `SetIcon()` is a no-op anyway)
- On macOS: use emoji if it renders properly, otherwise fall back to icon approach like Windows
- Remove emoji prefix from menu item text on platforms using icons to avoid duplicate indicators

## Impact

- Affected specs: tray (new capability)
- Affected code:
  - `tray/icons.go` (new) - Icon generation functions
  - `tray/tray.go` - Platform-specific indicator logic in `UpdatePrices()`
  - `movement/types.go` - Add icon generator calls

## Platform Behavior Matrix

| Platform | Menu Item Text | Menu Item Icon | Approach |
|----------|---------------|----------------|----------|
| Linux | `🟢 ETH $3,450` | (not supported) | Emoji in text |
| Windows | `ETH $3,450` | Green/red/grey circle | `SetIcon()` with generated PNG |
| macOS | `🟢 ETH $3,450` OR `ETH $3,450` | (optional icon) | Emoji preferred, fallback to `SetIcon()` if needed |
