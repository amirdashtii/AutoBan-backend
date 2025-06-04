package repository

import (
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"

	"gorm.io/gorm"
)

// UserRepository handles user data operations
// UserRepository عملیات داده‌های کاربر را مدیریت می‌کند
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
// NewUserRepository یک نمونه جدید از UserRepository ایجاد می‌کند
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.ConnectDatabase(),
	}
}

// CreateUser creates a new user in the database
// CreateUser یک کاربر جدید در پایگاه داده ایجاد می‌کند
func (r *UserRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

// GetUserByID retrieves a user by their ID
// GetUserByID یک کاربر را بر اساس ID بازیابی می‌کند
func (r *UserRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user
// UpdateUser یک کاربر موجود را به‌روزرسانی می‌کند
func (r *UserRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

// DeleteUser deletes a user by their ID
// DeleteUser یک کاربر را بر اساس ID حذف می‌کند
func (r *UserRepository) DeleteUser(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

// GetUserByPhoneNumber retrieves a user by their phone number
// GetUserByPhoneNumber یک کاربر را بر اساس شماره تلفن بازیابی می‌کند
func (r *UserRepository) GetUserByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
