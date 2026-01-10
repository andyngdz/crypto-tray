# Proposal: Update Settings Window Branding

## Summary

Update the settings window to display the correct application name ("CryptoTray") and show the app icon in the window titlebar/taskbar.

## Problem

Currently, when opening the settings window:
- The window title shows "Crypto Tray Settings" instead of "CryptoTray"
- No window icon is displayed, causing the Linux taskbar to show a generic icon or the binary name "crypto-tray-dev-linux-amd64"

## Solution

1. Update the window title from "Crypto Tray Settings" to "CryptoTray" in both `main.go` and `main_bindings.go`
2. Embed `frontend/src/assets/images/logo.png` and pass it to Wails via the `Icon` option

## Scope

- **In scope**: Window title and icon for the settings window
- **Out of scope**: System tray icon (already configured), app binary name, build icons

## Files Affected

| File | Change |
|------|--------|
| `main.go` | Update title, embed and set window icon |
| `main_bindings.go` | Update title, embed and set window icon |

## Risks

None - this is a cosmetic change with no functional impact.
