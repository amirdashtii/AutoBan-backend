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
	FindByPhoneNumber(ctx context.Context, user *entity.User) error
	FindByID(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
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
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errors.ErrUserAlreadyExists
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *authRepository) FindByPhoneNumber(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Where("phone_number = ?", user.PhoneNumber).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrUserNotFound
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *authRepository) FindByID(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Where("id = ?", user.ID).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrUserNotFound
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *authRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Model(&entity.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return errors.ErrInternalServerError
	}
	return nil
}
