## ADDED Requirements

### Requirement: Release Pipeline
The system SHALL provide automated builds on merge to release branch.

#### Scenario: Trigger on release branch
- **WHEN** code is pushed to `release` branch
- **THEN** the CI pipeline starts building for all platforms

#### Scenario: Build artifacts
- **WHEN** the pipeline completes successfully
- **THEN** the following artifacts are produced:
  - Linux amd64: binary, .tar.gz, .deb, .rpm, .AppImage
  - Linux arm64: binary, .tar.gz, .deb, .rpm, .AppImage
  - macOS amd64: .dmg (Intel Mac)
  - macOS arm64: .dmg (Apple Silicon)
  - Windows amd64: .exe, installer .exe
  - checksums.txt (SHA256 for all artifacts)

### Requirement: GitHub Release
The system SHALL create a GitHub Release with all build artifacts.

#### Scenario: Auto-versioning
- **WHEN** a release is created
- **THEN** it uses date-based versioning (v{YYYY.MM.DD}-{short-sha})

#### Scenario: Artifact download
- **WHEN** users visit the GitHub Release page
- **THEN** they can download binaries for their platform
