use std::collections::HashMap;
use std::sync::Arc;
use tokio::sync::RwLock;

use crate::error::{Error, Result};
use crate::providers::binance::BinanceProvider;
use crate::providers::coingecko::CoinGeckoProvider;
use crate::providers::traits::Provider;
use crate::providers::types::{PriceData, ProviderInfo, SymbolInfo};

pub struct ProviderRegistry {
    providers: HashMap<String, Arc<RwLock<dyn Provider>>>,
}

impl ProviderRegistry {
    pub fn new() -> Self {
        let mut providers: HashMap<String, Arc<RwLock<dyn Provider>>> = HashMap::new();

        providers.insert(
            "binance".to_string(),
            Arc::new(RwLock::new(BinanceProvider::new())),
        );
        providers.insert(
            "coingecko".to_string(),
            Arc::new(RwLock::new(CoinGeckoProvider::new())),
        );

        Self { providers }
    }

    pub fn list(&self) -> Vec<ProviderInfo> {
        // We need to block here since info() is sync but we have RwLock
        self.providers
            .values()
            .map(|p| {
                // Use try_read to avoid blocking
                p.try_read().map(|guard| guard.info()).unwrap_or_else(|_| ProviderInfo {
                    id: "unknown".to_string(),
                    name: "Unknown".to_string(),
                    requires_api_key: false,
                })
            })
            .collect()
    }

    pub async fn get_symbols(&self, provider_id: &str) -> Result<Vec<SymbolInfo>> {
        let provider = self
            .providers
            .get(provider_id)
            .ok_or_else(|| Error::Provider(format!("Provider '{}' not found", provider_id)))?;

        provider.read().await.fetch_symbols().await
    }

    pub async fn fetch_prices(
        &self,
        provider_id: &str,
        symbols: &[String],
    ) -> Result<Vec<PriceData>> {
        let provider = self
            .providers
            .get(provider_id)
            .ok_or_else(|| Error::Provider(format!("Provider '{}' not found", provider_id)))?;

        provider.read().await.fetch_prices(symbols).await
    }

    pub async fn get_default_coin_ids(&self, provider_id: &str) -> Result<Vec<String>> {
        let provider = self
            .providers
            .get(provider_id)
            .ok_or_else(|| Error::Provider(format!("Provider '{}' not found", provider_id)))?;

        Ok(provider.read().await.default_coin_ids())
    }

    pub async fn set_api_key(&self, provider_id: &str, key: String) -> Result<()> {
        let provider = self
            .providers
            .get(provider_id)
            .ok_or_else(|| Error::Provider(format!("Provider '{}' not found", provider_id)))?;

        provider.write().await.set_api_key(key);
        Ok(())
    }
}
