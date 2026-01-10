# Tasks: Add Number Format Setting

## 1. Backend Configuration
- [x] 1.1 Add `DefaultNumberFormat = "us"` constant to `config/types.go`
- [x] 1.2 Add `NumberFormat string` field to Config struct
- [x] 1.3 Add `NumberFormat: DefaultNumberFormat` to `defaultConfig()`

## 2. Price Formatter
- [x] 2.1 Update `FormatPrice` in `services/formatter.go` to accept format parameter
- [x] 2.2 Map format values to `language.Tag` (us→English, european→German)

## 3. Tray Manager
- [x] 3.1 Add `numberFormat` field to Manager struct in `tray/types.go`
- [x] 3.2 Update `tray.New` to accept numberFormat parameter
- [x] 3.3 Add `SetNumberFormat` method to Manager
- [x] 3.4 Update `UpdatePrices` to pass format to `FormatPrice`

## 4. App Callbacks
- [x] 4.1 Add `onNumberFormatChanged` callback field to App struct
- [x] 4.2 Add `setOnNumberFormatChanged` setter method
- [x] 4.3 Call callback in `SaveConfig` when format changes

## 5. Main Wiring
- [x] 5.1 Pass `cfg.NumberFormat` to `tray.New` in `main.go`
- [x] 5.2 Wire up number format change callback to tray

## 6. Frontend Types
- [x] 6.1 Add `number_format: string` to Config interface
- [x] 6.2 Update `isConfig` type guard

## 7. Frontend UI
- [x] 7.1 Create `numberFormatOptions.ts` constants
- [x] 7.2 Create `useSettingsNumberFormat.ts` hook
- [x] 7.3 Create `SettingsNumberFormat.tsx` component
- [x] 7.4 Add "Formatting" section to `SettingsView.tsx`

## 8. Verification
- [x] 8.1 Run `wails dev -tags webkit2_41` and verify dropdown appears
- [x] 8.2 Select European format and verify tray updates (e.g., `$97.000`)
- [x] 8.3 Test format persistence (change, close, reopen)
