package entity

import (
	"time"

	"github.com/google/uuid"
)

// VehicleType represents the main category of vehicles
type VehicleType struct {
	BaseModel

	NameFa        string `gorm:"index;not null;unique"`
	NameEn        string `gorm:"index;not null;unique"`
	DescriptionFa string
	DescriptionEn string

	// Relationships
	VehicleBrands []VehicleBrand `gorm:"foreignKey:VehicleTypeID;constraint:OnDelete:CASCADE"`
}

// VehicleBrand represents the manufacturer of vehicles
type VehicleBrand struct {
	BaseModel

	VehicleTypeID uint64 `gorm:"not null;uniqueIndex:idx_type_name"`
	NameFa        string `gorm:"index;not null;uniqueIndex:idx_type_name"`
	NameEn        string `gorm:"index;not null;uniqueIndex:idx_type_name"`
	DescriptionFa string
	DescriptionEn string

	// Relationships
	VehicleModels []VehicleModel `gorm:"foreignKey:BrandID;constraint:OnDelete:CASCADE"`
}

// VehicleModel represents specific models of vehicles
type VehicleModel struct {
	BaseModel

	BrandID       uint64 `gorm:"not null;uniqueIndex:idx_brand_name"`
	NameFa        string `gorm:"index;not null;uniqueIndex:idx_brand_name"`
	NameEn        string `gorm:"index;not null;uniqueIndex:idx_brand_name"`
	DescriptionFa string
	DescriptionEn string

	// Relationships
	VehicleGenerations []VehicleGeneration `gorm:"foreignKey:ModelID;constraint:OnDelete:CASCADE"`
}

// VehicleGeneration represents different generations/versions of a model
type VehicleGeneration struct {
	BaseModel

	ModelID       uint64 `gorm:"not null;uniqueIndex:idx_model_name"`
	NameFa        string `gorm:"index;not null;uniqueIndex:idx_model_name"`
	NameEn        string `gorm:"index;not null;uniqueIndex:idx_model_name"`
	DescriptionFa string
	DescriptionEn string
	StartYear     int    // First year of production
	EndYear       int    // Last year of production (0 if still in production)
	BodyStyleFa   string // سدان، هاچبک، کروس اور، ...
	BodyStyleEn   string // Sedan, Hatchback, Crossover, etc.
	Engine        string // Engine e.g. "1.6 TDI"
	EngineVolume  int    // in CC (e.g., 1600)
	Cylinders     int    // Number of cylinders
	DrivetrainFa  string // دودیفرانسیل
	DrivetrainEn  string // 4WD
	Gearbox       string // Automatic, Manual, CVT, etc.
	FuelType      string // Gasoline, Diesel, Hybrid, Electric
	Battery       string
	Seller        string
	AssemblyType  string
	Assembler     string
}

// UserVehicle represents vehicles owned by users
type UserVehicle struct {
	BaseModel

	UserID         uuid.UUID `gorm:"type:uuid;not null"`
	GenerationID   uint64    `gorm:"not null"`
	Name           string    `gorm:"not null"`
	ProductionYear int
	Color          string
	LicensePlate   string
	VIN            string // Vehicle Identification Number
	CurrentMileage int
	PurchaseDate   time.Time
}
