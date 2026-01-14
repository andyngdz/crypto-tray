use std::sync::Arc;

use crate::config::ConfigManager;
use crate::providers::PriceData;

use super::fetcher::Fetcher;

pub struct Converter {
    fetcher: Arc<Fetcher>,
    config_manager: Arc<ConfigManager>,
}

impl Converter {
    pub fn new(fetcher: Arc<Fetcher>, config_manager: Arc<ConfigManager>) -> Self {
        Self {
            fetcher,
            config_manager,
        }
    }

    pub async fn convert_prices(&self, data: &mut [PriceData]) {
        let cfg = self.config_manager.get();
        let rates = self.fetcher.get_rates().await;
        let mut rate = 1.0;
        let currency = cfg.display_currency.clone();

        if let Some(ref exchange_rates) = rates {
            if let Some(&r) = exchange_rates.rates.get(&cfg.display_currency) {
                rate = r;
            }
        }

        for price_data in data.iter_mut() {
            price_data.converted_price = price_data.price * rate;
            price_data.currency = currency.clone();
        }
    }
}
