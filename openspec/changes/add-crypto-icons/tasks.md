## 1. Setup

- [x] 1.1 Install `cryptocurrency-icons` package in frontend (switched from `react-crypto-icons` due to webpack/vite incompatibility)
- [x] 1.2 Copy `assets/logo.png` to `frontend/src/assets/logo.png` for fallback

## 2. Implementation

- [x] 2.1 Create `CryptoIcon` wrapper component with fallback handling
- [x] 2.2 Update `SettingsSymbols` to render icons in dropdown options using `renderOption`
- [x] 2.3 Update `SettingsSymbols` to render icons in selected chips using `renderTags`

## 3. Verification

- [x] 3.1 Test icons display for common coins (BTC, ETH, USDT, etc.)
- [x] 3.2 Test fallback icon displays for coins without icons
- [x] 3.3 Verify visual alignment in both dropdown and chips
