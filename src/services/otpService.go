package services

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/constants"
	"github.com/mohar9h/golang-clear-web-api/data/cache"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
	"github.com/mohar9h/golang-clear-web-api/services/errors"
	"time"
)

type OtpService struct {
	logger      logging.Logger
	config      *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOTPService(config *config.Config) *OtpService {
	logger := logging.NewLogger(config)
	redisClient := cache.GetRedisClient()
	return &OtpService{logger: logger, config: config, redisClient: redisClient}
}

func (service *OtpService) SetOtpCode(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	val := &OtpDto{Value: otp, Used: false}

	result, err := cache.Get[OtpDto](service.redisClient, key)
	if err == nil {
		if !result.Used {
			return &errors.ServiceErrors{EndUserMessage: errors.OtpExists}
		}
	}

	err = cache.Set(service.redisClient, key, val, service.config.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (service *OtpService) ValidateOtpCode(mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, mobileNumber)
	result, err := cache.Get[OtpDto](service.redisClient, key)
	if err != nil {
		return err
	} else if result.Used {
		return &errors.ServiceErrors{EndUserMessage: errors.OtpExists}
	} else if !result.Used && result.Value != otp {
		return &errors.ServiceErrors{EndUserMessage: errors.OtpNotValid}
	} else if !result.Used && result.Value == otp {
		result.Used = true
		err = cache.Set(service.redisClient, key, result, service.config.Otp.ExpireTime*time.Second)
	}

	return nil
}
