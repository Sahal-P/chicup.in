package utils

import (
	"fmt"
)

// FormatPrice formats a price to include a currency symbol.
func FormatPrice(price float64) string {
	return fmt.Sprintf("$%.2f", price)
}
