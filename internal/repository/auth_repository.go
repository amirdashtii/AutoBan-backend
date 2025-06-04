package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(ctx context.Context, user *entity.User) error
	FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository() AuthRepository {
	db := database.ConnectDatabase()
	return &authRepository{db: db}
}

func (r *authRepository) Register(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return errors.ErrUserAlreadyExists
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *authRepository) FindByPhoneNumber(ctx context.Context, phoneNumber string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("phone_number = ?", phoneNumber).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrInternalServerError
	}
	return &user, nil
}

func (r *authRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.ErrInternalServerError
	}
	return &user, nil
}
