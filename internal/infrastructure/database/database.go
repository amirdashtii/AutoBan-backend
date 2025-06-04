package database

import (
	"fmt"
	"sync"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

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
		err = db.AutoMigrate(&entity.User{}) // Add other models as needed
		if err != nil {
			logger.Error(err, "failed to migrate database")
		}

	})
	return db
}
