mod autostart;
mod config;
mod error;
mod exchange;
mod movement;
mod providers;
mod services;
mod tray;

use std::sync::Arc;

use config::{Config, ConfigManager};
use error::Result;
use exchange::{Converter, Fetcher};
use providers::{PriceData, ProviderInfo, ProviderRegistry, SymbolInfo};
use services::PriceService;
use tauri::{Manager, RunEvent, State};
use tokio::sync::{watch, Mutex};

struct ShutdownSignals {
    price: watch::Sender<bool>,
    exchange: watch::Sender<bool>,
}

// Config commands
#[tauri::command]
fn get_config<'state>(config_manager: State<'state, Arc<ConfigManager>>) -> Result<Config> {
    Ok(config_manager.get())
}

#[tauri::command]
async fn save_config<'state>(
    mut config: Config,
    config_manager: State<'state, Arc<ConfigManager>>,
    registry: State<'state, Arc<ProviderRegistry>>,
    fetcher: State<'state, Arc<Fetcher>>,
    price_service: State<'state, Mutex<Option<PriceService>>>,
) -> Result<()> {
    // Handle autostart change
    let current_config = config_manager.get();
    if config.auto_start != current_config.auto_start {
        if let Err(e) = autostart::set_enabled(config.auto_start) {
            eprintln!("Failed to update autostart: {}", e);
        }
    }

    if config.provider_id != current_config.provider_id {
        config.symbols = registry.get_default_coin_ids(&config.provider_id).await?;
    }

    if let Some(key) = config.api_keys.get(&config.provider_id).filter(|k| !k.is_empty()) {
        registry
            .set_api_key(&config.provider_id, key.clone())
            .await?;
    }

    config_manager.save(config)?;

    fetcher.refresh_now().await;

    // Trigger refresh to apply changes immediately
    if let Some(service) = price_service.lock().await.as_ref() {
        service.trigger_refresh().await;
    }

    Ok(())
}

// Provider commands
#[tauri::command]
fn get_available_providers<'state>(
    registry: State<'state, Arc<ProviderRegistry>>,
) -> Vec<ProviderInfo> {
    registry.list()
}

#[tauri::command]
async fn get_available_symbols<'state>(
    config_manager: State<'state, Arc<ConfigManager>>,
    registry: State<'state, Arc<ProviderRegistry>>,
) -> Result<Vec<SymbolInfo>> {
    let config = config_manager.get();
    registry.get_symbols(&config.provider_id).await
}

// Price commands
#[tauri::command]
async fn fetch_prices<'state>(
    symbols: Vec<String>,
    config_manager: State<'state, Arc<ConfigManager>>,
    registry: State<'state, Arc<ProviderRegistry>>,
) -> Result<Vec<PriceData>> {
    let config = config_manager.get();
    registry.fetch_prices(&config.provider_id, &symbols).await
}

#[tauri::command]
async fn refresh_prices<'state>(
    fetcher: State<'state, Arc<Fetcher>>,
    price_service: State<'state, Mutex<Option<PriceService>>>,
) -> Result<()> {
    fetcher.refresh_now().await;

    if let Some(service) = price_service.lock().await.as_ref() {
        service.trigger_refresh().await;
    }

    Ok(())
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    let config_manager =
        Arc::new(ConfigManager::new().expect("Failed to initialize config manager"));
    let provider_registry = Arc::new(ProviderRegistry::new());

    // Sync autostart state on startup
    let config = config_manager.get();
    if config.auto_start {
        if let Ok(is_enabled) = autostart::is_enabled() {
            if !is_enabled {
                if let Err(e) = autostart::set_enabled(true) {
                    eprintln!("Failed to enable autostart: {}", e);
                }
            }
        }
    }

    // Get initial symbols from config
    let initial_symbols = config.symbols.clone();

    let (price_shutdown_tx, price_shutdown_rx) = watch::channel(false);
    let (exchange_shutdown_tx, exchange_shutdown_rx) = watch::channel(false);

    // Clone for setup closure
    let config_manager_clone = config_manager.clone();
    let provider_registry_clone = provider_registry.clone();

    let app = tauri::Builder::default()
        .manage(ShutdownSignals {
            price: price_shutdown_tx.clone(),
            exchange: exchange_shutdown_tx.clone(),
        })
        .plugin(tauri_plugin_opener::init())
        .manage(config_manager)
        .manage(provider_registry)
        .manage(Mutex::new(None::<PriceService>))
        .invoke_handler(tauri::generate_handler![
            get_config,
            save_config,
            get_available_providers,
            get_available_symbols,
            fetch_prices,
            refresh_prices,
        ])
        .setup(move |app| {
            let tray_state = tray::setup_tray(app, initial_symbols, config_manager_clone.clone())?;
            app.manage(tray_state.clone());

            let fetcher = Arc::new(Fetcher::new(
                app.handle().clone(),
                config_manager_clone.clone(),
                exchange_shutdown_rx.clone(),
            ));
            app.manage(fetcher.clone());

            let converter = Arc::new(Converter::new(fetcher, config_manager_clone.clone()));
            let price_service = PriceService::new(
                app.handle().clone(),
                config_manager_clone.clone(),
                provider_registry_clone.clone(),
                tray_state,
                converter,
                price_shutdown_rx.clone(),
            );

            let state = app.state::<Mutex<Option<PriceService>>>();
            tauri::async_runtime::block_on(async {
                *state.lock().await = Some(price_service);
            });

            Ok(())
        })
        .build(tauri::generate_context!())
        .expect("error while building tauri application");

    app.run(|app_handle, event| match event {
        RunEvent::ExitRequested { .. } | RunEvent::Exit => {
            let signals = app_handle.state::<ShutdownSignals>();
            if let Err(e) = signals.price.send(true) {
                eprintln!("Failed to signal price shutdown: {}", e);
            }
            if let Err(e) = signals.exchange.send(true) {
                eprintln!("Failed to signal exchange shutdown: {}", e);
            }
        }
        _ => {}
    });
}
