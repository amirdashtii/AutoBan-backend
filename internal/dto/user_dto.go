package dto

// UserRegisterDTO is used for user registration
// UserRegisterDTO برای ثبت‌نام کاربر استفاده می‌شود

type UserRegisterDTO struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// UserLoginDTO is used for user login
// UserLoginDTO برای ورود کاربر استفاده می‌شود

type UserLoginDTO struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
