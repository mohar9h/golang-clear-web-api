package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenExpiresAt  int64  `json:"access_token_expires_at"`
	RefreshTokenExpiresAt int64  `json:"refresh_token_expires_at"`
}

type RegisterUserByUsernameRequest struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=30"`
	LastName  string `json:"last_name" binding:"required,min=3,max=30"`
	Email     string `json:"email" binding:"required,email,min=6,max=30"`
	Password  string `json:"password" binding:"required,min=6,max=30"`
	Username  string `json:"username" binding:"required,min=3,max=30"`
}

type RegisterLoginByMobileNumberRequest struct {
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
	Otp          string `json:"otp" binding:"required"`
}

type LoginByUserNameRequest struct {
	UserName string `json:"user_name" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=6"`
}
