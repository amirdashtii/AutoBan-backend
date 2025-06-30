package dto

import (
	"time"

	"github.com/google/uuid"
)

// Type
// CreateVehicleTypeRequest represents the request for creating a new vehicle type
// @Description Vehicle type creation request
type CreateVehicleTypeRequest struct {
	// Name of the vehicle type
	Name string `json:"name" validate:"required" example:"car"`
	// Description of the vehicle type
	Description string `json:"description" example:"A standard passenger car"`
}

// UpdateVehicleTypeRequest represents the request for updating vehicle type
// @Description Vehicle type update request
type UpdateVehicleTypeRequest struct {
	// Name of the vehicle type
	Name *string `json:"name" example:"sedan"`
	// Description of the vehicle type
	Description *string `json:"description" example:"A standard passenger car"`
}

// VehicleTypeResponse represents the response for vehicle type data
// @Description Vehicle type response
type VehicleTypeResponse struct {
	// ID of the vehicle type
	ID uint64 `json:"id"`
	// Name of the vehicle type
	Name string `json:"name"`
	// Description of the vehicle type
	Description string `json:"description"`
}

// Brand
// CreateVehicleBrandRequest represents the request for creating a new vehicle brand
// @Description Vehicle brand creation request
type CreateVehicleBrandRequest struct {
	// Name of the vehicle brand
	Name string `json:"name" validate:"required" example:"Toyota"`
	// Description of the vehicle brand
	Description string `json:"description" example:"A popular Japanese car brand"`
}

// UpdateVehicleBrandRequest represents the request for updating vehicle brand
// @Description Vehicle brand update request
type UpdateVehicleBrandRequest struct {
	// ID of the vehicle type
	VehicleTypeID *uint64 `json:"vehicle_type_id" example:"1"`
	// Name of the vehicle brand
	Name *string `json:"name" example:"Toyota"`
	// Description of the vehicle brand
	Description *string `json:"description" example:"A popular Japanese car brand"`
}

// VehicleBrandResponse represents the response for vehicle brand data
// @Description Vehicle brand response
type VehicleBrandResponse struct {
	// ID of the vehicle brand
	ID uint64 `json:"id"`
	// Name of the vehicle brand
	Name string `json:"name"`
	// Description of the vehicle brand
	Description string `json:"description"`
	// ID of the vehicle type
	VehicleTypeID uint64 `json:"vehicle_type_id"`
	// Vehicle type
	Type VehicleTypeResponse `json:"type"`
}

// Model
// CreateVehicleModelRequest represents the request for creating a new vehicle model
// @Description Vehicle model creation request
type CreateVehicleModelRequest struct {
	// Name of the vehicle model
	Name string `json:"name" validate:"required" example:"Camry"`
	// Description of the vehicle model
	Description string `json:"description" example:"A mid-size sedan"`
	// Start year of the vehicle model
	StartYear int `json:"start_year" validate:"omitempty,year" example:"2020"`
	// End year of the vehicle model
	EndYear int `json:"end_year" validate:"omitempty,year" example:"2022"`
}

// UpdateVehicleModelRequest represents the request for updating vehicle model
// @Description Vehicle model update request
type UpdateVehicleModelRequest struct {
	// ID of the vehicle brand
	BrandID *uint64 `json:"brand_id" example:"1"`
	// Name of the vehicle model
	Name *string `json:"name" example:"Camry"`
	// Description of the vehicle model
	Description *string `json:"description" example:"A mid-size sedan"`
	// Start year of the vehicle model
	StartYear *int `json:"start_year" validate:"omitempty,year" example:"2020"`
	// End year of the vehicle model
	EndYear *int `json:"end_year" validate:"omitempty,year" example:"2022"`
}

// VehicleModelResponse represents the response for vehicle model data
// @Description Vehicle model response
type VehicleModelResponse struct {
	// ID of the vehicle model
	ID uint64 `json:"id"`
	// Name of the vehicle model
	Name string `json:"name"`
	// Description of the vehicle model
	Description string `json:"description"`
	// ID of the vehicle brand
	BrandID uint64 `json:"brand_id"`
	// Start year of the vehicle model
	StartYear int `json:"start_year"`
	// End year of the vehicle model
	EndYear int `json:"end_year"`
	// Vehicle brand
	Brand VehicleBrandResponse `json:"brand"`
}

