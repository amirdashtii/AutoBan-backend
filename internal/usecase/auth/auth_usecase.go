package auth

import (
	"AutoBan/config"
	"AutoBan/internal/errors"
	"AutoBan/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// AuthUseCase interface defines the methods for authentication operations
// اینترفیس AuthUseCase متدهای مربوط به عملیات‌های احراز هویت را تعریف می‌کند

type AuthUseCase interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(token string) (bool, error)
	ValidateRefreshToken(token string) (bool, error)
}

// authUseCase struct implements the AuthUseCase interface
// ساختار authUseCase اینترفیس AuthUseCase را پیاده‌سازی می‌کند

type authUseCase struct {
	secretKey string
}

// NewAuthUseCase creates a new instance of authUseCase
// تابع NewAuthUseCase یک نمونه جدید از authUseCase ایجاد می‌کند

func NewAuthUseCase() AuthUseCase {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error(err, "Failed to get config")
		return nil
	}
	return &authUseCase{secretKey: cfg.JWT.Secret}
}

// ValidateToken validates a given JWT token
// تابع ValidateToken یک توکن JWT داده شده را اعتبارسنجی می‌کند

func (a *authUseCase) ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(a.secretKey), nil
	})

	if err != nil {
		logger.Error(err, "Failed to validate token")
		return false, errors.ErrInternalServerError
	}

	return token.Valid, nil
}

// GenerateAccessToken generates a new access token for a given user ID
// تابع GenerateAccessToken یک اکسس توکن جدید برای یک شناسه کاربری تولید می‌کند

func (a *authUseCase) GenerateAccessToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // 15 minutes expiration
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		logger.Error(err, "Failed to generate access token")
		return "", errors.ErrInternalServerError
	}

	return tokenString, nil
}

// GenerateRefreshToken generates a new refresh token for a given user ID
// تابع GenerateRefreshToken یک رفرش توکن جدید برای یک شناسه کاربری تولید می‌کند

func (a *authUseCase) GenerateRefreshToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiration
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		logger.Error(err, "Failed to generate refresh token")
		return "", errors.ErrInternalServerError
	}

	return tokenString, nil
}

// ValidateAccessToken validates a given access token
// تابع ValidateAccessToken یک اکسس توکن داده شده را اعتبارسنجی می‌کند

func (a *authUseCase) ValidateAccessToken(tokenString string) (bool, error) {
	return a.ValidateToken(tokenString)
}

// ValidateRefreshToken validates a given refresh token
// تابع ValidateRefreshToken یک رفرش توکن داده شده را اعتبارسنجی می‌کند

func (a *authUseCase) ValidateRefreshToken(tokenString string) (bool, error) {
	return a.ValidateToken(tokenString)
}
