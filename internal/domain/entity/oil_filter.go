package entity

import (
	"time"

	"github.com/google/uuid"
)

// OilFilter represents the oil filter change entity
type OilFilter struct {
	BaseModel

	UserID            uuid.UUID `gorm:"type:uuid;not null"`
	UserVehicleID     uint64    `gorm:"not null"`
	ServiceVisitID    uuid.UUID `gorm:"type:uuid;not null"`
	FilterName        string    `gorm:"not null"`
	FilterBrand       string
	FilterType        string
	FilterPartNumber  string
	ChangeMileage     uint      `gorm:"not null"`
	ChangeDate        time.Time `gorm:"not null"`
	NextChangeMileage uint
	NextChangeDate    time.Time
	ServiceCenter     string
	Notes             string
}
