package database

import (
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"gorm.io/gorm"
)

// RunMigrations runs all database migrations and creates indexes
func RunMigrations(db *gorm.DB) error {
	logger.Info("Starting database migrations...")

	// Auto migrate all tables
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Session{},
		&entity.VehicleType{},
		&entity.VehicleBrand{},
		&entity.VehicleModel{},
		&entity.VehicleGeneration{},
		&entity.UserVehicle{},
		&entity.ServiceVisit{},
		&entity.OilChange{},
		&entity.OilFilter{},
	)
	if err != nil {
		logger.Error(err, "Failed to run auto migrations")
		return err
	}

	// Create performance indexes
	err = createPerformanceIndexes(db)
	if err != nil {
		logger.Error(err, "Failed to create performance indexes")
		return err
	}

	logger.Info("Database migrations completed successfully")
	return nil
}

// createPerformanceIndexes creates indexes for better query performance
func createPerformanceIndexes(db *gorm.DB) error {
	logger.Info("Creating performance indexes...")

	// User table indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_users_phone_number ON users(phone_number)").Error; err != nil {
		logger.Error(err, "Failed to create index on users.phone_number")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_users_status ON users(status)").Error; err != nil {
		logger.Error(err, "Failed to create index on users.status")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_users_role ON users(role)").Error; err != nil {
		logger.Error(err, "Failed to create index on users.role")
		return err
	}

	// Vehicle hierarchy indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_vehicle_brands_type_id ON vehicle_brands(vehicle_type_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on vehicle_brands.vehicle_type_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_vehicle_models_brand_id ON vehicle_models(brand_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on vehicle_models.brand_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_vehicle_generations_model_id ON vehicle_generations(model_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on vehicle_generations.model_id")
		return err
	}

	// User vehicles indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_user_vehicles_user_id ON user_vehicles(user_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on user_vehicles.user_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_user_vehicles_generation_id ON user_vehicles(generation_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on user_vehicles.generation_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_user_vehicles_license_plate ON user_vehicles(license_plate)").Error; err != nil {
		logger.Error(err, "Failed to create index on user_vehicles.license_plate")
		return err
	}

	// Service visits indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_service_visits_user_vehicle_id ON service_visits(user_vehicle_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on service_visits.user_vehicle_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_service_visits_service_date ON service_visits(service_date)").Error; err != nil {
		logger.Error(err, "Failed to create index on service_visits.service_date")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_service_visits_service_mileage ON service_visits(service_mileage)").Error; err != nil {
		logger.Error(err, "Failed to create index on service_visits.service_mileage")
		return err
	}

	// Oil changes indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_oil_changes_service_visit_id ON oil_changes(service_visit_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on oil_changes.service_visit_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_oil_changes_change_date ON oil_changes(change_date)").Error; err != nil {
		logger.Error(err, "Failed to create index on oil_changes.change_date")
		return err
	}

	// Oil filters indexes
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_oil_filters_service_visit_id ON oil_filters(service_visit_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on oil_filters.service_visit_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_oil_filters_change_date ON oil_filters(change_date)").Error; err != nil {
		logger.Error(err, "Failed to create index on oil_filters.change_date")
		return err
	}

	// Sessions indexes (for Redis-like behavior in case of fallback)
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on sessions.user_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_sessions_device_id ON sessions(device_id)").Error; err != nil {
		logger.Error(err, "Failed to create index on sessions.device_id")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_sessions_is_active ON sessions(is_active)").Error; err != nil {
		logger.Error(err, "Failed to create index on sessions.is_active")
		return err
	}

	// Composite indexes for common query patterns
	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_user_vehicles_user_id_created_at ON user_vehicles(user_id, created_at DESC)").Error; err != nil {
		logger.Error(err, "Failed to create composite index on user_vehicles")
		return err
	}

	if err := db.Exec("CREATE INDEX IF NOT EXISTS idx_service_visits_vehicle_date ON service_visits(user_vehicle_id, service_date DESC)").Error; err != nil {
		logger.Error(err, "Failed to create composite index on service_visits")
		return err
	}

	logger.Info("Performance indexes created successfully")
	return nil
}
