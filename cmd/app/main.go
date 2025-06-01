package main

import (
	"AutoBan/config"
	"AutoBan/internal/interface/controller"
	"AutoBan/pkg/logger"

	_ "AutoBan/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           AutoBan API
// @version         1.0
// @description     This is a sample server for AutoBan.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey  BearerAuth
// @in                         header
// @name                       Authorization
// @description               Type "Bearer" followed by a space and JWT token.

// @schemes http https
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(config.Server.Address + ":" + config.Server.Port) // listen and serve on specified address and port
}
