use std::collections::HashMap;
use std::sync::Arc;
use tauri::{
    menu::{Menu, MenuBuilder, MenuItemBuilder, PredefinedMenuItem},
    tray::TrayIconBuilder,
    AppHandle, Manager, Runtime, Wry,
};
use tokio::sync::{Mutex, RwLock};

use crate::config::ConfigManager;
use crate::error::Result;
use crate::movement::Direction;
use crate::providers::PriceData;
use crate::services::{format_price_with_currency, format_tray_title, PriceService};

const ERROR_SUFFIX: &str = "$???";

/// TrayState holds the mutable tray state for dynamic updates
pub struct TrayState {
    config_manager: Arc<ConfigManager>,
    price_items: RwLock<Vec<tauri::menu::MenuItem<Wry>>>,
    symbols: RwLock<Vec<String>>,
}

impl TrayState {
    pub async fn set_error(&self, app: &AppHandle<Wry>) {
        let symbols = self.symbols.read().await;
        let display_symbols: Vec<String> = symbols
            .iter()
            .map(|s| s.strip_suffix("USDT").unwrap_or(s).to_string())
            .collect();

        let title = format_tray_title(&display_symbols, ERROR_SUFFIX);
        if let Some(tray) = app.tray_by_id("main") {
            if let Err(e) = tray.set_title(Some(&title)) {
                eprintln!("Failed to set tray title: {}", e);
            }
            if let Err(e) = tray.set_tooltip(Some(&title)) {
                eprintln!("Failed to set tray tooltip: {}", e);
            }
        }

        let price_items = self.price_items.read().await;
        for (slot_idx, item) in price_items.iter().enumerate() {
            if slot_idx >= display_symbols.len() {
                break;
            }
            if let Err(e) = item.set_text(&format!("{} Error", display_symbols[slot_idx])) {
                eprintln!("Failed to set tray item text: {}", e);
            }
        }
    }

    /// Rebuild the tray menu with new symbols
    pub async fn rebuild_menu(&self, app: &AppHandle<Wry>, symbols: Vec<String>) -> Result<()> {
        let menu = build_menu(app, &symbols)?;

        // Update tray menu
        if let Some(tray) = app.tray_by_id("main") {
            tray.set_menu(Some(menu.clone()))?;
        }

        // Update symbols
        *self.symbols.write().await = symbols.clone();

        // Update price_items with new menu items
        let mut price_items = self.price_items.write().await;
        price_items.clear();
        for (slot_idx, coin_id) in symbols.iter().enumerate() {
            if let Some(item) = menu.get(format!("price_{}", slot_idx).as_str()) {
                if let tauri::menu::MenuItemKind::MenuItem(menu_item) = item {
                    price_items.push(menu_item);
                }
            }

            // Update loading state label
            let symbol = coin_id.strip_suffix("USDT").unwrap_or(coin_id);
            if let Some(item) = price_items.get(slot_idx) {
                if let Err(e) = item.set_text(&format!("{} $--,---", symbol)) {
                    eprintln!("Failed to set tray item text: {}", e);
                }
            }
        }

        // Update tray title
        let display_symbols: Vec<String> = symbols
            .iter()
            .map(|s| s.strip_suffix("USDT").unwrap_or(s).to_string())
            .collect();
        let title = format_tray_title(&display_symbols, "$--,---");
        if let Some(tray) = app.tray_by_id("main") {
            if let Err(e) = tray.set_title(Some(&title)) {
                eprintln!("Failed to set tray title: {}", e);
            }
            if let Err(e) = tray.set_tooltip(Some("CryptoTray - Loading...")) {
                eprintln!("Failed to set tray tooltip: {}", e);
            }
        }

        Ok(())
    }

