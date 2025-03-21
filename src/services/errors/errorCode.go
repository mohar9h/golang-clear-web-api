package errors

const (
	// EmailExists UserExists
	EmailExists    string = "Email already exists"
	UsernameExists string = "Username already exists"
	// Unexpected JwtToken
	Unexpected    string = "Unexpected error"
	ClaimNotFound string = "Claim not found"
	// OtpExists OTP
	OtpExists   string = "otp already exists"
	OtpUsed     string = "otp used"
	OtpNotValid string = "otp not valid"
)
