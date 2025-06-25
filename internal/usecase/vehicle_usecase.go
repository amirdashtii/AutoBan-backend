package usecase

import (
	"context"
	"strconv"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/internal/validation"
	"github.com/amirdashtii/AutoBan/pkg/logger"
)

type VehicleUseCase interface {
	// Vehicle Types
	ListVehicleTypes(ctx context.Context) (*dto.ListVehicleTypesResponse, error)
	GetVehicleType(ctx context.Context, id string) (*dto.VehicleTypeResponse, error)
	CreateVehicleType(ctx context.Context, request dto.CreateVehicleTypeRequest) (*dto.VehicleTypeResponse, error)
	UpdateVehicleType(ctx context.Context, id string, request dto.UpdateVehicleTypeRequest) (*dto.VehicleTypeResponse, error)
	DeleteVehicleType(ctx context.Context, id string) error

	// Brands
	ListBrands(ctx context.Context) (*dto.ListVehicleBrandsResponse, error)
	GetBrand(ctx context.Context, id string) (*dto.VehicleBrandResponse, error)
	ListBrandsByType(ctx context.Context, typeID string) (*dto.ListVehicleBrandsResponse, error)
	CreateBrand(ctx context.Context, request dto.CreateVehicleBrandRequest) (*dto.VehicleBrandResponse, error)
	UpdateBrand(ctx context.Context, id string, request dto.UpdateVehicleBrandRequest) (*dto.VehicleBrandResponse, error)
	DeleteBrand(ctx context.Context, id string) error

	// Models
	ListModels(ctx context.Context) (*dto.ListVehicleModelsResponse, error)
	GetModel(ctx context.Context, id string) (*dto.VehicleModelResponse, error)
	ListModelsByBrand(ctx context.Context, brandID string) (*dto.ListVehicleModelsResponse, error)
	CreateModel(ctx context.Context, request dto.CreateVehicleModelRequest) (*dto.VehicleModelResponse, error)
	UpdateModel(ctx context.Context, id string, request dto.UpdateVehicleModelRequest) (*dto.VehicleModelResponse, error)
	DeleteModel(ctx context.Context, id string) error

	// Generations
	ListGenerations(ctx context.Context) (*dto.ListVehicleGenerationsResponse, error)
	GetGeneration(ctx context.Context, id string) (*dto.VehicleGenerationResponse, error)
	ListGenerationsByModel(ctx context.Context, modelID string) (*dto.ListVehicleGenerationsResponse, error)
	CreateGeneration(ctx context.Context, request dto.CreateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error)
	UpdateGeneration(ctx context.Context, id string, request dto.UpdateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error)
	DeleteGeneration(ctx context.Context, id string) error

	// User Vehicles
	AddUserVehicle(ctx context.Context, userID string, request *dto.CreateUserVehicleRequest) (*dto.UserVehicleResponse, error)
	ListUserVehicles(ctx context.Context, userID string) (*dto.ListUserVehiclesResponse, error)
	GetUserVehicle(ctx context.Context, userID, vehicleID string) (*dto.UserVehicleResponse, error)
	UpdateUserVehicle(ctx context.Context, userID, vehicleID string, request *dto.UpdateUserVehicleRequest) (*dto.UserVehicleResponse, error)
	DeleteUserVehicle(ctx context.Context, userID, vehicleID string) error
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
			ID:          vehicleType.ID,
			Name:        vehicleType.Name,
			Description: vehicleType.Description,
		})
	}

	return &dto.ListVehicleTypesResponse{
		Types: vehicleTypesResponse,
	}, nil
}

func (uc *vehicleUseCase) GetVehicleType(ctx context.Context, id string) (*dto.VehicleTypeResponse, error) {
	vehicleType := entity.VehicleType{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	vehicleType.ID = uint(uintID)

	err = uc.vehicleRepository.GetVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to get vehicle type")
		return nil, errors.ErrFailedToGetVehicleType
	}
	return &dto.VehicleTypeResponse{
		ID:          vehicleType.ID,
		Name:        vehicleType.Name,
		Description: vehicleType.Description,
	}, nil
}

