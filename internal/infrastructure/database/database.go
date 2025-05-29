package database

import (
	"log"
	"sync"

	"AutoBan/internal/domain/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// ConnectDatabase initializes the database connection and performs migrations
func ConnectDatabase() *gorm.DB {
	once.Do(func() {
		dsn := "host=localhost user=autoban password=autoban dbname=autoban port=5432 sslmode=disable"
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("failed to connect database: ", err)
		}

		// Perform migrations
		db.AutoMigrate(&entity.User{}) // Add other models as needed
	})
	return db
}
