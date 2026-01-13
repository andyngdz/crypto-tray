# Design: Upgrade from Wails v2 to Wails v3

## Context

CryptoTray is a cross-platform desktop application built with Wails v2 that displays cryptocurrency prices in the system tray. The current implementation has issues with the systray not appearing on Linux and macOS due to the forked `fyne.io/systray` library not properly initializing when used with `systray.Register()`.

Wails v3 provides native systray support, eliminating the need for external libraries and their associated integration issues.

### Stakeholders
- End users (Linux, macOS, Windows)
- Maintainers (reduced code complexity)

### Constraints
- Wails v3 is currently in alpha (may have breaking changes)
- Must maintain feature parity with current v2 implementation
- Must work on all three platforms

## Goals / Non-Goals

### Goals
- Migrate to Wails v3 framework
- Replace forked systray with native v3 systray
- Maintain all existing features
- Reduce codebase complexity (~3500 lines removed)
- Fix systray issues on Linux and macOS

### Non-Goals
- Adding new features during migration (scope creep)
- Supporting Wails v2 and v3 simultaneously
- Changing the frontend UI design
- Modifying business logic (price fetching, exchange rates)

## Decisions

### Decision 1: Use Service Pattern for Bindings

**What:** Convert the current `App` struct to Wails v3 Services.

**Why:** 
- v3 requires services with `ServiceName()`, `ServiceStartup()`, `ServiceShutdown()` methods
- Services don't require storing context (cleaner code)
- Services can be tested independently without Wails runtime

**v2 Pattern (current):**
```go
type App struct {
    ctx context.Context  // Must store context
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx  // Required for all runtime calls
}

func (a *App) SomeMethod() {
    runtime.EventsEmit(a.ctx, "event", data)  // Need context
}
```

**v3 Pattern (new):**
```go
type AppService struct {
    app           *application.App  // Injected after creation
    configManager *config.Manager
    priceService  *PriceService
}

func (s *AppService) ServiceName() string { 
    return "AppService" 
}

func (s *AppService) ServiceStartup(ctx context.Context, opts application.ServiceOptions) error {
    // Initialize resources, start background tasks
    return s.priceService.Start(ctx)
}

func (s *AppService) ServiceShutdown() error {
    // Cleanup resources, stop services
    return s.priceService.Stop()
}

func (s *AppService) SomeMethod() {
    s.app.EmitEvent("event", data)  // No context threading needed
}
```

**Alternatives considered:**
- Single monolithic service: Rejected - harder to maintain
- Keep v2 pattern with adapter: Rejected - adds complexity without benefit

---

### Decision 2: Native Systray with Window Attachment

**What:** Use `app.SystemTray.New()` with `AttachWindow()` for the settings popup.

**Why:**
- Native integration with Wails event loop
- No external dependencies
- Built-in window attachment feature
- Platform-specific icon handling (template icons for macOS)

**Implementation:**
```go
systray := app.SystemTray.New()

// Platform-specific icon setup (via build tags in icons_*.go files)
SetIcon(systray)  // Calls platform-specific function

// Label for pricing (Linux/macOS only - Windows shows tooltip)
systray.SetLabel("BTC $45,000")

// Attach window - shows on click, hides on focus loss
systray.AttachWindow(window).WindowOffset(5)
```

**Alternatives considered:**
- Keep external systray library: Rejected - source of current bugs
- Custom window positioning: Rejected - v3 has built-in `AttachWindow()`

---

### Decision 3: Simplified Tray Manager with Platform-Specific Icons

**What:** Replace complex tray manager (~500 lines, 6 files) with simplified version using build tags for platform-specific icon handling.

**Why:**
- v3 native systray handles most platform differences internally
- Separate icon files allow compile-time platform selection
- Template icons on macOS require different handling than Windows/Linux
- Menu creation is simpler with `app.Menu.New()`

**Current structure (to delete):**
```
tray/
├── tray.go           # Main logic
├── types.go          # Types and constants
├── tray_windows.go   # Windows-specific
├── tray_other.go     # Unix-specific
├── icon_unix.go      # Unix icon embedding
├── icon_windows.go   # Windows icon embedding
├── icon.png          # Keep
└── icon.ico          # Keep
```

**New structure:**
```
tray/
├── manager.go        # Core systray logic (platform-agnostic)
├── icons_darwin.go   # //go:build darwin - SetTemplateIcon with icon.png
├── icons_windows.go  # //go:build windows - SetIcon with icon.ico
├── icons_linux.go    # //go:build linux - SetIcon with icon.png
├── icon.png          # Keep (used by darwin, linux)
└── icon.ico          # Keep (used by windows)
```

**Platform-specific icon handling:**

`icons_darwin.go`:
```go
//go:build darwin

package tray

import _ "embed"

//go:embed icon.png
var iconData []byte

func SetIcon(systray *application.SystemTray) {
    systray.SetTemplateIcon(iconData)  // macOS auto-adapts to light/dark
}
```

`icons_windows.go`:
```go
//go:build windows

package tray

import _ "embed"

//go:embed icon.ico
var iconData []byte

func SetIcon(systray *application.SystemTray) {
    systray.SetIcon(iconData)
}
```

`icons_linux.go`:
```go
//go:build linux

package tray

import _ "embed"

//go:embed icon.png
var iconData []byte

func SetIcon(systray *application.SystemTray) {
    systray.SetIcon(iconData)
}
```

---

### Decision 4: Type-Safe Event Migration

**What:** Replace string-based context events with type-safe app-level events.

