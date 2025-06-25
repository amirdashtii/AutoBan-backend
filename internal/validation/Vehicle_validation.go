package validation

import (
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/go-playground/validator/v10"
)

func ValidateVehicleTypeCreateRequest(request dto.CreateVehicleTypeRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleTypeCreateRequest
	}
	return nil
}

func ValidateVehicleTypeUpdateRequest(request dto.UpdateVehicleTypeRequest) error {
	if request.Name == nil && request.Description == nil {
		return errors.ErrInvalidVehicleTypeUpdateRequest
	}

	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleTypeUpdateRequest
	}

	return nil
}

func ValidateVehicleBrandCreateRequest(request dto.CreateVehicleBrandRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleBrandCreateRequest
	}
	return nil
}

func ValidateVehicleBrandUpdateRequest(request dto.UpdateVehicleBrandRequest) error {
	// Check if at least one field has a value
	if request.VehicleTypeID == nil && request.Name == nil && request.Description == nil {
		return errors.ErrInvalidVehicleBrandUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleBrandUpdateRequest
	}

	return nil
}

func ValidateVehicleModelCreateRequest(request dto.CreateVehicleModelRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleModelCreateRequest
	}
	return nil
}

func ValidateVehicleModelUpdateRequest(request dto.UpdateVehicleModelRequest) error {
	// Check if at least one field has a value
	if request.BrandID == nil && request.Name == nil && request.Description == nil && request.StartYear == nil && request.EndYear == nil {
		return errors.ErrInvalidVehicleModelUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleModelUpdateRequest
	}

	return nil
}

func ValidateVehicleGenerationCreateRequest(request dto.CreateVehicleGenerationRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleGenerationCreateRequest
	}
	return nil
}

func ValidateVehicleGenerationUpdateRequest(request dto.UpdateVehicleGenerationRequest) error {
	// Check if at least one field has a value
	if request.Name == nil && request.Description == nil && request.ModelID == nil &&
		request.StartYear == nil && request.EndYear == nil && request.EngineType == nil &&
		request.AssemblyType == nil && request.Assembler == nil && request.Transmission == nil &&
		request.EngineSize == nil && request.BodyStyle == nil && request.SpecialFeatures == nil {
		return errors.ErrInvalidVehicleGenerationUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleGenerationUpdateRequest
	}

	return nil
}

func ValidateUserVehicleCreateRequest(request dto.CreateUserVehicleRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidUserVehicleCreateRequest
	}
	return nil
}

func ValidateUserVehicleUpdateRequest(request dto.UpdateUserVehicleRequest) error {
	// Check if at least one field has a value
	if request.Name == nil && request.GenerationID == nil && request.ProductionYear == nil &&
		request.Color == nil && request.LicensePlate == nil && request.VIN == nil &&
		request.CurrentMileage == nil && request.PurchaseDate == nil {
		return errors.ErrInvalidUserVehicleUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidUserVehicleUpdateRequest
	}

	return nil
}
