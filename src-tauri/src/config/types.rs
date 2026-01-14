use serde::{Deserialize, Serialize};
use std::collections::HashMap;

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Config {
    pub provider_id: String,
    pub api_keys: HashMap<String, String>,
    pub refresh_seconds: u32,
    pub symbols: Vec<String>,
    pub number_format: String,
    pub display_currency: String,
    pub auto_start: bool,
}

impl Default for Config {
    fn default() -> Self {
        Self {
            provider_id: "binance".to_string(),
            api_keys: HashMap::new(),
            refresh_seconds: 15,
            symbols: vec![
                "BTCUSDT".to_string(),
                "ETHUSDT".to_string(),
                "SOLUSDT".to_string(),
            ],
            number_format: "us".to_string(),
            display_currency: "usd".to_string(),
            auto_start: true,
        }
    }
}

impl Config {
    pub fn validate(&self) -> Result<(), String> {
        if self.refresh_seconds < 10 {
            return Err("Refresh interval must be at least 5 seconds".to_string());
        }
        if self.refresh_seconds > 3600 {
            return Err("Refresh interval must be at most 3600 seconds".to_string());
        }
        if self.symbols.is_empty() {
            return Err("At least one symbol is required".to_string());
        }
        if self.symbols.len() > 10 {
            return Err("Maximum 10 symbols allowed".to_string());
        }
        Ok(())
    }
}
