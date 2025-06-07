package dto

// GetProfileResponse represents the response for getting user profile
// @Description User profile information response
type GetProfileResponse struct {
	// User's first name (optional)
	// @Example John
	FirstName *string `json:"first_name" example:"John"`
	// User's last name (optional)
	// @Example Doe
	LastName *string `json:"last_name" example:"Doe"`
	// User's email address (optional)
	// @Example john.doe@example.com
	Email *string `json:"email" example:"john.doe@example.com"`
	// User's birthday (optional)
	// @Example 1990-01-01
	Birthday *string `json:"birthday" example:"1990-01-01"`
}

// UpdateProfileRequest represents the request for updating user profile
// @Description User profile update request
type UpdateProfileRequest struct {
	// User's first name (optional)
	// @Example John
	FirstName *string `json:"first_name" example:"John"`
	// User's last name (optional)
	// @Example Doe
	LastName *string `json:"last_name" example:"Doe"`
	// User's email address (optional)
	// @Example john.doe@example.com
	Email *string `json:"email" example:"john.doe@example.com"`
	// User's birthday (optional)
	// @Example 1990-01-01
	Birthday *string `json:"birthday" example:"1990-01-01"`
}

// UpdateProfileResponse represents the response for updating user profile
// @Description User profile update response
type UpdateProfileResponse struct {
	// User's first name (optional)
	// @Example John
	FirstName *string `json:"first_name" example:"John"`
	// User's last name (optional)
	// @Example Doe
	LastName *string `json:"last_name" example:"Doe"`
	// User's email address (optional)
	// @Example john.doe@example.com
	Email *string `json:"email" example:"john.doe@example.com"`
	// User's birthday (optional)
	// @Example 1990-01-01
	Birthday *string `json:"birthday" example:"1990-01-01"`
}

// UpdatePasswordRequest represents the request for updating user password
// @Description User password update request
type UpdatePasswordRequest struct {
	Password string `json:"password" example:"password"`
}


