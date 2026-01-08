package services

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FormatPrice formats a price with thousand separators and no decimals
// format: "us", "european", or "asian"
func FormatPrice(price float64, format string) string {
	var lang language.Tag
	switch format {
	case "european":
		lang = language.German
	default: // "us", "asian", or any other value
		lang = language.English
	}
	p := message.NewPrinter(lang)
	return p.Sprintf("$%.0f", price)
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
