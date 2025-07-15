package entity

import (
	"time"

	"github.com/google/uuid"
)

// ServiceVisit represents a visit to a service center for a vehicle
type ServiceVisit struct {
	BaseEntity

	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	UserVehicleID  uint64    `gorm:"not null"`
	ServiceMileage uint      `gorm:"not null"`
	ServiceDate    time.Time `gorm:"not null"`
	ServiceCenter  string
	Notes          string
	OilChange      OilChange `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	OilFilter      OilFilter `gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	// other services
}
