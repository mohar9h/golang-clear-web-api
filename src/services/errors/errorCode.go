package errors

const (
	// EmailExists UserExists
	EmailExists    string = "Email already exists"
	UsernameExists string = "Username already exists"
	// Unexpected JwtToken
	Unexpected      string = "Unexpected error"
	ClaimNotFound   string = "Claim not found"
	TokenRequired   string = "Token required"
	TokenExpired    string = "Token Expired"
	TokenInvalid    string = "Token invalid"
	Forbidden       string = "forbidden"
	UnAuthenticated string = "unauthenticated"
	// OtpExists OTP
	OtpExists   string = "otp already exists"
	OtpUsed     string = "otp used"
	OtpNotValid string = "otp not valid"
)
