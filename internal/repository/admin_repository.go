package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"gorm.io/gorm"
)

type AdminRepository interface {
	ListUsers(ctx context.Context, users *[]entity.User) error
	GetUserById(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository() AdminRepository {
	db := database.ConnectDatabase()
	return &adminRepository{db: db}
}

func (r *adminRepository) ListUsers(ctx context.Context, users *[]entity.User) error {
	if err := r.db.WithContext(ctx).Find(users).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetUserById(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) DeleteUser(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
