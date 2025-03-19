package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/handlers"
	"github.com/mohar9h/golang-clear-web-api/config"
)

func User(router *gin.RouterGroup, config *config.Config) {
	handler := handlers.NewUsersHandler(config)

	router.POST("/send-otp", handler.SendOtp)
}
