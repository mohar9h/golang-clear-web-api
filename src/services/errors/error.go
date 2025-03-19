package errors

type ServiceErrors struct {
	EndUserMessage   string `json:"end_user_message"`
	TechnicalMessage string `json:"technical_message"`
	Err              error  `json:"error"`
}

func (service *ServiceErrors) Error() string {
	return service.EndUserMessage
}
