package database

import (
	"context"
	"sync"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

// GetRedisClient returns a singleton Redis client instance
func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		redisClient = ConnectRedis()
	})
	return redisClient
}

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

	logger.Info("Connected to Redis successfully")
	return client
}
