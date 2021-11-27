package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/russelshane/hyperbank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		// check currency if supported
		return util.IsCurrencySupported(currency)
	}

	return false

}