package entity

import (
	"time"

	"gorm.io/gorm"
)

// OilFilter represents the oil filter change entity
type OilFilter struct {
	gorm.Model
	
	UserVehicleID     uint      `gorm:"not null" json:"user_vehicle_id"`
	FilterName        string    `gorm:"not null" json:"filter_name"`
	FilterBrand       string    `gorm:"not null" json:"filter_brand"`
	FilterType        string    `gorm:"not null" json:"filter_type"`
	FilterPartNumber  string    `json:"filter_part_number"`
	ChangeMileage     int       `gorm:"not null" json:"change_mileage"`
	ChangeDate        time.Time `gorm:"not null" json:"change_date"`
	NextChangeMileage int       `json:"next_change_mileage"`
	NextChangeDate    time.Time `json:"next_change_date"`
	ServiceCenter     string    `json:"service_center"`
	Notes             string    `json:"notes"`
}
