package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetProfile(ctx context.Context, user *entity.User) error
	UpdateProfile(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	db := database.ConnectDatabase()
	return &userRepository{db: db}
}

func (r *userRepository) GetProfile(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).First(&user).Error
}

func (r *userRepository) UpdateProfile(ctx context.Context, user *entity.User) error {
	err := r.db.WithContext(ctx).Updates(&user).Error
	if err != nil {
		// Check if it's a unique constraint violation for email
		if err == gorm.ErrDuplicatedKey {
			return errors.ErrEmailAlreadyExists
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *userRepository) ChangePassword(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Model(&user).Update("password", user.Password).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Delete(&user).Error
}
