package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/internal/validation"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/google/uuid"
)

type VehicleUseCase interface {
	// Vehicle Types
	ListVehicleTypes(ctx context.Context) (*dto.ListVehicleTypesResponse, error)
	GetVehicleType(ctx context.Context, typeID string) (*dto.VehicleTypeResponse, error)
	CreateVehicleType(ctx context.Context, request dto.CreateVehicleTypeRequest) (*dto.VehicleTypeResponse, error)
	UpdateVehicleType(ctx context.Context, typeID string, request dto.UpdateVehicleTypeRequest) (*dto.VehicleTypeResponse, error)
	DeleteVehicleType(ctx context.Context, typeID string) error

	// Brands
	GetBrand(ctx context.Context, typeID, brandID string) (*dto.VehicleBrandResponse, error)
	ListBrands(ctx context.Context, typeID string) (*dto.ListVehicleBrandsResponse, error)
	CreateBrand(ctx context.Context, typeID string, request dto.CreateVehicleBrandRequest) (*dto.VehicleBrandResponse, error)
	UpdateBrand(ctx context.Context, typeID, brandID string, request dto.UpdateVehicleBrandRequest) (*dto.VehicleBrandResponse, error)
	DeleteBrand(ctx context.Context, typeID, brandID string) error

	// Models
	GetModel(ctx context.Context, typeID, brandID, modelID string) (*dto.VehicleModelResponse, error)
	ListModels(ctx context.Context, typeID, brandID string) (*dto.ListVehicleModelsResponse, error)
	CreateModel(ctx context.Context, typeID, brandID string, request dto.CreateVehicleModelRequest) (*dto.VehicleModelResponse, error)
	UpdateModel(ctx context.Context, typeID, brandID, modelID string, request dto.UpdateVehicleModelRequest) (*dto.VehicleModelResponse, error)
	DeleteModel(ctx context.Context, typeID, brandID, modelID string) error

	// Generations
	GetGeneration(ctx context.Context, typeID, brandID, modelID, generationID string) (*dto.VehicleGenerationResponse, error)
	ListGenerations(ctx context.Context, typeID, brandID, modelID string) (*dto.ListVehicleGenerationsResponse, error)
	CreateGeneration(ctx context.Context, typeID, brandID, modelID string, request dto.CreateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error)
	UpdateGeneration(ctx context.Context, typeID, brandID, modelID, generationID string, request dto.UpdateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error)
	DeleteGeneration(ctx context.Context, typeID, brandID, modelID, generationID string) error

	// User Vehicles
	AddUserVehicle(ctx context.Context, userID string, request *dto.CreateUserVehicleRequest) (*dto.UserVehicleResponse, error)
	ListUserVehicles(ctx context.Context, userID string) (*dto.ListUserVehiclesResponse, error)
	GetUserVehicle(ctx context.Context, userID, vehicleID string) (*dto.UserVehicleResponse, error)
	UpdateUserVehicle(ctx context.Context, userID, vehicleID string, request *dto.UpdateUserVehicleRequest) (*dto.UserVehicleResponse, error)
	DeleteUserVehicle(ctx context.Context, userID, vehicleID string) error

	// Complete hierarchy
	GetCompleteVehicleHierarchy(ctx context.Context) (*dto.CompleteVehicleHierarchyResponse, error)
}

type vehicleUseCase struct {
	vehicleRepository repository.VehicleRepository
}

func NewVehicleUseCase() VehicleUseCase {
	vehicleRepository := repository.NewVehicleRepository()
	return &vehicleUseCase{vehicleRepository: vehicleRepository}
}

// Vehicle Types
func (uc *vehicleUseCase) ListVehicleTypes(ctx context.Context) (*dto.ListVehicleTypesResponse, error) {
	vehicleTypes := []entity.VehicleType{}
	err := uc.vehicleRepository.ListVehicleTypes(ctx, &vehicleTypes)
	if err != nil {
		logger.Error(err, "Failed to list vehicle types")
		return nil, errors.ErrFailedToListVehicleTypes
	}
	vehicleTypesResponse := []dto.VehicleTypeResponse{}
	for _, vehicleType := range vehicleTypes {
		vehicleTypesResponse = append(vehicleTypesResponse, dto.VehicleTypeResponse{
			ID:            vehicleType.ID,
			NameFa:        vehicleType.NameFa,
			NameEn:        vehicleType.NameEn,
			DescriptionFa: vehicleType.DescriptionFa,
			DescriptionEn: vehicleType.DescriptionEn,
		})
	}

	return &dto.ListVehicleTypesResponse{
		Types: vehicleTypesResponse,
	}, nil
}

