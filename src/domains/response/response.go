package responses

import (
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"net/http"
)

// Response defines the standard API response structure.
type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"domains,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

// Success creates a successful response.
func Success(code helpers.ResultCode, data interface{}, message string) Response {
	return Response{
		Success: true,
		Code:    int(code),
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

// Error creates an error response.
func Error(code helpers.ResultCode, err error) Response {
	return Response{
		Success: false,
		Code:    int(code),
		Message: http.StatusText(int(code)),
		Data:    nil,
		Errors:  map[string]interface{}{"message": err.Error()},
	}
}

// ValidationError creates a validation error response.
func ValidationError(errors map[string]string) Response {
	return Response{
		Success: false,
		Code:    http.StatusUnprocessableEntity,
		Message: "Validation failed",
		Data:    nil,
		Errors:  errors,
	}
}
