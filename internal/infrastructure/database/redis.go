package database

import (
	"context"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error(err, "Failed to get config")
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		logger.Error(err, "Failed to connect to Redis")
		return nil
	}

	return client
}
