package services

import (
	"github.com/mohar9h/golang-clear-web-api/api/dto"
	"github.com/mohar9h/golang-clear-web-api/common"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	config     *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(config *config.Config) *UserService {
	database := db.GetDBClient()
	logger := logging.NewLogger(config)
	return &UserService{
		logger:     logger,
		config:     config,
		database:   database,
		otpService: NewOTPService(config),
	}
}

func (service *UserService) SendOtp(request *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := service.otpService.SetOtpCode(request.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
