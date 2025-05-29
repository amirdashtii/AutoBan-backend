package validation

import (
	"errors"
	"regexp"
)

// ValidatePhoneNumber checks if the phone number is a valid Iranian number
func ValidatePhoneNumber(phoneNumber string) error {
	// Regular expression for Iranian phone numbers
	iranPhoneRegex := regexp.MustCompile(`^\+98\d{10}$`)
	if !iranPhoneRegex.MatchString(phoneNumber) {
		return errors.New("invalid Iranian phone number")
	}
	return nil
}

// ValidatePassword checks if the password meets the required criteria
func ValidatePassword(password string) error {
	// Regular expression for password with at least 8 characters, including uppercase, lowercase, and numbers
	passwordRegex := regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`)
	if !passwordRegex.MatchString(password) {
		return errors.New("password must be at least 8 characters long and include uppercase, lowercase, and numbers")
	}
	return nil
}
