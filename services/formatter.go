package services

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// FormatPrice formats a price with thousand separators and no decimals
func FormatPrice(price float64) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("$%.0f", price)
}
