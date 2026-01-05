## 1. Dependencies

- [x] 1.1 Install MUI packages: `@mui/material`, `@emotion/react`, `@emotion/styled`
- [x] 1.2 Remove HeroUI packages: `@heroui/react`, `@heroui/styles`
- [x] 1.3 Remove `framer-motion` (was used by HeroUI)

## 2. Theme Setup

- [x] 2.1 Add MUI ThemeProvider and CssBaseline to `main.tsx`
- [x] 2.2 Configure dark theme as default

## 3. Component Migration

- [x] 3.1 Migrate `SettingsApiKey` to MUI TextField
- [x] 3.2 Migrate `SettingsProvider` to MUI Select
- [x] 3.3 Migrate `SettingsRefreshInterval` to MUI Select
- [x] 3.4 Migrate `SettingsSymbols` to MUI Autocomplete with `multiple` prop
- [x] 3.5 Migrate `SettingsView` Button to MUI Button

## 4. State Hook Updates

- [x] 4.1 Update `useSettingsSymbols` - remove inputValue, simplify for Autocomplete API
- [x] 4.2 Update `useSettingsProvider` - change Key type to string
- [x] 4.3 Update `useSettingsRefreshInterval` - change Key type to string

## 5. Config Cleanup

- [x] 5.1 Remove HeroUI plugin from `tailwind.config.js`
- [x] 5.2 Remove HeroUI paths from tailwind content array
- [x] 5.3 Remove `@heroui/styles` import from `style.css`

## 6. Verification

- [x] 6.1 Build frontend successfully
- [ ] 6.2 Test all settings components manually
