use async_trait::async_trait;
use reqwest::Client;
use serde::Deserialize;

use crate::error::Result;
use crate::providers::traits::Provider;
use crate::providers::types::{PriceData, ProviderInfo, SymbolInfo};

const BASE_URL: &str = "https://api.binance.com";

#[derive(Debug, Deserialize)]
struct TickerResponse {
    symbol: String,
    #[serde(rename = "lastPrice")]
    last_price: String,
    #[serde(rename = "priceChangePercent")]
    price_change_percent: String,
}

#[derive(Debug, Deserialize)]
struct ExchangeInfoResponse {
    symbols: Vec<ExchangeSymbol>,
}

#[derive(Debug, Deserialize)]
struct ExchangeSymbol {
    symbol: String,
    #[serde(rename = "baseAsset")]
    base_asset: String,
    #[serde(rename = "quoteAsset")]
    quote_asset: String,
    status: String,
}

pub struct BinanceProvider {
    client: Client,
}

impl BinanceProvider {
    pub fn new() -> Self {
        Self {
            client: Client::new(),
        }
    }
}

#[async_trait]
impl Provider for BinanceProvider {
    fn info(&self) -> ProviderInfo {
        ProviderInfo {
            id: "binance".to_string(),
            name: "Binance".to_string(),
            requires_api_key: false,
        }
    }

    fn set_api_key(&mut self, _key: String) {
        // Binance doesn't require API key for public endpoints
    }

    fn default_coin_ids(&self) -> Vec<String> {
        vec![
            "BTCUSDT".to_string(),
            "ETHUSDT".to_string(),
            "SOLUSDT".to_string(),
        ]
    }

    async fn fetch_symbols(&self) -> Result<Vec<SymbolInfo>> {
        let url = format!("{}/api/v3/exchangeInfo", BASE_URL);
        let response: ExchangeInfoResponse = self
            .client
            .get(&url)
            .send()
            .await?
            .json()
            .await?;

        let symbols = response
            .symbols
            .into_iter()
            .filter(|s| s.status == "TRADING" && s.quote_asset == "USDT")
            .map(|s| SymbolInfo {
                id: s.symbol,
                symbol: s.base_asset.clone(),
                name: s.base_asset,
            })
            .collect();

        Ok(symbols)
    }

    async fn fetch_prices(&self, symbols: &[String]) -> Result<Vec<PriceData>> {
        let url = format!("{}/api/v3/ticker/24hr", BASE_URL);
        let response: Vec<TickerResponse> = self
            .client
            .get(&url)
            .send()
            .await?
            .json()
            .await?;

        let prices = response
            .into_iter()
            .filter(|t| symbols.contains(&t.symbol))
            .map(|t| {
                let price = t.last_price.parse::<f64>().unwrap_or(0.0);
                let change = t.price_change_percent.parse::<f64>().unwrap_or(0.0);
                // Convert coinID (e.g., "BTCUSDT") to display symbol (e.g., "BTC")
                let display_symbol = t.symbol.strip_suffix("USDT").unwrap_or(&t.symbol).to_string();
                PriceData {
                    coin_id: t.symbol.clone(),
                    symbol: display_symbol,
                    name: String::new(),
                    price,
                    change_24h: change,
                    converted_price: 0.0, // Binance returns USD prices directly
                    currency: "usd".to_string(),
                }
            })
            .collect();

        Ok(prices)
    }
}
