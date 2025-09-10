package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	ListVehicleTypes(ctx context.Context, vehicleTypes *[]entity.VehicleType) error
	GetVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error
	CreateVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error
	UpdateVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error
	DeleteVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error

	ListBrands(ctx context.Context, brands *[]entity.VehicleBrand) error
	GetBrand(ctx context.Context, brand *entity.VehicleBrand) error
	ListBrandsByType(ctx context.Context, brands *[]entity.VehicleBrand, typeID uint64) error
	CreateBrand(ctx context.Context, brand *entity.VehicleBrand) error
	UpdateBrand(ctx context.Context, brand *entity.VehicleBrand) error
	DeleteBrand(ctx context.Context, brand *entity.VehicleBrand) error

	ListModels(ctx context.Context, models *[]entity.VehicleModel) error
	GetModel(ctx context.Context, model *entity.VehicleModel) error
	ListModelsByBrand(ctx context.Context, models *[]entity.VehicleModel, brandID uint64) error
	CreateModel(ctx context.Context, model *entity.VehicleModel) error
	UpdateModel(ctx context.Context, model *entity.VehicleModel) error
	DeleteModel(ctx context.Context, model *entity.VehicleModel) error

	ListGenerations(ctx context.Context, generations *[]entity.VehicleGeneration) error
	GetGeneration(ctx context.Context, generation *entity.VehicleGeneration) error
	ListGenerationsByModel(ctx context.Context, generations *[]entity.VehicleGeneration, modelID uint64) error
	CreateGeneration(ctx context.Context, generation *entity.VehicleGeneration) error
	UpdateGeneration(ctx context.Context, generation *entity.VehicleGeneration) error
	DeleteGeneration(ctx context.Context, generation *entity.VehicleGeneration) error

	CreateUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error
	ListUserVehicles(ctx context.Context, userID uuid.UUID, userVehicles *[]entity.UserVehicle) error
	GetUserVehicle(ctx context.Context, userID uuid.UUID, vehicleId uint64, userVehicle *entity.UserVehicle) error
	UpdateUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error
	DeleteUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error

	// Complete hierarchy methods
	GetCompleteVehicleHierarchy(ctx context.Context, vehicleTypes *[]entity.VehicleType) error
}

type vehicleRepository struct {
	db    *gorm.DB
	cache CacheRepository
}

func NewVehicleRepository() VehicleRepository {
	db := database.ConnectDatabase()
	cache := NewCacheRepository()
	return &vehicleRepository{
		db:    db,
		cache: cache,
	}
}

// Vehicle Types
func (r *vehicleRepository) ListVehicleTypes(ctx context.Context, vehicleTypes *[]entity.VehicleType) error {
	return r.db.WithContext(ctx).Find(vehicleTypes).Error
}

func (r *vehicleRepository) GetVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error {
	return r.db.WithContext(ctx).Where("id = ?", vehicleType.ID).First(vehicleType).Error
}

func (r *vehicleRepository) CreateVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error {
	return r.db.WithContext(ctx).Create(vehicleType).Error
}

func (r *vehicleRepository) UpdateVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error {
	return r.db.WithContext(ctx).Model(vehicleType).Updates(vehicleType).Error
}

func (r *vehicleRepository) DeleteVehicleType(ctx context.Context, vehicleType *entity.VehicleType) error {
	return r.db.WithContext(ctx).Delete(vehicleType).Error
}

// Brands
func (r *vehicleRepository) ListBrands(ctx context.Context, brands *[]entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Find(brands).Error
}

func (r *vehicleRepository) ListBrandsByType(ctx context.Context, brands *[]entity.VehicleBrand, typeID uint64) error {
	return r.db.WithContext(ctx).Where("vehicle_type_id = ?", typeID).Find(brands).Error
}

func (r *vehicleRepository) GetBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Where("id = ?", brand.ID).First(brand).Error
}

func (r *vehicleRepository) CreateBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Create(brand).Error
}

func (r *vehicleRepository) UpdateBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Model(brand).Updates(brand).Error
}

func (r *vehicleRepository) DeleteBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Delete(brand).Error
}

// Models
func (r *vehicleRepository) ListModels(ctx context.Context, models *[]entity.VehicleModel) error {
	return r.db.WithContext(ctx).Find(models).Error
}

func (r *vehicleRepository) ListModelsByBrand(ctx context.Context, models *[]entity.VehicleModel, brandID uint64) error {
	return r.db.WithContext(ctx).Where("brand_id = ?", brandID).Find(models).Error
}

func (r *vehicleRepository) GetModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Where("id = ?", model.ID).First(model).Error
}

func (r *vehicleRepository) CreateModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *vehicleRepository) UpdateModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Model(model).Updates(model).Error
}

func (r *vehicleRepository) DeleteModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Delete(model).Error
}

// Generations
func (r *vehicleRepository) ListGenerations(ctx context.Context, generations *[]entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Find(generations).Error
}

