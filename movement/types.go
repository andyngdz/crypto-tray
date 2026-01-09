package movement

// Direction represents price movement direction
type Direction int

const (
	Neutral Direction = iota
	Up
	Down
)

// Indicator symbols for price movement
const (
	IndicatorUp      = "🟢" // Green circle - price increased
	IndicatorDown    = "🔴" // Red circle - price decreased
	IndicatorNeutral = "⚪" // White circle - first fetch / no change
)

// Indicator returns the emoji symbol for this direction
func (d Direction) Indicator() string {
	switch d {
	case Up:
		return IndicatorUp
	case Down:
		return IndicatorDown
	default:
		return IndicatorNeutral
	}
}
