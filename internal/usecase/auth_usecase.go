package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/internal/validation"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthUseCase interface defines the methods for authentication operations
type AuthUseCase interface {
	Register(request *dto.RegisterRequest) error
	Login(request *dto.LoginRequest) (*dto.TokenResponse, error)
	Logout(request *dto.LogoutRequest, userID string) error
	RefreshToken(request *dto.RefreshTokenRequest) (*dto.TokenResponse, error)
	GenerateAccessToken(user *entity.User) (string, error)
	GenerateRefreshToken(userID string, deviceID string) (string, error)
	GetUserSessions(userID string) ([]*entity.Session, error)
	LogoutAllDevices(userID string) error
}

// authUseCase struct implements the AuthUseCase interface
type authUseCase struct {
	authRepository    repository.AuthRepository
	sessionRepository repository.SessionRepository
	secretKey         string
}

// NewAuthUseCase creates a new instance of authUseCase
func NewAuthUseCase() AuthUseCase {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error(err, "Failed to get config")
		return nil
	}
	authRepository := repository.NewAuthRepository()
	sessionRepository := repository.NewSessionRepository()
	return &authUseCase{
		authRepository:    authRepository,
		sessionRepository: sessionRepository,
		secretKey:         cfg.JWT.Secret,
	}
}

func (a *authUseCase) Register(request *dto.RegisterRequest) error {
	err := validation.ValidateRegisterRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate register request")
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err, "Failed to hash password")
		return errors.ErrInternalServerError
	}

	user := entity.NewUser(request.PhoneNumber, string(hashedPassword))

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

func (a *authUseCase) Login(request *dto.LoginRequest) (*dto.TokenResponse, error) {
	err := validation.ValidateLoginRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate login request")
		return nil, err
	}

	user, err := a.authRepository.FindByPhoneNumber(request.PhoneNumber)
	if err != nil {
		logger.Error(err, "Failed to find user")
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		logger.Error(err, "Failed to compare hash and password")
		return nil, errors.ErrInvalidPhoneNumberOrPassword
	}

	deviceID := generateDeviceID()
	tokens, err := a.GenerateTokens(user, deviceID)
	if err != nil {
		return nil, errors.ErrInternalServerError
	}

	// ذخیره نشست در Redis
	session := entity.NewSession(user.ID.String(), deviceID, tokens.RefreshToken)
	err = a.sessionRepository.SaveSession(context.Background(), session)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (a *authUseCase) Logout(request *dto.LogoutRequest, userID string) error {
	// پارس کردن توکن برای دریافت شناسه کاربر و دستگاه
	token, err := jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidToken
		}
		return []byte(a.secretKey), nil
	})

	if err != nil {
		logger.Error(err, "Failed to parse refresh token")
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.Error(nil, "Failed to get token claims")
		return errors.ErrInvalidToken
	}
	if userID != claims["user_id"].(string) {
		logger.Error(nil, "User ID does not match")
		return errors.ErrInvalidToken
	}

	deviceID := claims["device_id"].(string)

	// حذف نشست از Redis
	err = a.sessionRepository.DeleteSession(context.Background(), userID, deviceID)
	if err != nil {
		return err
	}

	return nil
}

func (a *authUseCase) RefreshToken(request *dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	// چک کردن اعتبار توکن در وایت‌لیست
	if !a.sessionRepository.IsRefreshTokenValid(context.Background(), request.RefreshToken) {
		logger.Error(nil, "Token is not in whitelist")
		return nil, errors.ErrInvalidToken
	}
	// پارس کردن توکن
	token, err := jwt.Parse(request.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrInvalidToken
		}
		return []byte(a.secretKey), nil
	})

	if err != nil {
		logger.Error(err, "Failed to parse refresh token")
		return nil, errors.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logger.Error(nil, "Failed to get token claims")
		return nil, errors.ErrInvalidToken
	}

	userID := claims["user_id"].(string)
	deviceID := claims["device_id"].(string)

	// چک کردن وجود نشست در Redis
	session, err := a.sessionRepository.GetSession(context.Background(), userID, deviceID)
	if err != nil {
		logger.Error(err, "Failed to get session")
		return nil, errors.ErrInvalidToken
	}

	if !session.IsActive {
		logger.Error(nil, "Session is not active")
		return nil, errors.ErrInvalidToken
	}

	// دریافت اطلاعات کاربر
	user, err := a.authRepository.FindByID(userID)
	if err != nil {
		logger.Error(err, "Failed to find user")
		return nil, err
	}

	// ایجاد توکن‌های جدید
	tokens, err := a.GenerateTokens(user, deviceID)
	if err != nil {
		logger.Error(err, "Failed to generate new access token")
		return nil, errors.ErrInternalServerError
	}

	// بروزرسانی نشست در Redis
	session.RefreshToken = tokens.RefreshToken
	session.LastUsed = time.Now()
	err = a.sessionRepository.SaveSession(context.Background(), session)
	if err != nil {
		logger.Error(err, "Failed to update session")
		return nil, errors.ErrInternalServerError
	}

	return &tokens, nil
}

func generateDeviceID() string {
	return fmt.Sprintf("dev_%s", uuid.New().String())
}

func (a *authUseCase) GenerateTokens(user *entity.User, deviceID string) (dto.TokenResponse, error) {
	accessToken, err := a.GenerateAccessToken(user)
	if err != nil {
		return dto.TokenResponse{}, err
	}

	refreshToken, err := a.GenerateRefreshToken(user.ID.String(), deviceID)
	if err != nil {
		return dto.TokenResponse{}, err
	}

	return dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authUseCase) GenerateAccessToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      user.ID.String(),
		"role":         user.Role,
		"phone_number": user.PhoneNumber,
		"status":       user.Status,
		"exp":          time.Now().Add(time.Minute * 15).Unix(), // 15 minutes expiration
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		logger.Error(err, "Failed to generate access token")
		return "", errors.ErrInternalServerError
	}

	return tokenString, nil
}

func (a *authUseCase) GenerateRefreshToken(userID string, deviceID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"device_id": deviceID,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 روز
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secretKey))
}

func (a *authUseCase) GetUserSessions(userID string) ([]*entity.Session, error) {
	sessions, err := a.sessionRepository.GetAllSessions(context.Background(), userID)
	if err != nil {
		logger.Error(err, "Failed to get user sessions")
		return nil, errors.ErrInternalServerError
	}
	return sessions, nil
}

func (a *authUseCase) LogoutAllDevices(userID string) error {
	err := a.sessionRepository.DeleteAllSessions(context.Background(), userID)
	if err != nil {
		return err
	}
	return nil
}
