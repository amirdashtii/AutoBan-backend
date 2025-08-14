package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/redis/go-redis/v9"
)

// CacheRepository interface for caching operations
type CacheRepository interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	Delete(ctx context.Context, key string) error
	DeleteByPattern(ctx context.Context, pattern string) error
	Exists(ctx context.Context, key string) bool
	SetWithTags(ctx context.Context, key string, value interface{}, expiration time.Duration, tags []string) error
	InvalidateByTag(ctx context.Context, tag string) error
}

type cacheRepository struct {
	client *redis.Client
}

// NewCacheRepository creates a new cache repository
func NewCacheRepository() CacheRepository {
	client := database.GetRedisClient()
	return &cacheRepository{
		client: client,
	}
}

// Set stores a value in cache with expiration
func (r *cacheRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		logger.Error(err, "Failed to marshal value for caching")
		return err
	}

	err = r.client.Set(ctx, key, jsonValue, expiration).Err()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to set cache key: %s", key))
		return err
	}

	logger.Info(fmt.Sprintf("Cache set successfully for key: %s", key))
	return nil
}

// Get retrieves a value from cache
func (r *cacheRepository) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return fmt.Errorf("cache miss for key: %s", key)
		}
		logger.Error(err, fmt.Sprintf("Failed to get cache key: %s", key))
		return err
	}

	err = json.Unmarshal([]byte(val), dest)
	if err != nil {
		logger.Error(err, "Failed to unmarshal cached value")
		return err
	}

	logger.Info(fmt.Sprintf("Cache hit for key: %s", key))
	return nil
}

// Delete removes a key from cache
func (r *cacheRepository) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to delete cache key: %s", key))
		return err
	}

	logger.Info(fmt.Sprintf("Cache deleted for key: %s", key))
	return nil
}

// DeleteByPattern removes all keys matching a pattern
func (r *cacheRepository) DeleteByPattern(ctx context.Context, pattern string) error {
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to get keys for pattern: %s", pattern))
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	err = r.client.Del(ctx, keys...).Err()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to delete keys for pattern: %s", pattern))
		return err
	}

	logger.Info(fmt.Sprintf("Cache deleted for pattern: %s, keys count: %d", pattern, len(keys)))
	return nil
}

// Exists checks if a key exists in cache
func (r *cacheRepository) Exists(ctx context.Context, key string) bool {
	result, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to check existence of cache key: %s", key))
		return false
	}
	return result > 0
}

// SetWithTags stores a value with associated tags for invalidation
func (r *cacheRepository) SetWithTags(ctx context.Context, key string, value interface{}, expiration time.Duration, tags []string) error {
	// First set the main cache entry
	err := r.Set(ctx, key, value, expiration)
	if err != nil {
		return err
	}

	// Then associate the key with tags
	for _, tag := range tags {
		tagKey := fmt.Sprintf("tag:%s", tag)
		err = r.client.SAdd(ctx, tagKey, key).Err()
		if err != nil {
			logger.Error(err, fmt.Sprintf("Failed to add key to tag: %s", tag))
			continue
		}
		// Set expiration for tag set (longer than cache entries)
		r.client.Expire(ctx, tagKey, expiration+time.Hour)
	}

	return nil
}

// InvalidateByTag removes all cache entries associated with a tag
func (r *cacheRepository) InvalidateByTag(ctx context.Context, tag string) error {
	tagKey := fmt.Sprintf("tag:%s", tag)
	
	// Get all keys associated with this tag
	keys, err := r.client.SMembers(ctx, tagKey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil // No keys to invalidate
		}
		logger.Error(err, fmt.Sprintf("Failed to get members for tag: %s", tag))
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	// Delete all associated cache entries
	err = r.client.Del(ctx, keys...).Err()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to delete cache entries for tag: %s", tag))
		return err
	}

	// Delete the tag set itself
	err = r.client.Del(ctx, tagKey).Err()
	if err != nil {
		logger.Error(err, fmt.Sprintf("Failed to delete tag set: %s", tag))
	}

	logger.Info(fmt.Sprintf("Cache invalidated for tag: %s, keys count: %d", tag, len(keys)))
	return nil
}

// Cache key builders for consistency
func BuildCacheKey(prefix string, parts ...string) string {
	key := fmt.Sprintf("autoban:%s", prefix)
	for _, part := range parts {
		key = fmt.Sprintf("%s:%s", key, part)
	}
	return key
}

// Common cache keys
const (
	CacheKeyVehicleHierarchy = "vehicle:hierarchy"
	CacheKeyUserProfile      = "user:profile"
	CacheKeyUserVehicles     = "user:vehicles"
	CacheKeyVehicleTypes     = "vehicle:types"
	CacheKeyBrands           = "vehicle:brands"
	CacheKeyModels           = "vehicle:models"
	CacheKeyGenerations      = "vehicle:generations"
)

// Common cache tags
const (
	CacheTagVehicleData = "vehicle_data"
	CacheTagUserData    = "user_data"
	CacheTagStaticData  = "static_data"
)