func (uc *vehicleUseCase) CreateVehicleType(ctx context.Context, request dto.CreateVehicleTypeRequest) (*dto.VehicleTypeResponse, error) {
	err := validation.ValidateVehicleTypeCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle type create request")
		return nil, errors.ErrInvalidVehicleTypeCreateRequest
	}

	vehicleType := entity.VehicleType{
		Name:        request.Name,
		Description: request.Description,
	}
	err = uc.vehicleRepository.CreateVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to create vehicle type")
		return nil, errors.ErrFailedToCreateVehicleType
	}
	return &dto.VehicleTypeResponse{
		ID:          vehicleType.ID,
		Name:        vehicleType.Name,
		Description: vehicleType.Description,
	}, nil
}

func (uc *vehicleUseCase) UpdateVehicleType(ctx context.Context, id string, request dto.UpdateVehicleTypeRequest) (*dto.VehicleTypeResponse, error) {
	err := validation.ValidateVehicleTypeUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle type update request")
		return nil, errors.ErrInvalidVehicleTypeUpdateRequest
	}
	vehicleType := entity.VehicleType{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return nil, errors.ErrInvalidVehicleTypeID
	}
	vehicleType.ID = uint(uintID)

	if request.Name != nil {
		vehicleType.Name = *request.Name
	}
	if request.Description != nil {
		vehicleType.Description = *request.Description
	}

	err = uc.vehicleRepository.UpdateVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to update vehicle type")
		return nil, errors.ErrFailedToUpdateVehicleType
	}
	return &dto.VehicleTypeResponse{
		ID:          vehicleType.ID,
		Name:        vehicleType.Name,
		Description: vehicleType.Description,
	}, nil
}

func (uc *vehicleUseCase) DeleteVehicleType(ctx context.Context, id string) error {
	vehicleType := entity.VehicleType{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle type id")
		return errors.ErrInvalidVehicleTypeID
	}
	vehicleType.ID = uint(uintID)
	err = uc.vehicleRepository.DeleteVehicleType(ctx, &vehicleType)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle type")
		return errors.ErrFailedToDeleteVehicleType
	}
	return nil
}

// Brands
func (uc *vehicleUseCase) ListBrands(ctx context.Context) (*dto.ListVehicleBrandsResponse, error) {
	brands := []entity.VehicleBrand{}
	err := uc.vehicleRepository.ListBrands(ctx, &brands)
	if err != nil {
		logger.Error(err, "Failed to list vehicle brands")
		return nil, errors.ErrFailedToListVehicleBrands
	}
	brandsResponse := []dto.VehicleBrandResponse{}
	for _, brand := range brands {
		brandsResponse = append(brandsResponse, dto.VehicleBrandResponse{
			ID:          brand.ID,
			Name:        brand.Name,
			Description: brand.Description,
			VehicleType: brand.VehicleType.Name,
		})
	}
	return &dto.ListVehicleBrandsResponse{Brands: brandsResponse}, nil
}

func (uc *vehicleUseCase) GetBrand(ctx context.Context, id string) (*dto.VehicleBrandResponse, error) {
	brand := entity.VehicleBrand{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}
	brand.ID = uint(uintID)
	err = uc.vehicleRepository.GetBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to get vehicle brand")
		return nil, errors.ErrFailedToGetVehicleBrand
	}
	return &dto.VehicleBrandResponse{
		ID:          brand.ID,
		Name:        brand.Name,
		Description: brand.Description,
		VehicleType: brand.VehicleType.Name,
	}, nil
}

