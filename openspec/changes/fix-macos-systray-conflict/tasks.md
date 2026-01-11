# Tasks

## 1. Vendor systray locally

- [x] 1.1 Create `internal/systray/` directory
- [x] 1.2 Copy systray v1.2.2 source files from Go module cache
- [x] 1.3 Set write permissions on copied files

## 2. Rename AppDelegate class

- [x] 2.1 Rename `@interface AppDelegate` → `@interface SysTrayDelegate` (line 53)
- [x] 2.2 Rename `@implementation AppDelegate` → `@implementation SysTrayDelegate` (line 59)
- [x] 2.3 Rename `AppDelegate *delegate` → `SysTrayDelegate *delegate` (line 214)
- [x] 2.4 Rename `(AppDelegate*)[NSApp delegate]` → `(SysTrayDelegate*)[NSApp delegate]` (line 232)

## 3. Update Go module

- [x] 3.1 Add replace directive: `replace github.com/getlantern/systray => ./internal/systray`
- [x] 3.2 Run `go mod tidy`

## 4. Verify

- [ ] 4.1 Run `wails build -platform darwin/arm64` (via CI or locally)
- [ ] 4.2 Confirm no duplicate symbol errors
- [ ] 4.3 Test systray functionality (icon, title, menu items work correctly)
