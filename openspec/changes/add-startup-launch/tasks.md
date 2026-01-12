# Tasks: Add Start on System Startup

## 1. Backend Implementation

- [x] 1.1 Implement cross-platform autostart package (Note: `github.com/emersion/go-autostart` requires CGO which isn't available in this environment, so implemented pure-Go alternative)
- [x] 1.2 Add `AutoStart bool` field to `config/types.go` with default `true`
- [x] 1.3 Create `autostart/autostart.go` package with `SetEnabled()` and `IsEnabled()` functions
- [x] 1.4 Update `init.go` to enable auto-start on first run if configured
- [x] 1.5 Update `app.go` `SaveConfig()` to detect auto_start changes and sync with OS

## 2. Frontend Implementation

- [x] 2.1 Add `auto_start: boolean` to Config interface in `types/index.ts`
- [x] 2.2 Add `autoStart: true` default to `constants/defaults.ts`
- [x] 2.3 Create `useSettingsAutoStart.ts` state hook
- [x] 2.4 Create `SettingsAutoStart.tsx` checkbox component
- [x] 2.5 Add "System" section with checkbox to `SettingsView.tsx`

## 3. Testing & Validation

- [ ] 3.1 Test on Windows: verify .bat file created in Startup folder
- [ ] 3.2 Test on Linux: verify .desktop file created in autostart
- [ ] 3.3 Test toggle behavior: enable/disable updates OS registration
- [ ] 3.4 Test first-run behavior: auto-start enabled by default
- [x] 3.5 Build and verify no compilation errors across platforms
