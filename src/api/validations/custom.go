package validations

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {
	var validationErrors []ValidationError
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError) {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.Property = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			validationErrors = append(validationErrors, element)
		}
		return &validationErrors
	}
	return nil
}
