package usecase

import (
	"AutoBan/config"
	"AutoBan/internal/domain/entity"
	"AutoBan/internal/dto"
	"AutoBan/internal/errors"
	"AutoBan/internal/repository"
	"AutoBan/internal/validation"
	"AutoBan/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthUseCase interface defines the methods for authentication operations
// اینترفیس AuthUseCase متدهای مربوط به عملیات‌های احراز هویت را تعریف می‌کند

type AuthUseCase interface {
	Register(request *dto.RegisterRequest) error
	Login(request *dto.LoginRequest) (*dto.LoginResponse, error)
	Logout(request *dto.LogoutRequest) error
	RefreshToken(request *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(token string) (bool, error)
	ValidateRefreshToken(token string) (bool, error)
}

// authUseCase struct implements the AuthUseCase interface
// ساختار authUseCase اینترفیس AuthUseCase را پیاده‌سازی می‌کند

type authUseCase struct {
	authRepository repository.AuthRepository
	secretKey      string
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

func (a *authUseCase) Register(request *dto.RegisterRequest) error {
	err := validation.ValidateRegisterRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate register request")
		return err
	}

	user := &entity.User{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	}

	err = a.authRepository.Register(user)
	if err != nil {
		logger.Error(err, "Failed to register user")
		if err == errors.ErrUserAlreadyExists {
			return err
		}
		return errors.ErrInternalServerError
	}
	return nil
}

func (a *authUseCase) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
	err := validation.ValidateLoginRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate login request")
		return nil, err
	}

	user, err := a.authRepository.FindByPhoneNumber(request.PhoneNumber)
	if err != nil {
		logger.Error(err, "Failed to login user")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		logger.Error(err, "Failed to login user")
		return nil, errors.ErrInvalidPhoneNumberOrPassword
	}

	accessToken, err := a.GenerateAccessToken(user.ID.String())
	if err != nil {
		logger.Error(err, "Failed to generate access token")
		return nil, errors.ErrInternalServerError
	}

	refreshToken, err := a.GenerateRefreshToken(user.ID.String())
	if err != nil {
		logger.Error(err, "Failed to generate refresh token")
		return nil, errors.ErrInternalServerError
	}

	return &dto.LoginResponse{
		Token: dto.Token{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
func (a *authUseCase) Logout(request *dto.LogoutRequest) error {
	// todo: implement logout
	return nil
}

func (a *authUseCase) RefreshToken(request *dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	// todo: implement refresh token
	return nil, nil
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
