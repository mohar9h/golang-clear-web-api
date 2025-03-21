package usecases

import (
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
)

type UserUseCase struct {
	logger       logging.Logger
	cfg          *config.Config
	otpUseCase   *OtpUseCase
	tokenUseCase *TokenUseCase
	repository   repository.UserRepository
}
