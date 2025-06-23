package entity

import (
	"time"

	"gorm.io/gorm"
)

// VehicleType represents the main category of vehicles
type VehicleType struct {
	gorm.Model

	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}

// VehicleBrand represents the manufacturer of vehicles
type VehicleBrand struct {
	gorm.Model

	Name          string      `json:"name" gorm:"not null"`
	Description   string      `json:"description"`
	VehicleTypeID string      `json:"vehicle_type_id" gorm:"type:uuid;not null"`
	VehicleType   VehicleType `json:"vehicle_type" gorm:"foreignKey:VehicleTypeID"`
}

// VehicleModel represents specific models of vehicles
type VehicleModel struct {
	gorm.Model

	Name        string       `json:"name" gorm:"not null"`
	Description string       `json:"description"`
	BrandID     string       `json:"brand_id" gorm:"type:uuid;not null"`
	Brand       VehicleBrand `json:"brand" gorm:"foreignKey:BrandID"`
	StartYear   int          `json:"start_year"` // First year of production
	EndYear     int          `json:"end_year"`   // Last year of production (0 if still in production)
}

// VehicleGeneration represents different generations/versions of a model
type VehicleGeneration struct {
	gorm.Model

	Name         string       `json:"name" gorm:"not null"`
	Description  string       `json:"description"`
	ModelID      string       `json:"model_id" gorm:"type:uuid;not null"`
	ModelInfo    VehicleModel `json:"model" gorm:"foreignKey:ModelID"`
	StartYear    int          `json:"start_year"`
	EndYear      int          `json:"end_year"`                      // 0 if still in production
	EngineType   string       `json:"engine_type"`                   // Gasoline, Diesel, Hybrid, Electric
	AssemblyType string       `json:"assembly_type" gorm:"not null"` // Import, CKD, SKD, etc.
	Assembler    string       `json:"assembler"`                     // e.g., Kerman Motor, IKCO, Saipa

	// Technical specifications
	Transmission    string `json:"transmission"`     // MT, AT, CVT, etc.
	EngineSize      int    `json:"engine_size"`      // in CC (e.g., 1600)
	BodyStyle       string `json:"body_style"`       // Sedan, Hatchback, Crossover, etc.
	SpecialFeatures string `json:"special_features"` // e.g., Panoramic Roof, Sport Package
}

// UserVehicle represents vehicles owned by users
type UserVehicle struct {
	gorm.Model

	UserID         string            `json:"user_id" gorm:"type:uuid;not null"`
	Name           string            `json:"name" gorm:"not null"`
	GenerationID   string            `json:"generation_id" gorm:"type:uuid;not null"`
	Generation     VehicleGeneration `json:"generation" gorm:"foreignKey:GenerationID"`
	ProductionYear int               `json:"production_year" gorm:"not null"`
	Color          string            `json:"color"`
	LicensePlate   string            `json:"license_plate"`
	VIN            string            `json:"vin"` // Vehicle Identification Number
	CurrentMileage int               `json:"current_mileage"`
	PurchaseDate   time.Time         `json:"purchase_date"`
}
