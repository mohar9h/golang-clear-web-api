package common

import (
	"regexp"

	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/logging"
)

const iranianMobileNumberPattern string = `^09(1[0-9)|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

var logger = logging.NewLogger(config.GetConfig())

func IranianMobileValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(iranianMobileNumberPattern, mobileNumber)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
