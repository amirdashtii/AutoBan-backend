package user

import (
	"AutoBan/internal/domain/entity"
	"AutoBan/internal/dto"
	"AutoBan/internal/errors"
	"AutoBan/internal/repository"
	"AutoBan/internal/validation"
	"AutoBan/pkg/logger"

	"golang.org/x/crypto/bcrypt"
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

	// هش کردن رمز عبور
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err, "Failed to hash password")
		return nil, errors.ErrInternalServerError
	}

	// ذخیره‌سازی کاربر جدید
	user := entity.User{PhoneNumber: userDTO.PhoneNumber, Password: string(hashedPassword)}
	err = u.userRepo.CreateUser(&user)
	if err != nil {
		logger.Error(err, "Failed to save user")
		return nil, errors.ErrInternalServerError
	}

	return &user, nil
}

// Login logs in a user
// تابع Login یک کاربر را وارد سیستم می‌کند

func (u *userUseCase) Login(userDTO dto.UserLoginDTO) (*entity.User, error) {
	err := validation.ValidatePhoneNumber(userDTO.PhoneNumber)
	if err != nil {
		logger.Error(err, "Invalid phone number")
		return nil, errors.ErrInvalidPhoneNumber
	}

	err = validation.ValidatePassword(userDTO.Password)
	if err != nil {
		logger.Error(err, "Invalid password")
		return nil, errors.ErrInvalidPassword
	}

	// بررسی وجود کاربر
	user, err := u.userRepo.GetUserByPhoneNumber(userDTO.PhoneNumber)
	if err != nil {
		logger.Error(err, "User not found")
		return nil, errors.ErrUserNotFound
	}

	// تطبیق رمز عبور
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))
	if err != nil {
		logger.Error(err, "Invalid password")
		return nil, errors.ErrInvalidPassword
	}

	return user, nil
}
