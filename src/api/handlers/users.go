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

func (userHandler *UsersHandler) SendOtp(context *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	err = userHandler.services.SendOtp(req)
	if err != nil {
		context.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err), helpers.GenerateBaseResponseWithError(nil, false, -1, err))
	}

	context.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, 0))
}

// LoginByUsername godoc
// @Summary LoginByUsername
// @Description LoginByUsername
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-username [post]
func (userHandler *UsersHandler) LoginByUsername(context *gin.Context) {
	request := new(dto.LoginByUserNameRequest)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	token, err := userHandler.services.LoginByUsername(request)
	if err != nil {
		context.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	context.JSON(http.StatusCreated, helpers.GenerateBaseResponse(token, true, 1))
}

// RegisterByUsername godoc
// @Summary RegisterByUsername
// @Description RegisterByUsername
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/register-by-username [post]
func (userHandler *UsersHandler) RegisterByUsername(context *gin.Context) {
	request := new(dto.RegisterUserByUsernameRequest)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	err = userHandler.services.RegisterByUsername(request)
	if err != nil {
		context.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	context.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, 1))
}

// RegisterLoginByMobileNumber godoc
// @Summary RegisterLoginByMobileNumber
// @Description RegisterLoginByMobileNumber
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/users/login-by-mobile [post]
func (userHandler *UsersHandler) RegisterLoginByMobileNumber(context *gin.Context) {
	request := new(dto.RegisterLoginByMobileNumberRequest)
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationErrors(nil, false, -1, err))
		return
	}
	token, err := userHandler.services.RegisterLoginByMobileNumber(request)
	if err != nil {
		context.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	context.JSON(http.StatusCreated, helpers.GenerateBaseResponse(token, true, 1))
}
