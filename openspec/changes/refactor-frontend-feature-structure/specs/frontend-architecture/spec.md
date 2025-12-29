## ADDED Requirements

### Requirement: Feature-Based Directory Structure
The frontend SHALL organize code using a feature-based directory structure where each feature contains its own `presentations/`, `states/`, `services/`, `constants/`, and `types/` sub-folders.

#### Scenario: Settings feature organization
- **WHEN** viewing the settings feature directory
- **THEN** it SHALL contain:
  - `presentations/` folder with UI components (Settings.tsx, StatusDisplay.tsx)
  - `states/` folder with state management hooks (useConfig.ts)
  - `services/` folder with API abstraction (configService.ts)
  - `constants/` folder with feature constants (refreshOptions.ts)
  - `types/` folder with TypeScript definitions (index.ts)
  - `index.ts` barrel export file

### Requirement: Services Layer Abstraction
Each feature SHALL have a services layer that abstracts external API calls (Wails bindings) from state management logic.

#### Scenario: Config service provides API abstraction
- **WHEN** the settings feature needs to interact with the Go backend
- **THEN** it SHALL call functions from `configService.ts` instead of directly importing Wails bindings in hooks

#### Scenario: Service functions are testable
- **WHEN** writing tests for state management hooks
- **THEN** the service layer SHALL be mockable without modifying Wails bindings

### Requirement: Barrel Export Pattern
Each feature SHALL export its public API through a single `index.ts` file at the feature root.

#### Scenario: Clean imports from feature
- **WHEN** importing from the settings feature
- **THEN** consumers SHALL use `import { Settings } from './features/settings'` instead of deep paths

#### Scenario: Internal structure encapsulation
- **WHEN** the internal structure of a feature changes
- **THEN** consumers importing through the barrel export SHALL NOT need to update their imports
