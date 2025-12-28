# Tasks: Format Price Display

## Implementation Tasks

- [x] Add `golang.org/x/text` dependency
- [x] Create `formatPrice(price float64) string` helper in `tray/tray.go`
- [x] Update `UpdatePrice` to use `formatPrice`
- [x] Update placeholder text in `onReady` to match format (`$--,---`)
- [x] Test with various price values (hundreds, thousands, tens of thousands)
