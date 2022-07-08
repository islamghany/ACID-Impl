package api

import (
	"github/islmaghany/bank/util"

	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		// check if the currency is valid
		return util.CheckCurrency(currency)
	}

	return false
}
