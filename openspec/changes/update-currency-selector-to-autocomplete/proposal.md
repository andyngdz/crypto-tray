# Change: Update Currency Selector to Autocomplete with Tags

## Why
The current multi-select dropdown displays currencies as "BTC, Bitcoin" and "ETH, Ethereum" in a list format. An autocomplete search with tags provides a better UX: users can search/filter currencies quickly and see selected items as removable tags, making it easier to manage multiple selections.

## What Changes
- Replace `Select` component with `ComboBox` + `TagGroup` combination
- Add searchable input that filters available currencies by ID and name
- Display selected currencies as removable tags above the search input
- Exclude already-selected currencies from the dropdown options

## Impact
- Affected specs: `frontend-architecture`
- Affected code:
  - `frontend/src/features/settings/presentations/SettingsSymbols.tsx` - Replace component implementation
  - `frontend/src/features/settings/states/useSettingsSymbols.ts` - Add/remove handlers for tags
