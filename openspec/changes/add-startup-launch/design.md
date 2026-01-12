# Design: Start on System Startup

## Context
Users expect desktop tray applications to start automatically when their computer boots. This is a standard feature for system tray utilities. The implementation must work across Windows, macOS, and Linux without requiring users to manually configure OS-level startup settings.

## Goals / Non-Goals

**Goals:**
- Cross-platform auto-start support (Windows, macOS, Linux)
- Simple checkbox toggle in Settings UI
- Enabled by default for new installations
- Immediate effect when toggled (no restart required)

**Non-Goals:**
- Admin/elevated privilege startup (user-level only)
- Startup delay configuration
- Startup notification/splash screen

## Decisions

### Decision: Use custom pure-Go implementation
**Rationale:** Initially planned to use `github.com/emersion/go-autostart`, but it requires CGO on Windows which added complexity. Implemented a custom pure-Go solution that:
- Uses batch file (`.bat`) on Windows in the Startup folder
- Uses `.desktop` file on Linux in XDG autostart directory
- Uses LaunchAgent plist on macOS

**Alternatives considered:**
- `github.com/emersion/go-autostart`: Requires CGO on Windows for shortcut creation
- `github.com/ProtonMail/go-autostart`: Same CGO requirement

### Decision: Auto-start enabled by default
**Rationale:** User requested Option A - silently enable auto-start on first run without prompting. This is common for tray utilities.

### Decision: Store setting in existing config file
**Rationale:** Reuse existing `config.json` persistence mechanism. The `AutoStart` boolean will be stored alongside other settings.

### Decision: Sync OS registration on toggle
**Rationale:** When user toggles the checkbox, immediately call `autostart.Enable()` or `autostart.Disable()` to update OS-level registration. The config value tracks user preference; OS registration is the actual effect.

## Platform Behavior

| Platform | Mechanism | Location |
|----------|-----------|----------|
| Windows | Startup batch file (.bat) | `%APPDATA%\Microsoft\Windows\Start Menu\Programs\Startup` |
| macOS | LaunchAgent plist | `~/Library/LaunchAgents/` |
| Linux | XDG autostart | `~/.config/autostart/` (.desktop file) |

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                     Frontend (React)                         │
├─────────────────────────────────────────────────────────────┤
│  SettingsView.tsx                                            │
│    └── SettingsSection "System"                              │
│          └── SettingsAutoStart.tsx (checkbox)                │
│                └── useSettingsAutoStart.ts (hook)            │
│                      └── updateConfig({ auto_start: bool })  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                     Backend (Go)                             │
├─────────────────────────────────────────────────────────────┤
│  app.go                                                      │
│    └── SaveConfig() detects auto_start change                │
│          └── calls autostart.SetEnabled(bool)                │
│                                                              │
│  autostart/autostart.go (new package)                        │
│    └── Platform-specific implementations                     │
│    └── SetEnabled(bool) → creates/removes startup entry      │
│    └── IsEnabled() → checks if startup entry exists          │
│                                                              │
│  init.go                                                     │
│    └── On first run: if cfg.AutoStart && !IsEnabled()       │
│          └── autostart.SetEnabled(true)                      │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                     OS Layer                                 │
├─────────────────────────────────────────────────────────────┤
│  Custom implementations per platform:                        │
│  - Windows: batch file in Startup folder                     │
│  - macOS: plist in ~/Library/LaunchAgents/                   │
│  - Linux: .desktop file in ~/.config/autostart/              │
└─────────────────────────────────────────────────────────────┘
```

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| Permission errors on some systems | Graceful error handling; checkbox remains interactive |
| Executable path changes (app moved) | Re-register on startup if path differs |
| Pure-Go implementation | No external dependencies; full control over behavior |

## Migration Plan

1. Add `auto_start` field to config (defaults to `true`)
2. Existing users: On next launch, app registers for startup if `auto_start` is true
3. New users: Same behavior, enabled by default
4. No breaking changes to existing config files (new field with default)

## Open Questions

None - all requirements clarified with user.
