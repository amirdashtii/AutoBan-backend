package entity

import "time"

// Session represents a user session in Redis
type Session struct {
	UserID       string    `json:"user_id"`
	DeviceID     string    `json:"device_id"`
	RefreshToken string    `json:"refresh_token"`
	LastUsed     time.Time `json:"last_used"`
	IsActive     bool      `json:"is_active"`
}

// NewSession creates a new session
func NewSession(userID, deviceID, refreshToken string) *Session {
	return &Session{
		UserID:       userID,
		DeviceID:     deviceID,
		RefreshToken: refreshToken,
		LastUsed:     time.Now(),
		IsActive:     true,
	}
}
