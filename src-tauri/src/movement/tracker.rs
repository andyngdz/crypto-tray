use std::collections::HashMap;
use tokio::sync::RwLock;

use crate::providers::PriceData;

use super::types::Direction;

/// Tracker tracks price movements between fetches
pub struct Tracker {
    last_prices: RwLock<HashMap<String, f64>>,
    movements: RwLock<HashMap<String, Direction>>,
}

impl Tracker {
    /// Creates a new movement tracker
    pub fn new() -> Self {
        Self {
            last_prices: RwLock::new(HashMap::new()),
            movements: RwLock::new(HashMap::new()),
        }
    }

    /// Track compares current prices to last known and returns movement directions
    pub async fn track(&self, data: &[PriceData]) -> HashMap<String, Direction> {
        let mut last_prices = self.last_prices.write().await;
        let mut movements = self.movements.write().await;

        let mut result = HashMap::with_capacity(data.len());

        for price_data in data {
            let current_price = if price_data.converted_price != 0.0 {
                price_data.converted_price
            } else {
                price_data.price
            };

            let direction = if let Some(&last_price) = last_prices.get(&price_data.coin_id) {
                if current_price > last_price {
                    Direction::Up
                } else if current_price < last_price {
                    Direction::Down
                } else {
                    // Price unchanged - keep previous direction
                    movements
                        .get(&price_data.coin_id)
                        .copied()
                        .unwrap_or(Direction::Neutral)
                }
            } else {
                Direction::Neutral // First fetch
            };

            last_prices.insert(price_data.coin_id.clone(), current_price);
            movements.insert(price_data.coin_id.clone(), direction);
            result.insert(price_data.coin_id.clone(), direction);
        }

        result
    }
}

impl Default for Tracker {
    fn default() -> Self {
        Self::new()
    }
}
