# Change: Fix Windows Systray Icon Error

## Why

The Windows systray initialization fails with error:
```
ERROR systray: systray_windows.go:845 Unable to set icon: The operation completed successfully.
```

The code currently embeds `icon.png` (line 15-16 in `tray/tray.go`), but the Windows systray library requires `.ico` format, not `.png`. This causes the application to start but display an error when setting the tray icon.

## What Changes

- Convert `tray/icon.png` to `tray/icon.ico` using ImageMagick
- Create `tray/icon_windows.go` with Windows-specific icon embedding
- Create `tray/icon_unix.go` for macOS/Linux icon embedding  
- Remove generic icon embed from `tray/tray.go`
- Use Go build tags to select correct icon format per platform

## Impact

- Affected specs: `tray-display`
- Affected code: `tray/tray.go`, new `tray/icon_windows.go`, new `tray/icon_unix.go`
- Windows systray will display icon correctly without errors
- macOS and Linux behavior unchanged
