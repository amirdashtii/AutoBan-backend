package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"gorm.io/gorm"
)

type OilChangeRepository interface {
	CreateOilChange(ctx context.Context, oilChange *entity.OilChange) error
	GetOilChange(ctx context.Context, id uint, oilChange *entity.OilChange) error
	ListOilChanges(ctx context.Context, userVehicleID string, oilChanges *[]entity.OilChange) error
	UpdateOilChange(ctx context.Context, oilChange *entity.OilChange) error
	DeleteOilChange(ctx context.Context, oilChange *entity.OilChange) error
	GetLastOilChange(ctx context.Context, userVehicleID string, oilChange *entity.OilChange) error
}

type oilChangeRepository struct {
	db *gorm.DB
}

func NewOilChangeRepository() OilChangeRepository {
	db := database.ConnectDatabase()
	return &oilChangeRepository{db: db}
}

func (r *oilChangeRepository) CreateOilChange(ctx context.Context, oilChange *entity.OilChange) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Create(oilChange).Error
}

func (r *oilChangeRepository) GetOilChange(ctx context.Context, id uint, oilChange *entity.OilChange) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").First(oilChange, id).Error
}

func (r *oilChangeRepository) ListOilChanges(ctx context.Context, userVehicleID string, oilChanges *[]entity.OilChange) error {
	return r.db.WithContext(ctx).
		Preload("UserVehicle").
		Where("user_vehicle_id = ?", userVehicleID).
		Order("change_date DESC").
		Find(oilChanges).Error
}

func (r *oilChangeRepository) UpdateOilChange(ctx context.Context, oilChange *entity.OilChange) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Updates(oilChange).Error
}

func (r *oilChangeRepository) DeleteOilChange(ctx context.Context, oilChange *entity.OilChange) error {
	return r.db.WithContext(ctx).Delete(oilChange).Error
}

func (r *oilChangeRepository) GetLastOilChange(ctx context.Context, userVehicleID string, oilChange *entity.OilChange) error {
	return r.db.WithContext(ctx).
		Preload("UserVehicle").
		Where("user_vehicle_id = ?", userVehicleID).
		Order("change_date DESC").
		First(oilChange).Error
}
