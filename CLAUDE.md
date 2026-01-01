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

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

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

**Frontend**

- Modular design - separate presentation components from logic (hooks)
- Feature-based patterns - organize code by feature, not by type
- Use `!!value` instead of `value !== undefined` for filtering nullish values
