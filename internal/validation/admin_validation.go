package validation

import (
	"regexp"
	"time"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/go-playground/validator/v10"
)

// Custom validation for Iranian phone numbers
func AdminValidatePhoneNumber(fl validator.FieldLevel) bool {
	iranPhoneRegex := regexp.MustCompile(`^09\d{9}$`)
	return iranPhoneRegex.MatchString(fl.Field().String())
}

// Custom validation for password
func AdminValidatePassword(fl validator.FieldLevel) bool {
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

func AdminValidateDateTime(fl validator.FieldLevel) bool {
	dateTime, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return !dateTime.IsZero()
}

func AdminValidateEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}

func AdminValidateRole(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	return role == "User" || role == "Admin" || role == "SuperAdmin"
}

func AdminValidateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	return status == "Active" || status == "Deactivated" || status == "Deleted"
}

func AdminValidateUpdateProfileRequest(request dto.UpdateUserRequest) error {
	validate := validator.New()
	validate.RegisterValidation("phone", AdminValidatePhoneNumber)
	validate.RegisterValidation("datetime", AdminValidateDateTime)
	validate.RegisterValidation("email", AdminValidateEmail)
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "datetime":
				return errors.ErrInvalidBirthday
			case "email":
				return errors.ErrInvalidEmail
			case "phone":
				return errors.ErrInvalidPhoneNumber
			}
		}
	}
	return nil
}

func AdminValidateChangeUserRoleRequest(request dto.ChangeUserRoleRequest) error {
	validate := validator.New()
	validate.RegisterValidation("role", AdminValidateRole)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidRole
	}
	return nil
}

func AdminValidateChangeUserStatusRequest(request dto.ChangeUserStatusRequest) error {
	validate := validator.New()
	validate.RegisterValidation("status", AdminValidateStatus)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidStatus
	}
	return nil
}

func AdminValidateUpdatePasswordRequest(request dto.UpdatePasswordRequest) error {
	validate := validator.New()
	validate.RegisterValidation("password", AdminValidatePassword)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidPassword
	}
	return nil
}

func AdminValidateChangeUserPasswordRequest(request dto.ChangeUserPasswordRequest) error {
	validate := validator.New()
	validate.RegisterValidation("password", AdminValidatePassword)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidPassword
	}
	return nil
}
