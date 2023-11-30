package util

const (
	USD = "USD"
	KHR = "KHR"
	THB = "THB"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, KHR, THB:
		return true
	}
	return false
}