func (r *vehicleRepository) GetGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Where("id = ?", generation.ID).First(generation).Error
}

func (r *vehicleRepository) ListGenerationsByModel(ctx context.Context, generations *[]entity.VehicleGeneration, modelID uint64) error {
	return r.db.WithContext(ctx).Where("model_id = ?", modelID).Find(generations).Error
}
func (r *vehicleRepository) CreateGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Create(generation).Error
}

func (r *vehicleRepository) UpdateGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Model(generation).Updates(generation).Error
}

func (r *vehicleRepository) DeleteGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Delete(generation).Error
}

// User Vehicles with caching
func (r *vehicleRepository) CreateUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error {
	err := r.db.WithContext(ctx).Create(userVehicle).Error
	if err != nil {
		return err
	}

	// Invalidate user vehicles cache
	userCacheKey := BuildCacheKey(CacheKeyUserVehicles, userVehicle.UserID.String())
	r.cache.Delete(ctx, userCacheKey)
	
	return nil
}

func (r *vehicleRepository) ListUserVehicles(ctx context.Context, userID uuid.UUID, userVehicles *[]entity.UserVehicle) error {
	// Try cache first
	cacheKey := BuildCacheKey(CacheKeyUserVehicles, userID.String())
	
	err := r.cache.Get(ctx, cacheKey, userVehicles)
	if err == nil {
		logger.Info(fmt.Sprintf("User vehicles retrieved from cache for user: %s", userID.String()))
		return nil
	}

	// Cache miss - fetch from database
	err = r.db.WithContext(ctx).Where("user_id = ?", userID).Find(userVehicles).Error
	if err != nil {
		return err
	}

	// Cache for 5 minutes
	cacheErr := r.cache.SetWithTags(ctx, cacheKey, *userVehicles, 5*time.Minute, []string{CacheTagUserData})
	if cacheErr != nil {
		logger.Error(cacheErr, "Failed to cache user vehicles")
	}

	return nil
}

func (r *vehicleRepository) GetUserVehicle(ctx context.Context, userID uuid.UUID, vehicleId uint64, userVehicle *entity.UserVehicle) error {
	// Try cache first
	cacheKey := BuildCacheKey(CacheKeyUserVehicles, userID.String(), fmt.Sprintf("%d", vehicleId))
	
	err := r.cache.Get(ctx, cacheKey, userVehicle)
	if err == nil {
		logger.Info(fmt.Sprintf("User vehicle retrieved from cache: %d", vehicleId))
		return nil
	}

	// Cache miss - fetch from database
	err = r.db.WithContext(ctx).Where("user_id = ? AND id = ?", userID, vehicleId).First(userVehicle).Error
	if err != nil {
		return err
	}

	// Cache for 10 minutes
	cacheErr := r.cache.SetWithTags(ctx, cacheKey, *userVehicle, 10*time.Minute, []string{CacheTagUserData})
	if cacheErr != nil {
		logger.Error(cacheErr, "Failed to cache user vehicle")
	}

	return nil
}

func (r *vehicleRepository) UpdateUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error {
	err := r.db.WithContext(ctx).Model(userVehicle).Updates(userVehicle).Error
	if err != nil {
		return err
	}

	// Invalidate related caches
	userCacheKey := BuildCacheKey(CacheKeyUserVehicles, userVehicle.UserID.String())
	vehicleCacheKey := BuildCacheKey(CacheKeyUserVehicles, userVehicle.UserID.String(), fmt.Sprintf("%d", userVehicle.ID))
	
	r.cache.Delete(ctx, userCacheKey)
	r.cache.Delete(ctx, vehicleCacheKey)
	
	return nil
}

func (r *vehicleRepository) DeleteUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error {
	err := r.db.WithContext(ctx).Delete(userVehicle).Error
	if err != nil {
		return err
	}

	// Invalidate related caches
	userCacheKey := BuildCacheKey(CacheKeyUserVehicles, userVehicle.UserID.String())
	vehicleCacheKey := BuildCacheKey(CacheKeyUserVehicles, userVehicle.UserID.String(), fmt.Sprintf("%d", userVehicle.ID))
	
	r.cache.Delete(ctx, userCacheKey)
	r.cache.Delete(ctx, vehicleCacheKey)
	
	return nil
}

// Complete hierarchy methods with caching
func (r *vehicleRepository) GetCompleteVehicleHierarchy(ctx context.Context, vehicleTypes *[]entity.VehicleType) error {

	
	err := r.db.WithContext(ctx).
		Preload("VehicleBrands").
		Preload("VehicleBrands.VehicleModels").
		Preload("VehicleBrands.VehicleModels.VehicleGenerations").
		Find(vehicleTypes).Error
	
	if err != nil {
		logger.Error(err, "Failed to fetch vehicle hierarchy from database")
		return err
	}

	logger.Info("Vehicle hierarchy fetched from database and cached")
	return nil
}
