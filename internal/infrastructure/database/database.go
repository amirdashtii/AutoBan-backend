package database

import (
	"fmt"
	"sync"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// createSuperAdmin creates a super admin user if it doesn't exist
func createSuperAdmin(db *gorm.DB) error {
	var count int64
	if err := db.Model(&entity.User{}).Where("phone_number = ?", "09000000000").Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Admin123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		superAdmin := entity.NewUser("09000000000", string(hashedPassword))
		superAdmin.Role = entity.SuperAdminRole

		if err := db.Create(superAdmin).Error; err != nil {
			return err
		}

		logger.Info("Super admin user created successfully")
	}

	return nil
}

// ConnectDatabase initializes the database connection and performs migrations
func ConnectDatabase() *gorm.DB {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error(err, "Failed to get config")
		return nil
	}

	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Error(err, "failed to connect database")
		}

		// Perform migrations
		err = db.AutoMigrate(
			&entity.User{},
			&entity.VehicleType{},
			&entity.VehicleBrand{},
			&entity.VehicleModel{},
			&entity.VehicleGeneration{},
			&entity.UserVehicle{},
			&entity.OilChange{},
		// Add other models as needed
		)
		if err != nil {
			logger.Error(err, "failed to migrate database")
		}

		// Create super admin user
		if err := createSuperAdmin(db); err != nil {
			logger.Error(err, "failed to create super admin user")
		}
	})
	return db
}