    pub async fn update_prices<R: Runtime>(
        &self,
        app: &tauri::AppHandle<R>,
        prices: &[PriceData],
        movements: &HashMap<String, Direction>,
    ) {
        if prices.is_empty() {
            return;
        }

        let config = self.config_manager.get();
        let symbols = self.symbols.read().await;
        let price_items = self.price_items.read().await;
        let mut title_parts: Vec<String> = Vec::new();

        // Update each price item
        for (slot_idx, price_item) in price_items.iter().enumerate() {
            if slot_idx >= symbols.len() {
                break;
            }

            let coin_id = &symbols[slot_idx];

            // Find price data for this coin
            if let Some(price_data) = prices.iter().find(|p| &p.coin_id == coin_id) {
                let direction = movements
                    .get(coin_id)
                    .copied()
                    .unwrap_or(Direction::Neutral);
                let indicator = direction.indicator();

                let price = if price_data.converted_price != 0.0 {
                    price_data.converted_price
                } else {
                    price_data.price
                };

                let price_text = format_price_with_currency(
                    price,
                    &config.number_format,
                    &price_data.currency,
                );
                let label = format!("{} {} {}", indicator, price_data.symbol, price_text);

                if let Err(e) = price_item.set_text(&label) {
                    eprintln!("Failed to set tray item text: {}", e);
                }

                // Add to title
                title_parts.push(format!("{} {} {}", indicator, price_data.symbol, price_text));
            }
        }

        // Update tray title
        if !title_parts.is_empty() {
            let title = title_parts.join(" | ");
            if let Some(tray) = app.tray_by_id("main") {
                if let Err(e) = tray.set_title(Some(&title)) {
                    eprintln!("Failed to set tray title: {}", e);
                }
                if let Err(e) = tray.set_tooltip(Some(&title)) {
                    eprintln!("Failed to set tray tooltip: {}", e);
                }
            }
        }
    }
}

/// Build the tray menu with given symbols
fn build_menu<M: Manager<Wry>>(app: &M, symbols: &[String]) -> Result<Menu<Wry>> {
    let settings = MenuItemBuilder::with_id("settings", "Open Settings").build(app)?;
    let refresh = MenuItemBuilder::with_id("refresh", "Refresh Now").build(app)?;
    let quit = MenuItemBuilder::with_id("quit", "Quit").build(app)?;
    let separator1 = PredefinedMenuItem::separator(app)?;
    let separator2 = PredefinedMenuItem::separator(app)?;

    let mut menu_builder = MenuBuilder::new(app);

    // Add price items
    for (slot_idx, coin_id) in symbols.iter().enumerate() {
        let symbol = coin_id.strip_suffix("USDT").unwrap_or(coin_id);
        let label = format!("{} $--,---", symbol);
        let item = MenuItemBuilder::with_id(format!("price_{}", slot_idx), &label).build(app)?;
        menu_builder = menu_builder.item(&item);
    }

    // Add separators and actions
    let menu = menu_builder
        .item(&separator1)
        .item(&settings)
        .item(&refresh)
        .item(&separator2)
        .item(&quit)
        .build()?;

    Ok(menu)
}

pub fn setup_tray(
    app: &tauri::App<Wry>,
    initial_symbols: Vec<String>,
    config_manager: Arc<ConfigManager>,
) -> Result<Arc<TrayState>> {
    let menu = build_menu(app, &initial_symbols)?;

    // Collect price items from menu
    let mut price_items: Vec<tauri::menu::MenuItem<Wry>> = Vec::with_capacity(initial_symbols.len());
    for slot_idx in 0..initial_symbols.len() {
        if let Some(item) = menu.get(format!("price_{}", slot_idx).as_str()) {
            if let tauri::menu::MenuItemKind::MenuItem(menu_item) = item {
                price_items.push(menu_item);
            }
        }
    }

    // Initial title
    let display_symbols: Vec<String> = initial_symbols
        .iter()
        .map(|s| s.strip_suffix("USDT").unwrap_or(s).to_string())
        .collect();
    let initial_title = format_tray_title(&display_symbols, "$--,---");

    let tray_builder = TrayIconBuilder::with_id("main")
        .icon(app.default_window_icon().unwrap().clone())
        .menu(&menu)
        .title(&initial_title)
        .tooltip("CryptoTray - Loading...")
        .show_menu_on_left_click(false);

    #[cfg(target_os = "macos")]
    let tray_builder = tray_builder.icon_as_template(true);

    tray_builder.on_menu_event(move |app, event| {
            let id = event.id().as_ref();

            // Ignore price item clicks
            if id.starts_with("price_") {
                return;
            }

            match id {
                "settings" => {
                    if let Some(window) = app.get_webview_window("main") {
                        if let Err(e) = window.show() {
                            eprintln!("Failed to show window: {}", e);
                        }
                        if let Err(e) = window.set_focus() {
                            eprintln!("Failed to focus window: {}", e);
                        }
                    }
                }
                "refresh" => {
                    let app_handle = app.clone();
                    tauri::async_runtime::spawn(async move {
                        let state = app_handle.state::<Mutex<Option<PriceService>>>();
                        let guard = state.lock().await;
                        if let Some(service) = guard.as_ref() {
                            service.trigger_refresh().await;
                        }
                    });
                }
                "quit" => {
                    app.exit(0);
                }
                _ => {}
            }
        })
        .build(app)?;

    let tray_state = Arc::new(TrayState {
        config_manager,
        price_items: RwLock::new(price_items),
        symbols: RwLock::new(initial_symbols),
    });

    Ok(tray_state)
}
