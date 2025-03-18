package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/handlers"
)

// Health is a function to handle health check
// @Summary Show the health status of the service
// @Description Show the health status of the service
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /v1/health [get]
// @Security Bearer
func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
}
