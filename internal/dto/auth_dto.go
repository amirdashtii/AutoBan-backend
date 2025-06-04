package dto

// RegisterRequest represents the request body for user registration
// @Description User registration request
type RegisterRequest struct {
	// Iranian phone number in format 09XXXXXXXXX
	// @Example 09123456789
	PhoneNumber string `validate:"required,iranphone" json:"phone_number" example:"09123456789"`
	// Password must be at least 8 characters long and include uppercase, lowercase, and numbers
	// @Example Password123
	Password string `validate:"required,min=8,password" json:"password" example:"Password123"`
}

// LoginRequest represents the request body for user login
// @Description User login request
type LoginRequest struct {
	// Iranian phone number in format 09XXXXXXXXX
	// @Example 09123456789
	PhoneNumber string `validate:"required,iranphone" json:"phone_number" example:"09123456789"`
	// Password must be at least 8 characters long and include uppercase, lowercase, and numbers
	// @Example Password123
	Password string `validate:"required,min=8,password" json:"password" example:"Password123"`
}

// LoginResponse represents the response body for successful login
// @Description User login response containing access and refresh tokens
type TokenResponse struct {
	// JWT access token
	// @Example eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	AccessToken string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	// JWT refresh token
	// @Example eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// LogoutRequest represents the request body for user logout
// @Description User logout request
type LogoutRequest struct {
	// JWT refresh token
	// @Example eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	RefreshToken string `validate:"required" json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// RefreshTokenRequest represents the request body for token refresh
// @Description Token refresh request
type RefreshTokenRequest struct {
	// JWT refresh token
	// @Example eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
	RefreshToken string `validate:"required" json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// SessionResponse represents a user session in the response
// @Description User session information
type SessionResponse struct {
	DeviceID string `json:"device_id" example:"dev_1234567890"`
	LastUsed string `json:"last_used" example:"2024-03-15T14:30:00Z"`
	IsActive bool   `json:"is_active" example:"true"`
}

// GetSessionsResponse represents the response body for get sessions
// @Description Response containing list of user sessions
type GetSessionsResponse struct {
	Sessions []SessionResponse `json:"sessions"`
}