func (uc *vehicleUseCase) GetVehicleType(ctx context.Context, typeID string) (*dto.VehicleTypeResponse, error) {
	vehicleType := entity.VehicleType{}
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	vehicleType.ID = uintTypeID

	err = uc.vehicleRepository.GetVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to get vehicle type")
		return nil, errors.ErrFailedToGetVehicleType
	}
	return &dto.VehicleTypeResponse{
		ID:            vehicleType.ID,
		NameFa:        vehicleType.NameFa,
		NameEn:        vehicleType.NameEn,
		DescriptionFa: vehicleType.DescriptionFa,
		DescriptionEn: vehicleType.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) CreateVehicleType(ctx context.Context, request dto.CreateVehicleTypeRequest) (*dto.VehicleTypeResponse, error) {
	err := validation.ValidateVehicleTypeCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle type create request")
		return nil, errors.ErrInvalidVehicleTypeCreateRequest
	}

	vehicleType := entity.VehicleType{
		NameFa:        request.NameFa,
		NameEn:        request.NameEn,
		DescriptionFa: request.DescriptionFa,
		DescriptionEn: request.DescriptionEn,
	}
	err = uc.vehicleRepository.CreateVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to create vehicle type")
		return nil, errors.ErrFailedToCreateVehicleType
	}
	return &dto.VehicleTypeResponse{
		ID:            vehicleType.ID,
		NameFa:        vehicleType.NameFa,
		NameEn:        vehicleType.NameEn,
		DescriptionFa: vehicleType.DescriptionFa,
		DescriptionEn: vehicleType.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) UpdateVehicleType(ctx context.Context, typeID string, request dto.UpdateVehicleTypeRequest) (*dto.VehicleTypeResponse, error) {
	err := validation.ValidateVehicleTypeUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle type update request")
		return nil, errors.ErrInvalidVehicleTypeUpdateRequest
	}
	vehicleType := entity.VehicleType{}
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	vehicleType.ID = uintTypeID

	if request.NameFa != nil {
		vehicleType.NameFa = *request.NameFa
	}
	if request.DescriptionFa != nil {
		vehicleType.DescriptionFa = *request.DescriptionFa
	}
	if request.NameEn != nil {
		vehicleType.NameEn = *request.NameEn
	}
	if request.DescriptionEn != nil {
		vehicleType.DescriptionEn = *request.DescriptionEn
	}

	err = uc.vehicleRepository.UpdateVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to update vehicle type")
		return nil, errors.ErrFailedToUpdateVehicleType
	}
	return &dto.VehicleTypeResponse{
		ID:            vehicleType.ID,
		NameFa:        vehicleType.NameFa,
		NameEn:        vehicleType.NameEn,
		DescriptionFa: vehicleType.DescriptionFa,
		DescriptionEn: vehicleType.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) DeleteVehicleType(ctx context.Context, typeID string) error {
	vehicleType := entity.VehicleType{}
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return errors.ErrInvalidVehicleTypeID
	}
	vehicleType.ID = uintTypeID
	err = uc.vehicleRepository.DeleteVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle type")
		return errors.ErrFailedToDeleteVehicleType
	}
	return nil
}

// Brands
func (uc *vehicleUseCase) GetBrand(ctx context.Context, typeID, brandID string) (*dto.VehicleBrandResponse, error) {
	brand := entity.VehicleBrand{}
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}
	brand.ID = uintBrandID
	brand.VehicleTypeID = uintTypeID
	err = uc.vehicleRepository.GetBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to get vehicle brand")
		return nil, errors.ErrFailedToGetVehicleBrand
	}
	return &dto.VehicleBrandResponse{
		ID:            brand.ID,
		VehicleTypeID: brand.VehicleTypeID,
		NameFa:        brand.NameFa,
		NameEn:        brand.NameEn,
		DescriptionFa: brand.DescriptionFa,
		DescriptionEn: brand.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) ListBrands(ctx context.Context, typeID string) (*dto.ListVehicleBrandsResponse, error) {
	brands := []entity.VehicleBrand{}
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	err = uc.vehicleRepository.ListBrandsByType(ctx, &brands, uintTypeID)
	if err != nil {
		logger.Error(err, "Failed to list vehicle brands by type")
		return nil, errors.ErrFailedToListVehicleBrandsByType
	}
	brandsResponse := []dto.VehicleBrandResponse{}
	for _, brand := range brands {
		brandsResponse = append(brandsResponse, dto.VehicleBrandResponse{
			ID:            brand.ID,
			VehicleTypeID: brand.VehicleTypeID,
			NameFa:        brand.NameFa,
			NameEn:        brand.NameEn,
			DescriptionFa: brand.DescriptionFa,
			DescriptionEn: brand.DescriptionEn,
		})
	}
	return &dto.ListVehicleBrandsResponse{Brands: brandsResponse}, nil
}

