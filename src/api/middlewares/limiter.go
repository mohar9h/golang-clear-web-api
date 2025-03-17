package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
)

func LimitByRequest() gin.HandlerFunc {
	limit := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(limit, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helpers.GenerateBaseResponseWithError(nil, false, -100, err))
			return
		}
	}
}
