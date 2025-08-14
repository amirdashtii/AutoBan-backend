package dto

import (
	"time"

	"github.com/google/uuid"
)

// Type
// CreateVehicleTypeRequest represents the request for creating a new vehicle type
// @Description Vehicle type creation request
type CreateVehicleTypeRequest struct {
	// Name of the vehicle type in Persian
	NameFa string `json:"name_fa" validate:"required" example:"خودرو"`
	// Name of the vehicle type in English
	NameEn string `json:"name_en" validate:"required" example:"Car"`
	// Description of the vehicle type in Persian
	DescriptionFa string `json:"description_fa" example:"خودروی سواری استاندارد"`
	// Description of the vehicle type in English
	DescriptionEn string `json:"description_en" example:"A standard passenger car"`
}

// UpdateVehicleTypeRequest represents the request for updating vehicle type
// @Description Vehicle type update request
type UpdateVehicleTypeRequest struct {
	// Name of the vehicle type in Persian
	NameFa *string `json:"name_fa" example:"خودرو"`
	// Name of the vehicle type in English
	NameEn *string `json:"name_en" example:"Car"`
	// Description of the vehicle type in Persian
	DescriptionFa *string `json:"description_fa" example:"خودروی سواری استاندارد"`
	// Description of the vehicle type in English
	DescriptionEn *string `json:"description_en" example:"A standard passenger car"`
}

// VehicleTypeResponse represents the response for vehicle type data
// @Description Vehicle type response
type VehicleTypeResponse struct {
	// ID of the vehicle type
	ID uint64 `json:"id"`
	// Name of the vehicle type in Persian
	NameFa string `json:"name_fa"`
	// Name of the vehicle type in English
	NameEn string `json:"name_en"`
	// Description of the vehicle type in Persian
	DescriptionFa string `json:"description_fa"`
	// Description of the vehicle type in English
	DescriptionEn string `json:"description_en"`
}

// Brand
// CreateVehicleBrandRequest represents the request for creating a new vehicle brand
// @Description Vehicle brand creation request
type CreateVehicleBrandRequest struct {
	// Name of the vehicle brand
	NameFa string `json:"name_fa" validate:"required" example:"تویوتا"`
	// Name of the vehicle brand
	NameEn string `json:"name_en" validate:"required" example:"Toyota"`
	// Description of the vehicle brand
	DescriptionFa string `json:"description_fa" example:"برند خودرویی از کشور ژاپن"`
	// Description of the vehicle brand
	DescriptionEn string `json:"description_en" example:"A popular Japanese car brand"`
}

// UpdateVehicleBrandRequest represents the request for updating vehicle brand
// @Description Vehicle brand update request
type UpdateVehicleBrandRequest struct {
	// ID of the vehicle type
	VehicleTypeID *uint64 `json:"vehicle_type_id" example:"1"`
	// Name of the vehicle brand
	NameFa *string `json:"name_fa" example:"تویوتا"`
	// Name of the vehicle brand
	NameEn *string `json:"name_en" example:"Toyota"`
	// Description of the vehicle brand
	DescriptionFa *string `json:"description_fa" example:"برند خودرویی از کشور ژاپن"`
	// Description of the vehicle brand
	DescriptionEn *string `json:"description_en" example:"A popular Japanese car brand"`
}

// VehicleBrandResponse represents the response for vehicle brand data
// @Description Vehicle brand response
type VehicleBrandResponse struct {
	// ID of the vehicle brand
	ID uint64 `json:"id"`
	// ID of the vehicle type
	VehicleTypeID uint64 `json:"vehicle_type_id"`
	// Name of the vehicle brand
	NameFa string `json:"name_fa"`
	// Name of the vehicle brand
	NameEn string `json:"name_en"`
	// Description of the vehicle brand
	DescriptionFa string `json:"description_fa"`
	// Description of the vehicle brand
	DescriptionEn string `json:"description_en"`
}

// Model
// CreateVehicleModelRequest represents the request for creating a new vehicle model
// @Description Vehicle model creation request
type CreateVehicleModelRequest struct {
	// Name of the vehicle model
	NameFa string `json:"name_fa" validate:"required" example:"Camry"`
	// Name of the vehicle model
	NameEn string `json:"name_en" validate:"required" example:"Camry"`
	// Description of the vehicle model
	DescriptionFa string `json:"description_fa" example:"A mid-size sedan"`
	// Description of the vehicle model
	DescriptionEn string `json:"description_en" example:"A mid-size sedan"`
}