func (uc *vehicleUseCase) CreateBrand(ctx context.Context, typeID string, request dto.CreateVehicleBrandRequest) (*dto.VehicleBrandResponse, error) {
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	err = validation.ValidateVehicleBrandCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle brand create request")
		return nil, errors.ErrInvalidVehicleBrandCreateRequest
	}

	brand := entity.VehicleBrand{
		NameFa:        request.NameFa,
		NameEn:        request.NameEn,
		DescriptionFa: request.DescriptionFa,
		DescriptionEn: request.DescriptionEn,
		VehicleTypeID: uintTypeID,
	}
	err = uc.vehicleRepository.CreateBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to create vehicle brand")
		return nil, errors.ErrFailedToCreateVehicleBrand
	}
	return &dto.VehicleBrandResponse{
		ID:            brand.ID,
		VehicleTypeID: brand.VehicleTypeID,
		NameFa:        brand.NameFa,
		NameEn:        brand.NameEn,
		DescriptionFa: brand.DescriptionFa,
		DescriptionEn: brand.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) UpdateBrand(ctx context.Context, typeID, brandID string, request dto.UpdateVehicleBrandRequest) (*dto.VehicleBrandResponse, error) {
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}

	err = validation.ValidateVehicleBrandUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle brand update request")
		return nil, errors.ErrInvalidVehicleBrandUpdateRequest
	}

	brand := entity.VehicleBrand{}
	brand.ID = uintBrandID
	brand.VehicleTypeID = uintTypeID

	if request.NameFa != nil {
		brand.NameFa = *request.NameFa
	}
	if request.DescriptionFa != nil {
		brand.DescriptionFa = *request.DescriptionFa
	}
	if request.NameEn != nil {
		brand.NameEn = *request.NameEn
	}
	if request.DescriptionEn != nil {
		brand.DescriptionEn = *request.DescriptionEn
	}
	if request.VehicleTypeID != nil {
		brand.VehicleTypeID = *request.VehicleTypeID
	}

	err = uc.vehicleRepository.UpdateBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to update vehicle brand")
		return nil, errors.ErrFailedToUpdateVehicleBrand
	}

	return &dto.VehicleBrandResponse{
		ID:            brand.ID,
		VehicleTypeID: brand.VehicleTypeID,
		NameFa:        brand.NameFa,
		NameEn:        brand.NameEn,
		DescriptionFa: brand.DescriptionFa,
		DescriptionEn: brand.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) DeleteBrand(ctx context.Context, typeID, brandID string) error {
	brand := entity.VehicleBrand{}
	uintTypeID, err := strconv.ParseUint(typeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return errors.ErrInvalidVehicleTypeID
	}
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return errors.ErrInvalidVehicleBrandID
	}
	brand.ID = uintBrandID
	brand.VehicleTypeID = uintTypeID
	err = uc.vehicleRepository.DeleteBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle brand")
		return errors.ErrFailedToDeleteVehicleBrand
	}
	return nil
}

// Models
func (uc *vehicleUseCase) ListModels(ctx context.Context, typeID, brandID string) (*dto.ListVehicleModelsResponse, error) {
	models := []entity.VehicleModel{}
	_ = typeID
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}
	err = uc.vehicleRepository.ListModelsByBrand(ctx, &models, uintBrandID)
	if err != nil {
		logger.Error(err, "Failed to list vehicle models by brand")
		return nil, errors.ErrFailedToListVehicleModelsByBrand
	}
	modelsResponse := []dto.VehicleModelResponse{}
	for _, model := range models {
		modelsResponse = append(modelsResponse, dto.VehicleModelResponse{
			ID:            model.ID,
			BrandID:       model.BrandID,
			NameFa:        model.NameFa,
			NameEn:        model.NameEn,
			DescriptionFa: model.DescriptionFa,
			DescriptionEn: model.DescriptionEn,
		})
	}
	return &dto.ListVehicleModelsResponse{Models: modelsResponse}, nil
}

