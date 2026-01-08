package services

import (
	"fmt"
	"strings"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// GetCurrencySymbol returns the display symbol for a currency code using dynamic lookup
func GetCurrencySymbol(currencyCode string) string {
	unit, err := currency.ParseISO(strings.ToUpper(currencyCode))

	if err != nil {
		return strings.ToUpper(currencyCode)
	}

	return fmt.Sprint(currency.Symbol(unit))
}

// FormatPrice formats a price with thousand separators and no decimals
// format: "us", "european", or "asian"
func FormatPrice(price float64, format string) string {
	return FormatPriceWithCurrency(price, format, "usd")
}

// FormatPriceWithCurrency formats a price with the specified currency symbol
func FormatPriceWithCurrency(price float64, format string, currency string) string {
	var lang language.Tag
	switch format {
	case "european":
		lang = language.German
	default: // "us", "asian", or any other value
		lang = language.English
	}
	p := message.NewPrinter(lang)
	symbol := GetCurrencySymbol(currency)
	return p.Sprintf("%s%.0f", symbol, price)
}

// FormatTrayTitle builds a tray title from symbols with a suffix for each
// Example: FormatTrayTitle([]string{"ETH", "BTC"}, "$--,---") => "ETH $--,--- | BTC $--,---"
func FormatTrayTitle(symbols []string, suffix string) string {
	if len(symbols) == 0 {
		return ""
	}
	parts := make([]string, len(symbols))
	for i, symbol := range symbols {
		parts[i] = symbol + " " + suffix
	}
	return strings.Join(parts, " | ")
}
