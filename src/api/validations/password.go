package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/mohar9h/golang-clear-web-api/common"
)

func PasswordValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if !ok {
		field.Param()
		return false
	}
	return common.CheckPassword(value)
}
