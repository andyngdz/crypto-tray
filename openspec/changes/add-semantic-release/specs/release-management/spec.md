# Release Management

## ADDED Requirements

### Requirement: Semantic Versioning
The system SHALL use semantic versioning (MAJOR.MINOR.PATCH) for all releases based on conventional commit analysis.

#### Scenario: Feature commit triggers minor version bump
- **WHEN** a commit with prefix `feat:` is pushed to main
- **THEN** the minor version SHALL increment (e.g., 1.0.0 → 1.1.0)

#### Scenario: Fix commit triggers patch version bump
- **WHEN** a commit with prefix `fix:` is pushed to main
- **THEN** the patch version SHALL increment (e.g., 1.0.0 → 1.0.1)

#### Scenario: Breaking change triggers major version bump
- **WHEN** a commit contains `BREAKING CHANGE:` or uses `!` suffix (e.g., `feat!:`)
- **THEN** the major version SHALL increment (e.g., 1.0.0 → 2.0.0)

### Requirement: Git Tag Based Versioning
The system SHALL use git tags as the source of truth for version numbers.

#### Scenario: Version tag is created on release
- **WHEN** semantic-release determines a new version
- **THEN** a git tag SHALL be created in the format `vX.Y.Z`

#### Scenario: Build reads version from tag
- **WHEN** the release workflow runs
- **THEN** it SHALL read the version from the latest git tag

### Requirement: GitHub Release Notes
The system SHALL automatically generate release notes in GitHub Releases.

#### Scenario: Release notes are generated
- **WHEN** a new version is released
- **THEN** a GitHub Release SHALL be created with notes generated from commit messages

### Requirement: Two-Stage Release Process
The system SHALL use a two-stage release process separating versioning from building.

#### Scenario: Version tagging on main branch
- **WHEN** commits are pushed to the main branch
- **THEN** semantic-release SHALL analyze commits and create version tags
- **AND** no build artifacts SHALL be created

#### Scenario: Build triggered on release branch
- **WHEN** changes are merged to the release branch
- **THEN** the build workflow SHALL execute
- **AND** the version from the git tag SHALL be used for all packages

### Requirement: Consistent Package Versions
All build artifacts SHALL use the version from the git tag.

#### Scenario: Debian package uses correct version
- **WHEN** a .deb package is built
- **THEN** the package control file SHALL contain `Version: <TAG_VERSION>`

#### Scenario: RPM package uses correct version
- **WHEN** an .rpm package is built
- **THEN** the spec file SHALL contain `Version: <TAG_VERSION>`
