# Change: Switch from internal systray fork to fyne-io/systray

## Why

The current internal fork of `getlantern/systray` has a bug in Windows icon handling where `iconToBitmap()` uses `CreateCompatibleBitmap` which doesn't support alpha transparency. This causes menu item icons to display with a dark/black background instead of being transparent.

The `fyne-io/systray` library (also a fork of `getlantern/systray`) has already fixed this issue by using `CreateDIBSection` with 32-bit color depth, which properly supports alpha transparency.

## What Changes

- Remove the `internal/systray` folder (internal fork)
- Add `fyne.io/systray` as a Go module dependency
- Update all imports from `github.com/getlantern/systray` to `fyne.io/systray`
- Regenerate ICO files with transparent backgrounds (no longer need background color workaround)

## Impact

- Affected specs: tray
- Affected code:
  - `go.mod` - Add fyne.io/systray dependency
  - `tray/tray.go` - Update import
  - `tray/tray_windows.go` - Update import
  - `tray/tray_other.go` - Update import
  - `tray/types.go` - Update import
  - `tray/*.ico` - Regenerate with transparent backgrounds
- Removed:
  - `internal/systray/` - Entire folder removed

## Benefits

1. **Fixes transparency issue** - Icons will have proper transparent backgrounds on Windows
2. **No internal fork to maintain** - Uses external well-maintained library
3. **Actively maintained** - Fyne team maintains this fork
4. **Better Linux support** - Uses DBus/StatusNotifierItem (modern approach)

## Risks

- **Linux compatibility**: `fyne-io/systray` removes GTK dependency and uses DBus only. May not work on very old desktop environments (but this is rare).
- **API changes**: Both libraries have the same API (same origin), so minimal risk.

## Platform Behavior After Change

| Platform | Menu Item Icons | Expected Result |
|----------|-----------------|-----------------|
| Windows | SetIcon() with transparent ICO | Colored circles with transparent background |
| Linux | Emoji in text | Same as before (unchanged) |
| macOS | Emoji in text | Same as before (unchanged) |
