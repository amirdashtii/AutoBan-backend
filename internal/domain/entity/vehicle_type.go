package entity

import (
	"time"

	"github.com/google/uuid"
)

// VehicleType represents the main category of vehicles
type VehicleType struct {
	BaseModel

	Name        string `json:"name" gorm:"index;not null;unique"`
	Description string `json:"description"`

	// Relationships
	VehicleBrands []VehicleBrand `json:"vehicle_brands,omitempty" gorm:"foreignKey:VehicleTypeID;constraint:OnDelete:CASCADE"`
}

// VehicleBrand represents the manufacturer of vehicles
type VehicleBrand struct {
	BaseModel

	Name          string      `json:"name" gorm:"index;not null;uniqueIndex:idx_type_name"`
	Description   string      `json:"description"`
	VehicleTypeID uint64      `json:"vehicle_type_id" gorm:"not null;uniqueIndex:idx_type_name"`

	// Relationships
	VehicleModels []VehicleModel `json:"vehicle_models,omitempty" gorm:"foreignKey:BrandID;constraint:OnDelete:CASCADE"`
}

// VehicleModel represents specific models of vehicles
type VehicleModel struct {
	BaseModel

	Name        string       `json:"name" gorm:"index;not null;uniqueIndex:idx_brand_name"`
	Description string       `json:"description"`
	BrandID     uint64       `json:"brand_id" gorm:"not null;uniqueIndex:idx_brand_name"`
	StartYear   int          `json:"start_year"` // First year of production
	EndYear     int          `json:"end_year"`   // Last year of production (0 if still in production)

	// Relationships
	VehicleGenerations []VehicleGeneration `json:"vehicle_generations,omitempty" gorm:"foreignKey:ModelID;constraint:OnDelete:CASCADE"`
}

// VehicleGeneration represents different generations/versions of a model
type VehicleGeneration struct {
	BaseModel

	Name         string       `json:"name" gorm:"index;not null;uniqueIndex:idx_model_name"`
	Description  string       `json:"description"`
	ModelID      uint64       `json:"model_id" gorm:"not null;uniqueIndex:idx_model_name"`
	StartYear    int          `json:"start_year"`    // First year of production
	EndYear      int          `json:"end_year"`      // Last year of production (0 if still in production)
	EngineType   string       `json:"engine_type"`   // Gasoline, Diesel, Hybrid, Electric
	AssemblyType string       `json:"assembly_type"` // Import, CKD, SKD, etc.
	Assembler    string       `json:"assembler"`     // e.g., Kerman Motor, IKCO, Saipa

	// Technical specifications
	Transmission    string `json:"transmission"`     // MT, AT, CVT, etc.
	EngineSize      int    `json:"engine_size"`      // in CC (e.g., 1600)
	BodyStyle       string `json:"body_style"`       // Sedan, Hatchback, Crossover, etc.
	SpecialFeatures string `json:"special_features"` // e.g., Panoramic Roof, Sport Package
}

// UserVehicle represents vehicles owned by users
type UserVehicle struct {
	BaseModel

	UserID         uuid.UUID         `json:"user_id" gorm:"type:uuid;not null"`
	Name           string            `json:"name" gorm:"not null"`
	GenerationID   uint64            `json:"generation_id" gorm:"not null"`
	ProductionYear int               `json:"production_year"`
	Color          string            `json:"color"`
	LicensePlate   string            `json:"license_plate"`
	VIN            string            `json:"vin"` // Vehicle Identification Number
	CurrentMileage int               `json:"current_mileage"`
	PurchaseDate   time.Time         `json:"purchase_date"`
}
