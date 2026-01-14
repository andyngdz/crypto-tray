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

crypto-tray is a desktop application built with Tauri v2 (Rust) + React/TypeScript (Vite).

## Development Commands

```bash
# Install dependencies
pnpm install

# Run frontend dev server (for quick UI iteration)
pnpm dev

# Run full Tauri app in dev mode
pnpm tauri dev

# Build production app bundles
pnpm tauri build
```

## Architecture

**Frontend (React/TypeScript)**

- `src/main.tsx` - React entry point
- `src/App.tsx` - Main application component

**Backend (Rust / Tauri)**

- `src-tauri/src/main.rs` - Tauri app entry point
- `src-tauri/tauri.conf.json` - Tauri configuration

**IPC Bridge**

- Frontend calls Rust commands via `@tauri-apps/api` (e.g. `invoke`)
- Rust exposes commands via `#[tauri::command]`

## Build Output

- `dist/` - Frontend build output (Vite)
- `src-tauri/target/` - Rust build artifacts (Cargo)
- `src-tauri/target/release/bundle/` - Packaged application bundles

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

**Frontend**

- Modular design - separate presentation components from logic (hooks)
- No logic in presentation components (`presentations/`) - put all logic (formatting, calculations, transformations) in hooks (`states/`)
- Feature-based patterns - organize code by feature, not by type
- Use `!!value` instead of `value !== undefined` for filtering nullish values
