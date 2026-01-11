# Change: Fix macOS Build - Duplicate AppDelegate Symbol

## Why

The macOS build fails with linker errors:
```
duplicate symbol '_OBJC_METACLASS_$_AppDelegate'
duplicate symbol '_OBJC_CLASS_$_AppDelegate'
```

Both `github.com/wailsapp/wails/v2` and `github.com/getlantern/systray` define an Objective-C `AppDelegate` class. When compiled together for macOS, the linker cannot resolve the duplicate symbols.

This is a known issue documented in:
- [Wails Issue #3003](https://github.com/wailsapp/wails/issues/3003)
- [Wails Discussion #4514](https://github.com/wailsapp/wails/discussions/4514)

## What Changes

- Copy `github.com/getlantern/systray@v1.2.2` to `internal/systray/`
- Rename `AppDelegate` → `SysTrayDelegate` in `systray_darwin.m` (4 occurrences)
- Add `replace` directive in `go.mod` to use local copy

## Impact

- Affected code: `go.mod`, new `internal/systray/` directory
- No behavioral changes to the application
- macOS builds will succeed
- Linux and Windows builds unaffected
