package common

import (
	logging2 "github.com/mohar9h/golang-clear-web-api/pkg/logging"
	"regexp"

	"github.com/mohar9h/golang-clear-web-api/config"
)

const iranianMobileNumberPattern string = `^09(1[0-9)|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

var logger = logging2.NewLogger(config.GetConfig())

func IranianMobileValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(iranianMobileNumberPattern, mobileNumber)
	if err != nil {
		logger.Error(logging2.Validation, logging2.MobileValidation, err.Error(), nil)
	}
	return res
}
