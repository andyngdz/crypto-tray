## Context

The app displays crypto prices in the system tray. Users want to see at a glance whether prices have gone up or down since the last refresh, without needing to remember previous values.

## Goals / Non-Goals

**Goals:**
- Show visual indicator for price movement direction
- Track movement between consecutive fetches
- Work cross-platform in native menus

**Non-Goals:**
- Persist movement state across app restarts
- Show percentage change amounts
- Historical trend data

## Decisions

### Decision: Emoji Indicators
Use Unicode emoji (🟢🔴⚪) instead of colored text.
- **Why**: Emoji work cross-platform in native menus; colored text requires platform-specific code (NSAttributedString on macOS, not supported on Linux GTK)
- **Alternative**: Custom popup with HTML/CSS - rejected as too complex for simple indicators

### Decision: Separate Movement Package
Create `movement/` package rather than adding to price package.
- **Why**: Clear separation of concerns, follows existing pattern (exchange/, price/, tray/)
- **Alternative**: Embed in tray package - rejected because tracking logic is independent of display

### Decision: In-Memory Tracking Only
Store last prices in memory, reset on app restart.
- **Why**: Simple implementation, no persistence complexity, fresh start behavior is intuitive
- **Alternative**: Persist to config file - adds complexity without clear benefit

### Decision: Compare to Last Fetch
Compare current price to immediately previous fetch, not to session start or 24h ago.
- **Why**: Shows real-time movement, more responsive to market changes
- **Alternative**: Use existing Change24h field - shows different information (24h change vs recent movement)

### Decision: Indicator Constants
Define emoji as named constants in types.go.
- **Why**: Clear naming (IndicatorUp, IndicatorDown, IndicatorNeutral), easy to change symbols later
- **Alternative**: Inline strings - harder to maintain and understand

## Risks / Trade-offs

- **Emoji Rendering**: Some systems may render emoji differently → Acceptable: All modern systems support these basic emoji
- **State Loss on Restart**: Users see neutral indicators after restart → Acceptable: Intuitive "fresh start" behavior

## Migration Plan

No migration needed. New feature with no config changes - existing behavior preserved, indicators appear on next price update.

## Open Questions

None - implementation is straightforward.
