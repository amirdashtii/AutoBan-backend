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
	db := r.db.WithContext(ctx)
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	if serviceVisit.OilChangeID != nil {	
		oilChange := entity.OilChange{}
		oilChange.ID = *serviceVisit.OilChangeID
		err := tx.Delete(&oilChange).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if serviceVisit.OilFilterID != nil {
		oilFilter := entity.OilFilter{}
		oilFilter.ID = *serviceVisit.OilFilterID
		err := tx.Delete(&oilFilter).Error	
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	
	if err := tx.Delete(serviceVisit).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *ServiceVisitRepositoryImpl) GetLastServiceVisit(ctx context.Context, serviceVisit *entity.ServiceVisit) error {
	return r.db.WithContext(ctx).Preload("OilChange").Preload("OilFilter").Where("user_vehicle_id = ?", serviceVisit.UserVehicleID).Order("service_date DESC").First(serviceVisit).Error
}
