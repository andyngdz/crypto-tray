package services

import (
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FormatPrice formats a price with thousand separators and no decimals
func FormatPrice(price float64) string {
	p := message.NewPrinter(language.English)
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
