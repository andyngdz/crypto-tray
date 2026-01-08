## Context

The app displays crypto prices in USDT (the provider's quote currency). Users in different regions need prices in their local fiat currency. We integrate with fawazahmed0/exchange-api to fetch exchange rates for currency conversion.

## Goals / Non-Goals

**Goals:**
- Fetch exchange rates from external API
- Convert prices before display (backend conversion)
- Cache rates to survive temporary API failures
- Support 300+ currencies from the exchange API

**Non-Goals:**
- UI for selecting display currency (future scope)
- Historical exchange rate data
- Multiple simultaneous display currencies

## Decisions

### Decision: Separate Exchange Fetcher
Create `exchange/fetcher.go` as a standalone component rather than extending the price fetcher.
- **Why**: Separation of concerns - exchange rates and crypto prices are different data sources
- **Alternative**: Integrate into price fetcher - rejected because it couples unrelated APIs

### Decision: Backend Conversion
Convert prices in Go before sending to frontend/tray.
- **Why**: Single source of truth, tray and frontend receive identical data
- **Alternative**: Frontend conversion - rejected because tray would need duplicate logic

### Decision: Same Refresh Interval
Use the existing `RefreshSeconds` config for exchange rate refresh.
- **Why**: Simpler config, rates and prices stay in sync
- **Alternative**: Separate interval - adds complexity without clear benefit

### Decision: In-Memory Caching
Cache last successful rates in the fetcher struct, not file-based.
- **Why**: Exchange rates change slowly, memory cache survives brief API outages
- **Alternative**: Disk persistence - overkill for rates that refresh frequently

## Risks / Trade-offs

- **API Unavailability** → Mitigation: Fallback URL + cached rates + graceful USDT fallback
- **Rate Staleness** → Acceptable: Exchange rates don't change rapidly; 15-second old data is fine

## Migration Plan

No migration needed. New `DisplayCurrency` field defaults to "usd" - existing behavior preserved.

## Open Questions

None - implementation is straightforward.
