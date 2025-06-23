package dto

import (
	"time"
)

// VehicleTypeRequest represents the request for vehicle type data
type CreateVehicleTypeRequest struct {
	//Vehicle type name
	// @Example car
	Name string `json:"name" binding:"required" example:"car"`
	//Vehicle type description
	// @Example cars
	Description string `json:"description" binding:"required" example:"cars"`
}

// UpdateVehicleTypeRequest represents the request for updating vehicle type
// @Description Vehicle type update request
type UpdateVehicleTypeRequest struct {
	//Vehicle type name (optional)
	// @Example Car
	Name *string `json:"name" example:"car"`
	//Vehicle type description (optional)
	// @Example cars
	Description *string `json:"description" example:"cars"`
}

// VehicleTypeResponse represents the response for vehicle type data
type VehicleTypeResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// VehicleBrandResponse represents the response for vehicle brand data
type VehicleBrandResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VehicleType string `json:"vehicle_type"`
}

// VehicleBrandRequest represents the request for vehicle brand data
type VehicleBrandRequest struct {
	VehicleTypeID string `json:"vehicle_type_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"required"`
	StartYear     int    `json:"start_year"`
	EndYear       int    `json:"end_year"`
}

// VehicleModelRequest represents the request for vehicle model data
type VehicleModelRequest struct {
	BrandID     string `json:"brand_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
}

// VehicleModelResponse represents the response for vehicle model data
type VehicleModelResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BrandID     string `json:"brand_id"`
	StartYear   int    `json:"start_year"`
	EndYear     int    `json:"end_year"`
}

type VehicleGenerationRequest struct {
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

// VehicleGenerationResponse represents the response for vehicle generation data
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

// UserVehicleRequest represents the request for adding a new vehicle
type UserVehicleRequest struct {
	GenerationID   string    `json:"generation_id" binding:"required"`
	ProductionYear int       `json:"production_year" binding:"required"`
	Color          string    `json:"color"`
	LicensePlate   string    `json:"license_plate"`
	VIN            string    `json:"vin"`
	CurrentMileage int       `json:"current_mileage"`
	PurchaseDate   time.Time `json:"purchase_date"`
}

// UserVehicleResponse represents the response for user vehicle data
type UserVehicleResponse struct {
	ID     uint   `json:"id"`
	UserID string `json:"user_id"`
	// VehicleType    string    `json:"vehicle_type"`
	// Brand          string    `json:"brand"`
	// Model          string    `json:"model"`
	// Generation     string    `json:"generation"`
	GenerationID   string    `json:"generation_id"`
	ProductionYear int       `json:"production_year"`
	Color          string    `json:"color"`
	LicensePlate   string    `json:"license_plate"`
	VIN            string    `json:"vin"`
	CurrentMileage int       `json:"current_mileage"`
	PurchaseDate   time.Time `json:"purchase_date"`
	// Technical specifications from generation
	Transmission    string `json:"transmission"`
	EngineType      string `json:"engine_type"`
	EngineSize      int    `json:"engine_size"`
	BodyStyle       string `json:"body_style"`
	SpecialFeatures string `json:"special_features"`
	AssemblyType    string `json:"assembly_type"`
	Assembler       string `json:"assembler"`
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
