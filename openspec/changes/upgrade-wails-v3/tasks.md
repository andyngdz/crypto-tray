# Tasks: Upgrade from Wails v2 to Wails v3

## Prerequisites

- [x] 0.1 Verify Wails v3 CLI is installed (`wails3 version`)
- [x] 0.2 Create feature branch `feat/wails-v3-migration`
- [x] 0.3 Document current v2 behavior for testing reference

## 1. Update Go Dependencies

- [x] 1.1 Update `go.mod` - replace `github.com/wailsapp/wails/v2` with `github.com/wailsapp/wails/v3`
- [x] 1.2 Remove unused dependencies (godbus/dbus if no longer needed externally)
- [x] 1.3 Run `go mod tidy`
- [x] 1.4 Verify module downloads successfully

**Commit:** `chore: update to Wails v3 dependencies`

## 2. Create Service Layer

- [x] 2.1 Create `services/` directory
- [x] 2.2 Create `services/app_service.go` with:
  - [x] ServiceName() method
  - [x] ServiceStartup(ctx, options) method
  - [x] ServiceShutdown() method
  - [x] Migrate GetConfig, SaveConfig, GetAvailableProviders, GetAvailableSymbols methods
  - [x] Migrate ShowWindow, HideWindow, QuitApp methods
  - [x] Migrate FetchPrices, RefreshPrices methods
- [x] 2.3 Migrate price service with EventEmitter interface:
  - [x] Update event emission to use `app.EmitEvent()`
- [x] 2.4 Migrate exchange service with EventEmitter interface:
  - [x] Update event emission to use `app.EmitEvent()`
- [x] 2.5 Remove context provider pattern (replaced with EventEmitter interface)

**Commit:** `refactor: convert to Wails v3 service pattern`

## 3. Migrate Main Entry Point

- [x] 3.1 Update imports to use `github.com/wailsapp/wails/v3/pkg/application`
- [x] 3.2 Replace `wails.Run()` with `application.New()`:
  - [x] Configure Name, Description
  - [x] Configure Assets with `application.AssetOptions`
  - [x] Register services with `application.NewService()`
  - [x] Configure Mac options (ActivationPolicy: Accessory)
  - [x] Configure Windows options (DisableQuitOnLastWindowClosed: true)
- [x] 3.3 Create window with `app.NewWebviewWindowWithOptions()`:
  - [x] Set Name, Title, Width, Height, MinWidth, MinHeight
  - [x] Set Frameless: true
  - [x] Set Hidden: true (equivalent to StartHidden)
  - [x] Configure Windows-specific options (HiddenOnTaskbar)
  - [x] Configure BackgroundColour
- [x] 3.4 Add `app.Run()` at the end
- [x] 3.5 Delete `main_bindings.go` (v3 uses different approach)
- [x] 3.6 Delete `app.go` (replaced by services/app_service.go)

**Commit:** `refactor: migrate main.go to Wails v3`

## 4. Implement Native Systray

- [x] 4.1 Delete `internal/systray/` directory entirely (~3000 lines removed)
- [x] 4.2 Create `tray/manager.go` (core logic, platform-agnostic):
  - [x] Define Manager struct with systray, app, window fields
  - [x] Implement Setup() using `app.NewSystemTray()`
  - [x] Call platform-specific `SetIcon(systray)` function
  - [x] Implement `SetLabel()` for price display
  - [x] Implement icon updates for movement indicators
- [x] 4.3 Create `tray/menu.go` for menu creation logic
- [x] 4.4 Create platform-specific icon files:
  - [x] `tray/icons_darwin.go` - `//go:build darwin`, embed icon.png, use SetTemplateIcon()
  - [x] `tray/icons_windows.go` - `//go:build windows`, embed icon.ico, use SetIcon()
  - [x] `tray/icons_linux.go` - `//go:build linux`, embed icon.png, use SetIcon()
- [x] 4.5 Create systray menu:
  - [x] Add price display items (dynamic)
  - [x] Add separator
  - [x] Add "Open Settings" menu item with OnClick handler
  - [x] Add "Refresh Now" menu item with OnClick handler
  - [x] Add separator
  - [x] Add "Quit" menu item with OnClick handler
- [x] 4.6 Delete old tray files:
  - [x] `tray/tray.go`
  - [x] `tray/tray_windows.go`
  - [x] `tray/tray_other.go`
  - [x] `tray/icon_unix.go`
  - [x] `tray/icon_windows.go`
- [x] 4.7 Keep icon assets:
  - [x] `tray/icon.png`
  - [x] `tray/icon.ico`
- [x] 4.8 Update `tray/types.go` for Wails v3 types

**Commit:** `refactor: replace forked systray with Wails v3 native`

