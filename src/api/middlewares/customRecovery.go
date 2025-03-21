package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"net/http"
)

func ErrorHandler(context *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helpers.GenerateBaseResponseWithError(nil, false, -500, err.(error))
		context.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helpers.GenerateBaseResponseWithAnyError(nil, false, -500, err)
	context.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
