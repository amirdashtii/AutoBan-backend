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
		return errors.New("the phone number provided is not a valid Iranian number. Please ensure it starts with +98 and contains 10 digits")
	}
	return nil
}

// ValidatePassword checks if the password meets the required criteria
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("the password must be at least 8 characters long")
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return errors.New("the password must include at least one lowercase letter")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return errors.New("the password must include at least one uppercase letter")
	}
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return errors.New("the password must include at least one number")
	}
	return nil
}
