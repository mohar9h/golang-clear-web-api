package usecases

import (
	"github.com/go-redis/redis/v7"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
)

type OtpUseCase struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}
