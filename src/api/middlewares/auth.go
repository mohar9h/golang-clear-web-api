package middlewares

import (
	errors2 "errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/constants"
	responses "github.com/mohar9h/golang-clear-web-api/domains/response"
	"github.com/mohar9h/golang-clear-web-api/services"
	"github.com/mohar9h/golang-clear-web-api/services/errors"
	"net/http"
	"strings"
)

func Authentication(config *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(config)

	return func(context *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := context.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" {
			err = &errors.ServiceErrors{EndUserMessage: errors.TokenRequired}
		} else {
			claimMap, err = tokenService.ParseTokenWithClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &errors.ServiceErrors{EndUserMessage: errors.TokenExpired}
				default:
					err = &errors.ServiceErrors{EndUserMessage: errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized,
				responses.Error(helpers.AuthError, errors2.New(errors.UnAuthenticated)))
			return
		}

		context.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		context.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		context.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		context.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		context.Set(constants.EmailKey, claimMap[constants.EmailKey])
		context.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		context.Set(constants.RolesKey, claimMap[constants.RolesKey])
		context.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		context.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(context *gin.Context) {
		if len(context.Keys) == 0 {
			context.AbortWithStatusJSON(http.StatusForbidden, responses.Error(helpers.ForbiddenError, errors2.New(errors.Forbidden)))
			return
		}
		rolesValue := context.Keys[constants.RolesKey]
		fmt.Println(rolesValue)
		if rolesValue == nil {
			context.AbortWithStatusJSON(http.StatusForbidden,
				responses.Error(helpers.ForbiddenError, errors2.New(errors.Forbidden)))
			return
		}
		roles := rolesValue.([]interface{})
		value := map[string]int{}
		for _, item := range roles {
			value[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := value[item]; ok {
				context.Next()
				return
			}
		}
		context.AbortWithStatusJSON(http.StatusForbidden,
			responses.Error(helpers.ForbiddenError, errors2.New(errors.Forbidden)))
	}
}