// Generation
// CreateVehicleGenerationRequest represents the request for creating a new vehicle generation
// @Description Vehicle generation creation request
type CreateVehicleGenerationRequest struct {
	// Name of the vehicle generation
	Name string `json:"name" validate:"required" example:"Generation Name"`
	// Description of the vehicle generation
	Description string `json:"description" example:"A brief description of the generation"`
	// Start year of the vehicle generation
	StartYear int `json:"start_year" validate:"omitempty,year" example:"2020"`
	// End year of the vehicle generation
	EndYear int `json:"end_year" validate:"omitempty,year" example:"2022"`
	// Engine type of the vehicle generation
	EngineType string `json:"engine_type" example:"V6"`
	// Assembly type of the vehicle generation
	AssemblyType string `json:"assembly_type" example:"CKD"`
	// Assembler of the vehicle generation
	Assembler string `json:"assembler" example:"Toyota"`
	// Transmission of the vehicle generation
	Transmission string `json:"transmission" example:"Automatic"`
	// Engine size of the vehicle generation
	EngineSize int `json:"engine_size" example:"3000"`
	// Body style of the vehicle generation
	BodyStyle string `json:"body_style" example:"Sedan"`
	// Special features of the vehicle generation
	SpecialFeatures string `json:"special_features" example:"Leather seats, Sunroof"`
}

// UpdateVehicleGenerationRequest represents the request for updating vehicle generation
// @Description Vehicle generation update request
type UpdateVehicleGenerationRequest struct {
	// Name of the vehicle generation
	Name *string `json:"name" example:"Generation Name"`
	// Description of the vehicle generation
	Description *string `json:"description" example:"A brief description of the generation"`
	// ID of the vehicle model
	ModelID *uint64 `json:"model_id" example:"1"`
	// Start year of the vehicle generation
	StartYear *int `json:"start_year" validate:"omitempty,year" example:"2020"`
	// End year of the vehicle generation
	EndYear *int `json:"end_year" validate:"omitempty,year" example:"2022"`
	// Engine type of the vehicle generation
	EngineType *string `json:"engine_type" example:"V6"`
	// Assembly type of the vehicle generation
	AssemblyType *string `json:"assembly_type" example:"CKD"`
	// Assembler of the vehicle generation
	Assembler *string `json:"assembler" example:"Toyota"`
	// Transmission of the vehicle generation
	Transmission *string `json:"transmission" example:"Automatic"`
	// Engine size of the vehicle generation
	EngineSize *int `json:"engine_size" example:"3000"`
	// Body style of the vehicle generation
	BodyStyle *string `json:"body_style" example:"Sedan"`
	// Special features of the vehicle generation
	SpecialFeatures *string `json:"special_features" example:"Leather seats, Sunroof"`
}

// VehicleGenerationResponse represents the response for vehicle generation data
// @Description Vehicle generation response
type VehicleGenerationResponse struct {
	// ID of the vehicle generation
	ID uint64 `json:"id"`
	// Name of the vehicle generation
	Name string `json:"name"`
	// Description of the vehicle generation
	Description string `json:"description"`
	// ID of the vehicle model
	ModelID uint64 `json:"model_id"`
	// Start year of the vehicle generation
	StartYear int `json:"start_year"`
	// End year of the vehicle generation
	EndYear int `json:"end_year"`
	// Engine type of the vehicle generation
	EngineType string `json:"engine_type"`
	// Assembly type of the vehicle generation
	AssemblyType string `json:"assembly_type"`
	// Assembler of the vehicle generation
	Assembler string `json:"assembler"`
	// Transmission of the vehicle generation
	Transmission string `json:"transmission"`
	// Engine size of the vehicle generation
	EngineSize int `json:"engine_size"`
	// Body style of the vehicle generation
	BodyStyle string `json:"body_style"`
	// Special features of the vehicle generation
	SpecialFeatures string `json:"special_features"`
	// Vehicle model
	ModelInfo VehicleModelResponse `json:"model_info"`
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
	LicensePlate string `json:"license_plate" validate:"omitempty,iranian_license_plate" example:"۱۲الف۳۴۵۶۸"`
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
	// Vehicle generation
	Generation VehicleGenerationResponse `json:"generation"`
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