func (uc *vehicleUseCase) ListBrandsByType(ctx context.Context, typeID string) (*dto.ListVehicleBrandsResponse, error) {
	brands := []entity.VehicleBrand{}
	err := uc.vehicleRepository.ListBrandsByType(ctx, &brands, typeID)
	if err != nil {
		logger.Error(err, "Failed to list vehicle brands by type")
		return nil, errors.ErrFailedToListVehicleBrandsByType
	}
	brandsResponse := []dto.VehicleBrandResponse{}
	for _, brand := range brands {
		brandsResponse = append(brandsResponse, dto.VehicleBrandResponse{
			ID:          brand.ID,
			Name:        brand.Name,
			Description: brand.Description,
			VehicleType: brand.VehicleType.Name,
		})
	}
	return &dto.ListVehicleBrandsResponse{Brands: brandsResponse}, nil
}

func (uc *vehicleUseCase) CreateBrand(ctx context.Context, request dto.CreateVehicleBrandRequest) (*dto.VehicleBrandResponse, error) {
	err := validation.ValidateVehicleBrandCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle brand create request")
		return nil, errors.ErrInvalidVehicleBrandCreateRequest
	}

	brand := entity.VehicleBrand{
		Name:          request.Name,
		Description:   request.Description,
		VehicleTypeID: request.VehicleTypeID,
	}
	err = uc.vehicleRepository.CreateBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to create vehicle brand")
		return nil, errors.ErrFailedToCreateVehicleBrand
	}
	return &dto.VehicleBrandResponse{
		ID:          brand.ID,
		Name:        brand.Name,
		Description: brand.Description,
		VehicleType: brand.VehicleType.Name,
	}, nil
}

func (uc *vehicleUseCase) UpdateBrand(ctx context.Context, id string, request dto.UpdateVehicleBrandRequest) (*dto.VehicleBrandResponse, error) {
	err := validation.ValidateVehicleBrandUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle brand update request")
		return nil, errors.ErrInvalidVehicleBrandUpdateRequest
	}

	brand := entity.VehicleBrand{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return nil, errors.ErrInvalidVehicleBrandID
	}
	brand.ID = uint(uintID)

	if request.Name != nil {
		brand.Name = *request.Name
	}
	if request.Description != nil {
		brand.Description = *request.Description
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
		ID:          brand.ID,
		Name:        brand.Name,
		Description: brand.Description,
		VehicleType: brand.VehicleType.Name,
	}, nil
}

func (uc *vehicleUseCase) DeleteBrand(ctx context.Context, id string) error {
	brand := entity.VehicleBrand{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle brand id")
		return errors.ErrInvalidVehicleBrandID
	}
	brand.ID = uint(uintID)
	err = uc.vehicleRepository.DeleteBrand(ctx, &brand)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle brand")
		return errors.ErrFailedToDeleteVehicleBrand
	}
	return nil
}

// Models
func (uc *vehicleUseCase) ListModels(ctx context.Context) (*dto.ListVehicleModelsResponse, error) {
	models := []entity.VehicleModel{}
	err := uc.vehicleRepository.ListModels(ctx, &models)
	if err != nil {
		logger.Error(err, "Failed to list vehicle models")
		return nil, errors.ErrFailedToListVehicleModels
	}
	modelsResponse := []dto.VehicleModelResponse{}
	for _, model := range models {
		modelsResponse = append(modelsResponse, dto.VehicleModelResponse{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			BrandID:     model.BrandID,
			StartYear:   model.StartYear,
			EndYear:     model.EndYear,
		})
	}
	return &dto.ListVehicleModelsResponse{Models: modelsResponse}, nil
}

func (uc *vehicleUseCase) ListModelsByBrand(ctx context.Context, brandID string) (*dto.ListVehicleModelsResponse, error) {
	models := []entity.VehicleModel{}
	err := uc.vehicleRepository.ListModelsByBrand(ctx, &models, brandID)
	if err != nil {
		logger.Error(err, "Failed to list vehicle models by brand")
		return nil, errors.ErrFailedToListVehicleModelsByBrand
	}
	modelsResponse := []dto.VehicleModelResponse{}
	for _, model := range models {
		modelsResponse = append(modelsResponse, dto.VehicleModelResponse{
			ID:          model.ID,
			Name:        model.Name,
			Description: model.Description,
			BrandID:     model.BrandID,
			StartYear:   model.StartYear,
			EndYear:     model.EndYear,
		})
	}
	return &dto.ListVehicleModelsResponse{Models: modelsResponse}, nil
}

