# Proposal: Format Price Display

## Summary
Add thousand-separator formatting to cryptocurrency price displays in the system tray. Currently prices show as `$97000`, this change formats them as `$97,000`.

## Motivation
- Improves readability of large numbers
- Standard money formatting convention
- Better UX for price monitoring

## Scope
- **In scope**: Format price in system tray title and menu item
- **Out of scope**: Decimal places, currency symbol customization, locale-aware formatting

## Approach
Add a `formatPrice` helper function using Go's `golang.org/x/text/message` package for locale-aware thousand separators. Format as whole numbers (no decimal places).

## Files Affected
- `tray/tray.go` - Add formatting helper and update `UpdatePrice` function

## Risks
- None significant - purely cosmetic change
