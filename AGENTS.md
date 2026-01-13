<!-- OPENSPEC:START -->
# OpenSpec Instructions

These instructions are for AI assistants working in this project.

Always open `@/openspec/AGENTS.md` when the request:
- Mentions planning or proposals (words like proposal, spec, change, plan)
- Introduces new capabilities, breaking changes, architecture shifts, or big performance/security work
- Sounds ambiguous and you need the authoritative spec before coding

Use `@/openspec/AGENTS.md` to learn:
- How to create and apply change proposals
- Spec format and conventions
- Project structure and guidelines

Keep this managed block so 'openspec update' can refresh the instructions.

<!-- OPENSPEC:END -->

# AGENTS.md

This file provides guidance to AI assistants when working with code in this repository.

## Project Overview

crypto-tray is a desktop application built with Wails v2 - a Go + React/TypeScript framework for cross-platform desktop apps. Currently based on the Wails React-TypeScript template.

## Development Commands

```bash
# Run in development mode with hot reload
wails dev

# Build production binary
wails build

# Frontend-only commands (from frontend/ directory)
npm install     # Install dependencies
npm run dev     # Vite dev server
npm run build   # Build frontend assets
```

### Ubuntu

On Ubuntu, use the `-tags webkit2_41` flag for webkit2gtk 4.1 compatibility:

```bash
wails dev -tags webkit2_41
wails build -tags webkit2_41
```

## Architecture

**Backend (Go)**

- `main.go` - Wails app initialization, window configuration, frontend asset embedding
- `app.go` - App struct with lifecycle hooks and methods exposed to frontend

**Frontend (React/TypeScript)**

- `frontend/src/main.tsx` - React entry point
- `frontend/src/App.tsx` - Main application component
- `frontend/wailsjs/` - Auto-generated Wails bindings (do not edit manually)

**IPC Bridge**

- Go methods on the `App` struct with public (capitalized) names are automatically exposed to frontend
- Wails generates TypeScript bindings in `frontend/wailsjs/go/main/App.js`
- Frontend imports and calls Go methods directly: `import { Greet } from '../wailsjs/go/main/App'`

## Build Output

- `build/bin/` - Compiled binaries
- `build/darwin/` - macOS build configuration (Info.plist)
- `build/windows/` - Windows build configuration (manifests, installer, icons)

## Code Style

**General**

- DRY (Don't Repeat Yourself)
- Type-safe without `any`
- One function, one responsibility
- Each file should be a space for a specific service/concern
- Each file should not exceed 150 lines
- Modular design - separate concerns into distinct packages/modules
- Use blank lines to separate logical blocks (imports, variable declarations, control flow, return statements)
- Add blank lines between code blocks for readability
- Follow existing patterns in the codebase - check how similar code is written before adding new code
- Only add defensive code (nil checks, empty checks, fallbacks) when necessary - trace the data flow to verify the check is actually needed

**Go**

- Avoid `for _, item := range` on slices/arrays - use `for idx := range` and access `arr[idx]` instead (exception: maps don't have meaningful indices, so `for _, value := range` is acceptable)
- Use meaningful index names instead of `i` (e.g., `symbolIdx`, `slotIdx`, `coinIdx`)

**Frontend**

- Modular design - separate presentation components from logic (hooks)
- No logic in presentation components (`presentations/`) - put all logic (formatting, calculations, transformations) in hooks (`states/`)
- Feature-based patterns - organize code by feature, not by type
- Use `!!value` instead of `value !== undefined` for filtering nullish values
