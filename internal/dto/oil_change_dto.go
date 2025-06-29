package dto

import "time"

// CreateOilChangeRequest - Request to create an oil change
// @Description Request to create an oil change
type CreateOilChangeRequest struct {
	// User vehicle ID
	UserVehicleID uint64 `json:"user_vehicle_id" validate:"required" example:"1"`
	// Oil name
	OilName string `json:"oil_name" validate:"required" example:"تکتاز"`
	// Oil brand
	OilBrand string `json:"oil_brand" example:"بهران"`
	// Oil type
	OilType string `json:"oil_type" example:"مینرال، سنتتیک، نیمه سنتتیک"`
	// Oil viscosity
	OilViscosity string `json:"oil_viscosity" example:"5W-30, 10W-40, etc."`
	// Change mileage
	ChangeMileage uint `json:"change_mileage" validate:"required" example:"10000"`
	// Change date
	ChangeDate string `json:"change_date" validate:"required,date" example:"2021-01-01"`
	// Oil capacity
	OilCapacity float64 `json:"oil_capacity" example:"5"`
	// Next change mileage
	NextChangeMileage uint `json:"next_change_mileage" example:"20000"`
	// Next change date
	NextChangeDate string `json:"next_change_date" validate:"omitempty,date" example:"2021-01-01"`
	// Service center
	ServiceCenter string `json:"service_center" example:"اتوبان سرویس"`
	// Notes
	Notes string `json:"notes" example:"تعویض روغن"`
}

// UpdateOilChangeRequest - Request to update an oil change
// @Description Request to update an oil change
type UpdateOilChangeRequest struct {
	// Oil name
	OilName *string `json:"oil_name" example:"تکتاز"`
	// Oil brand
	OilBrand *string `json:"oil_brand" example:"بهران"`
	// Oil type
	OilType *string `json:"oil_type" example:"مینرال، سنتتیک، نیمه سنتتیک"`
	// Oil viscosity
	OilViscosity *string `json:"oil_viscosity" example:"5W-30, 10W-40, etc."`
	// Change mileage
	ChangeMileage *uint `json:"change_mileage" example:"10000"`
	// Change date
	ChangeDate *string `json:"change_date" validate:"omitempty,date" example:"2021-01-01"`
	// Oil capacity
	OilCapacity *float64 `json:"oil_capacity" example:"5"`
	// Next change mileage
	NextChangeMileage *uint `json:"next_change_mileage" example:"20000"`
	// Next change date
	NextChangeDate *string `json:"next_change_date" validate:"omitempty,date" example:"2021-01-01"`
	// Service center
	ServiceCenter *string `json:"service_center" example:"اتوبان سرویس"`
	// Notes
	Notes *string `json:"notes" example:"تعویض روغن"`
}

// OilChangeResponse - Oil change response
// @Description Oil change response
type OilChangeResponse struct {
	// ID
	ID uint64 `json:"id" example:"1"`
	// User vehicle ID
	UserVehicleID uint64 `json:"user_vehicle_id" example:"1"`
	// Oil name
	OilName string `json:"oil_name" example:"تکتاز"`
	// Oil brand
	OilBrand string `json:"oil_brand" example:"بهران"`
	// Oil type
	OilType string `json:"oil_type" example:"مینرال، سنتتیک، نیمه سنتتیک"`
	// Oil viscosity
	OilViscosity string `json:"oil_viscosity" example:"5W-30, 10W-40, etc."`
	// Change mileage
	ChangeMileage uint `json:"change_mileage" example:"10000"`
	// Change date
	ChangeDate time.Time `json:"change_date" example:"2021-01-01"`
	// Oil capacity
	OilCapacity float64 `json:"oil_capacity" example:"5"`
	// Next change mileage
	NextChangeMileage uint `json:"next_change_mileage" example:"20000"`
	// Next change date
	NextChangeDate time.Time `json:"next_change_date" example:"2021-01-01"`
	// Service center
	ServiceCenter string `json:"service_center" example:"اتوبان سرویس"`
	// Notes
	Notes string `json:"notes" example:"تعویض روغن"`
}

// ListOilChangesResponse - Oil change list response
// @Description Oil change list response
type ListOilChangesResponse struct {
	// Oil changes
	OilChanges []OilChangeResponse `json:"oil_changes"`
}
