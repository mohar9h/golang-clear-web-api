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
	"github.com/mohar9h/golang-clear-web-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(config *config.Config) {
	r := gin.New()

	shouldReturn := RegisterValidator()
	if shouldReturn {
		return
	}
	r.Use(middlewares.Cors(config))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequest())

	RegisterRoutes(r)
	RegisterSwagger(r, config)

	err := r.Run(fmt.Sprintf(":%s", config.Server.Port))
	if err != nil {
		return
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("health")
		testRouter := v1.Group("test")
		routers.TestRouter(testRouter)
		routers.Health(health)
	}
}

func RegisterValidator() bool {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			return true
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			return true
		}

	}
	return false
}

func RegisterSwagger(routes *gin.Engine, config *config.Config) {
	docs.SwaggerInfo.Title = "Golang Clear Web API"
	docs.SwaggerInfo.Description = "This is a sample server for Golang Clear Web API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", config.Server.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
