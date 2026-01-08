# Change: Add Number Format Setting

## Why
Users in different regions use different number formatting conventions. Currently prices are hardcoded to US format (1,234.56). This change adds a settings dropdown to let users choose their preferred number format.

## What Changes
- Add new "Formatting" section to settings UI
- Add number format dropdown with US, European, and Asian options
- Store format preference in config
- Update price formatting to use selected format

## Impact
- Affected specs: settings (new capability)
- Affected code:
  - `config/types.go` - Add NumberFormat field
  - `services/formatter.go` - Use configured format
  - `frontend/src/features/settings/` - New UI components
