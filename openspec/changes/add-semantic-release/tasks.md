# Tasks: Add Semantic Release

## 1. Setup
- [x] 1.1 Create `.releaserc.json` with minimal plugin configuration
- [x] 1.2 Create root `package.json` with semantic-release dependency

## 2. Workflows
- [x] 2.1 Create `.github/workflows/version.yml` for semantic-release on main branch
- [x] 2.2 Update `.github/workflows/release.yml` to read version from git tag
- [x] 2.3 Update .deb package build to use dynamic version from tag
- [x] 2.4 Update .rpm package build to use dynamic version from tag

## 3. Validation
- [ ] 3.1 Test version.yml workflow with a feat: commit on main
- [ ] 3.2 Verify git tag is created with correct semver format (vX.Y.Z)
- [ ] 3.3 Verify GitHub Release is created with release notes
- [ ] 3.4 Test release.yml by merging to release branch
- [ ] 3.5 Verify .deb and .rpm packages have correct version
