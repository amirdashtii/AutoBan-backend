package main

import (
	"AutoBan/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	r := gin.Default()

	r.Run(config.Server.Address + ":" + config.Server.Port) // listen and serve on specified address and port}
}
