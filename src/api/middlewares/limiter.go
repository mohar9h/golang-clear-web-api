package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
import "github.com/didip/tollbooth"

func LimitByRequest() gin.HandlerFunc {
	limit := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(limit, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": err,
			})
			return
		}
	}
}
