package main

import (
	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/interface/controller"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	_ "github.com/amirdashtii/AutoBan/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           AutoBan API
// @version         1.0
// @description     This is a sample server for AutoBan.

// @license.name   GNU General Public License v3.0
// @license.url    https://www.gnu.org/licenses/gpl-3.0.en.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Type "Bearer" followed by a space and JWT token.

// @schemes http https

// @tag.name        Authentication
// @tag.description Authentication operations

// @tag.name        Users
// @tag.description User management operations

// @tag.name        User - Vehicles
// @tag.description User vehicle management

// @tag.name        Types
// @tag.description Vehicle types management

// @tag.name        Brands
// @tag.description Vehicle brands management

// @tag.name        Models
// @tag.description Vehicle models management

// @tag.name        Generations
// @tag.description Vehicle generations management

// @tag.name        Admin - Users
// @tag.description Admin user management operations

// @tag.name        Admin - Types
// @tag.description Admin vehicle type management operations

// @tag.name        Admin - Brands
// @tag.description Admin vehicle brand management operations

// @tag.name        Admin - Models
// @tag.description Admin vehicle model management operations

// @tag.name        Admin - Generations
// @tag.description Admin vehicle generation management operations

// @tag.name        Admin - UserVehicles
// @tag.description Admin user vehicle management operations

// @tag.name        Service Visits
// @tag.description Service visit management operations

// @tag.name        Oil Changes
// @tag.description Oil change management operations

// @tag.name        Oil Filters
// @tag.description Oil filter management operations

func main() {
	logger.InitLogger()
	config, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	// Setup routes
	controller.AuthRoutes(r)
	controller.UserRoutes(r)
	controller.AdminRoutes(r)
	controller.VehicleRoutes(r)
	controller.ServiceVisitRoutes(r)
	controller.OilChangeRoutes(r)
	controller.OilFilterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(config.Server.Address + ":" + config.Server.Port) // listen and serve on specified address and port
}
