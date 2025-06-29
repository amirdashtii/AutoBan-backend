package repository

import (
	"context"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
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
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository() VehicleRepository {
	db := database.ConnectDatabase()
	return &vehicleRepository{db: db}
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
	return r.db.WithContext(ctx).Preload("VehicleType").Find(brands).Error
}

func (r *vehicleRepository) ListBrandsByType(ctx context.Context, brands *[]entity.VehicleBrand, typeID uint64) error {
	return r.db.WithContext(ctx).Preload("VehicleType").Where("vehicle_type_id = ?", typeID).Find(brands).Error
}

func (r *vehicleRepository) GetBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Preload("VehicleType").Where("id = ?", brand.ID).First(brand).Error
}

func (r *vehicleRepository) CreateBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Preload("VehicleType").Create(brand).Error
}

func (r *vehicleRepository) UpdateBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Preload("VehicleType").Model(brand).Updates(brand).Error
}

func (r *vehicleRepository) DeleteBrand(ctx context.Context, brand *entity.VehicleBrand) error {
	return r.db.WithContext(ctx).Delete(brand).Error
}

// Models
func (r *vehicleRepository) ListModels(ctx context.Context, models *[]entity.VehicleModel) error {
	return r.db.WithContext(ctx).Preload("Brand").Preload("Brand.VehicleType").Find(models).Error
}

func (r *vehicleRepository) ListModelsByBrand(ctx context.Context, models *[]entity.VehicleModel, brandID uint64) error {
	return r.db.WithContext(ctx).Preload("Brand").Preload("Brand.VehicleType").Where("brand_id = ?", brandID).Find(models).Error
}

func (r *vehicleRepository) GetModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Preload("Brand").Preload("Brand.VehicleType").Where("id = ?", model.ID).First(model).Error
}

func (r *vehicleRepository) CreateModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Preload("Brand").Preload("Brand.VehicleType").Create(model).Error
}

func (r *vehicleRepository) UpdateModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Preload("Brand").Preload("Brand.VehicleType").Model(model).Updates(model).Error
}

func (r *vehicleRepository) DeleteModel(ctx context.Context, model *entity.VehicleModel) error {
	return r.db.WithContext(ctx).Delete(model).Error
}

// Generations
func (r *vehicleRepository) ListGenerations(ctx context.Context, generations *[]entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Preload("ModelInfo").Preload("ModelInfo.Brand").Preload("ModelInfo.Brand.VehicleType").Find(generations).Error
}

func (r *vehicleRepository) GetGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Preload("ModelInfo").Preload("ModelInfo.Brand").Preload("ModelInfo.Brand.VehicleType").Where("id = ?", generation.ID).First(generation).Error
}

func (r *vehicleRepository) ListGenerationsByModel(ctx context.Context, generations *[]entity.VehicleGeneration, modelID uint64) error {
	return r.db.WithContext(ctx).Preload("ModelInfo").Preload("ModelInfo.Brand").Preload("ModelInfo.Brand.VehicleType").Where("model_id = ?", modelID).Find(generations).Error
}
func (r *vehicleRepository) CreateGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Preload("ModelInfo").Preload("ModelInfo.Brand").Preload("ModelInfo.Brand.VehicleType").Create(generation).Error
}

func (r *vehicleRepository) UpdateGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Preload("ModelInfo").Preload("ModelInfo.Brand").Preload("ModelInfo.Brand.VehicleType").Model(generation).Updates(generation).Error
}

func (r *vehicleRepository) DeleteGeneration(ctx context.Context, generation *entity.VehicleGeneration) error {
	return r.db.WithContext(ctx).Delete(generation).Error
}

// User Vehicles
func (r *vehicleRepository) CreateUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error {
	return r.db.WithContext(ctx).Preload("Generation").Preload("Generation.ModelInfo").Preload("Generation.ModelInfo.Brand").Preload("Generation.ModelInfo.Brand.VehicleType").Create(userVehicle).Error
}

func (r *vehicleRepository) ListUserVehicles(ctx context.Context, userID uuid.UUID, userVehicles *[]entity.UserVehicle) error {
	return r.db.WithContext(ctx).Preload("Generation").Preload("Generation.ModelInfo").Preload("Generation.ModelInfo.Brand").Preload("Generation.ModelInfo.Brand.VehicleType").Where("user_id = ?", userID).Find(&userVehicles).Error
}

func (r *vehicleRepository) GetUserVehicle(ctx context.Context, userID uuid.UUID, vehicleId uint64, userVehicle *entity.UserVehicle) error {
	return r.db.WithContext(ctx).Preload("Generation").Preload("Generation.ModelInfo").Preload("Generation.ModelInfo.Brand").Preload("Generation.ModelInfo.Brand.VehicleType").Where("user_id = ? AND id = ?", userID, vehicleId).First(userVehicle).Error
}

func (r *vehicleRepository) UpdateUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error {
	return r.db.WithContext(ctx).Preload("Generation").Preload("Generation.ModelInfo").Preload("Generation.ModelInfo.Brand").Preload("Generation.ModelInfo.Brand.VehicleType").Model(userVehicle).Updates(userVehicle).Error
}

func (r *vehicleRepository) DeleteUserVehicle(ctx context.Context, userVehicle *entity.UserVehicle) error {
	return r.db.WithContext(ctx).Delete(userVehicle).Error
}
