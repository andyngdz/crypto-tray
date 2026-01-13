# Tasks: Upgrade from Wails v2 to Wails v3

## Prerequisites

- [ ] 0.1 Verify Wails v3 CLI is installed (`wails3 version`)
- [ ] 0.2 Create feature branch `feat/wails-v3-migration`
- [ ] 0.3 Document current v2 behavior for testing reference

## 1. Update Go Dependencies

- [ ] 1.1 Update `go.mod` - replace `github.com/wailsapp/wails/v2` with `github.com/wailsapp/wails/v3`
- [ ] 1.2 Remove unused dependencies (godbus/dbus if no longer needed externally)
- [ ] 1.3 Run `go mod tidy`
- [ ] 1.4 Verify module downloads successfully

**Commit:** `chore: update to Wails v3 dependencies`

## 2. Create Service Layer

- [ ] 2.1 Create `services/` directory
- [ ] 2.2 Create `services/app_service.go` with:
  - [ ] ServiceName() method
  - [ ] ServiceStartup(ctx, options) method
  - [ ] ServiceShutdown() method
  - [ ] Migrate GetConfig, SaveConfig, GetAvailableProviders, GetAvailableSymbols methods
  - [ ] Migrate ShowWindow, HideWindow, QuitApp methods
  - [ ] Migrate FetchPrices, RefreshPrices methods
- [ ] 2.3 Create `services/price_service.go` with:
  - [ ] ServiceName() method
  - [ ] ServiceStartup(ctx, options) method
  - [ ] ServiceShutdown() method
  - [ ] Migrate price fetching logic from `price/service.go`
  - [ ] Update event emission to use `app.EmitEvent()`
- [ ] 2.4 Create `services/exchange_service.go` with:
  - [ ] ServiceName() method
  - [ ] ServiceStartup(ctx, options) method
  - [ ] ServiceShutdown() method
  - [ ] Migrate exchange rate logic from `exchange/service.go`
  - [ ] Update event emission to use `app.EmitEvent()`
- [ ] 2.5 Remove context provider pattern (no longer needed in v3)

**Commit:** `refactor: convert to Wails v3 service pattern`

## 3. Migrate Main Entry Point

- [ ] 3.1 Update imports to use `github.com/wailsapp/wails/v3/pkg/application`
- [ ] 3.2 Replace `wails.Run()` with `application.New()`:
  - [ ] Configure Name, Description
  - [ ] Configure Assets with `application.BundledAssetFileServer()`
  - [ ] Register services with `application.NewService()`
  - [ ] Configure Mac options (ActivationPolicy: Accessory)
  - [ ] Configure Windows options (DisableQuitOnLastWindowClosed: true)
- [ ] 3.3 Create window with `app.Window.NewWithOptions()`:
  - [ ] Set Name, Title, Width, Height, MinWidth, MinHeight
  - [ ] Set Frameless: true
  - [ ] Set Hidden: true (equivalent to StartHidden)
  - [ ] Configure Windows-specific options (HiddenOnTaskbar)
  - [ ] Configure BackgroundColour
- [ ] 3.4 Add `app.Run()` at the end
- [ ] 3.5 Delete `main_bindings.go` (v3 uses different approach)

**Commit:** `refactor: migrate main.go to Wails v3`

## 4. Implement Native Systray

- [ ] 4.1 Delete `internal/systray/` directory entirely
- [ ] 4.2 Create `tray/manager.go` (core logic, platform-agnostic):
  - [ ] Define Manager struct with systray, app, window fields
  - [ ] Implement Setup() using `app.SystemTray.New()`
  - [ ] Call platform-specific `SetIcon(systray)` function
  - [ ] Implement `SetLabel()` for price display
  - [ ] Implement icon updates for movement indicators
- [ ] 4.3 Create platform-specific icon files:
  - [ ] `tray/icons_darwin.go` - `//go:build darwin`, embed icon.png, use SetTemplateIcon()
  - [ ] `tray/icons_windows.go` - `//go:build windows`, embed icon.ico, use SetIcon()
  - [ ] `tray/icons_linux.go` - `//go:build linux`, embed icon.png, use SetIcon()
- [ ] 4.4 Create systray menu with `app.Menu.New()`:
  - [ ] Add price display items (dynamic)
  - [ ] Add separator
  - [ ] Add "Open Settings" menu item with OnClick handler
  - [ ] Add "Refresh Now" menu item with OnClick handler
  - [ ] Add separator
  - [ ] Add "Quit" menu item with OnClick handler
- [ ] 4.5 Configure window attachment:
  - [ ] Use `systray.AttachWindow(window)`
  - [ ] Set WindowOffset for positioning