**Why:**
- v2 events are pure strings with `any` data - no compile-time type checking
- v3 supports `application.RegisterEvent[T]()` for typed events
- Generated TypeScript types provide autocomplete and type safety in frontend
- Reduces runtime bugs from mismatched event data structures

**v2 (string-based, no type safety):**
```go
// Go - any data, no type checking
runtime.EventsEmit(ctx, "price:update", data)

// JavaScript - data is 'any'
import { EventsOn, EventsOff } from '@wailsjs/runtime/runtime'
EventsOn("price:update", (data) => {
    // No autocomplete, no type checking
})
```

**v3 (type-safe events):**
```go
// Go - Define event types in events/events.go
package events

import "github.com/wailsapp/wails/v3/pkg/application"

// Event data structures
type PriceUpdateData struct {
    Prices []PriceData `json:"prices"`
}

type ExchangeUpdateData struct {
    Rates map[string]float64 `json:"rates"`
}

// Register typed events (called during init)
func Register() {
    application.RegisterEvent[PriceUpdateData]("price:update")
    application.RegisterEvent[ExchangeUpdateData]("exchange:update")
}

// Go - Emit typed event
app.EmitEvent("price:update", events.PriceUpdateData{Prices: prices})
```

```typescript
// Frontend - Import generated types (using absolute path alias)
import { On } from '@wailsio/runtime/events'
import { PriceUpdate } from '@bindings/events'

On(PriceUpdate, (event) => {
    // event.data is typed as PriceUpdateData
    // Full autocomplete: event.data.Prices[0].Symbol
    console.log(event.data.Prices)
})
```

**Events to migrate:**

| Event Name | Data Type | Direction |
|------------|-----------|-----------|
| `price:update` | `PriceUpdateData` | Go → Frontend |
| `exchange:update` | `ExchangeUpdateData` | Go → Frontend |

---

### Decision 5: Binding Import Path Strategy

**What:** Update frontend imports to new v3 binding structure with absolute path alias.

**Why:**
- Consistent with existing `@` alias pattern in the project
- Avoids brittle relative paths (`./bindings`, `../bindings`)
- Easier refactoring when moving files

**v2 paths:**
```javascript
import { GetConfig } from '@wailsjs/go/main/App'
import { EventsOn } from '@wailsjs/runtime/runtime'
```

**v3 paths (with `@bindings` alias):**
```javascript
import { GetConfig } from '@bindings/cryptotray/appservice'
import { PriceUpdate } from '@bindings/events'
import { On } from '@wailsio/runtime/events'
```

**Vite config update required:**
```typescript
// vite.config.ts
resolve: {
  alias: {
    '@': path.resolve(__dirname, './src'),
    '@bindings': path.resolve(__dirname, './bindings'),
  },
},
```

**Note:** Exact service/event paths will be determined after running `wails3 generate bindings`.

## Risks / Trade-offs

| Risk | Impact | Probability | Mitigation |
|------|--------|-------------|------------|
| v3 alpha instability | High | Medium | Pin to specific version, keep v2 branch for rollback |
| Breaking API changes | Medium | Medium | Monitor Wails releases, test thoroughly |
| Platform-specific bugs | Medium | Low | Test on all platforms before merge |
| Frontend binding issues | Low | Low | Regenerate bindings, test incrementally |
| Build/CI failures | Medium | Low | Update workflows incrementally |

### Trade-offs Accepted

1. **Alpha framework risk** vs **Fixing systray bugs**: We accept the alpha risk because the current v2 systray is fundamentally broken on Linux/macOS.

2. **Migration effort** vs **Maintenance burden**: One-time migration effort is preferable to maintaining 3500 lines of forked systray code.

3. **Learning curve** vs **Modern patterns**: v3's service pattern is cleaner and worth the initial learning.

## Migration Plan

### Phase 1: Preparation
1. Create feature branch
2. Install Wails v3 CLI
3. Document current behavior

### Phase 2: Backend Migration (Go)
1. Update dependencies
2. Create services
3. Migrate main.go
4. Implement native systray
5. Update events

### Phase 3: Frontend Migration
1. Generate new bindings
2. Update imports
3. Update event handlers
4. Remove old wailsjs directory

### Phase 4: Configuration & Testing
1. Update wails.json
2. Update CI/CD
3. Test all platforms
4. Fix issues

### Rollback Plan

If critical issues are found:
1. Keep `main` branch on v2 until v3 is verified
2. Feature branch can be abandoned if needed
3. v2 systray fix (`RunWithExternalLoop`) can be applied as temporary solution

## Open Questions

1. **Wails v3 version**: Should we pin to a specific alpha version or use latest?
   - Recommendation: Pin to specific version for stability

2. **Frontend runtime package**: Is `@wailsio/runtime` a separate npm package or bundled?
   - Need to verify during implementation

3. **Binding output path**: What is the exact path structure after `wails3 generate bindings`?
   - Will be determined during implementation

4. **CI/CD Wails v3 installation**: How to install Wails v3 CLI in GitHub Actions?
   - May need `go install github.com/wailsapp/wails/v3/cmd/wails3@latest`

## References

- [Wails v3 Alpha Docs](https://v3alpha.wails.io/)
- [v2 to v3 Migration Guide](https://v3alpha.wails.io/migration/v2-to-v3/)
- [Wails v3 SystemTray](https://v3alpha.wails.io/features/menus/systray)
- [Wails v3 Services](https://v3alpha.wails.io/features/bindings/services)
- [Wails v3 Events](https://v3alpha.wails.io/features/events)