func (uc *vehicleUseCase) GetModel(ctx context.Context, id string) (*dto.VehicleModelResponse, error) {
	model := entity.VehicleModel{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	model.ID = uint(uintID)
	err = uc.vehicleRepository.GetModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to get vehicle model")
		return nil, errors.ErrFailedToGetVehicleModel
	}
	return &dto.VehicleModelResponse{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		BrandID:     model.BrandID,
		StartYear:   model.StartYear,
		EndYear:     model.EndYear,
	}, nil
}

func (uc *vehicleUseCase) CreateModel(ctx context.Context, request dto.CreateVehicleModelRequest) (*dto.VehicleModelResponse, error) {
	err := validation.ValidateVehicleModelCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle model create request")
		return nil, errors.ErrInvalidVehicleModelCreateRequest
	}

	model := entity.VehicleModel{
		Name:        request.Name,
		Description: request.Description,
		BrandID:     request.BrandID,
		StartYear:   request.StartYear,
		EndYear:     request.EndYear,
	}
	err = uc.vehicleRepository.CreateModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to create vehicle model")
		return nil, errors.ErrFailedToCreateVehicleModel
	}
	return &dto.VehicleModelResponse{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		BrandID:     model.BrandID,
		StartYear:   model.StartYear,
		EndYear:     model.EndYear,
	}, nil
}

func (uc *vehicleUseCase) UpdateModel(ctx context.Context, id string, request dto.UpdateVehicleModelRequest) (*dto.VehicleModelResponse, error) {
	err := validation.ValidateVehicleModelUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle model update request")
		return nil, errors.ErrInvalidVehicleModelUpdateRequest
	}

	model := entity.VehicleModel{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return nil, errors.ErrInvalidVehicleModelID
	}
	model.ID = uint(uintID)

	if request.Name != nil {
		model.Name = *request.Name
	}
	if request.Description != nil {
		model.Description = *request.Description
	}
	if request.BrandID != nil {
		model.BrandID = *request.BrandID
	}
	if request.StartYear != nil {
		model.StartYear = *request.StartYear
	}
	if request.EndYear != nil {
		model.EndYear = *request.EndYear
	}

	err = uc.vehicleRepository.UpdateModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to update vehicle model")
		return nil, errors.ErrFailedToUpdateVehicleModel
	}
	return &dto.VehicleModelResponse{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
		BrandID:     model.BrandID,
		StartYear:   model.StartYear,
		EndYear:     model.EndYear,
	}, nil
}

func (uc *vehicleUseCase) DeleteModel(ctx context.Context, id string) error {
	model := entity.VehicleModel{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle model id")
		return errors.ErrInvalidVehicleModelID
	}
	model.ID = uint(uintID)
	err = uc.vehicleRepository.DeleteModel(ctx, &model)
	if err != nil {
		logger.Error(err, "Failed to delete vehicle model")
		return errors.ErrFailedToDeleteVehicleModel
	}
	return nil
}

