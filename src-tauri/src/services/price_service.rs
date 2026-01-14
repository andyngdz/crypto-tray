use std::sync::Arc;
use std::time::Duration;
use tauri::{AppHandle, Emitter, Wry};
use tokio::sync::{mpsc, watch};
use tokio::time::interval;

use crate::config::ConfigManager;
use crate::exchange::Converter;
use crate::movement::Tracker;
use crate::providers::ProviderRegistry;
use crate::tray::TrayState;

pub struct PriceService {
    refresh_tx: mpsc::Sender<()>,
}

impl PriceService {
    pub fn new(
        app: AppHandle<Wry>,
        config_manager: Arc<ConfigManager>,
        registry: Arc<ProviderRegistry>,
        tray_state: Arc<TrayState>,
        converter: Arc<Converter>,
        shutdown_rx: watch::Receiver<bool>,
    ) -> Self {
        let (refresh_tx, refresh_rx) = mpsc::channel::<()>(1);

        tauri::async_runtime::spawn(Self::run_polling_loop(
            app,
            config_manager,
            registry,
            tray_state,
            converter,
            refresh_rx,
            shutdown_rx,
        ));

        Self { refresh_tx }
    }

    pub async fn trigger_refresh(&self) {
        if let Err(e) = self.refresh_tx.send(()).await {
            eprintln!("Failed to send refresh signal: {}", e);
        }
    }

    async fn run_polling_loop(
        app: AppHandle<Wry>,
        config_manager: Arc<ConfigManager>,
        registry: Arc<ProviderRegistry>,
        tray_state: Arc<TrayState>,
        converter: Arc<Converter>,
        mut refresh_rx: mpsc::Receiver<()>,
        mut shutdown_rx: watch::Receiver<bool>,
    ) {
        let movement_tracker = Tracker::new();

        let mut current_symbols: Vec<String> = Vec::new();
        let mut current_provider_id = String::new();
        let mut current_refresh_seconds: u64;

        loop {
            if *shutdown_rx.borrow() {
                return;
            }

            let config = config_manager.get();

            if config.provider_id != current_provider_id {
                if let Err(e) = registry.get_symbols(&config.provider_id).await {
                    eprintln!("Failed to fetch symbols: {}", e);
                }
                current_provider_id = config.provider_id.clone();
            }

            // Rebuild tray menu if symbols changed
            if config.symbols != current_symbols {
                if let Err(e) = tray_state.rebuild_menu(&app, config.symbols.clone()).await {
                    eprintln!("Failed to rebuild tray menu: {}", e);
                }
                current_symbols = config.symbols.clone();
            }

            // Recreate interval if refresh rate changed
            let refresh_seconds = config.refresh_seconds as u64;
            let mut poll_interval = interval(Duration::from_secs(refresh_seconds));
            current_refresh_seconds = refresh_seconds;

            Self::fetch_and_emit(&app, &config_manager, &registry, &tray_state, &converter, &movement_tracker)
                .await;

            loop {
                tokio::select! {
                    _ = poll_interval.tick() => {
                        Self::fetch_and_emit(&app, &config_manager, &registry, &tray_state, &converter, &movement_tracker).await;
                    }
                    Some(()) = refresh_rx.recv() => {
                        let new_config = config_manager.get();
                        if new_config.symbols != current_symbols {
                            if let Err(e) = tray_state.rebuild_menu(&app, new_config.symbols.clone()).await {
                                eprintln!("Failed to rebuild tray menu: {}", e);
                            }
                            current_symbols = new_config.symbols.clone();
                        }
                        if new_config.provider_id != current_provider_id {
                            break;
                        }
                        if new_config.refresh_seconds as u64 != current_refresh_seconds {
                            break;
                        }
                        Self::fetch_and_emit(&app, &config_manager, &registry, &tray_state, &converter, &movement_tracker).await;
                    }
                    result = shutdown_rx.changed() => {
                        if result.is_err() || *shutdown_rx.borrow() {
                            return;
                        }
                    }
                }
            }
        }
    }

    async fn fetch_and_emit(
        app: &AppHandle<Wry>,
        config_manager: &Arc<ConfigManager>,
        registry: &Arc<ProviderRegistry>,
        tray_state: &Arc<TrayState>,
        converter: &Arc<Converter>,
        movement_tracker: &Tracker,
    ) {
        let config = config_manager.get();

        match registry
            .fetch_prices(&config.provider_id, &config.symbols)
            .await
        {
            Ok(mut prices) => {
                converter.convert_prices(&mut prices).await;
                let movements = movement_tracker.track(&prices).await;
                if let Err(e) = app.emit("price:update", &prices) {
                    eprintln!("Failed to emit price update: {}", e);
                }
                tray_state.update_prices(app, &prices, &movements).await;
            }
            Err(e) => {
                eprintln!("Failed to fetch prices: {}", e);
                tray_state.set_error(app).await;
            }
        }
    }
}
