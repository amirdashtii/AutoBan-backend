package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"

	"github.com/redis/go-redis/v9"
)

type VehicleCacheRepository interface {
	GetVehicleHierarchy(ctx context.Context, vehicleTypes *[]entity.VehicleType) error
	SetVehicleHierarchy(ctx context.Context, vehicleTypes []entity.VehicleType) error
	InvalidateVehicleHierarchy(ctx context.Context) error
}

type vehicleCacheRepository struct {
	client *redis.Client
}

func NewVehicleCacheRepository() VehicleCacheRepository {
	return &vehicleCacheRepository{
		client: database.ConnectRedis(),
	}
}

// ساخت کلید برای ذخیره hierarchy در Redis
func makeVehicleHierarchyKey() string {
	return "vehicle:hierarchy:complete"
}

func (r *vehicleCacheRepository) GetVehicleHierarchy(ctx context.Context, vehicleTypes *[]entity.VehicleType) error {
	key := makeVehicleHierarchyKey()
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil
	}
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), &vehicleTypes); err != nil {
		return err
	}

	return nil
}

func (r *vehicleCacheRepository) SetVehicleHierarchy(ctx context.Context, vehicleTypes []entity.VehicleType) error {
	vehicleTypesData, err := json.Marshal(vehicleTypes)
	if err != nil {
		return err
	}

	key := makeVehicleHierarchyKey()
	// Set cache with 24 hour expiration
	err = r.client.Set(ctx, key, vehicleTypesData, 24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *vehicleCacheRepository) InvalidateVehicleHierarchy(ctx context.Context) error {
	key := makeVehicleHierarchyKey()
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
