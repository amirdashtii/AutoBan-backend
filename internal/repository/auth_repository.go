package repository

import (
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *entity.User) error
	FindByPhoneNumber(phoneNumber string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	db := database.ConnectDatabase()
	return &authRepository{db: db}
}

func (r *authRepository) Register(user *entity.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return errors.ErrUserAlreadyExists
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *authRepository) FindByPhoneNumber(phoneNumber string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrInternalServerError
	}
	return &user, nil
}

func (r *authRepository) FindByID(id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrInternalServerError
	}
	return &user, nil
}
