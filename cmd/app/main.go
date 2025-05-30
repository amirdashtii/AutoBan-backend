package main

import (
	"AutoBan/config"
	"AutoBan/internal/interface/controller"
	"AutoBan/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.InitLogger()
	config, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()
	
	controller.AuthRoutes(r)

	r.Run(config.Server.Address + ":" + config.Server.Port) // listen and serve on specified address and port}
}
