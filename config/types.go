package config

// Configuration constants
const (
	MinRefreshSeconds      = 10
	MaxRefreshSeconds      = 3600
	DefaultRefreshSeconds  = 15
	DefaultSymbol          = "bitcoin"
	DefaultProviderID      = "coingecko"
	DefaultNumberFormat    = "us"
	DefaultDisplayCurrency = "usd"
	DefaultAutoStart       = true
	configFileName         = "config.json"
	appDirName             = "CryptoTray"
)

// Config holds the application configuration
type Config struct {
	ProviderID      string            `json:"provider_id"`
	APIKeys         map[string]string `json:"api_keys"`
	RefreshSeconds  int               `json:"refresh_seconds"`
	Symbols         []string          `json:"symbols"`
	NumberFormat    string            `json:"number_format"`
	DisplayCurrency string            `json:"display_currency"`
	AutoStart       bool              `json:"auto_start"`
}

func defaultConfig() *Config {
	return &Config{
		ProviderID:      DefaultProviderID,
		APIKeys:         make(map[string]string),
		RefreshSeconds:  DefaultRefreshSeconds,
		Symbols:         []string{DefaultSymbol},
		NumberFormat:    DefaultNumberFormat,
		DisplayCurrency: DefaultDisplayCurrency,
		AutoStart:       DefaultAutoStart,
	}
}
