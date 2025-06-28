package dto

import "github.com/google/uuid"

// GetProfileResponse represents the response for getting user profile
// @Description User profile information response
type GetProfileResponse struct {
	// User's ID
	ID uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	// User's first name (optional)
	FirstName string `json:"first_name" example:"John"`
	// User's last name (optional)
	LastName string `json:"last_name" example:"Doe"`
	// User's email address (optional)
	Email string `json:"email" example:"john.doe@example.com"`
	// User's birthday (optional)
	Birthday string `json:"birthday" example:"1990-01-01"`
}

// UpdateProfileRequest represents the request for updating user profile
// @Description User profile update request
type UpdateProfileRequest struct {
	// User's first name (optional)
	FirstName *string `json:"first_name" example:"John"`
	// User's last name (optional)
	LastName *string `json:"last_name" example:"Doe"`
	// User's email address (optional)
	Email *string `validate:"omitempty,email" json:"email" example:"john.doe@example.com"`
	// User's birthday (optional)
	Birthday *string `validate:"omitempty,datetime" json:"birthday" example:"1990-01-01"`
}

// UpdateProfileResponse represents the response for updating user profile
// @Description User profile update response
type UpdateProfileResponse struct {
	// User's ID
	ID uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	// User's first name (optional)
	FirstName string `json:"first_name" example:"John"`
	// User's last name (optional)
	LastName string `json:"last_name" example:"Doe"`
	// User's email address (optional)
	Email string `json:"email" example:"john.doe@example.com"`
	// User's birthday (optional)
	Birthday string `json:"birthday" example:"1990-01-01"`
}

// UpdatePasswordRequest represents the request for updating user password
// @Description User password update request
type UpdatePasswordRequest struct {
	Password string `validate:"password" json:"password" example:"password"`
}
