package dto

import (
	"github.com/google/uuid"
)

// ServiceVisitOilChange - add oil change to create service visit request
// @Description add oil change to create service visit
type ServiceVisitOilChange struct {
	// Oil name
	OilName string `json:"oil_name" validate:"required" example:"تکتاز"`
	// Oil brand
	OilBrand string `json:"oil_brand" example:"بهران"`
	// Oil type
	OilType string `json:"oil_type" example:"مینرال، سنتتیک، نیمه سنتتیک"`
	// Oil viscosity
	OilViscosity string `json:"oil_viscosity" example:"5W-30, 10W-40, etc."`
	// Oil capacity
	OilCapacity float64 `json:"oil_capacity" example:"5"`
	// Next change mileage
	NextChangeMileage uint `json:"next_change_mileage" validate:"omitempty,min=0" example:"20000"`
	// Next change date
	NextChangeDate string `json:"next_change_date" validate:"omitempty,date" example:"2021-01-01"`
	// Notes
	Notes string `json:"notes" example:"تعویض روغن"`
}

// ServiceVisitOilFilter - add oil filter to create service visit request
// @Description Requst to add oil filter to create service visit
type ServiceVisitOilFilter struct {
	// Name of the oil filter
	FilterName string `json:"filter_name" validate:"required" example:"Oil Filter"`
	// Brand of the oil filter
	FilterBrand string `json:"filter_brand" example:"Mann"`
	// Type of the oil filter
	FilterType string `json:"filter_type" example:"Cartridge"`
	// Part number of the oil filter
	FilterPartNumber string `json:"filter_part_number" example:"HU816x"`
	// Next change mileage
	NextChangeMileage uint `json:"next_change_mileage" validate:"omitempty,min=0" example:"30000"`
	// Next change date
	NextChangeDate string `json:"next_change_date" validate:"omitempty,date" example:"2024-07-15"`
	// Additional notes
	Notes string `json:"notes" example:"Changed with oil change"`
}

// CreateServiceVisitRequest represents the request for creating a new service visit
// @Description Service visit creation request
type CreateServiceVisitRequest struct {
	// ID of the user vehicle
	UserVehicleID uint64 `json:"user_vehicle_id" validate:"required" example:"1"`
	// Mileage when service was performed
	ServiceMileage uint `json:"service_mileage" validate:"required,min=0" example:"15000"`
	// Date when service was performed
	ServiceDate string `json:"service_date" validate:"required,date" example:"2024-01-15"`
	// Service center where service was performed
	ServiceCenter string `json:"service_center" example:"Auto Service Center"`
	// Additional notes
	Notes string `json:"notes" example:"Regular maintenance service"`
	// Oil change information (optional)
	OilChange *ServiceVisitOilChange `json:"oil_change,omitempty"`
	// Oil filter information (optional)
	OilFilter *ServiceVisitOilFilter `json:"oil_filter,omitempty"`
	// Other services can be added here in the future
}

// UpdateServiceVisitOilChange - add oil change to update service visit request
// @Description add oil change to update service visit request
type UpdateServiceVisitOilChange struct {
	// Oil name
	OilName *string `json:"oil_name" example:"تکتاز"`
	// Oil brand
	OilBrand *string `json:"oil_brand" example:"بهران"`
	// Oil type
	OilType *string `json:"oil_type" example:"مینرال، سنتتیک، نیمه سنتتیک"`
	// Oil viscosity
	OilViscosity *string `json:"oil_viscosity" example:"5W-30, 10W-40, etc."`
	// Oil capacity
	OilCapacity *float64 `json:"oil_capacity" example:"5"`
	// Next change mileage
	NextChangeMileage *uint `json:"next_change_mileage" validate:"omitempty,min=0" example:"20000"`
	// Next change date
	NextChangeDate *string `json:"next_change_date" validate:"omitempty,date" example:"2021-01-01"`
	// Notes
	Notes *string `json:"notes" example:"تعویض روغن"`
}

