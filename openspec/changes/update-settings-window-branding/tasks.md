# Tasks: Update Settings Window Branding

## Implementation Tasks

- [x] Update window title in `main.go` from "Crypto Tray Settings" to "CryptoTray"
- [x] Add embed directive for `frontend/src/assets/images/logo.png` in `main.go`
- [x] Add `Linux: &linux.Options{Icon: windowIcon}` to Wails options in `main.go`
- [x] Update window title in `main_bindings.go` from "Crypto Tray Settings" to "CryptoTray"
- [x] Add embed directive for `frontend/src/assets/images/logo.png` in `main_bindings.go`
- [x] Add `Linux: &linux.Options{Icon: windowIcon}` to Wails options in `main_bindings.go`

## Verification

- [x] Run `wails dev -tags webkit2_41` and open settings window
- [x] Verify window title displays "CryptoTray"
- [x] Verify window icon appears in taskbar with purple/pink gradient logo (deferred: GNOME Wayland issue tracked separately)
