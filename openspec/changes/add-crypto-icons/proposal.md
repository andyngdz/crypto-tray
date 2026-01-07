# Change: Add Cryptocurrency Icons to Settings Modal

## Why
The currency selector currently displays coins as text only (e.g., "BTC - Bitcoin"). Adding visual coin icons improves recognition and provides a more polished user experience, making it easier to quickly identify and select currencies.

## What Changes
- Install `cryptocurrency-icons` npm package for static SVG cryptocurrency icons (switched from `react-crypto-icons` due to webpack/vite incompatibility)
- Create a reusable `CryptoIcon` wrapper component with fallback support using Vite's glob import
- Update the currency selector autocomplete to display icons in:
  - Dropdown options (icon + symbol + name)
  - Selected value chips (icon + symbol)
- Use project logo (`assets/logo.png`) as fallback for coins without icons

## Impact
- Affected specs: `frontend-architecture`
- Affected code:
  - `frontend/package.json` - Add react-crypto-icons dependency
  - `frontend/src/assets/logo.png` - Copy fallback icon
  - `frontend/src/components/CryptoIcon.tsx` - New wrapper component
  - `frontend/src/features/settings/presentations/SettingsSymbols.tsx` - Add icon rendering
