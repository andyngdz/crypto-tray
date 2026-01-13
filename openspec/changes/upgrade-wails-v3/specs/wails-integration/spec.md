# Wails Integration Specification Delta

## ADDED Requirements

### Requirement: Wails v3 Framework

The application SHALL use Wails v3 framework for cross-platform desktop functionality.

#### Scenario: Application initialization
- **WHEN** the application starts
- **THEN** it SHALL initialize using `application.New()` with configured options
- **AND** register all services via `application.NewService()`
- **AND** create the main window via `app.Window.NewWithOptions()`
- **AND** call `app.Run()` to start the event loop

#### Scenario: Service registration
- **WHEN** the application initializes
- **THEN** it SHALL register AppService for configuration and window management
- **AND** register PriceService for cryptocurrency price fetching
- **AND** register ExchangeService for currency exchange rates

---

### Requirement: Native System Tray

The application SHALL use Wails v3 native system tray instead of external libraries.

#### Scenario: Systray creation
- **WHEN** the application starts
- **THEN** it SHALL create a system tray icon using `app.SystemTray.New()`
- **AND** the icon SHALL be visible in the system notification area

#### Scenario: Systray label display
- **WHEN** price data is available
- **THEN** the systray SHALL display cryptocurrency prices in the label (Linux/macOS)
- **AND** the label SHALL update when prices change

#### Scenario: Platform-specific icon handling
- **WHEN** running on macOS
- **THEN** the systray SHALL use `SetTemplateIcon()` for light/dark mode support
- **WHEN** running on Windows or Linux
- **THEN** the systray SHALL use `SetIcon()` and optionally `SetDarkModeIcon()`

#### Scenario: Systray menu
- **WHEN** user right-clicks the systray icon (or clicks on some platforms)
- **THEN** a menu SHALL appear with options:
  - Price display items
  - "Open Settings" - shows the settings window
  - "Refresh Now" - triggers immediate price refresh
  - "Quit" - exits the application

#### Scenario: Window attachment
- **WHEN** the systray is configured
- **THEN** the settings window SHALL be attached via `systray.AttachWindow()`
- **AND** clicking the systray icon SHALL toggle window visibility
- **AND** the window SHALL hide when it loses focus

---

### Requirement: Wails v3 Service Pattern

All bound Go structs SHALL implement the Wails v3 Service interface.

#### Scenario: Service interface implementation
- **WHEN** a service is created
- **THEN** it SHALL implement `ServiceName() string`
- **AND** it SHALL implement `ServiceStartup(ctx, options) error`
- **AND** it SHALL implement `ServiceShutdown() error`

#### Scenario: Service independence
- **WHEN** a service method is called
- **THEN** it SHALL NOT require a stored Wails context
- **AND** application access SHALL be through injected `*application.App` reference

---

### Requirement: Type-Safe Events

The application SHALL use Wails v3 type-safe event API for Go-to-frontend communication.

#### Scenario: Event type registration
- **WHEN** the application initializes
- **THEN** it SHALL register typed events using `application.RegisterEvent[T]()`
- **AND** define event data structures in a dedicated `events` package

#### Scenario: Event emission from Go
- **WHEN** price data is updated
- **THEN** the service SHALL emit typed events using `app.EmitEvent("price:update", events.PriceUpdateData{...})`
- **AND** NOT use `runtime.EventsEmit(ctx, ...)`
- **AND** NOT use untyped data structures

#### Scenario: Event subscription in frontend
- **WHEN** the frontend needs to receive events
- **THEN** it SHALL import typed event from `@bindings/events` (e.g., `PriceUpdate`)
- **AND** use `On(PriceUpdate, handler)` from `@wailsio/runtime/events`
- **AND** handler SHALL receive typed `event.data` with full TypeScript autocomplete
- **AND** NOT use string-based `EventsOn()` from old runtime

---

## REMOVED Requirements

### Requirement: External Systray Library

**Reason:** Wails v3 provides native systray support, eliminating the need for external libraries.

**Migration:** Delete `internal/systray/` directory and use `app.SystemTray.New()`.

---

### Requirement: Context Provider Pattern

**Reason:** Wails v3 services don't require stored context for runtime operations.

**Migration:** Remove `ContextProvider` interface and `GetContext()` methods. Inject `*application.App` directly where needed.

---

### Requirement: Wails v2 Binding Pattern

**Reason:** Wails v3 uses service-based bindings instead of struct binding with `Bind: []interface{}{}`.

**Migration:** Convert `App` struct to services implementing `ServiceName()`, `ServiceStartup()`, `ServiceShutdown()`.
