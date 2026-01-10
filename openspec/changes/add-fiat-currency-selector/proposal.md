# Change: Add Fiat Currency Selector Dropdown

## Why

Users cannot select their preferred fiat currency from the settings UI. While the backend already supports `display_currency` configuration, the frontend is missing the dropdown to select it.

## What Changes

- Add `display_currency` field to TypeScript Config interface
- Create a fiat currency options constant with top world currencies:
  - Major: USD, EUR, GBP, JPY, CHF, CAD, AUD, NZD, CNY, HKD, SGD
  - European: SEK, NOK, DKK
  - Asian: KRW, TWD, INR, THB, VND
  - Americas: BRL, MXN
  - Other: ZAR
- Create settings hook and presentation component for fiat currency selection
- Position dropdown below the "Currencies" (crypto symbols) dropdown in settings

## Impact

- Affected specs: settings-ui (new capability)
- Affected code:
  - `frontend/src/features/settings/types/index.ts` - Add display_currency field
  - `frontend/src/features/settings/constants/defaults.ts` - Add displayCurrency default
  - `frontend/src/features/settings/presentations/SettingsView.tsx` - Add component
- New files:
  - `frontend/src/features/settings/constants/fiatCurrencyOptions.ts`
  - `frontend/src/features/settings/states/useSettingsFiatCurrency.ts`
  - `frontend/src/features/settings/presentations/SettingsFiatCurrency.tsx`
