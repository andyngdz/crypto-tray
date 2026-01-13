package main

import (
	"crypto-tray/autostart"
	"crypto-tray/config"
	"crypto-tray/providers"
)

// AppDependencies holds all initialized application dependencies
type AppDependencies struct {
	ConfigManager *config.Manager
	Registry      *providers.Registry
}

// InitApp creates and initializes all application dependencies
func InitApp() (*AppDependencies, error) {
	configManager, err := config.NewManager()
	if err != nil {
		return nil, err
	}

	// Enable auto-start on first run if configured
	cfg := configManager.Get()
	if cfg.AutoStart && !autostart.IsEnabled() {
		autostart.SetEnabled(true)
	}

	registry := providers.NewRegistry()
	registry.Register(providers.NewCoinGecko())
	registry.Register(providers.NewBinance())

	return &AppDependencies{
		ConfigManager: configManager,
		Registry:      registry,
	}, nil
}
