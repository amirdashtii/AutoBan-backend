package entity

import (
	"time"

	"github.com/google/uuid"
)

// ServiceVisit represents a visit to a service center for a vehicle
type ServiceVisit struct {
	BaseEntity

	UserID         uuid.UUID  `json:"user_id" gorm:"not null"`
	UserVehicleID  uint64       `gorm:"not null" json:"user_vehicle_id"`
	ServiceMileage uint       `gorm:"not null" json:"service_mileage"`
	ServiceDate    time.Time  `gorm:"not null" json:"service_date"`
	ServiceCenter  string     `json:"service_center"`
	Notes          string     `json:"notes"`
	OilChangeID    *uint64      `json:"oil_change_id"`
	OilChange      *OilChange `gorm:"foreignKey:OilChangeID" json:"oil_change,omitempty"`
	OilFilterID    *uint64      `json:"oil_filter_id"`
	OilFilter      *OilFilter `gorm:"foreignKey:OilFilterID" json:"oil_filter,omitempty"`
	// other services
}
