# Change: Refactor Frontend to Feature-Based Structure

## Why
The current frontend uses a flat structure (`components/`, `hooks/`, `types/`, `constants/`) which doesn't scale well as features grow. A feature-based structure improves code organization, maintainability, and makes it easier to add new features without cross-contamination.

## What Changes
- Restructure frontend from flat folders to feature-based folders
- Introduce `features/settings/` with sub-folders: `presentations/`, `states/`, `services/`, `constants/`, `types/`
- Extract Wails API calls into a dedicated services layer
- Create barrel exports for clean imports
- **No functional changes** - this is a pure restructuring

## Impact
- Affected specs: `frontend-architecture` (new capability)
- Affected code:
  - `frontend/src/components/Settings.tsx` → `frontend/src/features/settings/presentations/Settings.tsx`
  - `frontend/src/components/StatusDisplay.tsx` → `frontend/src/features/settings/presentations/StatusDisplay.tsx`
  - `frontend/src/hooks/useConfig.ts` → `frontend/src/features/settings/states/useConfig.ts`
  - `frontend/src/types/config.ts` → `frontend/src/features/settings/types/index.ts`
  - `frontend/src/constants/settings.ts` → `frontend/src/features/settings/constants/refreshOptions.ts`
  - `frontend/src/App.tsx` - Update import paths
  - New file: `frontend/src/features/settings/services/configService.ts`
  - New file: `frontend/src/features/settings/index.ts` (barrel export)
