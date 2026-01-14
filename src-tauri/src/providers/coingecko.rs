use async_trait::async_trait;
use reqwest::Client;
use serde::Deserialize;
use std::collections::HashMap;
use tokio::sync::RwLock;

use crate::error::Result;
use crate::providers::traits::Provider;
use crate::providers::types::{PriceData, ProviderInfo, SymbolInfo};

const BASE_URL: &str = "https://api.coingecko.com/api/v3";

#[derive(Debug, Deserialize)]
struct MarketCoin {
    id: String,
    symbol: String,
    name: String,
}

#[derive(Debug, Deserialize)]
struct PriceResponse {
    usd: Option<f64>,
    usd_24h_change: Option<f64>,
}

pub struct CoinGeckoProvider {
    client: Client,
    api_key: Option<String>,
    symbol_map: RwLock<HashMap<String, String>>,
}

impl CoinGeckoProvider {
    pub fn new() -> Self {
        Self {
            client: Client::new(),
            api_key: None,
            symbol_map: RwLock::new(HashMap::new()),
        }
    }
}

#[async_trait]
impl Provider for CoinGeckoProvider {
    fn info(&self) -> ProviderInfo {
        ProviderInfo {
            id: "coingecko".to_string(),
            name: "CoinGecko".to_string(),
            requires_api_key: false,
        }
    }

    fn set_api_key(&mut self, key: String) {
        self.api_key = Some(key);
    }

    fn default_coin_ids(&self) -> Vec<String> {
        vec![
            "bitcoin".to_string(),
            "ethereum".to_string(),
            "solana".to_string(),
        ]
    }

    async fn fetch_symbols(&self) -> Result<Vec<SymbolInfo>> {
        let url = format!(
            "{}/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1",
            BASE_URL
        );

        let mut request = self.client.get(&url);
        if let Some(ref key) = self.api_key {
            request = request.header("x-cg-demo-api-key", key);
        }

        let response: Vec<MarketCoin> = request.send().await?.json().await?;

        let mut symbol_map = HashMap::with_capacity(response.len());
        let symbols = response
            .into_iter()
            .map(|c| {
                symbol_map.insert(c.id.clone(), c.symbol.to_uppercase());
                SymbolInfo {
                    id: c.id,
                    symbol: c.symbol.to_uppercase(),
                    name: c.name,
                }
            })
            .collect();

        *self.symbol_map.write().await = symbol_map;

        Ok(symbols)
    }

    async fn fetch_prices(&self, symbols: &[String]) -> Result<Vec<PriceData>> {
        // CoinGecko uses coin IDs, not symbols
        // For simplicity, we'll treat symbols as IDs here
        let ids = symbols.join(",");
        let url = format!(
            "{}/simple/price?ids={}&vs_currencies=usd&include_24hr_change=true",
            BASE_URL, ids
        );

        let mut request = self.client.get(&url);
        if let Some(ref key) = self.api_key {
            request = request.header("x-cg-demo-api-key", key);
        }

        let response: HashMap<String, PriceResponse> = request.send().await?.json().await?;

        let symbol_map = self.symbol_map.read().await;

        let prices = response
            .into_iter()
            .map(|(id, data)| {
                let symbol = symbol_map
                    .get(&id)
                    .cloned()
                    .unwrap_or_else(|| id.to_uppercase());

                PriceData {
                    coin_id: id.clone(),
                    symbol,
                    name: String::new(),
                    price: data.usd.unwrap_or(0.0),
                    change_24h: data.usd_24h_change.unwrap_or(0.0),
                    converted_price: 0.0,
                    currency: "usd".to_string(),
                }
            })
            .collect();

        Ok(prices)
    }
}
