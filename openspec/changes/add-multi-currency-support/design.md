# Design: Multi-Currency Support

## Context
The application tracks cryptocurrency prices and displays them in the system tray. Currently limited to one currency. Users want to watch multiple currencies simultaneously.

## Goals / Non-Goals
- **Goals**:
  - Support tracking multiple cryptocurrencies
  - Display each currency separately in tray menu for future styling (red/green highlighting)
  - Efficient batch fetching when provider supports it
  - No hard limit on number of currencies
- **Non-Goals**:
  - Price change highlighting (future feature)
  - Reordering currencies
  - Per-currency refresh intervals

## Decisions

### 1. Config Structure
**Decision**: Change `Symbol string` to `Symbols []string`

**Rationale**: Simple array is sufficient. No need for complex objects per symbol since all currencies share the same provider and refresh interval.

**Migration**: On load, if old `symbol` field exists, migrate to `symbols: [symbol]`.

### 2. Provider Interface - Batch Fetch with Fallback
**Decision**: Add optional batch method with automatic fallback

```go
type Provider interface {
    // Existing single-fetch (required)
    FetchPrice(ctx context.Context, symbol string) (*PriceData, error)
    
    // New batch fetch (optional - has default implementation)
    FetchPrices(ctx context.Context, symbols []string) ([]*PriceData, error)
    SupportsBatchFetch() bool
    
    // Existing
    ID() string
    Name() string
    RequiresAPIKey() bool
    SetAPIKey(key string)
}
```

**Rationale**: 
- CoinGecko supports batch (multiple IDs in one API call)
- Future providers may not support batch
- Default implementation calls `FetchPrice` in a loop
- Providers override `FetchPrices` and `SupportsBatchFetch` for efficiency

### 3. Tray Display
**Decision**: Separate menu item per currency, first currency in tray title

```
Tray Title: "BTC $97,000"
─────────────────────────
Menu:
  BTC  $97,000
  ETH  $3,400
  SOL  $180
  ─────────────
  Open Settings
  Refresh Now
  ─────────────
  Quit
```

**Rationale**:
- Each currency as separate `*systray.MenuItem` allows individual styling
- Enables future red/green highlighting based on price changes
- Tray title has limited space, show only primary (first) currency

### 4. Available Symbols
**Decision**: Hardcode supported symbols in provider, expose via `GetAvailableSymbols()` binding

**Rationale**:
- CoinGecko has finite list of supported coins with ID mapping
- Dynamic fetching adds complexity and API calls
- Start with popular coins: BTC, ETH, SOL, ADA, DOT, LINK, AVAX, MATIC, ATOM, XRP

### 5. UI Component
**Decision**: Use HeroUI `Select` with `selectionMode="multiple"`

**Rationale**:
- Consistent with existing provider dropdown
- Built-in multi-select support
- Familiar UX pattern

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| Too many currencies slows refresh | Batch fetch minimizes API calls; CoinGecko rate limits apply regardless |
| Tray menu gets too long | No hard limit, but UI naturally discourages >10 selections |
| Config migration breaks existing users | Graceful migration: detect old format, convert automatically |

## Migration Plan

1. On config load, check for old `symbol` string field
2. If exists, convert to `symbols: [symbol]`
3. Save migrated config
4. Old config files remain compatible

## Open Questions
- None - design decisions finalized based on discussion