func (uc *vehicleUseCase) GetModel(ctx context.Context, typeID, brandID, modelID string) (*dto.VehicleModelResponse, error) {
	_ = typeID
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}
	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	model := entity.VehicleModel{}
	model.ID = uintModelID
	model.BrandID = uintBrandID
	err = uc.vehicleRepository.GetModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to get vehicle model")
		return nil, errors.ErrFailedToGetVehicleModel
	}
	return &dto.VehicleModelResponse{
		ID:            model.ID,
		BrandID:       model.BrandID,
		NameFa:        model.NameFa,
		NameEn:        model.NameEn,
		DescriptionFa: model.DescriptionFa,
		DescriptionEn: model.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) CreateModel(ctx context.Context, typeID, brandID string, request dto.CreateVehicleModelRequest) (*dto.VehicleModelResponse, error) {
	err := validation.ValidateVehicleModelCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle model create request")
		return nil, errors.ErrInvalidVehicleModelCreateRequest
	}
	_ = typeID
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}

	model := entity.VehicleModel{
		NameFa:        request.NameFa,
		NameEn:        request.NameEn,
		DescriptionFa: request.DescriptionFa,
		DescriptionEn: request.DescriptionEn,
		BrandID:       uintBrandID,
	}
	err = uc.vehicleRepository.CreateModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to create vehicle model")
		return nil, errors.ErrFailedToCreateVehicleModel
	}
	return &dto.VehicleModelResponse{
		ID:            model.ID,
		BrandID:       model.BrandID,
		NameFa:        model.NameFa,
		NameEn:        model.NameEn,
		DescriptionFa: model.DescriptionFa,
		DescriptionEn: model.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) UpdateModel(ctx context.Context, typeID, brandID, modelID string, request dto.UpdateVehicleModelRequest) (*dto.VehicleModelResponse, error) {
	err := validation.ValidateVehicleModelUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle model update request")
		return nil, errors.ErrInvalidVehicleModelUpdateRequest
	}

	_ = typeID
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}
	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	model := entity.VehicleModel{}
	model.ID = uintModelID
	model.BrandID = uintBrandID

	if request.NameFa != nil {
		model.NameFa = *request.NameFa
	}
	if request.DescriptionFa != nil {
		model.DescriptionFa = *request.DescriptionFa
	}
	if request.NameEn != nil {
		model.NameEn = *request.NameEn
	}
	if request.DescriptionEn != nil {
		model.DescriptionEn = *request.DescriptionEn
	}
	if request.BrandID != nil {
		model.BrandID = *request.BrandID
	}

	err = uc.vehicleRepository.UpdateModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to update vehicle model")
		return nil, errors.ErrFailedToUpdateVehicleModel
	}
	return &dto.VehicleModelResponse{
		ID:            model.ID,
		BrandID:       model.BrandID,
		NameFa:        model.NameFa,
		NameEn:        model.NameEn,
		DescriptionFa: model.DescriptionFa,
		DescriptionEn: model.DescriptionEn,
	}, nil
}

func (uc *vehicleUseCase) DeleteModel(ctx context.Context, typeID, brandID, modelID string) error {
	_ = typeID
	uintBrandID, err := strconv.ParseUint(brandID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return errors.ErrInvalidVehicleModelID
	}
	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return errors.ErrInvalidVehicleModelID
	}
	model := entity.VehicleModel{}
	model.ID = uintModelID
	model.BrandID = uintBrandID
	err = uc.vehicleRepository.DeleteModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle model")
		return errors.ErrFailedToDeleteVehicleModel
	}
	return nil
}

// Generations
func (uc *vehicleUseCase) GetGeneration(ctx context.Context, typeID, brandID, modelID, generationID string) (*dto.VehicleGenerationResponse, error) {
	_ = typeID
	_ = brandID
	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	uintGenerationID, err := strconv.ParseUint(generationID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle generation id")
		return nil, errors.ErrInvalidVehicleGenerationID
	}
	generation := entity.VehicleGeneration{}
	generation.ID = uintGenerationID
	generation.ModelID = uintModelID

	err = uc.vehicleRepository.GetGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to get vehicle generation")
		return nil, errors.ErrFailedToGetVehicleGeneration
	}
	return &dto.VehicleGenerationResponse{
		ID:            generation.ID,
		NameFa:        generation.NameFa,
		NameEn:        generation.NameEn,
		DescriptionFa: generation.DescriptionFa,
		DescriptionEn: generation.DescriptionEn,
		ModelID:       generation.ModelID,
		StartYear:     generation.StartYear,
		EndYear:       generation.EndYear,
		BodyStyleFa:   generation.BodyStyleFa,
		BodyStyleEn:   generation.BodyStyleEn,
		Engine:        generation.Engine,
		EngineVolume:  generation.EngineVolume,
		Cylinders:     generation.Cylinders,
		DrivetrainFa:  generation.DrivetrainFa,
		DrivetrainEn:  generation.DrivetrainEn,
		Gearbox:       generation.Gearbox,
		FuelType:      generation.FuelType,
		Battery:       generation.Battery,
		Seller:        generation.Seller,
		AssemblyType:  generation.AssemblyType,
		Assembler:     generation.Assembler,
	}, nil
}