// UpdateVehicleModelRequest represents the request for updating vehicle model
// @Description Vehicle model update request
type UpdateVehicleModelRequest struct {
	// ID of the vehicle brand
	BrandID *uint64 `json:"brand_id" example:"1"`
	// Name of the vehicle model
	NameFa *string `json:"name_fa" example:"Camry"`
	// Name of the vehicle model
	NameEn *string `json:"name_en" example:"Camry"`
	// Description of the vehicle model
	DescriptionFa *string `json:"description_fa" example:"A mid-size sedan"`
	// Description of the vehicle model
	DescriptionEn *string `json:"description_en" example:"A mid-size sedan"`
}

// VehicleModelResponse represents the response for vehicle model data
// @Description Vehicle model response
type VehicleModelResponse struct {
	// ID of the vehicle model
	ID uint64 `json:"id"`
	// ID of the vehicle brand
	BrandID uint64 `json:"brand_id"`
	// Name of the vehicle model
	NameFa string `json:"name_fa"`
	// Name of the vehicle model
	NameEn string `json:"name_en"`
	// Description of the vehicle model
	DescriptionFa string `json:"description_fa"`
	// Description of the vehicle model
	DescriptionEn string `json:"description_en"`
}

// Generation
// CreateVehicleGenerationRequest represents the request for creating a new vehicle generation
// @Description Vehicle generation creation request
type CreateVehicleGenerationRequest struct {
	// Name of the vehicle generation
	NameFa string `json:"name_fa" validate:"required" example:"Generation Name"`
	// Name of the vehicle generation
	NameEn string `json:"name_en" validate:"required" example:"Generation Name"`
	// Description of the vehicle generation
	DescriptionFa string `json:"description_fa" example:"A brief description of the generation"`
	// Description of the vehicle generation
	DescriptionEn string `json:"description_en" example:"A brief description of the generation"`
	// Start year of the vehicle generation
	StartYear int `json:"start_year" validate:"omitempty,year" example:"2020"`
	// End year of the vehicle generation
	EndYear int `json:"end_year" validate:"omitempty,year" example:"2022"`
	// Body style of the vehicle generation
	BodyStyleFa string `json:"body_style_fa" example:"سدان"`
	BodyStyleEn string `json:"body_style_en" example:"Sedan"`
	// Engine of the vehicle generation
	Engine string `json:"engine" example:"1.6 TDI"`
	// Engine volume of the vehicle generation
	EngineVolume int `json:"engine_volume" example:"1600"`
	// Cylinders of the vehicle generation
	Cylinders int `json:"cylinders" example:"4"`
	// Drivetrain of the vehicle generation
	DrivetrainFa string `json:"drivetrain_fa" example:"دودیفرانسیل"`
	DrivetrainEn string `json:"drivetrain_en" example:"4WD"`
	// Gearbox of the vehicle generation
	Gearbox string `json:"gearbox" example:"Automatic"`
	// Fuel type of the vehicle generation
	FuelType string `json:"fuel_type" example:"Gasoline"`
	// Battery of the vehicle generation
	Battery string `json:"battery" example:"Li-ion"`
	// Seller of the vehicle generation
	Seller string `json:"seller" example:"Toyota"`
	// Assembly type of the vehicle generation
	AssemblyType string `json:"assembly_type" example:"CKD"`
	// Assembler of the vehicle generation
	Assembler string `json:"assembler" example:"Toyota"`
}

// UpdateVehicleGenerationRequest represents the request for updating vehicle generation
// @Description Vehicle generation update request
type UpdateVehicleGenerationRequest struct {
	// ID of the vehicle model
	ModelID *uint64 `json:"model_id" example:"1"`
	// Name of the vehicle generation
	NameFa *string `json:"name_fa" example:"Generation Name"`
	// Name of the vehicle generation
	NameEn *string `json:"name_en" example:"Generation Name"`
	// Description of the vehicle generation
	DescriptionFa *string `json:"description_fa" example:"A brief description of the generation"`
	// Description of the vehicle generation
	DescriptionEn *string `json:"description_en" example:"A brief description of the generation"`
	// Start year of the vehicle generation
	StartYear *int `json:"start_year" validate:"omitempty,year" example:"2020"`
	// End year of the vehicle generation
	EndYear *int `json:"end_year" validate:"omitempty,year" example:"2022"`
	// Body style of the vehicle generation
	BodyStyleFa *string `json:"body_style_fa" example:"سدان"`
	BodyStyleEn *string `json:"body_style_en" example:"Sedan"`
	// Engine of the vehicle generation
	Engine *string `json:"engine" example:"1.6 TDI"`
	// Engine volume of the vehicle generation
	EngineVolume *int `json:"engine_volume" example:"1600"`
	// Cylinders of the vehicle generation
	Cylinders *int `json:"cylinders" example:"4"`
	// Drivetrain of the vehicle generation
	DrivetrainFa *string `json:"drivetrain_fa" example:"دودیفرانسیل"`
	DrivetrainEn *string `json:"drivetrain_en" example:"4WD"`
	// Gearbox of the vehicle generation
	Gearbox *string `json:"gearbox" example:"Automatic"`
	// Fuel type of the vehicle generation
	FuelType *string `json:"fuel_type" example:"Gasoline"`
	// Battery of the vehicle generation
	Battery *string `json:"battery" example:"Li-ion"`
	// Seller of the vehicle generation
	Seller *string `json:"seller" example:"Toyota"`
	// Assembly type of the vehicle generation
	AssemblyType *string `json:"assembly_type" example:"CKD"`
	// Assembler of the vehicle generation
	Assembler *string `json:"assembler" example:"Toyota"`
}

