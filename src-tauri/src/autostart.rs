use auto_launch::AutoLaunchBuilder;
use std::env;

use crate::error::{Error, Result};

pub fn is_enabled() -> Result<bool> {
    let auto_launch = create_auto_launch()?;
    auto_launch
        .is_enabled()
        .map_err(|e| Error::AutoLaunch(e.to_string()))
}

pub fn set_enabled(enabled: bool) -> Result<()> {
    let auto_launch = create_auto_launch()?;

    if enabled {
        auto_launch
            .enable()
            .map_err(|e| Error::AutoLaunch(e.to_string()))?;
    } else {
        auto_launch
            .disable()
            .map_err(|e| Error::AutoLaunch(e.to_string()))?;
    }

    Ok(())
}

fn create_auto_launch() -> Result<auto_launch::AutoLaunch> {
    let exe_path = env::current_exe()
        .map_err(|e| Error::AutoLaunch(format!("Failed to get executable path: {}", e)))?;

    let exe_path_str = exe_path
        .to_str()
        .ok_or_else(|| Error::AutoLaunch("Invalid executable path".to_string()))?;

    AutoLaunchBuilder::new()
        .set_app_name("CryptoTray")
        .set_app_path(exe_path_str)
        .set_args(&["--minimized"])
        .build()
        .map_err(|e| Error::AutoLaunch(e.to_string()))
}
