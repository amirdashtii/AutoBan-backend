package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"gorm.io/gorm"
)

type OilFilterRepository interface {
	CreateOilFilter(ctx context.Context, oilFilter *entity.OilFilter) error
	GetOilFilter(ctx context.Context, id uint64, oilFilter *entity.OilFilter) error
	ListOilFilters(ctx context.Context, userVehicleID uint64, oilFilters *[]entity.OilFilter) error
	UpdateOilFilter(ctx context.Context, oilFilter *entity.OilFilter) error
	DeleteOilFilter(ctx context.Context, oilFilter *entity.OilFilter) error
	GetLastOilFilter(ctx context.Context, userVehicleID uint64, oilFilter *entity.OilFilter) error
}

type OilFilterRepositoryImpl struct {
	db *gorm.DB
}

func NewOilFilterRepository() OilFilterRepository {
	db := database.ConnectDatabase()
	return &OilFilterRepositoryImpl{db: db}
}

func (r *OilFilterRepositoryImpl) CreateOilFilter(ctx context.Context, oilFilter *entity.OilFilter) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Create(oilFilter).Error
}

func (r *OilFilterRepositoryImpl) GetOilFilter(ctx context.Context, id uint64, oilFilter *entity.OilFilter) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Where("id = ?", id).First(oilFilter).Error
}

func (r *OilFilterRepositoryImpl) ListOilFilters(ctx context.Context, userVehicleID uint64, oilFilters *[]entity.OilFilter) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Where("user_vehicle_id = ?", userVehicleID).Order("change_date DESC").Find(oilFilters).Error
}

func (r *OilFilterRepositoryImpl) UpdateOilFilter(ctx context.Context, oilFilter *entity.OilFilter) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Save(oilFilter).Error
}

func (r *OilFilterRepositoryImpl) DeleteOilFilter(ctx context.Context, oilFilter *entity.OilFilter) error {
	return r.db.WithContext(ctx).Delete(oilFilter).Error
}

func (r *OilFilterRepositoryImpl) GetLastOilFilter(ctx context.Context, userVehicleID uint64, oilFilter *entity.OilFilter) error {
	return r.db.WithContext(ctx).Preload("UserVehicle").Where("user_vehicle_id = ?", userVehicleID).Order("change_date DESC").First(oilFilter).Error
}