// VehicleGenerationResponse represents the response for vehicle generation data
// @Description Vehicle generation response
type VehicleGenerationResponse struct {
	// ID of the vehicle generation
	ID uint64 `json:"id"`
	// ID of the vehicle model
	ModelID uint64 `json:"model_id"`
	// Name of the vehicle generation
	NameFa string `json:"name_fa"`
	// Name of the vehicle generation
	NameEn string `json:"name_en"`
	// Description of the vehicle generation
	DescriptionFa string `json:"description_fa"`
	// Description of the vehicle generation
	DescriptionEn string `json:"description_en"`
	// Start year of the vehicle generation
	StartYear int `json:"start_year"`
	// End year of the vehicle generation
	EndYear int `json:"end_year"`
	// Body style of the vehicle generation
	BodyStyleFa string `json:"body_style_fa"`
	// Body style of the vehicle generation
	BodyStyleEn string `json:"body_style_en"`
	// Engine of the vehicle generation
	Engine string `json:"engine"`
	// Engine volume of the vehicle generation
	EngineVolume int `json:"engine_volume"`
	// Cylinders of the vehicle generation
	Cylinders int `json:"cylinders"`
	// Drivetrain of the vehicle generation
	DrivetrainFa string `json:"drivetrain_fa"`
	DrivetrainEn string `json:"drivetrain_en"`
	// Gearbox of the vehicle generation
	Gearbox string `json:"gearbox"`
	// Fuel type of the vehicle generation
	FuelType string `json:"fuel_type"`
	// Battery of the vehicle generation
	Battery string `json:"battery"`
	// Seller of the vehicle generation
	Seller string `json:"seller"`
	// Assembly type of the vehicle generation
	AssemblyType string `json:"assembly_type"`
	// Assembler of the vehicle generation
	Assembler string `json:"assembler"`
}

// UserVehicle
// CreateUserVehicleRequest represents the request for adding a new user vehicle
// @Description User vehicle creation request
type CreateUserVehicleRequest struct {
	// Name of the user vehicle
	Name string `json:"name" validate:"required" example:"My Car"`
	// ID of the vehicle generation
	GenerationID uint64 `json:"generation_id" validate:"required" example:"1"`
	// Production year of the user vehicle
	ProductionYear int `json:"production_year" validate:"omitempty,year" example:"2020"`
	// Color of the user vehicle
	Color string `json:"color" example:"Red"`
	// License plate of the user vehicle
	LicensePlate string `json:"license_plate" validate:"omitempty,iranian_license_plate" example:"12آلف345-67"`
	// VIN of the user vehicle
	VIN string `json:"vin" validate:"omitempty,iranian_vin" example:"1HGCM82633A123456"`
	// Current mileage of the user vehicle
	CurrentMileage int `json:"current_mileage" validate:"omitempty,min=0" example:"15000"`
	// Purchase date of the user vehicle
	PurchaseDate string `json:"purchase_date" validate:"omitempty,date" example:"2020-01-01"`
}

// UpdateUserVehicleRequest represents the request for updating user vehicle
// @Description User vehicle update request
type UpdateUserVehicleRequest struct {
	// Name of the user vehicle
	Name *string `json:"name" example:"My Car"`
	// ID of the vehicle generation
	GenerationID *uint64 `json:"generation_id" example:"1"`
	// Production year of the user vehicle
	ProductionYear *int `json:"production_year" validate:"omitempty,year" example:"2020"`
	// Color of the user vehicle
	Color *string `json:"color" example:"Red"`
	// License plate of the user vehicle
	LicensePlate *string `json:"license_plate" validate:"omitempty,iranian_license_plate" example:"۱۲الف۳۴۵۶۸"`
	// VIN of the user vehicle
	VIN *string `json:"vin" validate:"omitempty,iranian_vin" example:"1HGCM82633A123456"`
	// Current mileage of the user vehicle
	CurrentMileage *int `json:"current_mileage" validate:"omitempty,min=0" example:"15000"`
	// Purchase date of the user vehicle
	PurchaseDate *string `json:"purchase_date" validate:"omitempty,date" example:"2020-01-01"`
}

