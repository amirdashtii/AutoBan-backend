package user

import (
	"AutoBan/internal/domain/entity"
	"AutoBan/internal/dto"
	"AutoBan/internal/errors"
	"AutoBan/pkg/logger"
)

// UserUseCase interface defines the methods for user operations
// اینترفیس UserUseCase متدهای مربوط به عملیات‌های کاربر را تعریف می‌کند

type UserUseCase interface {
	Register(userDTO dto.UserRegisterDTO) (*entity.User, error)
	Login(userDTO dto.UserLoginDTO) (*entity.User, error)
}

// userUseCase struct implements the UserUseCase interface
// ساختار userUseCase اینترفیس UserUseCase را پیاده‌سازی می‌کند

type userUseCase struct {
	// اینجا می‌توانید وابستگی‌های مورد نیاز را اضافه کنید
}

// NewUserUseCase creates a new instance of userUseCase
// تابع NewUserUseCase یک نمونه جدید از userUseCase ایجاد می‌کند

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

// Register registers a new user
// تابع Register یک کاربر جدید را ثبت‌نام می‌کند

func (u *userUseCase) Register(userDTO dto.UserRegisterDTO) (*entity.User, error) {
	if userDTO.PhoneNumber == "" || userDTO.Password == "" {
		logger.Error(errors.ErrPhoneNumberOrPasswordRequired, "Phone number or password is required")
		return nil, errors.ErrPhoneNumberOrPasswordRequired
	}
	// اینجا می‌توانید منطق ثبت‌نام را اضافه کنید
	return &entity.User{PhoneNumber: userDTO.PhoneNumber, Password: userDTO.Password}, nil
}

// Login logs in a user
// تابع Login یک کاربر را وارد سیستم می‌کند

func (u *userUseCase) Login(userDTO dto.UserLoginDTO) (*entity.User, error) {
	if userDTO.PhoneNumber == "" || userDTO.Password == "" {
		logger.Error(errors.ErrPhoneNumberOrPasswordRequired, "Phone number or password is required")
		return nil, errors.ErrPhoneNumberOrPasswordRequired
	}
	// اینجا می‌توانید منطق ورود را اضافه کنید
	return &entity.User{PhoneNumber: userDTO.PhoneNumber, Password: userDTO.Password}, nil
}
