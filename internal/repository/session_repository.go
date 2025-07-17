package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/redis/go-redis/v9"
)

type SessionRepository interface {
	SaveSession(ctx context.Context, session *entity.Session) error
	GetSession(ctx context.Context, session *entity.Session) error
	DeleteSession(ctx context.Context, session *entity.Session) error
	DeleteAllSessions(ctx context.Context, userID string) error
	IsRefreshTokenValid(ctx context.Context, token string) bool
	GetAllSessions(ctx context.Context, userID string, sessions *[]entity.Session) error
}

type sessionRepository struct {
	client *redis.Client
}

func NewSessionRepository() SessionRepository {
	return &sessionRepository{
		client: database.ConnectRedis(),
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

func (r *sessionRepository) GetSession(ctx context.Context, session *entity.Session) error {
	key := makeSessionKey(session.UserID, session.DeviceID)
	data, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return errors.ErrTokenNotFound
	}
	if err != nil {
		logger.Error(err, "Failed to get session from Redis")
		return errors.ErrInternalServerError
	}

	if err := json.Unmarshal([]byte(data), session); err != nil {
		logger.Error(err, "Failed to unmarshal session")
		return errors.ErrInternalServerError
	}

	return nil
}

func (r *sessionRepository) DeleteSession(ctx context.Context, session *entity.Session) error {
	key := makeSessionKey(session.UserID, session.DeviceID)
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

func (r *sessionRepository) GetAllSessions(ctx context.Context, userID string, sessions *[]entity.Session) error {
	pattern := fmt.Sprintf("session:%s:*", userID)
	keys, err := r.client.Keys(ctx, pattern).Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		sessionData, err := r.client.Get(ctx, key).Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return err
		}

		var session entity.Session
		err = json.Unmarshal([]byte(sessionData), &session)
		if err != nil {
			return err
		}

		*sessions = append(*sessions, session)
	}

	return nil
}
