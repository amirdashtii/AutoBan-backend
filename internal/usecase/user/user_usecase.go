package user

import (
	"AutoBan/internal/domain/entity"
	"AutoBan/internal/dto"
	"AutoBan/internal/errors"
	"AutoBan/internal/repository"
	"AutoBan/internal/validation"
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
	userRepo repository.UserRepository
}

// NewUserUseCase creates a new instance of userUseCase
// تابع NewUserUseCase یک نمونه جدید از userUseCase ایجاد می‌کند

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{userRepo: userRepo}
}

// Register registers a new user
// تابع Register یک کاربر جدید را ثبت‌نام می‌کند

func (u *userUseCase) Register(userDTO dto.UserRegisterDTO) (*entity.User, error) {

	err := validation.ValidatePhoneNumber(userDTO.PhoneNumber)
	if err != nil {
		logger.Error(err, "Invalid phone number")
		return nil, errors.ErrInvalidPhoneNumber
	}

	// اعتبارسنجی رمز عبور
	err = validation.ValidatePassword(userDTO.Password)
	if err != nil {
		logger.Error(err, "Invalid password")
		return nil, errors.ErrInvalidPassword
	}

	// ذخیره‌سازی کاربر جدید
	user := entity.User{PhoneNumber: userDTO.PhoneNumber, Password: userDTO.Password}
	err = u.userRepo.CreateUser(&user)
	if err != nil {
		logger.Error(err, "Failed to save user")
		return nil, err
	}

	return &user, nil
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
