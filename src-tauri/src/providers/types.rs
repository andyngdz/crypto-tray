use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct ProviderInfo {
    pub id: String,
    pub name: String,
    pub requires_api_key: bool,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct SymbolInfo {
    #[serde(rename = "coinId")]
    pub id: String,
    pub symbol: String,
    pub name: String,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct PriceData {
    pub coin_id: String,
    pub symbol: String,
    pub name: String,
    pub price: f64,
    #[serde(rename = "change_24h")]
    pub change_24h: f64,
    pub converted_price: f64,
    pub currency: String,
}
