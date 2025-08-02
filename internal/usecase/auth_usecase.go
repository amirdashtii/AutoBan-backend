package usecase

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/infrastructure/http"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/internal/validation"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// AuthUseCase interface defines the methods for authentication operations
type AuthUseCase interface {
	Register(ctx context.Context, request *dto.RegisterRequest) (*dto.TokenResponse, error)
	Login(ctx context.Context, request *dto.LoginRequest) (*dto.TokenResponse, error)
	Logout(ctx context.Context, request *dto.LogoutRequest, userID string) error
	RefreshToken(ctx context.Context, request *dto.RefreshTokenRequest) (*dto.TokenResponse, error)
	SendVerificationCode(ctx context.Context, request *dto.VerifyPhoneRequest) error
	VerifyCode(ctx context.Context, request *dto.VerifyCodeRequest) (*dto.TokenResponse, error)
	GenerateAccessToken(ctx context.Context, user *entity.User) (string, error)
	GenerateRefreshToken(ctx context.Context, userID string, deviceID string) (string, error)
	GetUserSessions(ctx context.Context, userID string) ([]dto.SessionResponse, error)
	LogoutAllDevices(ctx context.Context, userID string) error
}

// authUseCase struct implements the AuthUseCase interface
type authUseCase struct {
	authRepository         repository.AuthRepository
	sessionRepository      repository.SessionRepository
	verificationRepository repository.VerificationRepository
	smsService             http.SMSService
	secretKey              string
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
	verificationRepository := repository.NewVerificationRepository()
	smsService := http.NewSMSService(cfg.SMS.BaseURL, cfg.SMS.XAPIKey)
	return &authUseCase{
		authRepository:         authRepository,
		sessionRepository:      sessionRepository,
		verificationRepository: verificationRepository,
		smsService:             smsService,
		secretKey:              cfg.JWT.Secret,
	}
}

func (a *authUseCase) Register(ctx context.Context, request *dto.RegisterRequest) (*dto.TokenResponse, error) {
	err := validation.ValidateRegisterRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate register request")
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err, "Failed to hash password")
		return nil, errors.ErrInternalServerError
	}

	user := entity.NewUser(request.PhoneNumber, string(hashedPassword))

	err = a.authRepository.Register(ctx, user)
	if err != nil {
		logger.Error(err, "Failed to register user")
		if err == errors.ErrUserAlreadyExists {
			return nil, err
		}
		return nil, errors.ErrInternalServerError
	}

	// login user
	tokens, err := a.Login(ctx, &dto.LoginRequest{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
	})
	if err != nil {
		return nil, errors.TokenGenerationFailed
	}

	return tokens, nil
}

func (a *authUseCase) Login(ctx context.Context, request *dto.LoginRequest) (*dto.TokenResponse, error) {
	err := validation.ValidateLoginRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate login request")
		return nil, err
	}

	var user entity.User
	user.PhoneNumber = request.PhoneNumber
	err = a.authRepository.FindByPhoneNumber(ctx, &user)
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
	tokens, err := a.GenerateTokens(ctx, &user, deviceID)
	if err != nil {
		logger.Error(err, "Failed to generate tokens")
		return nil, errors.ErrInternalServerError
	}

	// ذخیره نشست در Redis
	session := entity.NewSession(user.ID.String(), deviceID, tokens.RefreshToken)
	err = a.sessionRepository.SaveSession(ctx, session)
	if err != nil {
		return nil, err
	}

	return &tokens, nil
}

func (a *authUseCase) Logout(ctx context.Context, request *dto.LogoutRequest, userID string) error {
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
	var session entity.Session
	session.UserID = userID
	session.DeviceID = deviceID
	err = a.sessionRepository.DeleteSession(ctx, &session)
	if err != nil {
		return err
	}

	return nil
}

func (a *authUseCase) RefreshToken(ctx context.Context, request *dto.RefreshTokenRequest) (*dto.TokenResponse, error) {
	// چک کردن اعتبار توکن در وایت‌لیست
	if !a.sessionRepository.IsRefreshTokenValid(ctx, request.RefreshToken) {
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

	var session entity.Session
	// چک کردن وجود نشست در Redis
	session.UserID = userID
	session.DeviceID = deviceID
	err = a.sessionRepository.GetSession(ctx, &session)
	if err != nil {
		logger.Error(err, "Failed to get session")
		return nil, errors.ErrInvalidToken
	}

	if !session.IsActive {
		logger.Error(nil, "Session is not active")
		return nil, errors.ErrInvalidToken
	}

	// دریافت اطلاعات کاربر
	var user entity.User
	user.ID = uuid.MustParse(userID)
	err = a.authRepository.FindByID(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to find user")
		return nil, err
	}

	// ایجاد توکن‌های جدید
	tokens, err := a.GenerateTokens(ctx, &user, deviceID)
	if err != nil {
		logger.Error(err, "Failed to generate new access token")
		return nil, errors.ErrInternalServerError
	}

	// بروزرسانی نشست در Redis
	session.RefreshToken = tokens.RefreshToken
	session.LastUsed = time.Now()
	err = a.sessionRepository.SaveSession(ctx, &session)
	if err != nil {
		logger.Error(err, "Failed to update session")
		return nil, errors.ErrInternalServerError
	}

	return &tokens, nil
}

func (a *authUseCase) SendVerificationCode(ctx context.Context, request *dto.VerifyPhoneRequest) error {
	err := validation.ValidateVerifyPhoneRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate verify phone request")
		return err
	}

	var user entity.User
	user.PhoneNumber = request.PhoneNumber
	err = a.authRepository.FindByPhoneNumber(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to find user")
		return err
	}

	code := generateCode()
	fmt.Println(code)

	// ذخیره کد تایید در Redis
	err = a.verificationRepository.SaveVerificationCode(ctx, request.PhoneNumber, code)
	if err != nil {
		logger.Error(err, "Failed to save verification code to Redis")
		return errors.ErrInternalServerError
	}
	logger.Info(fmt.Sprintf("Verification code saved for phone: %s", request.PhoneNumber))

	// ارسال کد تایید از طریق SMS
	err = a.smsService.SendVerificationCode(ctx, request.PhoneNumber, code)
	if err != nil {
		logger.Error(err, "Failed to send verification code via SMS")
		return errors.ErrInternalServerError
	}
	logger.Info(fmt.Sprintf("SMS verification code sent successfully to %s", request.PhoneNumber))

	return nil
}

