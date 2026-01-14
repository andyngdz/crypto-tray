use tauri::{image::Image, App, Wry};

use crate::error::Result;

#[cfg(target_os = "macos")]
pub fn tray_icon(_app: &App<Wry>) -> Result<Image<'static>> {
    let icon = Image::from_bytes(include_bytes!(concat!(
        env!("CARGO_MANIFEST_DIR"),
        "/icons/icon-template.png"
    )))?;

    Ok(icon)
}

#[cfg(not(target_os = "macos"))]
pub fn tray_icon<'app>(app: &'app App<Wry>) -> Result<Image<'app>> {
    Ok(app.default_window_icon().unwrap().clone())
}
