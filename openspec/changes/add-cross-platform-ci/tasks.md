## 1. Setup
- [x] 1.1 Create `.github/workflows/` directory
- [x] 1.2 Update `wails.json` to use pnpm

## 2. Workflow Implementation
- [x] 2.1 Create `release.yml` with trigger on `release` branch
- [x] 2.2 Add build-linux job (matrix: amd64/arm64, binary + .tar.gz + .deb + .rpm + .AppImage)
- [x] 2.3 Add build-macos job (matrix: amd64/arm64, .dmg disk images)
- [x] 2.4 Add build-windows job (NSIS installer)
- [x] 2.5 Add release job (collect artifacts, generate checksums.txt, create GitHub Release)

## 3. Verification
- [ ] 3.1 Push to test branch and verify workflow runs
- [ ] 3.2 Merge to `release` and verify GitHub Release created
