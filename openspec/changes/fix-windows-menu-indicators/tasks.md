## 1. Implementation
- [x] 1.1 Create ICO files using ImageMagick (up.ico, down.ico, neutral.ico)
- [x] 1.2 Create `tray/icons_windows.go` with embedded ICO files
- [x] 1.3 Create `tray/icons_other.go` stub for non-Windows platforms
- [x] 1.4 Create `IconForDirection()` helper function to map movement direction to icon bytes
- [x] 1.5 Update `tray/tray.go` `UpdatePrices()` with platform-specific logic
- [ ] 1.6 Test on Windows to verify icons display correctly
- [ ] 1.7 Test on Linux to verify emoji still works
- [ ] 1.8 Test on macOS to verify emoji behavior