func (uc *vehicleUseCase) ListGenerations(ctx context.Context, typeID, brandID, modelID string) (*dto.ListVehicleGenerationsResponse, error) {
	_ = typeID
	_ = brandID

	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	generations := []entity.VehicleGeneration{}
	err = uc.vehicleRepository.ListGenerationsByModel(ctx, &generations, uintModelID)
	if err != nil {
		logger.Error(err, "Failed to list vehicle generations by model")
		return nil, errors.ErrFailedToListVehicleGenerationsByModel
	}

	generationsResponse := []dto.VehicleGenerationResponse{}
	for _, generation := range generations {
		generationsResponse = append(generationsResponse, dto.VehicleGenerationResponse{
			ID:            generation.ID,
			NameFa:        generation.NameFa,
			NameEn:        generation.NameEn,
			DescriptionFa: generation.DescriptionFa,
			DescriptionEn: generation.DescriptionEn,
			ModelID:       generation.ModelID,
			StartYear:     generation.StartYear,
			EndYear:       generation.EndYear,
			BodyStyleFa:   generation.BodyStyleFa,
			BodyStyleEn:   generation.BodyStyleEn,
			Engine:        generation.Engine,
			EngineVolume:  generation.EngineVolume,
			Cylinders:     generation.Cylinders,
			DrivetrainFa:  generation.DrivetrainFa,
			DrivetrainEn:  generation.DrivetrainEn,
			Gearbox:       generation.Gearbox,
			FuelType:      generation.FuelType,
			Battery:       generation.Battery,
			Seller:        generation.Seller,
			AssemblyType:  generation.AssemblyType,
			Assembler:     generation.Assembler,
		})
	}
	return &dto.ListVehicleGenerationsResponse{Generations: generationsResponse}, nil
}

func (uc *vehicleUseCase) CreateGeneration(ctx context.Context, typeID, brandID, modelID string, request dto.CreateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error) {
	err := validation.ValidateVehicleGenerationCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle generation create request")
		return nil, errors.ErrInvalidVehicleGenerationCreateRequest
	}

	_ = typeID
	_ = brandID
	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}

	generation := entity.VehicleGeneration{
		ModelID:       uintModelID,
		NameFa:        request.NameFa,
		NameEn:        request.NameEn,
		DescriptionFa: request.DescriptionFa,
		DescriptionEn: request.DescriptionEn,
		StartYear:     request.StartYear,
		EndYear:       request.EndYear,
		BodyStyleFa:   request.BodyStyleFa,
		BodyStyleEn:   request.BodyStyleEn,
		Engine:        request.Engine,
		EngineVolume:  request.EngineVolume,
		Cylinders:     request.Cylinders,
		DrivetrainFa:  request.DrivetrainFa,
		DrivetrainEn:  request.DrivetrainEn,
		Gearbox:       request.Gearbox,
		FuelType:      request.FuelType,
		Battery:       request.Battery,
		Seller:        request.Seller,
		AssemblyType:  request.AssemblyType,
		Assembler:     request.Assembler,
	}
	err = uc.vehicleRepository.CreateGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to create vehicle generation")
		return nil, errors.ErrFailedToCreateVehicleGeneration
	}
	return &dto.VehicleGenerationResponse{
		ID:            generation.ID,
		NameFa:        generation.NameFa,
		NameEn:        generation.NameEn,
		DescriptionFa: generation.DescriptionFa,
		DescriptionEn: generation.DescriptionEn,
		ModelID:       generation.ModelID,
		StartYear:     generation.StartYear,
		EndYear:       generation.EndYear,
		BodyStyleFa:   generation.BodyStyleFa,
		BodyStyleEn:   generation.BodyStyleEn,
		Engine:        generation.Engine,
		EngineVolume:  generation.EngineVolume,
		Cylinders:     generation.Cylinders,
		DrivetrainFa:  generation.DrivetrainFa,
		DrivetrainEn:  generation.DrivetrainEn,
		Gearbox:       generation.Gearbox,
		FuelType:      generation.FuelType,
		Battery:       generation.Battery,
		AssemblyType:  generation.AssemblyType,
		Assembler:     generation.Assembler,
	}, nil
}

