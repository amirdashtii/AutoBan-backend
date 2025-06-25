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
