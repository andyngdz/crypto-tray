package config

// Configuration constants
const (
	MinRefreshSeconds     = 10
	MaxRefreshSeconds     = 3600
	DefaultRefreshSeconds = 15
	DefaultSymbol         = "bitcoin"
	DefaultProviderID     = "coingecko"
	DefaultNumberFormat   = "us"
	configFileName        = "config.json"
	appDirName            = "crypto-tray"
)

// Config holds the application configuration
type Config struct {
	ProviderID     string            `json:"provider_id"`
	APIKeys        map[string]string `json:"api_keys"`
	RefreshSeconds int               `json:"refresh_seconds"`
	Symbols        []string          `json:"symbols"`
	NumberFormat   string            `json:"number_format"`
}

func defaultConfig() *Config {
	return &Config{
		ProviderID:     DefaultProviderID,
		APIKeys:        make(map[string]string),
		RefreshSeconds: DefaultRefreshSeconds,
		Symbols:        []string{DefaultSymbol},
		NumberFormat:   DefaultNumberFormat,
	}
}
