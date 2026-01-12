# Tasks

## 1. Convert PNG to ICO format

- [x] 1.1 Run ImageMagick conversion: `magick tray/icon.png tray/icon.ico`
- [x] 1.2 Verify `tray/icon.ico` was created successfully

## 2. Create platform-specific icon files

- [x] 2.1 Create `tray/icon_windows.go` with Windows icon embed
- [x] 2.2 Create `tray/icon_unix.go` for macOS/Linux icon embed
- [x] 2.3 Ensure both files define `var iconData []byte`

## 3. Update main tray file

- [x] 3.1 Remove `//go:embed icon.png` from `tray/tray.go`
- [x] 3.2 Remove `var iconData []byte` from `tray/tray.go`

## 4. Verify

- [x] 4.1 Run `wails dev` on Windows
- [x] 4.2 Confirm no "Unable to set icon" error appears
- [x] 4.3 Verify systray icon displays correctly in Windows system tray
- [x] 4.4 Test macOS/Linux builds to ensure icons still work