## 5. Migrate to Type-Safe Events

- [x] 5.1 Create `events/events.go` for typed event definitions:
  - [x] Define event name constants
  - [x] (Note: v3 alpha doesn't have RegisterEvent[T] yet - using string events)
- [x] 5.2 Update price service event emission:
  - [x] Use `emitter.EmitEvent("price:update", data)` via EventEmitter interface
- [x] 5.3 Update exchange service event emission:
  - [x] Use `emitter.EmitEvent("exchange:update", data)` via EventEmitter interface
- [x] 5.4 Remove all `runtime.*` imports from Go code
- [x] 5.5 Remove context provider interface and implementations

**Commit:** `refactor: migrate to Wails v3 events`

## 6. Add Wails v3 Build System

- [x] 6.1 Create `Taskfile.yml` root orchestration
- [x] 6.2 Create `build/Taskfile.yml` common tasks (using pnpm)
- [x] 6.3 Create `build/config.yml` project configuration
- [x] 6.4 Create `build/linux/Taskfile.yml` Linux-specific tasks
- [x] 6.5 Create `build/darwin/Taskfile.yml` macOS-specific tasks
- [x] 6.6 Create `build/windows/Taskfile.yml` Windows-specific tasks
- [x] 6.7 Update `.gitignore` for v3 build outputs

**Commit:** `chore: add Wails v3 build system`

## 7. Migrate Frontend

- [x] 7.1 Update `wails.json` to v3 format:
  - [x] Change schema and structure to v3 format
  - [x] Configure frontend dev settings
- [x] 7.2 Add `@wailsio/runtime` package:
  - [x] `pnpm add @wailsio/runtime`
- [x] 7.3 Update `vite.config.ts`:
  - [x] Add Wails plugin `@wailsio/runtime/plugins/vite`
  - [x] Add `@bindings` alias for generated bindings
- [x] 7.4 Update `tsconfig.json`:
  - [x] Remove `@wailsjs/*` path mapping
  - [x] Add `@bindings/*` path mapping
- [x] 7.5 Update frontend service files:
  - [x] Update `configService.ts` imports (use `@bindings/services/appservice`)
  - [x] Update `priceService.ts` imports (use `@bindings/services/appservice`)
  - [x] Update `usePrices.ts` imports (use `Events` from `@wailsio/runtime`)
  - [x] Update `TitleBar.tsx` imports (use `Window` from `@wailsio/runtime`)
- [x] 7.6 Delete old `frontend/wailsjs/` directory
- [x] 7.7 Generate TypeScript bindings with `wails3 generate bindings -ts`
- [x] 7.8 Update `build/Taskfile.yml` to use `-ts` flag for bindings
- [x] 7.9 Add `build:dev` script to `package.json`
- [x] 7.10 Test build with `wails3 task linux:build`

**Commit:** `refactor: migrate frontend to Wails v3 bindings`

## 8. Testing

- [ ] 8.1 Build and test on Linux:
  - [x] App starts without errors (`wails3 task linux:build`)
  - [x] Systray icon appears
  - [x] Label displays pricing
  - [x] Frontend builds with TypeScript (no errors)
  - [ ] Menu items work
  - [ ] Window shows/hides on click
  - [ ] Settings save correctly
  - [ ] Price updates work
- [ ] 8.2 Build and test on macOS:
  - [ ] All above tests
  - [ ] Template icon respects light/dark mode
- [ ] 8.3 Build and test on Windows:
  - [ ] All above tests (except label - Windows limitation)
  - [ ] Icon displays correctly in system tray
- [ ] 8.4 Fix any platform-specific issues found

**Commit:** `fix: address platform-specific issues`

## 9. Cleanup

- [ ] 9.1 Remove any dead code
- [ ] 9.2 Update code comments
- [ ] 9.3 Run `go mod tidy`
- [ ] 9.4 Run `pnpm install` in frontend
- [ ] 9.5 Verify CI/CD pipeline passes

**Commit:** `chore: cleanup and finalize v3 migration`

## Verification Checklist

| Feature | Linux | macOS | Windows |
|---------|-------|-------|---------|
| App starts | [x] | [ ] | [ ] |
| Systray icon | [x] | [ ] | [ ] |
| Systray label | [x] | [ ] | N/A |
| Menu items | [ ] | [ ] | [ ] |
| Window toggle | [ ] | [ ] | [ ] |
| Settings save | [ ] | [ ] | [ ] |
| Price updates | [ ] | [ ] | [ ] |
| Movement icons | [ ] | [ ] | [ ] |
| Template icon (light/dark) | N/A | [ ] | N/A |
| Quit | [ ] | [ ] | [ ] |