func (uc *vehicleUseCase) UpdateGeneration(ctx context.Context, typeID, brandID, modelID, generationID string, request dto.UpdateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error) {
	err := validation.ValidateVehicleGenerationUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle generation update request")
		return nil, errors.ErrInvalidVehicleGenerationUpdateRequest
	}

	_ = typeID
	_ = brandID
	uintModelID, err := strconv.ParseUint(modelID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	uintGenerationID, err := strconv.ParseUint(generationID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle generation id")
		return nil, errors.ErrInvalidVehicleGenerationID
	}
	generation := entity.VehicleGeneration{}
	generation.ID = uintGenerationID
	generation.ModelID = uintModelID

	if request.NameFa != nil {
		generation.NameFa = *request.NameFa
	}
	if request.NameEn != nil {
		generation.NameEn = *request.NameEn
	}
	if request.DescriptionFa != nil {
		generation.DescriptionFa = *request.DescriptionFa
	}
	if request.DescriptionEn != nil {
		generation.DescriptionEn = *request.DescriptionEn
	}
	if request.ModelID != nil {
		generation.ModelID = *request.ModelID
	}
	if request.StartYear != nil {
		generation.StartYear = *request.StartYear
	}
	if request.EndYear != nil {
		generation.EndYear = *request.EndYear
	}
	if request.Engine != nil {
		generation.Engine = *request.Engine
	}
	if request.EngineVolume != nil {
		generation.EngineVolume = *request.EngineVolume
	}
	if request.Cylinders != nil {
		generation.Cylinders = *request.Cylinders
	}
	if request.DrivetrainFa != nil {
		generation.DrivetrainFa = *request.DrivetrainFa
	}
	if request.DrivetrainEn != nil {
		generation.DrivetrainEn = *request.DrivetrainEn
	}
	if request.Gearbox != nil {
		generation.Gearbox = *request.Gearbox
	}
	if request.FuelType != nil {
		generation.FuelType = *request.FuelType
	}
	if request.Battery != nil {
		generation.Battery = *request.Battery
	}
	if request.Seller != nil {
		generation.Seller = *request.Seller
	}
	if request.AssemblyType != nil {
		generation.AssemblyType = *request.AssemblyType
	}
	if request.Assembler != nil {
		generation.Assembler = *request.Assembler
	}

	err = uc.vehicleRepository.UpdateGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to update vehicle generation")
		return nil, errors.ErrFailedToUpdateVehicleGeneration
	}
	return &dto.VehicleGenerationResponse{
		ID:            generation.ID,
		ModelID:       generation.ModelID,
		NameFa:        generation.NameFa,
		NameEn:        generation.NameEn,
		DescriptionFa: generation.DescriptionFa,
		DescriptionEn: generation.DescriptionEn,
		StartYear:     generation.StartYear,
		EndYear:       generation.EndYear,
		BodyStyleFa:   generation.BodyStyleFa,
		BodyStyleEn:   generation.BodyStyleEn,
		Engine:        generation.Engine,
		EngineVolume:  generation.EngineVolume,
		Cylinders:     generation.Cylinders,
		DrivetrainFa:  generation.DrivetrainFa,
		DrivetrainEn:  generation.DrivetrainEn,
		Gearbox:       generation.Gearbox,
		FuelType:      generation.FuelType,
		Battery:       generation.Battery,
		Seller:        generation.Seller,
		AssemblyType:  generation.AssemblyType,
		Assembler:     generation.Assembler,
	}, nil
}

func (uc *vehicleUseCase) DeleteGeneration(ctx context.Context, typeID, brandID, modelID, generationID string) error {
	_ = typeID
	_ = brandID
	_ = modelID
	uintGenerationID, err := strconv.ParseUint(generationID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle generation id")
		return errors.ErrInvalidVehicleGenerationID
	}

	generation := entity.VehicleGeneration{}
	generation.ID = uintGenerationID
	err = uc.vehicleRepository.DeleteGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle generation")
		return errors.ErrFailedToDeleteVehicleGeneration
	}
	return nil
}

// User Vehicles
func (uc *vehicleUseCase) AddUserVehicle(ctx context.Context, userID string, request *dto.CreateUserVehicleRequest) (*dto.UserVehicleResponse, error) {
	err := validation.ValidateUserVehicleCreateRequest(*request)
	if err != nil {
		logger.Error(err, "Failed to validate user vehicle create request")
		return nil, errors.ErrInvalidUserVehicleCreateRequest
	}

	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}

	purchaseDate, err := time.Parse("2006-01-02", request.PurchaseDate)
	if err != nil {
		logger.Error(err, "Failed to parse purchase date")
		return nil, errors.ErrInvalidPurchaseDate
	}

	userVehicle := entity.UserVehicle{
		UserID:         uuidUserID,
		GenerationID:   request.GenerationID,
		Name:           request.Name,
		ProductionYear: request.ProductionYear,
		Color:          request.Color,
		LicensePlate:   request.LicensePlate,
		VIN:            request.VIN,
		CurrentMileage: request.CurrentMileage,
		PurchaseDate:   purchaseDate,
	}
	err = uc.vehicleRepository.CreateUserVehicle(ctx, &userVehicle)
	if err != nil {
		logger.Error(err, "Failed to create user vehicle")
		return nil, errors.ErrFailedToCreateUserVehicle
	}
	return &dto.UserVehicleResponse{
		ID:             userVehicle.ID,
		UserID:         userVehicle.UserID,
		GenerationID:   userVehicle.GenerationID,
		Name:           userVehicle.Name,
		ProductionYear: userVehicle.ProductionYear,
		Color:          userVehicle.Color,
		LicensePlate:   userVehicle.LicensePlate,
		VIN:            userVehicle.VIN,
		CurrentMileage: userVehicle.CurrentMileage,
		PurchaseDate:   userVehicle.PurchaseDate,
	}, nil
}

