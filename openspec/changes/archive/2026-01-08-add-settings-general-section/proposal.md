# Change: Add "General" Section to Settings UI

## Why
The settings modal currently displays all settings (API Provider, API Key, Currencies, Refresh Interval) as a flat list without any visual grouping. Adding section containers will improve organization and prepare the UI for future settings additions.

## What Changes
- Create a reusable `SettingsSection` component for grouping related settings with a title header
- Wrap API Provider, API Key (conditional), Currencies, and Refresh Interval in a "General" section

## Impact
- Affected specs: `settings-ui`
- Affected code:
  - `frontend/src/features/settings/presentations/SettingsSection.tsx` - New section wrapper component
  - `frontend/src/features/settings/presentations/SettingsView.tsx` - Wrap settings in General section
