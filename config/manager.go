package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// Manager handles configuration persistence
type Manager struct {
	config   *Config
	mu       sync.RWMutex
	filePath string
}

// NewManager creates a new configuration manager
func NewManager() (*Manager, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	appConfigDir := filepath.Join(configDir, appDirName)
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		return nil, err
	}

	m := &Manager{
		filePath: filepath.Join(appConfigDir, configFileName),
		config:   defaultConfig(),
	}

	// Try to load existing config, ignore if doesn't exist
	if err := m.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return m, nil
}

// Load reads configuration from disk
func (m *Manager) Load() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := os.ReadFile(m.filePath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, m.config); err != nil {
		return err
	}

	return nil
}

// Save writes configuration to disk
func (m *Manager) Save() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := json.MarshalIndent(m.config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(m.filePath, data, 0644)
}

// Get returns a copy of the current configuration
func (m *Manager) Get() Config {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return *m.config
}

// Update validates, updates the configuration and saves to disk
func (m *Manager) Update(cfg Config) error {
	if err := cfg.Validate(); err != nil {
		return err
	}
	m.mu.Lock()
	m.config = &cfg
	m.mu.Unlock()
	return m.Save()
}