// Generations
func (uc *vehicleUseCase) ListGenerations(ctx context.Context) (*dto.ListVehicleGenerationsResponse, error) {
	generations := []entity.VehicleGeneration{}
	err := uc.vehicleRepository.ListGenerations(ctx, &generations)
	if err != nil {
		logger.Error(err, "Failed to list vehicle generations")
		return nil, errors.ErrFailedToListVehicleGenerations
	}
	generationsResponse := []dto.VehicleGenerationResponse{}
	for _, generation := range generations {
		generationsResponse = append(generationsResponse, dto.VehicleGenerationResponse{
			ID:              generation.ID,
			Name:            generation.Name,
			Description:     generation.Description,
			ModelID:         generation.ModelID,
			StartYear:       generation.StartYear,
			EndYear:         generation.EndYear,
			EngineType:      generation.EngineType,
			AssemblyType:    generation.AssemblyType,
			Assembler:       generation.Assembler,
			Transmission:    generation.Transmission,
			EngineSize:      generation.EngineSize,
			BodyStyle:       generation.BodyStyle,
			SpecialFeatures: generation.SpecialFeatures,
		})
	}
	return &dto.ListVehicleGenerationsResponse{Generations: generationsResponse}, nil
}

func (uc *vehicleUseCase) GetGeneration(ctx context.Context, id string) (*dto.VehicleGenerationResponse, error) {
	generation := entity.VehicleGeneration{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle generation id")
		return nil, errors.ErrInvalidVehicleGenerationID
	}
	generation.ID = uint(uintID)
	err = uc.vehicleRepository.GetGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to get vehicle generation")
		return nil, errors.ErrFailedToGetVehicleGeneration
	}
	return &dto.VehicleGenerationResponse{
		ID:              generation.ID,
		Name:            generation.Name,
		Description:     generation.Description,
		ModelID:         generation.ModelID,
		StartYear:       generation.StartYear,
		EndYear:         generation.EndYear,
		EngineType:      generation.EngineType,
		AssemblyType:    generation.AssemblyType,
		Assembler:       generation.Assembler,
		Transmission:    generation.Transmission,
		EngineSize:      generation.EngineSize,
		BodyStyle:       generation.BodyStyle,
		SpecialFeatures: generation.SpecialFeatures,
	}, nil
}

func (uc *vehicleUseCase) ListGenerationsByModel(ctx context.Context, modelID string) (*dto.ListVehicleGenerationsResponse, error) {
	generations := []entity.VehicleGeneration{}
	err := uc.vehicleRepository.ListGenerationsByModel(ctx, &generations, modelID)
	if err != nil {
		logger.Error(err, "Failed to list vehicle generations by model")
		return nil, errors.ErrFailedToListVehicleGenerationsByModel
	}

	generationsResponse := []dto.VehicleGenerationResponse{}
	for _, generation := range generations {
		generationsResponse = append(generationsResponse, dto.VehicleGenerationResponse{
			ID:              generation.ID,
			Name:            generation.Name,
			Description:     generation.Description,
			ModelID:         generation.ModelID,
			StartYear:       generation.StartYear,
			EndYear:         generation.EndYear,
			EngineType:      generation.EngineType,
			AssemblyType:    generation.AssemblyType,
			Assembler:       generation.Assembler,
			Transmission:    generation.Transmission,
			EngineSize:      generation.EngineSize,
			BodyStyle:       generation.BodyStyle,
			SpecialFeatures: generation.SpecialFeatures,
		})
	}
	return &dto.ListVehicleGenerationsResponse{Generations: generationsResponse}, nil
}

func (uc *vehicleUseCase) CreateGeneration(ctx context.Context, request dto.CreateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error) {
	err := validation.ValidateVehicleGenerationCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle generation create request")
		return nil, errors.ErrInvalidVehicleGenerationCreateRequest
	}

	generation := entity.VehicleGeneration{
		Name:            request.Name,
		Description:     request.Description,
		ModelID:         request.ModelID,
		StartYear:       request.StartYear,
		EndYear:         request.EndYear,
		EngineType:      request.EngineType,
		AssemblyType:    request.AssemblyType,
		Assembler:       request.Assembler,
		Transmission:    request.Transmission,
		EngineSize:      request.EngineSize,
		BodyStyle:       request.BodyStyle,
		SpecialFeatures: request.SpecialFeatures,
	}
	err = uc.vehicleRepository.CreateGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to create vehicle generation")
		return nil, errors.ErrFailedToCreateVehicleGeneration
	}
	return &dto.VehicleGenerationResponse{
		ID:              generation.ID,
		Name:            generation.Name,
		Description:     generation.Description,
		ModelID:         generation.ModelID,
		StartYear:       generation.StartYear,
		EndYear:         generation.EndYear,
		EngineType:      generation.EngineType,
		AssemblyType:    generation.AssemblyType,
		Assembler:       generation.Assembler,
		Transmission:    generation.Transmission,
		EngineSize:      generation.EngineSize,
		BodyStyle:       generation.BodyStyle,
		SpecialFeatures: generation.SpecialFeatures,
	}, nil
}

