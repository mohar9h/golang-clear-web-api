package helpers

import (
	"github.com/mohar9h/golang-clear-web-api/services/errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	// Otp
	errors.OtpExists:   409,
	errors.OtpNotValid: 400,
	errors.OtpUsed:     409,

	//User
	errors.EmailExists:    409,
	errors.UsernameExists: 409,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
