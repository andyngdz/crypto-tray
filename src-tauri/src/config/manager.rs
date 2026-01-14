use std::fs;
use std::path::PathBuf;
use std::sync::RwLock;

use crate::config::types::Config;
use crate::error::{Error, Result};

pub struct ConfigManager {
    config: RwLock<Config>,
    config_path: PathBuf,
}

impl ConfigManager {
    pub fn new() -> Result<Self> {
        let config_dir = dirs::config_dir()
            .ok_or_else(|| Error::Config("Could not find config directory".to_string()))?
            .join("CryptoTray");

        fs::create_dir_all(&config_dir)?;

        let config_path = config_dir.join("config.json");
        let config = Self::load_from_file(&config_path)?;

        Ok(Self {
            config: RwLock::new(config),
            config_path,
        })
    }

    fn load_from_file(path: &PathBuf) -> Result<Config> {
        if path.exists() {
            let content = fs::read_to_string(path)?;
            let config: Config = serde_json::from_str(&content)?;
            Ok(config)
        } else {
            Ok(Config::default())
        }
    }

    pub fn get(&self) -> Config {
        self.config.read().unwrap().clone()
    }

    pub fn save(&self, config: Config) -> Result<()> {
        config
            .validate()
            .map_err(|e| Error::Config(e))?;

        let content = serde_json::to_string_pretty(&config)?;
        fs::write(&self.config_path, content)?;

        let mut current = self.config.write().unwrap();
        *current = config;

        Ok(())
    }
}
