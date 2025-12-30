package config

// Configuration constants
const (
	MinRefreshSeconds     = 10
	MaxRefreshSeconds     = 3600
	DefaultRefreshSeconds = 30
	DefaultSymbol         = "BTC"
	DefaultProviderID     = "coingecko"
	configFileName        = "config.json"
	appDirName            = "crypto-tray"
)

// Config holds the application configuration
type Config struct {
	ProviderID     string            `json:"provider_id"`
	APIKeys        map[string]string `json:"api_keys"`
	RefreshSeconds int               `json:"refresh_seconds"`
	Symbols        []string          `json:"symbols"`
}

func defaultConfig() *Config {
	return &Config{
		ProviderID:     DefaultProviderID,
		APIKeys:        make(map[string]string),
		RefreshSeconds: DefaultRefreshSeconds,
		Symbols:        []string{DefaultSymbol},
	}
}
