package entity

import (
	"time"

	"github.com/google/uuid"
)

type OilChange struct {
	BaseModel

	UserID            uuid.UUID `gorm:"type:uuid;not null"`
	UserVehicleID     uint64    `gorm:"not null"`
	ServiceVisitID    uuid.UUID `gorm:"type:uuid;not null"`
	OilName           string    `gorm:"not null"`
	OilBrand          string
	OilType           string
	OilViscosity      string
	ChangeMileage     uint      `gorm:"not null"`
	ChangeDate        time.Time `gorm:"not null"`
	OilCapacity       float64
	NextChangeMileage uint
	NextChangeDate    time.Time
	ServiceCenter     string
	Notes             string
}
