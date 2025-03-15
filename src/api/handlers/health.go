package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helpers.GenerateBaseResponse("working!", true, 0))
	return
}
