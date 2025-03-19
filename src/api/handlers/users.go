package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/dto"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/services"
	"net/http"
)

type UsersHandler struct {
	services *services.UserService
}

func NewUsersHandler(config *config.Config) *UsersHandler {
	service := services.NewUserService(config)
	return &UsersHandler{services: service}
}

func (handler *UsersHandler) SendOtp(context *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	err = handler.services.SendOtp(req)
	if err != nil {
		context.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err), helpers.GenerateBaseResponseWithError(nil, false, -1, err))
	}

	context.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, 0))
}
