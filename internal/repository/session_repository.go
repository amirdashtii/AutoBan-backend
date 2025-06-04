package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type SessionRepository interface {
	SaveSession(ctx context.Context, session *entity.Session) error
	GetSession(ctx context.Context, userID, deviceID string) (*entity.Session, error)
	DeleteSession(ctx context.Context, userID, deviceID string) error
	DeleteAllSessions(ctx context.Context, userID string) error
	IsRefreshTokenValid(ctx context.Context, token string) bool
	GetAllSessions(ctx context.Context, userID string) ([]*entity.Session, error)
}

type sessionRepository struct {
	client *redis.Client
}

func NewSessionRepository() SessionRepository {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to get config: %v", err)
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	// تست اتصال به Redis
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		logger.Fatalf("Failed to connect to Redis: %v", err)
		return nil
	}

	return &sessionRepository{
		client: client,
	}
}

// ساخت کلید برای ذخیره نشست در Redis
func makeSessionKey(userID, deviceID string) string {
	return fmt.Sprintf("session:%s:%s", userID, deviceID)
}

func (r *sessionRepository) SaveSession(ctx context.Context, session *entity.Session) error {
	sessionData, err := json.Marshal(session)
	if err != nil {
		logger.Error(err, "Failed to marshal session")
		return errors.ErrInternalServerError
	}

	key := makeSessionKey(session.UserID, session.DeviceID)
	err = r.client.Set(ctx, key, sessionData, 7*24*time.Hour).Err()
	if err != nil {
		logger.Error(err, "Failed to save session to Redis")
		return errors.ErrInternalServerError
	}

	return nil
}

func (r *sessionRepository) GetSession(ctx context.Context, userID, deviceID string) (*entity.Session, error) {
	key := makeSessionKey(userID, deviceID)
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, errors.ErrTokenNotFound
	}
	if err != nil {
		logger.Error(err, "Failed to get session from Redis")
		return nil, errors.ErrInternalServerError
	}

	var session entity.Session
	if err := json.Unmarshal([]byte(data), &session); err != nil {
		logger.Error(err, "Failed to unmarshal session")
		return nil, errors.ErrInternalServerError
	}

	return &session, nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, userID, deviceID string) error {
	key := makeSessionKey(userID, deviceID)
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		logger.Error(err, "Failed to delete session from Redis")
		return errors.ErrInternalServerError
	}
	return nil
}

func (r *sessionRepository) DeleteAllSessions(ctx context.Context, userID string) error {
	pattern := fmt.Sprintf("session:%s:*", userID)
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		logger.Error(err, "Failed to get user sessions from Redis")
		return errors.ErrInternalServerError
	}

	if len(keys) > 0 {
		err = r.client.Del(ctx, keys...).Err()
		if err != nil {
			logger.Error(err, "Failed to delete user sessions from Redis")
			return errors.ErrInternalServerError
		}
	}
	return nil
}

func (r *sessionRepository) IsRefreshTokenValid(ctx context.Context, token string) bool {
	pattern := "session:*"
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		logger.Error(err, "Failed to get sessions")
		return false
	}

	for _, key := range keys {
		sessionData, err := r.client.Get(ctx, key).Result()
		if err != nil {
			continue
		}

		var session entity.Session
		err = json.Unmarshal([]byte(sessionData), &session)
		if err != nil {
			continue
		}

		if session.RefreshToken == token && session.IsActive {
			return true
		}
	}

	return false
}

func (r *sessionRepository) GetAllSessions(ctx context.Context, userID string) ([]*entity.Session, error) {
	pattern := fmt.Sprintf("session:%s:*", userID)
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	var sessions []*entity.Session
	for _, key := range keys {
		sessionData, err := r.client.Get(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, err
		}

		var session entity.Session
		err = json.Unmarshal([]byte(sessionData), &session)
		if err != nil {
			return nil, err
		}

		sessions = append(sessions, &session)
	}

	return sessions, nil
}
