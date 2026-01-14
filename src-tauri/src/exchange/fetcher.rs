use std::sync::Arc;
use std::time::Duration;
use tauri::{AppHandle, Emitter, Wry};
use tokio::sync::{mpsc, watch, RwLock};
use tokio::time::interval;

use crate::config::ConfigManager;

use super::types::{ApiResponse, ExchangeRates, BASE_CURRENCY, FALLBACK_URL, PRIMARY_URL, TIMEOUT};

pub struct Fetcher {
    current_rates: Arc<RwLock<Option<ExchangeRates>>>,
    refresh_tx: mpsc::Sender<()>,
}

impl Fetcher {
    pub fn new(
        app: AppHandle<Wry>,
        config_manager: Arc<ConfigManager>,
        shutdown_rx: watch::Receiver<bool>,
    ) -> Self {
        let (refresh_tx, refresh_rx) = mpsc::channel::<()>(1);
        let current_rates = Arc::new(RwLock::new(None));

        let client = reqwest::Client::builder()
            .timeout(TIMEOUT)
            .build()
            .unwrap_or_default();

        let fetcher = Self {
            current_rates: current_rates.clone(),
            refresh_tx,
        };

        tauri::async_runtime::spawn(Self::run_loop(
            app,
            config_manager,
            current_rates,
            refresh_rx,
            client,
            shutdown_rx,
        ));

        fetcher
    }

    pub async fn refresh_now(&self) {
        if let Err(e) = self.refresh_tx.send(()).await {
            eprintln!("Failed to send exchange refresh signal: {}", e);
        }
    }

    pub async fn get_rates(&self) -> Option<ExchangeRates> {
        self.current_rates.read().await.clone()
    }

    async fn run_loop(
        app: AppHandle<Wry>,
        config_manager: Arc<ConfigManager>,
        current_rates: Arc<RwLock<Option<ExchangeRates>>>,
        mut refresh_rx: mpsc::Receiver<()>,
        client: reqwest::Client,
        mut shutdown_rx: watch::Receiver<bool>,
    ) {
        Self::fetch_once(&current_rates, &client).await;
        Self::emit_rates(&app, &current_rates).await;

        let cfg = config_manager.get();
        let mut current_interval = cfg.refresh_seconds;
        let mut poll_interval = interval(Duration::from_secs(current_interval as u64));

        loop {
            if *shutdown_rx.borrow() {
                return;
            }

            tokio::select! {
                _ = poll_interval.tick() => {
                    let cfg = config_manager.get();
                    Self::fetch_once(&current_rates, &client).await;
                    Self::emit_rates(&app, &current_rates).await;

                    if cfg.refresh_seconds != current_interval {
                        current_interval = cfg.refresh_seconds;
                        poll_interval = interval(Duration::from_secs(current_interval as u64));
                    }
                }
                Some(()) = refresh_rx.recv() => {
                    let cfg = config_manager.get();
                    Self::fetch_once(&current_rates, &client).await;
                    Self::emit_rates(&app, &current_rates).await;

                    current_interval = cfg.refresh_seconds;
                    poll_interval = interval(Duration::from_secs(current_interval as u64));
                }
                result = shutdown_rx.changed() => {
                    if result.is_err() || *shutdown_rx.borrow() {
                        return;
                    }
                }
            }
        }
    }

    async fn fetch_once(
        current_rates: &Arc<RwLock<Option<ExchangeRates>>>,
        client: &reqwest::Client,
    ) {
        let url = format!("{}/{}.json", PRIMARY_URL, BASE_CURRENCY);
        match Self::fetch_from_url(client, &url).await {
            Ok(rates) => {
                *current_rates.write().await = Some(rates);
                return;
            }
            Err(_) => {
                let fallback_url = format!("{}/{}.json", FALLBACK_URL, BASE_CURRENCY);
                if let Ok(rates) = Self::fetch_from_url(client, &fallback_url).await {
                    *current_rates.write().await = Some(rates);
                }
            }
        }
    }

    async fn emit_rates(app: &AppHandle<Wry>, current_rates: &Arc<RwLock<Option<ExchangeRates>>>) {
        let rates = current_rates.read().await;
        if let Some(ref rates) = *rates {
            if let Err(e) = app.emit("exchange:update", &rates.rates) {
                eprintln!("Failed to emit exchange update: {}", e);
            }
        }
    }

    async fn fetch_from_url(
        client: &reqwest::Client,
        url: &str,
    ) -> Result<ExchangeRates, reqwest::Error> {
        let resp: ApiResponse = client.get(url).send().await?.json().await?;

        Ok(ExchangeRates {
            date: resp.date,
            base: BASE_CURRENCY.to_string(),
            rates: resp.usdt,
        })
    }
}
