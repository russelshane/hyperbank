package util

// Constants for hyperbank supported currencies
const (
	PHP = "PHP"
	USD = "USD"
)

func IsCurrencySupported(currency string) bool {
	switch currency {
	case PHP, USD:
		return true
	}
	return false
}