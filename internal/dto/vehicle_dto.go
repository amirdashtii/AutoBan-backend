package dto

import (
	"time"
)

// Type
// CreateVehicleTypeRequest represents the request for creating a new vehicle type
// @Description Vehicle type creation request
type CreateVehicleTypeRequest struct {
	Name        string `json:"name" validate:"required" example:"car"`
	Description string `json:"description" example:"A standard passenger car"`
}

// UpdateVehicleTypeRequest represents the request for updating vehicle type
// @Description Vehicle type update request
type UpdateVehicleTypeRequest struct {
	Name        *string `json:"name" example:"sedan"`
	Description *string `json:"description" example:"A standard passenger car"`
}

// VehicleTypeResponse represents the response for vehicle type data
// @Description Vehicle type response
type VehicleTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Brand
// CreateVehicleBrandRequest represents the request for creating a new vehicle brand
// @Description Vehicle brand creation request
type CreateVehicleBrandRequest struct {
	VehicleTypeID string `json:"vehicle_type_id" validate:"required" example:"1"`
	Name          string `json:"name" validate:"required" example:"Toyota"`
	Description   string `json:"description" example:"A popular Japanese car brand"`
}

// UpdateVehicleBrandRequest represents the request for updating vehicle brand
// @Description Vehicle brand update request
type UpdateVehicleBrandRequest struct {
	VehicleTypeID *string `json:"vehicle_type_id" example:"1"`
	Name          *string `json:"name" example:"Toyota"`
	Description   *string `json:"description" example:"A popular Japanese car brand"`
}

// VehicleBrandResponse represents the response for vehicle brand data
// @Description Vehicle brand response
type VehicleBrandResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VehicleType string `json:"vehicle_type"`
}

// Model
// CreateVehicleModelRequest represents the request for creating a new vehicle model
// @Description Vehicle model creation request
type CreateVehicleModelRequest struct {
	BrandID     string `json:"brand_id" validate:"required" example:"1"`
	Name        string `json:"name" validate:"required" example:"Camry"`
	Description string `json:"description" example:"A mid-size sedan"`
	StartYear   int    `json:"start_year" example:"2020"`
	EndYear     int    `json:"end_year" example:"2022"`
}

// UpdateVehicleModelRequest represents the request for updating vehicle model
// @Description Vehicle model update request
type UpdateVehicleModelRequest struct {
	BrandID     *string `json:"brand_id" example:"1"`
	Name        *string `json:"name" example:"Camry"`
	Description *string `json:"description" example:"A mid-size sedan"`
	StartYear   *int    `json:"start_year" example:"2020"`
	EndYear     *int    `json:"end_year" example:"2022"`
}