func (a *authUseCase) VerifyCode(ctx context.Context, request *dto.VerifyCodeRequest) (*dto.TokenResponse, error) {
	// اعتبارسنجی درخواست
	if err := validation.ValidateVerifyCodeRequest(request); err != nil {
		logger.Error(err, "Failed to validate verify code request")
		return nil, err
	}
	var user entity.User
	user.PhoneNumber = request.PhoneNumber
	err := a.authRepository.FindByPhoneNumber(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to find user")
		return nil, err
	}

	// بررسی اعتبار کد تایید
	if !a.verificationRepository.IsVerificationCodeValid(ctx, request.PhoneNumber, request.Code) {
		logger.Error(errors.ErrInvalidVerificationCode, "Invalid verification code")
		return nil, errors.ErrInvalidVerificationCode
	}

	// حذف کد تایید پس از تایید موفق
	err = a.verificationRepository.DeleteVerificationCode(ctx, request.PhoneNumber)
	if err != nil {
		logger.Error(err, "Failed to delete verification code after successful verification")
		// این خطا نباید باعث شکست تایید شود
	}

	user.Status = entity.Active
	err = a.authRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to update user status")
		return nil, err
	}
	logger.Info(fmt.Sprintf("Verification code deleted for phone: %s", request.PhoneNumber))

	deviceID := generateDeviceID()
	tokens, err := a.GenerateTokens(ctx, &user, deviceID)
	if err != nil {
		logger.Error(err, "Failed to generate tokens")
		return nil, errors.ErrInternalServerError
	}

	session := entity.NewSession(user.ID.String(), deviceID, tokens.RefreshToken)
	err = a.sessionRepository.SaveSession(ctx, session)
	if err != nil {
		logger.Error(err, "Failed to save session")
		return nil, err
	}

	return &tokens, nil
}

func generateCode() string {
	return fmt.Sprintf("%d", rand.Intn(1000000))
}

func generateDeviceID() string {
	return fmt.Sprintf("dev_%s", uuid.New().String())
}

func (a *authUseCase) GenerateTokens(ctx context.Context, user *entity.User, deviceID string) (dto.TokenResponse, error) {
	accessToken, err := a.GenerateAccessToken(ctx, user)
	if err != nil {
		return dto.TokenResponse{}, err
	}

	refreshToken, err := a.GenerateRefreshToken(ctx, user.ID.String(), deviceID)
	if err != nil {
		return dto.TokenResponse{}, err
	}

	return dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *authUseCase) GenerateAccessToken(ctx context.Context, user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      user.ID.String(),
		"role":         user.Role,
		"phone_number": user.PhoneNumber,
		"status":       user.Status,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // 1 day expiration
	})

	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		logger.Error(err, "Failed to generate access token")
		return "", errors.ErrInternalServerError
	}

	return tokenString, nil
}

func (a *authUseCase) GenerateRefreshToken(ctx context.Context, userID string, deviceID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"device_id": deviceID,
		"exp":       time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days expiration
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.secretKey))
}

func (a *authUseCase) GetUserSessions(ctx context.Context, userID string) ([]dto.SessionResponse, error) {
	var sessions []entity.Session
	err := a.sessionRepository.GetAllSessions(ctx, userID, &sessions)
	if err != nil {
		logger.Error(err, "Failed to get user sessions")
		return nil, errors.ErrInternalServerError
	}
	// تبدیل نشست‌ها به مدل پاسخ
	var sessionResponses []dto.SessionResponse
	for _, session := range sessions {
		sessionResponses = append(sessionResponses, dto.SessionResponse{
			DeviceID: session.DeviceID,
			LastUsed: session.LastUsed.Format(time.RFC3339),
			IsActive: session.IsActive,
		})
	}
	return sessionResponses, nil
}

func (a *authUseCase) LogoutAllDevices(ctx context.Context, userID string) error {
	err := a.sessionRepository.DeleteAllSessions(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
