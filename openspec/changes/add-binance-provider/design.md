## Context

The application has a provider abstraction (`Provider` interface) that allows different cryptocurrency data sources. Currently only CoinGecko is implemented. Adding Binance requires implementing the same interface but with Binance-specific API integration.

Key constraint: Binance uses trading pairs (BTCUSDT) while CoinGecko uses coin IDs (bitcoin). These are incompatible, so switching providers requires resetting the user's symbol selection.

## Goals / Non-Goals

**Goals:**
- Add Binance as a selectable provider in settings
- Use USDT trading pairs for price data
- No API key required (public endpoints)
- Reset to BTC default when switching providers

**Non-Goals:**
- Symbol mapping between providers (too complex, edge cases)
- Support for multiple quote assets (only USDT)
- API key support for higher rate limits

## Decisions

### Decision 1: Use USDT as quote asset
Binance uses trading pairs. USDT is the most common and liquid quote asset.
- Alternatives: BTC, BUSD, USDC
- Rationale: USDT has the most pairs and highest liquidity

### Decision 2: Add DefaultCoinID() to Provider interface
Each provider needs to specify its default coin (BTC) in its own format.
- CoinGecko: "bitcoin"
- Binance: "BTCUSDT"
- Rationale: Clean abstraction, each provider owns its defaults

### Decision 3: Use symbol as name
Binance's exchangeInfo doesn't return full coin names. Use the symbol (baseAsset) directly as the name.
- Alternatives: Hardcoded map for common coins, or fetch from authenticated API
- Rationale: Simplest solution, no maintenance burden, consistent behavior

### Decision 4: Clear symbols on provider switch
When user changes provider, reset symbols to provider's BTC default.
- Alternatives: Map symbols via ticker (BTC exists on both)
- Rationale: Simpler, avoids edge cases where coins don't exist on both providers

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| Binance API unavailable in some regions | Document limitation; users can use CoinGecko |
| Coins show symbol instead of full name | Acceptable UX; symbol is clear and unambiguous |
| User loses selection on provider switch | Pre-select BTC so list isn't empty |

## Binance API Details

| Endpoint | Purpose | Rate Weight |
|----------|---------|-------------|
| `/api/v3/exchangeInfo?permissions=SPOT` | Get trading symbols | 20 |
| `/api/v3/ticker/24hr?symbols=[...]` | Get prices + 24h change | 2-80 |

Base URL: `https://api.binance.com`
Rate Limit: 1200 weight/minute (well within our usage)

## Open Questions

None - all decisions made with user input.