// VehicleModelResponse represents the response for vehicle model data
// @Description Vehicle model response
type VehicleModelResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BrandID     string `json:"brand_id"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
}

// Generation
// CreateVehicleGenerationRequest represents the request for creating a new vehicle generation
// @Description Vehicle generation creation request
type CreateVehicleGenerationRequest struct {
	Name            string `json:"name" validate:"required" example:"Generation Name"`
	Description     string `json:"description" example:"A brief description of the generation"`
	ModelID         string `json:"model_id" validate:"required" example:"1"`
	StartYear       int    `json:"start_year" example:"2020"`
	EndYear         int    `json:"end_year" example:"2022"`
	EngineType      string `json:"engine_type" example:"V6"`
	AssemblyType    string `json:"assembly_type" example:"CKD"`
	Assembler       string `json:"assembler" example:"Toyota"`
	Transmission    string `json:"transmission" example:"Automatic"`
	EngineSize      int    `json:"engine_size" example:"3000"`
	BodyStyle       string `json:"body_style" example:"Sedan"`
	SpecialFeatures string `json:"special_features" example:"Leather seats, Sunroof"`
}

// UpdateVehicleGenerationRequest represents the request for updating vehicle generation
// @Description Vehicle generation update request
type UpdateVehicleGenerationRequest struct {
	Name            *string `json:"name" example:"Generation Name"`
	Description     *string `json:"description" example:"A brief description of the generation"`
	ModelID         *string `json:"model_id" example:"1"`
	StartYear       *int    `json:"start_year" example:"2020"`
	EndYear         *int    `json:"end_year" example:"2022"`
	EngineType      *string `json:"engine_type" example:"V6"`
	AssemblyType    *string `json:"assembly_type" example:"CKD"`
	Assembler       *string `json:"assembler" example:"Toyota"`
	Transmission    *string `json:"transmission" example:"Automatic"`
	EngineSize      *int    `json:"engine_size" example:"3000"`
	BodyStyle       *string `json:"body_style" example:"Sedan"`
	SpecialFeatures *string `json:"special_features" example:"Leather seats, Sunroof"`
}

// VehicleGenerationResponse represents the response for vehicle generation data
// @Description Vehicle generation response
type VehicleGenerationResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	ModelID         string `json:"model_id"`
	StartYear       int    `json:"start_year"`
	EndYear         int    `json:"end_year"`
	EngineType      string `json:"engine_type"`
	AssemblyType    string `json:"assembly_type"`
	Assembler       string `json:"assembler"`
	Transmission    string `json:"transmission"`
	EngineSize      int    `json:"engine_size"`
	BodyStyle       string `json:"body_style"`
	SpecialFeatures string `json:"special_features"`
}

// UserVehicle
// CreateUserVehicleRequest represents the request for adding a new user vehicle
// @Description User vehicle creation request
type CreateUserVehicleRequest struct {
	Name           string    `json:"name" validate:"required" example:"My Car"`
	GenerationID   string    `json:"generation_id" validate:"required" example:"1"`
	ProductionYear int       `json:"production_year" example:"2020"`
	Color          string    `json:"color" example:"Red"`
	LicensePlate   string    `json:"license_plate" example:"ABC123"`
	VIN            string    `json:"vin" example:"1HGCM82633A123456"`
	CurrentMileage int       `json:"current_mileage" binding:"required" example:"15000"`
	PurchaseDate   time.Time `json:"purchase_date" binding:"required" example:"2020-01-01"`
}

// UpdateUserVehicleRequest represents the request for updating user vehicle
// @Description User vehicle update request
type UpdateUserVehicleRequest struct {
	Name           *string    `json:"name" example:"My Car"`
	GenerationID   *string    `json:"generation_id" example:"1"`
	ProductionYear *int       `json:"production_year" example:"2020"`
	Color          *string    `json:"color" example:"Red"`
	LicensePlate   *string    `json:"license_plate" example:"ABC123"`
	VIN            *string    `json:"vin" example:"1HGCM82633A123456"`
	CurrentMileage *int       `json:"current_mileage" example:"15000"`
	PurchaseDate   *time.Time `json:"purchase_date" example:"2020-01-01"`
}

// UserVehicleResponse represents the response for user vehicle data
// @Description User vehicle response
type UserVehicleResponse struct {
	ID             uint      `json:"id"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	GenerationID   string    `json:"generation_id"`
	ProductionYear int       `json:"production_year"`
	Color          string    `json:"color"`
	LicensePlate   string    `json:"license_plate"`
	VIN            string    `json:"vin"`
	CurrentMileage int       `json:"current_mileage"`
	PurchaseDate   time.Time `json:"purchase_date"`
}

// ListVehicleTypesResponse represents the response for listing vehicle types
type ListVehicleTypesResponse struct {
	Types []VehicleTypeResponse `json:"types"`
}

// ListVehicleBrandsResponse represents the response for listing vehicle brands
type ListVehicleBrandsResponse struct {
	Brands []VehicleBrandResponse `json:"brands"`
}

// ListVehicleModelsResponse represents the response for listing vehicle models
type ListVehicleModelsResponse struct {
	Models []VehicleModelResponse `json:"models"`
}

// ListVehicleGenerationsResponse represents the response for listing vehicle generations
type ListVehicleGenerationsResponse struct {
	Generations []VehicleGenerationResponse `json:"generations"`
}

// ListUserVehiclesResponse represents the response for listing user vehicles
type ListUserVehiclesResponse struct {
	Vehicles []UserVehicleResponse `json:"vehicles"`
}
