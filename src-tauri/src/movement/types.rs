/// Direction represents price movement direction
#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Direction {
    Neutral,
    Up,
    Down,
}

/// Indicator symbols for price movement
pub const INDICATOR_UP: &str = "🟢"; // Green circle - price increased
pub const INDICATOR_DOWN: &str = "🔴"; // Red circle - price decreased
pub const INDICATOR_NEUTRAL: &str = "⚪"; // White circle - first fetch / no change

impl Direction {
    /// Returns the emoji symbol for this direction
    pub fn indicator(&self) -> &'static str {
        match self {
            Direction::Up => INDICATOR_UP,
            Direction::Down => INDICATOR_DOWN,
            Direction::Neutral => INDICATOR_NEUTRAL,
        }
    }
}
