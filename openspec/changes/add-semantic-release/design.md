# Design: Semantic Release Implementation

## Context
The project uses conventional commits (`feat:`, `fix:`, `chore:`) but has no automated versioning. The CI pipeline builds for multiple platforms but uses hardcoded versions.

## Goals / Non-Goals
- **Goals**:
  - Automate version bumping based on commit history
  - Generate release notes automatically
  - Ensure all build artifacts use consistent versions
  - Maintain control over when releases are published
  - Avoid permission issues by not committing files to repo

- **Non-Goals**:
  - Embedding version in the application UI (future work)
  - Automating the merge to release branch
  - Maintaining a CHANGELOG.md file in the repo

## Decisions

### Git Tags as Source of Truth
**Decision**: Use git tags exclusively for version tracking. No VERSION file or CHANGELOG.md committed to repo.

**Rationale**: Avoids permission issues with committing files back to the repository. Git tags are created by semantic-release using the default `GITHUB_TOKEN`, which has permission to create tags.

**Alternatives considered**:
- VERSION file + CHANGELOG.md: Rejected - requires write permissions to commit files, can fail on protected branches
- Personal Access Token: Rejected - adds maintenance burden and security considerations

### Two-Stage Release Workflow
**Decision**: Separate version tagging (main) from building (release branch).

**Rationale**: This provides control over when builds are published. Merging to main creates the version tag, but the actual release only happens when deliberately merging to the release branch.

### Minimal Plugin Set
**Decision**: Use only three plugins:
1. `@semantic-release/commit-analyzer` - Determine version bump
2. `@semantic-release/release-notes-generator` - Generate release notes
3. `@semantic-release/github` - Create git tag and GitHub Release

**Rationale**: Fewer plugins = simpler setup, fewer failure points, no permission issues.

## Workflow Diagram

```
Developer pushes to main
        |
        v
+-------------------+
| version.yml       |
| (runs on main)    |
+-------------------+
        |
        v
semantic-release analyzes commits
        |
        v
Creates tag v1.2.0
Creates GitHub Release with notes
        |
        v
Developer merges main -> release
        |
        v
+-------------------+
| release.yml       |
| (runs on release) |
+-------------------+
        |
        v
Reads version from git tag (1.2.0)
        |
        v
Builds all platforms with version
- .deb Version: 1.2.0
- .rpm Version: 1.2.0
        |
        v
Uploads assets to GitHub Release v1.2.0
```

## Risks / Trade-offs
- **Trade-off**: No CHANGELOG.md in repo
  - Mitigation: Release notes are still available in GitHub Releases

- **Risk**: Release branch may not have the latest tag if not properly synced
  - Mitigation: Use `git describe --tags --abbrev=0` to get the latest tag from the repo

## Open Questions
None - requirements are clear.