func (uc *vehicleUseCase) ListUserVehicles(ctx context.Context, userID string) (*dto.ListUserVehiclesResponse, error) {
	userVehicles := []entity.UserVehicle{}
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	err = uc.vehicleRepository.ListUserVehicles(ctx, uuidUserID, &userVehicles)
	if err != nil {
		logger.Error(err, "Failed to list user vehicles")
		return nil, errors.ErrFailedToListUserVehicles
	}
	userVehiclesResponse := []dto.UserVehicleResponse{}
	for _, userVehicle := range userVehicles {
		userVehiclesResponse = append(userVehiclesResponse, dto.UserVehicleResponse{
			ID:             userVehicle.ID,
			UserID:         userVehicle.UserID,
			GenerationID:   userVehicle.GenerationID,
			Name:           userVehicle.Name,
			ProductionYear: userVehicle.ProductionYear,
			Color:          userVehicle.Color,
			LicensePlate:   userVehicle.LicensePlate,
			VIN:            userVehicle.VIN,
			CurrentMileage: userVehicle.CurrentMileage,
			PurchaseDate:   userVehicle.PurchaseDate,
		})
	}
	return &dto.ListUserVehiclesResponse{Vehicles: userVehiclesResponse}, nil
}

func (uc *vehicleUseCase) GetUserVehicle(ctx context.Context, userID, vehicleID string) (*dto.UserVehicleResponse, error) {
	userVehicle := entity.UserVehicle{}
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle id")
		return nil, errors.ErrInvalidVehicleID
	}
	err = uc.vehicleRepository.GetUserVehicle(ctx, uuidUserID, uintVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "Failed to get user vehicle")
		return nil, errors.ErrFailedToGetUserVehicle
	}

	return &dto.UserVehicleResponse{
		ID:             userVehicle.ID,
		UserID:         userVehicle.UserID,
		GenerationID:   userVehicle.GenerationID,
		Name:           userVehicle.Name,
		ProductionYear: userVehicle.ProductionYear,
		Color:          userVehicle.Color,
		LicensePlate:   userVehicle.LicensePlate,
		VIN:            userVehicle.VIN,
		CurrentMileage: userVehicle.CurrentMileage,
		PurchaseDate:   userVehicle.PurchaseDate,
	}, nil
}

func (uc *vehicleUseCase) UpdateUserVehicle(ctx context.Context, userID, vehicleID string, request *dto.UpdateUserVehicleRequest) (*dto.UserVehicleResponse, error) {
	err := validation.ValidateUserVehicleUpdateRequest(*request)
	if err != nil {
		logger.Error(err, "Failed to validate user vehicle update request")
		return nil, errors.ErrInvalidUserVehicleUpdateRequest
	}

	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}

	purchaseDate, err := time.Parse("2006-01-02", *request.PurchaseDate)
	if err != nil {
		logger.Error(err, "Failed to parse purchase date")
		return nil, errors.ErrInvalidPurchaseDate
	}

	userVehicle := entity.UserVehicle{
		UserID: uuidUserID,
	}

	// Validate and assign optional fields
	if request.Name != nil {
		userVehicle.Name = *request.Name
	}
	if request.GenerationID != nil {
		userVehicle.GenerationID = *request.GenerationID
	}
	if request.ProductionYear != nil {
		userVehicle.ProductionYear = *request.ProductionYear
	}
	if request.Color != nil {
		userVehicle.Color = *request.Color
	}
	if request.LicensePlate != nil {
		userVehicle.LicensePlate = *request.LicensePlate
	}
	if request.VIN != nil {
		userVehicle.VIN = *request.VIN
	}
	if request.CurrentMileage != nil {
		userVehicle.CurrentMileage = *request.CurrentMileage
	}
	if request.PurchaseDate != nil {
		userVehicle.PurchaseDate = purchaseDate
	}

	uintUserVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse user vehicle id")
		return nil, errors.ErrInvalidUserVehicleID
	}
	userVehicle.ID = uintUserVehicleID
	err = uc.vehicleRepository.UpdateUserVehicle(ctx, &userVehicle)
	if err != nil {
		logger.Error(err, "Failed to update user vehicle")
		return nil, errors.ErrFailedToUpdateUserVehicle
	}
	return &dto.UserVehicleResponse{
		ID:             userVehicle.ID,
		UserID:         userVehicle.UserID,
		GenerationID:   userVehicle.GenerationID,
		Name:           userVehicle.Name,
		ProductionYear: userVehicle.ProductionYear,
		Color:          userVehicle.Color,
		LicensePlate:   userVehicle.LicensePlate,
		VIN:            userVehicle.VIN,
		CurrentMileage: userVehicle.CurrentMileage,
		PurchaseDate:   userVehicle.PurchaseDate,
	}, nil
}

