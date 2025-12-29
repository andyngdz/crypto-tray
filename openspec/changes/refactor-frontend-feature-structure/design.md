## Context
The crypto-tray frontend currently uses a flat folder structure common in small React projects. As the application grows, this structure becomes harder to maintain. Moving to a feature-based structure aligns with React best practices for scalable applications.

**Current Structure:**
```
src/
├── components/     # All UI components
├── hooks/          # All custom hooks
├── types/          # All TypeScript types
├── constants/      # All constants
├── App.tsx
└── main.tsx
```

**Proposed Structure:**
```
src/
├── features/
│   └── settings/
│       ├── presentations/   # UI components
│       ├── states/          # Hooks and state management
│       ├── services/        # API calls (Wails bindings)
│       ├── constants/       # Feature-specific constants
│       ├── types/           # Feature-specific types
│       └── index.ts         # Barrel export
├── shared/                  # Cross-feature components (future)
├── App.tsx
└── main.tsx
```

## Goals / Non-Goals

**Goals:**
- Organize code by feature for better maintainability
- Separate concerns: UI (presentations), logic (states), API (services)
- Enable parallel feature development without merge conflicts
- Create clear boundaries between features

**Non-Goals:**
- Adding new functionality
- Changing existing behavior
- Introducing state management libraries (Redux, Zustand)
- Creating a shared component library (can be done later)

## Decisions

### Decision 1: Folder Naming Convention
Use `presentations/`, `states/`, `services/`, `constants/`, `types/` as sub-folder names.

**Alternatives considered:**
- `components/`, `hooks/` - More common in React, but doesn't separate concerns as clearly
- `ui/`, `logic/`, `api/` - Shorter but less descriptive

**Rationale:** The chosen names explicitly describe the purpose of each folder, making it easier for new developers to understand the structure.

### Decision 2: Services Layer for API Calls
Extract Wails binding calls into a `services/configService.ts` file.

**Rationale:**
- Separates API concerns from state management
- Makes it easier to mock API calls in tests
- Provides a single source of truth for API interactions

### Decision 3: Barrel Exports
Each feature exports public API through `index.ts`.

**Rationale:**
- Clean import paths: `import { Settings } from './features/settings'`
- Encapsulates internal structure
- Makes refactoring easier (internal changes don't affect importers)

## Risks / Trade-offs

| Risk | Mitigation |
|------|------------|
| Deeper import paths during development | Barrel exports simplify public API |
| Over-engineering for small app | Structure supports growth without overhead |
| Learning curve for contributors | Clear naming conventions reduce confusion |

## Migration Plan

1. Create new directory structure
2. Create services layer (new file)
3. Move and update each file one at a time
4. Update App.tsx imports
5. Delete old empty directories
6. Verify build succeeds

**Rollback:** Git revert if issues discovered during testing.

## Open Questions
None - structure is well-defined and follows established patterns.
