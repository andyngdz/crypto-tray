# Change: Add Semantic Release for Automated Versioning

## Why
The project currently lacks proper version management. Package versions are hardcoded to `1.0.0`, and git tags use date-based formats (`v2025.01.10-abc1234`). This makes it difficult to track releases and communicate changes to users.

## What Changes
- Add semantic-release to automate version bumps based on conventional commits
- Create a two-stage release workflow:
  - **main branch**: semantic-release creates git tags and GitHub Release
  - **release branch**: triggers builds using the version from the git tag
- Update .deb and .rpm package versions dynamically from git tag
- Release notes generated automatically in GitHub Release (no CHANGELOG.md file)

## Impact
- Affected specs: release-management (new capability)
- Affected code:
  - `.github/workflows/release.yml` - modify to use dynamic versions from git tag
  - `.github/workflows/version.yml` - new workflow for semantic-release
  - `.releaserc.json` - new config file
  - `package.json` (root) - new file for semantic-release dependency
