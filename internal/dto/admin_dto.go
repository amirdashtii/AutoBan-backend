package dto

// User represents a user in the system
// @Description User information model
type User struct {
	// Unique identifier of the user
	// @Example 123e4567-e89b-12d3-a456-426614174000
	ID string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	// User's first name
	// @Example John
	FirstName string `json:"first_name" example:"John"`
	// User's last name
	// @Example Doe
	LastName string `json:"last_name" example:"Doe"`
	// User's email address
	// @Example john.doe@example.com
	Email string `json:"email" example:"john.doe@example.com"`
	// User's phone number
	// @Example 09123456789
	Phone string `json:"phone" example:"09123456789"`
	// User's role (User, Admin, SuperAdmin)
	// @Example Admin
	Role string `json:"role" example:"Admin"`
	// User's status (Active, Deactivated, Deleted)
	// @Example Active
	Status string `json:"status" example:"Active"`
	// User's birthday in YYYY-MM-DD format
	// @Example 1990-01-01
	Birthday string `json:"birthday" example:"1990-01-01"`
}

// ListUsersResponse represents the response for listing users
// @Description Response containing list of users
type ListUsersResponse struct {
	// List of users
	Users []User `json:"users"`
}

// GetUserByIdResponse represents the response for getting a single user
// @Description Response containing user details
type GetUserByIdResponse struct {
	// User information
	User User `json:"user"`
}

// UpdateUserRequest represents the request for updating user information
// @Description Request to update user details
type UpdateUserRequest struct {
	// User's first name
	// @Example John
	FirstName string `json:"first_name" example:"John"`
	// User's last name
	// @Example Doe
	LastName string `json:"last_name" example:"Doe"`
	// User's email address
	// @Example john.doe@example.com
	Email string `validate:"email" json:"email" example:"john.doe@example.com"`
	// User's phone number
	// @Example 09123456789
	Phone string `validate:"phone" json:"phone" example:"09123456789"`
	// User's birthday in YYYY-MM-DD format
	// @Example 1990-01-01
	Birthday string `validate:"datetime" json:"birthday" example:"1990-01-01"`
}

// ChangeUserRoleRequest represents the request for changing user role
// @Description Request to change user role
type ChangeUserRoleRequest struct {
	// New role for the user (User, Admin, SuperAdmin)
	// @Example Admin
	Role string `validate:"role" json:"role" example:"Admin" enums:"User,Admin,SuperAdmin"`
}

// ChangeUserStatusRequest represents the request for changing user status
// @Description Request to change user status
type ChangeUserStatusRequest struct {
	// New status for the user (Active, Deactivated, Deleted)
	// @Example Active
	Status string `validate:"status" json:"status" example:"Active" enums:"Active,Deactivated,Deleted"`
}

// ChangeUserPasswordRequest represents the request for changing user password
// @Description Request to change user password
type ChangeUserPasswordRequest struct {
	// New password for the user (minimum 8 characters)
	// @Example NewPassword123
	NewPassword string `validate:"password" json:"new_password" example:"NewPassword123" minLength:"8"`
}
