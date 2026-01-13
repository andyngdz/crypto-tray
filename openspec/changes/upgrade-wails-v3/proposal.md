# Change: Upgrade from Wails v2 to Wails v3

## Why

The current implementation uses Wails v2 with a forked `fyne.io/systray` library (`internal/systray/`) that has several issues:

1. **Systray not working on Linux/macOS**: The forked `fyne.io/systray` uses `systray.Register()` which doesn't call `nativeStart()` on Linux (empty function) or macOS (returns early when `internalLoop == false`). This means the systray never initializes on these platforms.

2. **Maintenance burden**: The `internal/systray/` fork contains ~3000 lines of code that we must maintain, including platform-specific CGO code for macOS and Windows.

3. **AppDelegate conflict on macOS**: Required forking systray to rename `AppDelegate` to avoid conflicts with Wails.

Wails v3 provides **native systray support** built into the framework, eliminating all these issues.

## What Changes

### **BREAKING**: Complete framework upgrade

- **Go dependencies**: Replace `github.com/wailsapp/wails/v2` with `github.com/wailsapp/wails/v3`
- **Application initialization**: Replace `wails.Run(&options.App{})` with `application.New()` + `app.Run()`
- **Binding system**: Replace struct binding (`Bind: []interface{}{app}`) with Service pattern
- **Events API**: Replace `runtime.EventsEmit(ctx, name, data)` with `app.EmitEvent(name, data)`
- **Window API**: Replace `runtime.WindowShow/Hide(ctx)` with `window.Show()/Hide()`
- **Systray**: Replace forked `internal/systray/` with native `app.SystemTray.New()`
- **Frontend bindings**: Regenerate with new import paths (`@wailsio/runtime`, `./bindings/`)

### Files to DELETE (~3500 lines)

| Path | Lines | Reason |
|------|-------|--------|
| `internal/systray/*` | ~3000 | Native v3 systray |
| `tray/tray.go` | ~200 | Replaced by manager.go |
| `tray/types.go` | ~50 | Simplified |
| `tray/tray_windows.go` | ~30 | Replaced by icons_windows.go |
| `tray/tray_other.go` | ~20 | Replaced by icons_darwin.go/icons_linux.go |
| `tray/icon_unix.go` | ~20 | Replaced by icons_darwin.go/icons_linux.go |
| `tray/icon_windows.go` | ~20 | Replaced by icons_windows.go |
| `main_bindings.go` | ~50 | v3 uses different approach |

### Files to CREATE

| Path | Purpose |
|------|---------|
| `events/events.go` | Type-safe event definitions and registration |
| `tray/manager.go` | Core systray logic (platform-agnostic) |
| `tray/icons_darwin.go` | macOS icon embedding with SetTemplateIcon |
| `tray/icons_windows.go` | Windows icon embedding with SetIcon |
| `tray/icons_linux.go` | Linux icon embedding with SetIcon |

### Files to MODIFY

| Path | Changes |
|------|---------|
| `go.mod` | Update Wails v2 → v3 dependency |
| `main.go` | Complete rewrite (~200 lines) |
| `app.go` | Convert to Service pattern |
| `price/service.go` | Update events API |
| `exchange/service.go` | Update events API |
| `wails.json` | Update schema to v3 |
| `frontend/src/*` | Update imports for new bindings |

## Impact

- **Affected code**: All Go files with Wails imports, all frontend files using wailsjs bindings
- **Affected platforms**: Linux, macOS, Windows (all platforms)
- **Risk level**: High - complete framework upgrade
- **Rollback strategy**: Keep v2 code on `main` branch until v3 migration is verified

## Benefits

1. **Native systray support**: No external library, no forking, no CGO maintenance
2. **Cleaner codebase**: Remove ~3500 lines of forked/platform-specific code
3. **Type-safe events**: Compile-time type checking for Go→Frontend events, autocomplete in TypeScript
4. **New features**: Window attachment to systray, template icons on macOS, better click handlers
5. **Future-proof**: Aligned with Wails project roadmap
6. **Easier testing**: Service pattern has no implicit context dependencies

## References

- [Wails v3 Alpha Documentation](https://v3alpha.wails.io/)
- [v2 to v3 Migration Guide](https://v3alpha.wails.io/migration/v2-to-v3/)
- [Wails v3 SystemTray API](https://v3alpha.wails.io/features/menus/systray)
