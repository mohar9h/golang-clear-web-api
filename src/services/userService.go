package services

import (
	"github.com/mohar9h/golang-clear-web-api/api/dto"
	"github.com/mohar9h/golang-clear-web-api/common"
	"github.com/mohar9h/golang-clear-web-api/config"
	"github.com/mohar9h/golang-clear-web-api/constants"
	"github.com/mohar9h/golang-clear-web-api/data/db"
	"github.com/mohar9h/golang-clear-web-api/data/models"
	"github.com/mohar9h/golang-clear-web-api/pkg/logging"
	"github.com/mohar9h/golang-clear-web-api/services/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	config       *config.Config
	otpService   *OtpService
	tokenService *TokenService
	database     *gorm.DB
}

func NewUserService(config *config.Config) *UserService {
	database := db.GetDBClient()
	logger := logging.NewLogger(config)
	return &UserService{
		logger:       logger,
		config:       config,
		database:     database,
		otpService:   NewOTPService(config),
		tokenService: NewTokenService(config),
	}
}

// LoginByUsername Login by username
func (userService *UserService) LoginByUsername(request *dto.LoginByUserNameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := userService.database.
		Model(&models.User{}).
		Where("username = ?", request.UserName).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, err
	}
	jwtTokenDto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			jwtTokenDto.Roles = append(jwtTokenDto.Roles, ur.Role.Name)
		}
	}

	token, err := userService.tokenService.CreateToken(&jwtTokenDto)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (userService *UserService) RegisterByUsername(request *dto.RegisterUserByUsernameRequest) error {
	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
		Username:  request.Username,
	}

	exists, err := userService.existsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return &errors.ServiceErrors{EndUserMessage: errors.EmailExists}
	}
	exists, err = userService.existsByUsername(user.Username)
	if err != nil {
		return err
	}
	if exists {
		return &errors.ServiceErrors{EndUserMessage: errors.UsernameExists}
	}

	bytePassword := []byte(request.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		userService.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	user.Password = string(hashedPassword)
	roleId, err := userService.getDefaultRole()
	if err != nil {
		userService.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}
	transaction := userService.database.Begin()
	err = transaction.Create(&user).Error
	if err != nil {
		transaction.Rollback()
		userService.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	err = transaction.Create(&models.UserRole{RoleId: roleId, UserId: roleId}).Error
	if err != nil {
		transaction.Rollback()
		userService.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	transaction.Commit()
	return nil
}

func (userService *UserService) RegisterLoginByMobileNumber(request *dto.RegisterLoginByMobileNumberRequest) (*dto.TokenDetail, error) {
	err := userService.otpService.ValidateOtpCode(request.MobileNumber, request.Otp)
	if err != nil {
		return nil, err
	}
	exists, err := userService.existsByMobile(request.MobileNumber)
	if err != nil {
		return nil, err
	}

	u := models.User{MobileNumber: request.MobileNumber, Username: request.MobileNumber}

	if exists {
		var user models.User
		err = userService.database.
			Model(&models.User{}).
			Where("username = ?", u.Username).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			Find(&user).Error
		if err != nil {
			return nil, err
		}
		jwtTokenDto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
			Email: user.Email, MobileNumber: user.MobileNumber}

		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				jwtTokenDto.Roles = append(jwtTokenDto.Roles, ur.Role.Name)
			}
		}

		token, err := userService.tokenService.CreateToken(&jwtTokenDto)
		if err != nil {
			return nil, err
		}
		return token, nil

	}

	bcryptPassword := []byte(common.GeneratePassword())
	hashPassword, err := bcrypt.GenerateFromPassword(bcryptPassword, bcrypt.DefaultCost)
	if err != nil {
		userService.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return nil, err
	}
	u.Password = string(hashPassword)
	roleId, err := userService.getDefaultRole()
	if err != nil {
		userService.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return nil, err
	}

	transaction := userService.database.Begin()
	err = transaction.Create(&u).Error
	if err != nil {
		transaction.Rollback()
		userService.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	err = transaction.Create(&models.UserRole{RoleId: roleId, UserId: u.Id}).Error
	if err != nil {
		transaction.Rollback()
		userService.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}
	transaction.Commit()

	var user models.User
	err = userService.database.
		Model(&models.User{}).
		Where("username = ?", u.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}
	jwtTokenDto := tokenDto{UserId: user.Id, FirstName: user.FirstName, LastName: user.LastName,
		Email: user.Email, MobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			jwtTokenDto.Roles = append(jwtTokenDto.Roles, ur.Role.Name)
		}
	}

	token, err := userService.tokenService.CreateToken(&jwtTokenDto)
	if err != nil {
		return nil, err
	}
	return token, nil

}

func (userService *UserService) SendOtp(request *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := userService.otpService.SetOtpCode(request.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}

func (userService *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := userService.database.Model(&models.User{}).
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		userService.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (userService *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := userService.database.Model(&models.User{}).
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		userService.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (userService *UserService) existsByMobile(mobile string) (bool, error) {
	var exists bool
	if err := userService.database.Model(&models.User{}).
		Where("mobile_number = ?", mobile).
		Find(&exists).
		Error; err != nil {
		userService.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (userService *UserService) getDefaultRole() (roleId int, err error) {

	if err = userService.database.Model(&models.Role{}).
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
