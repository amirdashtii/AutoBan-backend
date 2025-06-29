package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"gorm.io/gorm"
)

type ServiceVisitRepository interface {
	CreateServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error
	GetServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error
	ListServiceVisits(ctx context.Context, userVehicleID string, serviceVisits *[]entity.ServiceVisit) error
	UpdateServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error
	DeleteServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error
	GetLastServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error
}

type ServiceVisitRepositoryImpl struct {
	db *gorm.DB
}

func NewServiceVisitRepository() ServiceVisitRepository {
	db := database.ConnectDatabase()
	return &ServiceVisitRepositoryImpl{db: db}
}

func (r *ServiceVisitRepositoryImpl) CreateServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Preload("OilChange").Preload("OilFilter").Create(serviceVisit).Error
}

func (r *ServiceVisitRepositoryImpl) GetServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Preload("OilChange").Preload("OilFilter").First(&serviceVisit).Error
}

func (r *ServiceVisitRepositoryImpl) ListServiceVisits(ctx context.Context, userVehicleID string, serviceVisits *[]entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Preload("OilChange").Preload("OilFilter").Where("user_vehicle_id = ?", userVehicleID).Order("service_date DESC").Find(serviceVisits).Error
}

func (r *ServiceVisitRepositoryImpl) UpdateServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Preload("OilChange").Preload("OilFilter").Save(serviceVisit).Error
}

func (r *ServiceVisitRepositoryImpl) DeleteServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Delete(serviceVisit).Error
}

func (r *ServiceVisitRepositoryImpl) GetLastServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Preload("OilChange").Preload("OilFilter").Where("user_vehicle_id = ?", serviceVisit.UserVehicleID).Order("service_date DESC").First(serviceVisit).Error
}
