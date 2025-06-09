package validation

import (
	"regexp"
	"time"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/go-playground/validator/v10"
)

// Custom validation for Iranian phone numbers
func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	iranPhoneRegex := regexp.MustCompile(`^09\d{9}$`)
	return iranPhoneRegex.MatchString(fl.Field().String())
}

// Custom validation for password
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false
	}
	if !regexp.MustCompile(`\d`).MatchString(password) {
		return false
	}
	return true
}

func validateDateTime(fl validator.FieldLevel) bool {
	dateTime, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return !dateTime.IsZero()
}

func validateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func ValidateUpdateProfileRequest(request dto.UpdateProfileRequest) error {
	validate := validator.New()
	validate.RegisterValidation("datetime", validateDateTime)
	validate.RegisterValidation("email", validateEmail)
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "datetime":
				return errors.ErrInvalidBirthday
			case "email":
				return errors.ErrInvalidEmail
			}
		}
	}
	return nil
}

func ValidateUpdatePasswordRequest(request dto.UpdatePasswordRequest) error {
	validate := validator.New()
	validate.RegisterValidation("password", ValidatePassword)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidPassword
	}
	return nil
}
