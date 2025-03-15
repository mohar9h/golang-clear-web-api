package helpers

import "github.com/mohar9h/golang-clear-web-api/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       int                            `json:"result_code"`
	ValidationErrors *[]validations.ValidationError `json:"validation_errors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result, Success: success, ResultCode: resultCode}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result, Success: success, ResultCode: resultCode, Error: err}
}

func GenerateBaseResponseWithValidationErrors(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result, Success: success, ResultCode: resultCode, ValidationErrors: validations.GetValidationErrors(err)}
}