// UserVehicleResponse represents the response for user vehicle data
// @Description User vehicle response
type UserVehicleResponse struct {
	// ID of the user vehicle
	ID uint64 `json:"id"`
	// ID of the user
	UserID uuid.UUID `json:"user_id"`
	// Name of the user vehicle
	Name string `json:"name"`
	// ID of the vehicle generation
	GenerationID uint64 `json:"generation_id"`
	// Production year of the user vehicle
	ProductionYear int `json:"production_year"`
	// Color of the user vehicle
	Color string `json:"color"`
	// License plate of the user vehicle
	LicensePlate string `json:"license_plate"`
	// VIN of the user vehicle
	VIN string `json:"vin"`
	// Current mileage of the user vehicle
	CurrentMileage int `json:"current_mileage"`
	// Purchase date of the user vehicle
	PurchaseDate time.Time `json:"purchase_date"`
	// Expanded path (optional)
	Type       *VehicleTypeResponse       `json:"type,omitempty"`
	Brand      *VehicleBrandResponse      `json:"brand,omitempty"`
	Model      *VehicleModelResponse      `json:"model,omitempty"`
	Generation *VehicleGenerationResponse `json:"generation,omitempty"`
}

// ListVehicleTypesResponse represents the response for listing vehicle types
// @Description List of vehicle types
type ListVehicleTypesResponse struct {
	// List of vehicle types
	Types []VehicleTypeResponse `json:"types"`
}

// ListVehicleBrandsResponse represents the response for listing vehicle brands
// @Description List of vehicle brands
type ListVehicleBrandsResponse struct {
	// List of vehicle brands
	Brands []VehicleBrandResponse `json:"brands"`
}

// ListVehicleModelsResponse represents the response for listing vehicle models
// @Description List of vehicle models
type ListVehicleModelsResponse struct {
	// List of vehicle models
	Models []VehicleModelResponse `json:"models"`
}

// ListVehicleGenerationsResponse represents the response for listing vehicle generations
// @Description List of vehicle generations
type ListVehicleGenerationsResponse struct {
	// List of vehicle generations
	Generations []VehicleGenerationResponse `json:"generations"`
}

// ListUserVehiclesResponse represents the response for listing user vehicles
// @Description List of user vehicles
type ListUserVehiclesResponse struct {
	// List of user vehicles
	Vehicles []UserVehicleResponse `json:"vehicles"`
}

// Hierarchical Vehicle DTOs for complete tree structure

// VehicleGenerationTreeResponse represents a vehicle generation in the tree structure
// @Description Vehicle generation in hierarchical tree
type VehicleGenerationTreeResponse struct {
	// ID of the vehicle generation
	ID uint64 `json:"id"`
	// Name of the vehicle generation
	NameFa string `json:"name_fa"`
	// Name of the vehicle generation
	NameEn string `json:"name_en"`
}

// VehicleModelTreeResponse represents a vehicle model in the tree structure
// @Description Vehicle model in hierarchical tree
type VehicleModelTreeResponse struct {
	// ID of the vehicle model
	ID uint64 `json:"id"`
	// Name of the vehicle model
	NameFa string `json:"name_fa"`
	// Name of the vehicle model
	NameEn string `json:"name_en"`
	// List of generations for this model
	Generations []VehicleGenerationTreeResponse `json:"generations"`
}

// VehicleBrandTreeResponse represents a vehicle brand in the tree structure
// @Description Vehicle brand in hierarchical tree
type VehicleBrandTreeResponse struct {
	// ID of the vehicle brand
	ID uint64 `json:"id"`
	// Name of the vehicle brand
	NameFa string `json:"name_fa"`
	// Name of the vehicle brand
	NameEn string `json:"name_en"`
	// List of models for this brand
	Models []VehicleModelTreeResponse `json:"models"`
}

// VehicleTypeTreeResponse represents a vehicle type in the tree structure
// @Description Vehicle type in hierarchical tree
type VehicleTypeTreeResponse struct {
	// ID of the vehicle type
	ID uint64 `json:"id"`
	// Name of the vehicle type
	NameFa string `json:"name_fa"`
	// Name of the vehicle type
	NameEn string `json:"name_en"`
	// List of brands for this type
	Brands []VehicleBrandTreeResponse `json:"brands"`
}

// CompleteVehicleHierarchyResponse represents the complete vehicle hierarchy
// @Description Complete vehicle hierarchy with all types, brands, models, and generations
type CompleteVehicleHierarchyResponse struct {
	// List of all vehicle types with their complete hierarchy
	VehicleTypes []VehicleTypeTreeResponse `json:"vehicle_types"`
	// Total count of vehicle types
	TotalTypes int `json:"total_types"`
	// Total count of vehicle brands
	TotalBrands int `json:"total_brands"`
	// Total count of vehicle models
	TotalModels int `json:"total_models"`
	// Total count of vehicle generations
	TotalGenerations int `json:"total_generations"`
}
