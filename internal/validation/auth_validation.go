package validation

import (
	"regexp"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"

	"github.com/go-playground/validator/v10"
)

// Custom validation for Iranian phone numbers
func iranPhone(fl validator.FieldLevel) bool {
	iranPhoneRegex := regexp.MustCompile(`^09\d{9}$`)
	return iranPhoneRegex.MatchString(fl.Field().String())
}

// Custom validation for password
func password(fl validator.FieldLevel) bool {
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

func ValidateRegisterRequest(request *dto.RegisterRequest) error {
	validate := validator.New()
	validate.RegisterValidation("iranphone", iranPhone)
	validate.RegisterValidation("password", password)

	err := validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func ValidateLoginRequest(request *dto.LoginRequest) error {
	validate := validator.New()
	validate.RegisterValidation("iranphone", iranPhone)
	validate.RegisterValidation("password", password)

	err := validate.Struct(request)
	if err != nil {
		return err
	}
	return nil
}

func ValidateVerifyPhoneRequest(request *dto.VerifyPhoneRequest) error {
	validate := validator.New()
	validate.RegisterValidation("iranphone", iranPhone)

	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "iranphone":
				return errors.ErrInvalidPhoneNumber
			}
		}
	}
	return nil
}

func ValidateResetPasswordRequest(request *dto.ResetPasswordRequest) error {
	validate := validator.New()
	validate.RegisterValidation("iranphone", iranPhone)
	validate.RegisterValidation("password", password)

	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "iranphone":
				return errors.ErrInvalidPhoneNumber
			case "password":
				return errors.ErrInvalidPassword
			case "len":
				return errors.ErrInvalidVerificationCode
			}
		}
	}
	return nil
}

func ValidateVerifyCodeRequest(request *dto.VerifyCodeRequest) error {
	validate := validator.New()
	validate.RegisterValidation("iranphone", iranPhone)

	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "iranphone":
				return errors.ErrInvalidPhoneNumber
			case "len":
				return errors.ErrInvalidVerificationCode
			}
		}
	}
	return nil
}
