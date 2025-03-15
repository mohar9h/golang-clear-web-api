package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/mohar9h/golang-clear-web-api/api/middlewares"
	"github.com/mohar9h/golang-clear-web-api/api/routers"
	"github.com/mohar9h/golang-clear-web-api/api/validations"
	"github.com/mohar9h/golang-clear-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranMobileValidator, true)
		if err != nil {
			return
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			return
		}

	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("health")
		testRouter := v1.Group("test")
		routers.TestRouter(testRouter)
		routers.Health(health)
	}

	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		return
	}
}