func (uc *vehicleUseCase) UpdateGeneration(ctx context.Context, id string, request dto.UpdateVehicleGenerationRequest) (*dto.VehicleGenerationResponse, error) {
	err := validation.ValidateVehicleGenerationUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate vehicle generation update request")
		return nil, errors.ErrInvalidVehicleGenerationUpdateRequest
	}

	generation := entity.VehicleGeneration{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle generation id")
		return nil, errors.ErrInvalidVehicleGenerationID
	}
	generation.ID = uint(uintID)

	if request.Name != nil {
		generation.Name = *request.Name
	}
	if request.Description != nil {
		generation.Description = *request.Description
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
	if request.EngineType != nil {
		generation.EngineType = *request.EngineType
	}
	if request.AssemblyType != nil {
		generation.AssemblyType = *request.AssemblyType
	}
	if request.Assembler != nil {
		generation.Assembler = *request.Assembler
	}
	if request.Transmission != nil {
		generation.Transmission = *request.Transmission
	}
	if request.EngineSize != nil {
		generation.EngineSize = *request.EngineSize
	}
	if request.BodyStyle != nil {
		generation.BodyStyle = *request.BodyStyle
	}
	if request.SpecialFeatures != nil {
		generation.SpecialFeatures = *request.SpecialFeatures
	}

	err = uc.vehicleRepository.UpdateGeneration(ctx, &generation)
	if err != nil {
		logger.Error(err, "Failed to update vehicle generation")
		return nil, errors.ErrFailedToUpdateVehicleGeneration
	}
	return &dto.VehicleGenerationResponse{
		ID:              generation.ID,
		Name:            generation.Name,
		Description:     generation.Description,
		ModelID:         generation.ModelID,
		StartYear:       generation.StartYear,
		EndYear:         generation.EndYear,
		EngineType:      generation.EngineType,
		AssemblyType:    generation.AssemblyType,
		Assembler:       generation.Assembler,
		Transmission:    generation.Transmission,
		EngineSize:      generation.EngineSize,
		BodyStyle:       generation.BodyStyle,
		SpecialFeatures: generation.SpecialFeatures,
	}, nil
}

func (uc *vehicleUseCase) DeleteGeneration(ctx context.Context, id string) error {
	generation := entity.VehicleGeneration{}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle generation id")
		return errors.ErrInvalidVehicleGenerationID
	}
	generation.ID = uint(uintID)
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

	userVehicle := entity.UserVehicle{
		UserID:         userID,
		GenerationID:   request.GenerationID,
		Name:           request.Name,
		ProductionYear: request.ProductionYear,
		Color:          request.Color,
		LicensePlate:   request.LicensePlate,
		VIN:            request.VIN,
		CurrentMileage: request.CurrentMileage,
		PurchaseDate:   request.PurchaseDate,
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
	err := uc.vehicleRepository.ListUserVehicles(ctx, userID, &userVehicles)
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
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, vehicleID, &userVehicle)
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

	userVehicle := entity.UserVehicle{
		UserID: userID,
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
		userVehicle.PurchaseDate = *request.PurchaseDate
	}

	uintUserVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse user vehicle id")
		return nil, errors.ErrInvalidUserVehicleID
	}
	userVehicle.ID = uint(uintUserVehicleID)
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
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, vehicleID, &userVehicle)
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
