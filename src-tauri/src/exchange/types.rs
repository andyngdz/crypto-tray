use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::time::Duration;

pub const PRIMARY_URL: &str =
    "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies";
pub const FALLBACK_URL: &str = "https://latest.currency-api.pages.dev/v1/currencies";
pub const BASE_CURRENCY: &str = "usdt";
pub const TIMEOUT: Duration = Duration::from_secs(10);

#[derive(Debug, Clone, Serialize)]
pub struct ExchangeRates {
    pub date: String,
    pub base: String,
    pub rates: HashMap<String, f64>,
}

#[derive(Debug, Deserialize)]
pub struct ApiResponse {
    pub date: String,
    pub usdt: HashMap<String, f64>,
}