func (uc *vehicleUseCase) DeleteUserVehicle(ctx context.Context, userID, vehicleID string) error {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return errors.ErrInvalidUserID
	}
	uintVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle id")
		return errors.ErrInvalidVehicleID
	}
	userVehicle := entity.UserVehicle{}
	err = uc.vehicleRepository.GetUserVehicle(ctx, uuidUserID, uintVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "Failed to get user vehicle")
		return errors.ErrFailedToGetUserVehicle
	}
	err = uc.vehicleRepository.DeleteUserVehicle(ctx, &userVehicle)
	if err != nil {
		logger.Error(err, "Failed to delete user vehicle")
		return errors.ErrFailedToDeleteUserVehicle
	}
	return nil
}

// Complete hierarchy
func (uc *vehicleUseCase) GetCompleteVehicleHierarchy(ctx context.Context) (*dto.CompleteVehicleHierarchyResponse, error) {
	vehicleTypes, err := uc.vehicleRepository.GetCompleteVehicleHierarchy(ctx)
	if err != nil {
		logger.Error(err, "Failed to get complete vehicle hierarchy")
		return nil, errors.ErrFailedToListVehicleTypes
	}

	// Convert entities to DTO tree structure
	vehicleTypeTrees := []dto.VehicleTypeTreeResponse{}
	totalTypes := 0
	totalBrands := 0
	totalModels := 0
	totalGenerations := 0

	for _, vehicleType := range vehicleTypes {
		totalTypes++

		brandTrees := []dto.VehicleBrandTreeResponse{}
		for _, brand := range vehicleType.VehicleBrands {
			totalBrands++

			modelTrees := []dto.VehicleModelTreeResponse{}
			for _, model := range brand.VehicleModels {
				totalModels++

				generationTrees := []dto.VehicleGenerationTreeResponse{}
				for _, generation := range model.VehicleGenerations {
					totalGenerations++

					generationTree := dto.VehicleGenerationTreeResponse{
						ID:            generation.ID,
						ModelID:       model.ID,
						NameFa:        generation.NameFa,
						NameEn:        generation.NameEn,
						DescriptionFa: generation.DescriptionFa,
						DescriptionEn: generation.DescriptionEn,
						StartYear:     generation.StartYear,
						EndYear:       generation.EndYear,
						Engine:        generation.Engine,
						EngineVolume:  generation.EngineVolume,
						Cylinders:     generation.Cylinders,
						DrivetrainFa:  generation.DrivetrainFa,
						DrivetrainEn:  generation.DrivetrainEn,
						Gearbox:       generation.Gearbox,
						FuelType:      generation.FuelType,
						Battery:       generation.Battery,
						Seller:        generation.Seller,
						AssemblyType:  generation.AssemblyType,
						Assembler:     generation.Assembler,
						BodyStyleFa:   generation.BodyStyleFa,
						BodyStyleEn:   generation.BodyStyleEn,
					}
					generationTrees = append(generationTrees, generationTree)
				}

				modelTree := dto.VehicleModelTreeResponse{
					ID:            model.ID,
					BrandID:       model.BrandID,
					NameFa:        model.NameFa,
					NameEn:        model.NameEn,
					DescriptionFa: model.DescriptionFa,
					DescriptionEn: model.DescriptionEn,
					Generations:   generationTrees,
				}
				modelTrees = append(modelTrees, modelTree)
			}

			brandTree := dto.VehicleBrandTreeResponse{
				ID:            brand.ID,
				NameFa:        brand.NameFa,
				NameEn:        brand.NameEn,
				DescriptionFa: brand.DescriptionFa,
				DescriptionEn: brand.DescriptionEn,
				Models:        modelTrees,
			}
			brandTrees = append(brandTrees, brandTree)
		}

		vehicleTypeTree := dto.VehicleTypeTreeResponse{
			ID:            vehicleType.ID,
			NameFa:        vehicleType.NameFa,
			NameEn:        vehicleType.NameEn,
			DescriptionFa: vehicleType.DescriptionFa,
			DescriptionEn: vehicleType.DescriptionEn,
			Brands:        brandTrees,
		}
		vehicleTypeTrees = append(vehicleTypeTrees, vehicleTypeTree)
	}

	return &dto.CompleteVehicleHierarchyResponse{
		VehicleTypes:     vehicleTypeTrees,
		TotalTypes:       totalTypes,
		TotalBrands:      totalBrands,
		TotalModels:      totalModels,
		TotalGenerations: totalGenerations,
	}, nil
}
