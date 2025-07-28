package dto

// RegisterRequest represents the request body for user registration
// @Description User registration request
type RegisterRequest struct {
	// Iranian phone number in format 09XXXXXXXXX
	PhoneNumber string `validate:"required,iranphone" json:"phone_number" example:"09123456789"`
	// Password must be at least 8 characters long and include uppercase, lowercase, and numbers
	Password string `validate:"required,min=8,password" json:"password" example:"Password123"`
}

// LoginRequest represents the request body for user login
// @Description User login request
type LoginRequest struct {
	// Iranian phone number in format 09XXXXXXXXX
	PhoneNumber string `validate:"required,iranphone" json:"phone_number" example:"09123456789"`
	// Password must be at least 8 characters long and include uppercase, lowercase, and numbers
	Password string `validate:"required,min=8,password" json:"password" example:"Password123"`
}

// LoginResponse represents the response body for successful login
// @Description User login response containing access and refresh tokens
type TokenResponse struct {
	// JWT access token
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// JWT refresh token
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// LogoutRequest represents the request body for user logout
// @Description User logout request
type LogoutRequest struct {
	// JWT refresh token
	RefreshToken string `validate:"required" json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// RefreshTokenRequest represents the request body for token refresh
// @Description Token refresh request
type RefreshTokenRequest struct {
	// JWT refresh token
	RefreshToken string `validate:"required" json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// SessionResponse represents a user session in the response
// @Description User session information
type SessionResponse struct {
	// Device ID
	DeviceID string `json:"device_id" example:"dev_1234567890"`
	// Last used time
	LastUsed string `json:"last_used" example:"2024-03-15T14:30:00Z"`
	// Is active
	IsActive bool `json:"is_active" example:"true"`
}

// GetSessionsResponse represents the response body for get sessions
// @Description Response containing list of user sessions
type GetSessionsResponse struct {
	// List of user sessions
	Sessions []SessionResponse `json:"sessions"`
}

type VerifyPhoneRequest struct {
	PhoneNumber string `validate:"required,iranphone" json:"phone_number" example:"09123456789"`
}

type VerifyCodeRequest struct {
	PhoneNumber string `validate:"required,iranphone" json:"phone_number" example:"09123456789"`
	Code        string `validate:"required,len=6" json:"code" example:"123456"`
}

type SmsIrRequest struct {
	Mobile     string `json:"mobile" validate:"required,iranphone"`
	TemplateId string `json:"templateId" validate:"required"`
	Parameters []struct {
		Name  string `json:"name" validate:"required"`
		Value string `json:"value" validate:"required"`
	} `json:"parameters" validate:"required,dive"`
}

type SmsIrResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		MessageID int     `json:"messageId"`
		Cost      float64 `json:"cost"`
	} `json:"data"`
}
