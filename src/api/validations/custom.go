package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/mohar9h/golang-clear-web-api/common"
	"log"
	"regexp"
)

func IranMobileValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	log.Print(value)
	if !ok {
		return false
	}

	matchString, err := regexp.MatchString(`^09(1[0-9)|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return matchString
}

func PasswordValidator(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)
	if !ok {
		field.Param()
		return false
	}
	return common.CheckPassword(value)
}
