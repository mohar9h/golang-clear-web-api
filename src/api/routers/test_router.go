package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/:id", h.TestById)
	r.GET("/header", h.HeaderBinder2)
}