// UpdateServiceVisitOilFilter - add oil filter to update service visit request
// @Description add oil filter to update service visit request
type UpdateServiceVisitOilFilter struct {
	// Name of the oil filter
	FilterName *string `json:"filter_name" example:"Oil Filter"`
	// Brand of the oil filter
	FilterBrand *string `json:"filter_brand" example:"Mann"`
	// Type of the oil filter
	FilterType *string `json:"filter_type" example:"Cartridge"`
	// Part number of the oil filter
	FilterPartNumber *string `json:"filter_part_number" example:"HU816x"`
	// Next change mileage
	NextChangeMileage *uint `json:"next_change_mileage" validate:"omitempty,min=0" example:"30000"`
	// Next change date
	NextChangeDate *string `json:"next_change_date" validate:"omitempty,date" example:"2024-07-15"`
	// Additional notes
	Notes *string `json:"notes" example:"Changed with oil change"`
}

// UpdateServiceVisitRequest represents the request for updating service visit
// @Description Service visit update request
type UpdateServiceVisitRequest struct {
	// Mileage when service was performed
	ServiceMileage *uint `json:"service_mileage" validate:"omitempty,min=0" example:"15000"`
	// Date when service was performed
	ServiceDate *string `json:"service_date" validate:"omitempty,date" example:"2024-01-15"`
	// Service center where service was performed
	ServiceCenter *string `json:"service_center" example:"Auto Service Center"`
	// Additional notes
	Notes *string `json:"notes" example:"Regular maintenance service"`
	// Oil change information (optional)
	OilChange *UpdateServiceVisitOilChange `json:"oil_change,omitempty"`
	// Oil filter information (optional)
	OilFilter *UpdateServiceVisitOilFilter `json:"oil_filter,omitempty"`
	// Other services can be added here in the future
}

// ServiceVisitOilChangeResponse - Oil change response
// @Description Oil change response
type ServiceVisitOilChangeResponse struct {
	// ID
	ID uint64 `json:"id" example:"1"`
	// Oil name
	OilName string `json:"oil_name" example:"تکتاز"`
	// Oil brand
	OilBrand string `json:"oil_brand" example:"بهران"`
	// Oil type
	OilType string `json:"oil_type" example:"مینرال، سنتتیک، نیمه سنتتیک"`
	// Oil viscosity
	OilViscosity string `json:"oil_viscosity" example:"5W-30, 10W-40, etc."`
	// Oil capacity
	OilCapacity float64 `json:"oil_capacity" example:"5"`
	// Next change mileage
	NextChangeMileage uint `json:"next_change_mileage" example:"20000"`
	// Next change date
	NextChangeDate string `json:"next_change_date" example:"2021-01-01"`
	// Notes
	Notes string `json:"notes" example:"تعویض روغن"`
}

// ServiceVisitOilFilterResponse - Oil filter response
// @Description Oil filter response
type ServiceVisitOilFilterResponse struct {
	// ID
	ID uint64 `json:"id" example:"1"`
	// Name of the oil filter
	FilterName string `json:"filter_name"`
	// Brand of the oil filter
	FilterBrand string `json:"filter_brand"`
	// Type of the oil filter
	FilterType string `json:"filter_type"`
	// Part number of the oil filter
	FilterPartNumber string `json:"filter_part_number"`
	// Next change mileage
	NextChangeMileage uint `json:"next_change_mileage"`
	// Next change date
	NextChangeDate string `json:"next_change_date"`
	// Additional notes
	Notes string `json:"notes"`
}

// ServiceVisitResponse represents the response for service visit data
// @Description Service visit response
type ServiceVisitResponse struct {
	// ID of the service visit
	ID uuid.UUID `json:"id"`
	// ID of the user vehicle
	UserVehicleID uint64 `json:"user_vehicle_id"`
	// Mileage when service was performed
	ServiceMileage uint `json:"service_mileage"`
	// Date when service was performed
	ServiceDate string `json:"service_date"`
	// Service center where service was performed
	ServiceCenter string `json:"service_center"`
	// Additional notes
	Notes string `json:"notes"`
	// Oil change information (if performed)
	OilChange *ServiceVisitOilChangeResponse `json:"oil_change,omitempty"`
	// Oil filter information (if performed)
	OilFilter *ServiceVisitOilFilterResponse `json:"oil_filter,omitempty"`
	// Other services can be added here in the future
}

// ListServiceVisitsResponse represents the response for listing service visits
// @Description List of service visits
type ListServiceVisitsResponse struct {
	// List of service visits
	ServiceVisits []ServiceVisitResponse `json:"service_visits"`
}
