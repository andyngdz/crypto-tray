## 1. Type Updates
- [x] 1.1 Add `display_currency` field to Config interface in `types/index.ts`
- [x] 1.2 Update `isConfig` type guard to check for `display_currency`
- [x] 1.3 Add `displayCurrency` default to `constants/defaults.ts`

## 2. Constants
- [x] 2.1 Create `fiatCurrencyOptions.ts` with top world currencies (USD, EUR, GBP, JPY, CHF, CAD, AUD, NZD, CNY, HKD, SGD, SEK, NOK, DKK, KRW, TWD, INR, THB, VND, BRL, MXN, ZAR)

## 3. State Hook
- [x] 3.1 Create `useSettingsFiatCurrency.ts` hook following `useSettingsNumberFormat.ts` pattern

## 4. UI Component
- [x] 4.1 Create `SettingsFiatCurrency.tsx` presentation component
- [x] 4.2 Add `SettingsFiatCurrency` to `SettingsView.tsx` below `SettingsSymbols`

## 5. Verification
- [ ] 5.1 Run `wails dev -tags webkit2_41` and verify dropdown appears
- [ ] 5.2 Test selecting different currencies and verify tray prices update
