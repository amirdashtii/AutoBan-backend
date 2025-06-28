package dto

import (
	"time"
)

// CreateOilFilterRequest represents the request for creating a new oil filter change
// @Description Oil filter change creation request
type CreateOilFilterRequest struct {
	// ID of the user vehicle
	UserVehicleID uint `json:"user_vehicle_id" validate:"required" example:"1"`
	// Name of the oil filter
	FilterName string `json:"filter_name" validate:"required" example:"Oil Filter"`
	// Brand of the oil filter
	FilterBrand string `json:"filter_brand" example:"Mann"`
	// Type of the oil filter
	FilterType string `json:"filter_type" example:"Cartridge"`
	// Part number of the oil filter
	FilterPartNumber string `json:"filter_part_number" example:"HU816x"`
	// Mileage when filter was changed
	ChangeMileage int `json:"change_mileage" validate:"required,min=0" example:"15000"`
	// Date when filter was changed
	ChangeDate string `json:"change_date" validate:"required,date" example:"2024-01-15"`
	// Next change mileage
	NextChangeMileage int `json:"next_change_mileage" validate:"omitempty,min=0" example:"30000"`
	// Next change date
	NextChangeDate string `json:"next_change_date" validate:"omitempty,date" example:"2024-07-15"`
	// Service center where filter was changed
	ServiceCenter string `json:"service_center" example:"Auto Service Center"`
	// Additional notes
	Notes string `json:"notes" example:"Changed with oil change"`
}

// UpdateOilFilterRequest represents the request for updating oil filter change
// @Description Oil filter change update request
type UpdateOilFilterRequest struct {
	// Name of the oil filter
	FilterName *string `json:"filter_name" example:"Oil Filter"`
	// Brand of the oil filter
	FilterBrand *string `json:"filter_brand" example:"Mann"`
	// Type of the oil filter
	FilterType *string `json:"filter_type" example:"Cartridge"`
	// Part number of the oil filter
	FilterPartNumber *string `json:"filter_part_number" example:"HU816x"`
	// Mileage when filter was changed
	ChangeMileage *int `json:"change_mileage" validate:"omitempty,min=0" example:"15000"`
	// Date when filter was changed
	ChangeDate *string `json:"change_date" validate:"omitempty,date" example:"2024-01-15"`
	// Next change mileage
	NextChangeMileage *int `json:"next_change_mileage" validate:"omitempty,min=0" example:"30000"`
	// Next change date
	NextChangeDate *string `json:"next_change_date" validate:"omitempty,date" example:"2024-07-15"`
	// Service center where filter was changed
	ServiceCenter *string `json:"service_center" example:"Auto Service Center"`
	// Additional notes
	Notes *string `json:"notes" example:"Changed with oil change"`
}

// OilFilterResponse represents the response for oil filter change data
// @Description Oil filter change response
type OilFilterResponse struct {
	// ID of the oil filter change
	ID uint `json:"id"`
	// ID of the user vehicle
	UserVehicleID uint `json:"user_vehicle_id"`
	// Name of the oil filter
	FilterName string `json:"filter_name"`
	// Brand of the oil filter
	FilterBrand string `json:"filter_brand"`
	// Type of the oil filter
	FilterType string `json:"filter_type"`
	// Part number of the oil filter
	FilterPartNumber string `json:"filter_part_number"`
	// Mileage when filter was changed
	ChangeMileage int `json:"change_mileage"`
	// Date when filter was changed
	ChangeDate time.Time `json:"change_date"`
	// Next change mileage
	NextChangeMileage int `json:"next_change_mileage"`
	// Next change date
	NextChangeDate time.Time `json:"next_change_date"`
	// Service center where filter was changed
	ServiceCenter string `json:"service_center"`
	// Additional notes
	Notes string `json:"notes"`
}

// ListOilFiltersResponse represents the response for listing oil filter changes
// @Description List of oil filter changes
type ListOilFiltersResponse struct {
	// List of oil filter changes
	OilFilters []OilFilterResponse `json:"oil_filters"`
}
