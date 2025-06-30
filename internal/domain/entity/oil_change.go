package entity

import (
	"time"

	"github.com/google/uuid"
)

type OilChange struct {
	BaseModel

	UserID            uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	UserVehicleID     uint64    `json:"user_vehicle_id" gorm:"not null"`
	OilName           string    `json:"oil_name" gorm:"not null"`
	OilBrand          string    `json:"oil_brand"`
	OilType           string    `json:"oil_type"`      // مینرال، سنتتیک، نیمه سنتتیک
	OilViscosity      string    `json:"oil_viscosity"` // 5W-30, 10W-40, etc.
	ChangeMileage     uint      `json:"change_mileage" gorm:"not null"`
	ChangeDate        time.Time `json:"change_date" gorm:"not null"`
	OilCapacity       float64   `json:"oil_capacity"` // لیتر
	NextChangeMileage uint      `json:"next_change_mileage"`
	NextChangeDate    time.Time `json:"next_change_date"`
	ServiceCenter     string    `json:"service_center"`
	Notes             string    `json:"notes"`
}
