package services

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mohar9h/golang-clear-web-api/api/dto"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
	"github.com/mohar9h/golang-clear-web-api/services/errors"
	"time"
)

type TokenService struct {
	logger logging.Logger
	config *config.Config
}

type tokenDto struct {
	UserId       int      `json:"user_id"`
	FirstName    string   `json:"first_name"`
	LastName     string   `json:"last_name"`
	Username     string   `json:"username"`
	MobileNumber string   `json:"mobile_number"`
	Email        string   `json:"email"`
	Roles        []string `json:"roles"`
}

func NewTokenService(config *config.Config) *TokenService {
	logger := logging.NewLogger(config)
	return &TokenService{
		logger: logger,
		config: config,
	}
}

func (service *TokenService) CreateToken(token *tokenDto) (*dto.TokenDetail, error) {
	accessToken := &dto.TokenDetail{}
	accessToken.AccessTokenExpiresAt = time.Now().Add(service.config.Jwt.AccessTokenExpireDuration * time.Minute).Unix()
	accessToken.RefreshTokenExpiresAt = time.Now().Add(service.config.Jwt.RefreshTokenExpireDuration * time.Minute).Unix()

	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["user_id"] = token.UserId
	accessTokenClaims["first_name"] = token.FirstName
	accessTokenClaims["last_name"] = token.LastName
	accessTokenClaims["username"] = token.Username
	accessTokenClaims["mobile_number"] = token.MobileNumber
	accessTokenClaims["email"] = token.Email
	accessTokenClaims["roles"] = token.Roles
	accessTokenClaims["expire_time"] = accessToken.AccessTokenExpiresAt

	accessTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error
	accessToken.AccessToken, err = accessTokenJwt.SignedString([]byte(service.config.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["user_id"] = token.UserId
	refreshTokenClaims["expire_time"] = accessToken.RefreshTokenExpiresAt

	refreshTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	accessToken.RefreshToken, err = refreshTokenJwt.SignedString([]byte(service.config.Jwt.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (service *TokenService) ParseToken(token string) (*jwt.Token, error) {
	accessToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &errors.ServiceErrors{EndUserMessage: errors.Unexpected}
		}
		return []byte(service.config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (service *TokenService) ParseTokenWithClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = make(map[string]interface{})

	parseToken, err := service.ParseToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if ok && parseToken.Valid {
		for key, value := range claims {
			claimMap[key] = value
		}
		return claimMap, nil
	}
	return nil, &errors.ServiceErrors{EndUserMessage: errors.ClaimNotFound}
}
