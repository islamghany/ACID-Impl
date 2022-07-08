package util

const (
	EGP = "EGP"
	USD = "USD"
	EUR = "EUR"
)

func CheckCurrency(currency string) bool {

	switch currency {
	case EGP, USD, EUR:
		return true
	}
	return false
}
