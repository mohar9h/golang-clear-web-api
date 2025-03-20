package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"github.com/mohar9h/golang-clear-web-api/config"
	ipLimiter "github.com/mohar9h/golang-clear-web-api/pkg/limiter"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"time"
)

func OtpLimiter(config *config.Config) gin.HandlerFunc {
	limiter := ipLimiter.NewIpLimiter(rate.Every(time.Duration(config.Otp.Limiter)*time.Second), 1)

	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(getIP(c.Request.RemoteAddr))

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helpers.GenerateBaseResponseWithError(nil, false, -1, errors.New("not allowed")))
			c.Abort()
		}
		c.Next()
	}
}

func getIP(remoteAddr string) string {
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		return remoteAddr
	}
	return ip
}
