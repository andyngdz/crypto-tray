# Change: Replace HeroUI with Material-UI

## Why
HeroUI's ComboBox only supports single selection - there's no `selectionMode="multiple"` option. The current workaround (ComboBox + TagGroup) works but is more complex than necessary. MUI's Autocomplete component provides search + multi-select + tags in a single component, offering a cleaner implementation and better UX for the currencies selector.

## What Changes
- Replace all HeroUI components with MUI equivalents
- Install `@mui/material`, `@emotion/react`, `@emotion/styled`
- Remove `@heroui/react`, `@heroui/styles`, `framer-motion`
- Configure MUI dark theme
- Simplify SettingsSymbols using MUI Autocomplete with `multiple` prop

## Impact
- Affected specs: None (no formal specs exist yet)
- Affected code:
  - `frontend/src/features/settings/presentations/SettingsSymbols.tsx` - Use MUI Autocomplete
  - `frontend/src/features/settings/presentations/SettingsProvider.tsx` - Use MUI Select
  - `frontend/src/features/settings/presentations/SettingsRefreshInterval.tsx` - Use MUI Select
  - `frontend/src/features/settings/presentations/SettingsApiKey.tsx` - Use MUI TextField
  - `frontend/src/features/settings/presentations/SettingsView.tsx` - Use MUI Button
  - `frontend/src/features/settings/states/useSettingsSymbols.ts` - Simplify (remove inputValue state)
  - `frontend/src/features/settings/states/useSettingsProvider.ts` - Update Key type to string
  - `frontend/src/features/settings/states/useSettingsRefreshInterval.ts` - Update Key type to string
  - `frontend/src/main.tsx` - Add MUI ThemeProvider with dark theme
  - `frontend/tailwind.config.js` - Remove HeroUI plugin
  - `frontend/package.json` - Update dependencies