- [ ] 4.6 Delete old tray files:
  - [ ] `tray/tray.go`
  - [ ] `tray/types.go`
  - [ ] `tray/tray_windows.go`
  - [ ] `tray/tray_other.go`
  - [ ] `tray/icon_unix.go`
  - [ ] `tray/icon_windows.go`
- [ ] 4.7 Keep icon assets:
  - [ ] `tray/icon.png`
  - [ ] `tray/icon.ico`

**Commit:** `refactor: replace forked systray with Wails v3 native`

## 5. Migrate to Type-Safe Events

- [ ] 5.1 Create `events/events.go` for typed event definitions:
  - [ ] Define `PriceUpdateData` struct matching frontend expectations
  - [ ] Define `ExchangeUpdateData` struct matching frontend expectations
  - [ ] Create `Register()` function to register all typed events
  - [ ] Call `application.RegisterEvent[T]()` for each event type
- [ ] 5.2 Update `main.go` to register events:
  - [ ] Import `events` package
  - [ ] Call `events.Register()` before `app.Run()`
- [ ] 5.3 Update price service event emission:
  - [ ] Replace `runtime.EventsEmit(ctx, "price:update", data)` with `app.EmitEvent("price:update", events.PriceUpdateData{...})`
- [ ] 5.4 Update exchange service event emission:
  - [ ] Replace `runtime.EventsEmit(ctx, "exchange:update", data)` with `app.EmitEvent("exchange:update", events.ExchangeUpdateData{...})`
- [ ] 5.5 Remove all `runtime.*` imports from Go code
- [ ] 5.6 Remove context provider interface and implementations

**Commit:** `refactor: migrate to Wails v3 type-safe events`

## 6. Migrate Frontend

- [ ] 6.1 Run `wails3 generate bindings` to regenerate bindings
  - [ ] Verify bindings are generated in `frontend/bindings/`
  - [ ] Verify typed event bindings are generated in `frontend/bindings/events/`
- [ ] 6.2 Update `vite.config.ts` with new alias:
  - [ ] Remove `@wailsjs` alias
  - [ ] Add `@bindings` alias pointing to `./bindings`
- [ ] 6.3 Install `@wailsio/runtime` package if needed:
  - [ ] `pnpm add @wailsio/runtime`
- [ ] 6.4 Update `configService.ts`:
  - [ ] Change import from `@wailsjs/go/main/App` to `@bindings/cryptotray/appservice`
- [ ] 6.5 Update `priceService.ts`:
  - [ ] Change import from `@wailsjs/go/main/App` to `@bindings/cryptotray/appservice`
- [ ] 6.6 Update `usePrices.ts` with typed events:
  - [ ] Import `On` from `@wailsio/runtime/events`
  - [ ] Import `PriceUpdate` from `@bindings/events`
  - [ ] Replace `EventsOn("price:update", handler)` with `On(PriceUpdate, handler)`
  - [ ] Update handler to use typed `event.data`
- [ ] 6.7 Update `useExchangeRates.ts` (if exists) with typed events:
  - [ ] Import `ExchangeUpdate` from `@bindings/events`
  - [ ] Replace string-based event with typed event
- [ ] 6.8 Delete old `frontend/wailsjs/` directory

**Commit:** `refactor: migrate frontend to Wails v3 bindings`

## 7. Update Configuration

- [ ] 7.1 Update `wails.json`:
  - [ ] Change schema to v3
  - [ ] Update bindings output path
- [ ] 7.2 Update `.github/workflows/release.yml`:
  - [ ] Update Wails CLI installation for v3
  - [ ] Update build commands (`wails3 build`)
- [ ] 7.3 Review `build/` directory files:
  - [ ] `build/darwin/Info.plist`
  - [ ] `build/windows/info.json`
  - [ ] `build/windows/wails.exe.manifest`

**Commit:** `chore: update configuration for Wails v3`

## 8. Testing

- [ ] 8.1 Build and test on Linux:
  - [ ] App starts without errors
  - [ ] Systray icon appears
  - [ ] Label displays pricing
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
| App starts | [ ] | [ ] | [ ] |
| Systray icon | [ ] | [ ] | [ ] |
| Systray label | [ ] | [ ] | N/A |
| Menu items | [ ] | [ ] | [ ] |
| Window toggle | [ ] | [ ] | [ ] |
| Settings save | [ ] | [ ] | [ ] |
| Price updates | [ ] | [ ] | [ ] |
| Movement icons | [ ] | [ ] | [ ] |
| Template icon (light/dark) | N/A | [ ] | N/A |
| Quit | [ ] | [ ] | [ ] |
