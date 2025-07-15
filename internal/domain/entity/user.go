package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type RoleType int

const (
	UserRole RoleType = iota
	SuperAdminRole
	AdminRole
)

func (r RoleType) String() string {
	switch r {
	case SuperAdminRole:
		return "SuperAdmin"
	case AdminRole:
		return "Admin"
	case UserRole:
		return "User"
	default:
		return "Unknown"
	}
}

func ParseRoleType(s string) RoleType {
	switch strings.ToLower(s) {
	case "superadmin":
		return SuperAdminRole
	case "admin":
		return AdminRole
	case "user":
		return UserRole
	default:
		return UserRole
	}
}

type StatusType int

const (
	Active StatusType = iota
	Deactivated
	Deleted
)

func (s StatusType) String() string {
	switch s {
	case Active:
		return "Active"
	case Deactivated:
		return "Deactivated"
	case Deleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

func ParseStatusType(s string) StatusType {
	switch strings.ToLower(s) {
	case "active":
		return Active
	case "deactivated":
		return Deactivated
	case "deleted":
		return Deleted
	default:
		return Active
	}
}

type User struct {
	BaseEntity

	PhoneNumber string     `gorm:"index;unique;not null"`
	Password    string     `gorm:"not null"`
	FirstName   string
	LastName    string
	Birthday    time.Time
	Email       string     `gorm:"unique"`
	Status      StatusType
	Role        RoleType
}

// NewUser creates a new user with default values
func NewUser(phoneNumber, password string) *User {
	return &User{
		BaseEntity: BaseEntity{
			ID: uuid.New(),
		},
		PhoneNumber: phoneNumber,
		Password:    password,
		Status:      Active,
		Role:        UserRole,
	}
}

// UpdateProfile updates the user's profile information
func (u *User) UpdateProfile(firstName, lastName, email string) {
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
}

// ChangePassword changes the user's password
func (u *User) ChangePassword(newPassword string) {
	u.Password = newPassword
}

// Deactivate deactivates the user
func (u *User) Deactivate() {
	u.Status = Deactivated
}

// Delete marks the user as deleted
func (u *User) Delete() {
	u.Status = Deleted
}

// IsActive checks if the user is active
func (u *User) IsActive() bool {
	return u.Status == Active
}

// IsAdmin checks if the user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == AdminRole || u.Role == SuperAdminRole
}

// IsSuperAdmin checks if the user is a super admin
func (u *User) IsSuperAdmin() bool {
	return u.Role == SuperAdminRole
}
