package exchange

import (
	"crypto-tray/config"
	"crypto-tray/providers"
)

// Converter handles currency conversion for price data
type Converter struct {
	fetcher       *Fetcher
	configManager *config.Manager
}

// NewConverter creates a new currency converter
func NewConverter(fetcher *Fetcher, configManager *config.Manager) *Converter {
	return &Converter{
		fetcher:       fetcher,
		configManager: configManager,
	}
}

// ConvertPrices applies exchange rate conversion to price data
func (c *Converter) ConvertPrices(data []*providers.PriceData) {
	cfg := c.configManager.Get()
	rates := c.fetcher.GetRates()
	rate := 1.0
	currency := cfg.DisplayCurrency

	if rates != nil {
		if r, ok := rates.Rates[cfg.DisplayCurrency]; ok {
			rate = r
		}
	}

	for dataIdx := range data {
		data[dataIdx].ConvertedPrice = data[dataIdx].Price * rate
		data[dataIdx].Currency = currency
	}
}
