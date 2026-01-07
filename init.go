package main

import (
	"crypto-tray/config"
	"crypto-tray/providers"
)

// AppDependencies holds all initialized application dependencies
type AppDependencies struct {
	ConfigManager *config.Manager
	Registry      *providers.Registry
	App           *App
}

// InitApp creates and initializes all application dependencies
func InitApp() (*AppDependencies, error) {
	configManager, err := config.NewManager()
	if err != nil {
		return nil, err
	}

	registry := providers.NewRegistry()
	registry.Register(providers.NewCoinGecko())
	registry.Register(providers.NewBinance())

	app := NewApp(configManager, registry)

	return &AppDependencies{
		ConfigManager: configManager,
		Registry:      registry,
		App:           app,
	}, nil
}
