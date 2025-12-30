package config

import "fmt"

// Validate checks if the configuration is valid
func (c *Config) Validate() error {
	if c.RefreshSeconds < MinRefreshSeconds || c.RefreshSeconds > MaxRefreshSeconds {
		return fmt.Errorf("refresh_seconds must be between %d and %d", MinRefreshSeconds, MaxRefreshSeconds)
	}
	if len(c.Symbols) == 0 {
		return fmt.Errorf("at least one symbol is required")
	}
	if c.ProviderID == "" {
		return fmt.Errorf("provider_id is required")
	}
	return nil
}
