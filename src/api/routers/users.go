package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/handlers"
	"github.com/mohar9h/golang-clear-web-api/api/middlewares"
	"github.com/mohar9h/golang-clear-web-api/config"
)

func User(router *gin.RouterGroup, config *config.Config) {
	handler := handlers.NewUsersHandler(config)

	router.POST("/send-otp", middlewares.OtpLimiter(config), handler.SendOtp)
	router.POST("/login-by-username", handler.LoginByUsername)
	router.POST("/register-by-username", handler.RegisterByUsername)
	router.POST("/login-by-mobile", handler.RegisterLoginByMobileNumber)
}
