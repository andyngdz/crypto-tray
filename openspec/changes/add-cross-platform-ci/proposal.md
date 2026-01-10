# Change: Add cross-platform CI pipeline

## Why
The project needs automated builds for distribution. Currently builds are manual.

## What Changes
- Add GitHub Actions workflow triggered on `release` branch
- Build for Linux (amd64), macOS (amd64 + arm64), Windows (amd64)
- Create Windows NSIS installer
- Auto-create GitHub Releases with all artifacts
- Update wails.json to use pnpm

## Impact
- Affected specs: New `ci-pipeline` capability
- Affected code: `.github/workflows/release.yml`, `wails.json`
