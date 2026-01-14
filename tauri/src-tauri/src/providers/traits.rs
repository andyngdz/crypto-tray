use async_trait::async_trait;

use crate::error::Result;
use crate::providers::types::{PriceData, ProviderInfo, SymbolInfo};

#[async_trait]
pub trait Provider: Send + Sync {
    fn info(&self) -> ProviderInfo;
    fn set_api_key(&mut self, key: String);
    fn default_coin_ids(&self) -> Vec<String>;
    async fn fetch_symbols(&self) -> Result<Vec<SymbolInfo>>;
    async fn fetch_prices(&self, symbols: &[String]) -> Result<Vec<PriceData>>;
}
