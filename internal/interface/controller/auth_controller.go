package controller

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/middleware"
	"github.com/amirdashtii/AutoBan/internal/usecase"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController() *AuthController {
	authUseCase := usecase.NewAuthUseCase()
	return &AuthController{authUseCase: authUseCase}
}

func AuthRoutes(router *gin.Engine) {
	c := NewAuthController()

	authGroup := router.Group("/api/v1/auth")
	{
		// Public routes
		authGroup.POST("/register", c.Register)
		authGroup.POST("/login", c.Login)
		authGroup.POST("/refresh", c.RefreshToken)
		authGroup.POST("/forgot-password", c.ForgotPassword)
		authGroup.POST("/reset-password", c.ResetPassword)

		// Protected routes
		protected := authGroup.Use(middleware.AuthMiddleware())
		protected.POST("/logout", c.Logout)
		protected.POST("/logout-all", c.LogoutAllDevices)
		protected.GET("/sessions", c.GetUserSessions)
		authGroup.POST("/send-verification-code", c.SendVerificationCode)
		authGroup.POST("/verify-phone", c.VerifyPhone)
	}
}

// Public routes

// @Summary     Register a new user
// @Description Register a new user with phone number and password
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.RegisterRequest true "User registration details"
// @Success     200 {object} dto.TokenResponse "User registered successfully"
// @Success     201 {object} errors.CustomError "User registered successfully but failed to generate tokens"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     409 {object} errors.CustomError "Conflict - User already exists"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	response, err := c.authUseCase.Register(ctx, &request)
	if err != nil {
		switch err {
		case errors.ErrBadRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserAlreadyExists:
			ctx.JSON(http.StatusConflict, gin.H{"error": err})
		case errors.TokenGenerationFailed:
			ctx.JSON(http.StatusCreated, gin.H{"message": err})
		case errors.TokenGenerationFailed:
			ctx.JSON(http.StatusCreated, gin.H{"message": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary     User login
// @Description Login a user with phone number and password
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.LoginRequest true "User login details"
// @Success     200 {object} dto.TokenResponse "Login successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     401 {object} errors.CustomError "Unauthorized - Invalid phone number or password"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	response, err := c.authUseCase.Login(ctx, &request)
	if err != nil {
		switch err {
		case errors.ErrBadRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserNotFound:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary     Refresh access token
// @Description Get new access and refresh tokens using a valid refresh token
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.RefreshTokenRequest true "Current refresh token"
// @Success     200 {object} dto.TokenResponse "Refresh token successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     401 {object} errors.CustomError "Unauthorized - Invalid token"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	response, err := c.authUseCase.RefreshToken(ctx, &request)
	if err != nil {
		switch err {
		case errors.ErrInvalidToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		case errors.ErrUserNotFound:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary     Forgot password
// @Description Get verification code to reset password
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.VerifyPhoneRequest true "Forgot password request"
// @Success     200 {object} map[string]string "Verification code sent successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     404 {object} errors.CustomError "Not Found - User not found"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/forgot-password [post]
func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var request dto.VerifyPhoneRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	err := c.authUseCase.SendVerificationCode(ctx, &request)
	if err != nil {
		switch err {
		case errors.ErrBadRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		case errors.ErrInvalidPhoneNumber:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Verification code sent successfully"})
}

// @Summary     Reset password
// @Description Reset password with verification code
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.ResetPasswordRequest true "Reset password request"
// @Success     200 {object} dto.TokenResponse "Password reset successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     401 {object} errors.CustomError "Unauthorized - Invalid verification code"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/reset-password [post]
func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var request dto.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	response, err := c.authUseCase.ResetPassword(ctx, &request)
	if err != nil {
		switch err {
		case errors.ErrBadRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserNotActive:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		case errors.ErrInvalidVerificationCode:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Protected routes

// @Summary     User logout
// @Description Logout a user by invalidating the refresh token
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.LogoutRequest true "Refresh token to invalidate"
// @Success     200 {object} map[string]string "Logout successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     401 {object} errors.CustomError "Unauthorized - Invalid token"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/logout [post]
func (c *AuthController) Logout(ctx *gin.Context) {
	var request dto.LogoutRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	userID := ctx.GetString("user_id")
	if userID == "" {
		logger.Error(nil, "User ID not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	if err := c.authUseCase.Logout(ctx, &request, userID); err != nil {
		switch err {
		case errors.ErrInvalidToken:
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}

// @Summary     Logout from all devices
// @Description Logs out the user from all devices by invalidating all refresh tokens
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} map[string]string "Logout from all devices successfully"
// @Failure     401 {object} errors.CustomError "Unauthorized - Invalid token"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/logout-all [post]
func (c *AuthController) LogoutAllDevices(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		logger.Error(nil, "User ID not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	err := c.authUseCase.LogoutAllDevices(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout from all devices successfully"})
}

// @Summary     Get user sessions
// @Description Returns all active sessions for the authenticated user
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.GetSessionsResponse "Get user sessions successfully"
// @Failure     401 {object} errors.CustomError "Unauthorized - Invalid token"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/sessions [get]
func (c *AuthController) GetUserSessions(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		logger.Error(nil, "User ID not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	sessionResponses, err := c.authUseCase.GetUserSessions(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, sessionResponses)
}

// @Summary     Send verification code
// @Description Send verification code to user's phone number
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} map[string]string "Send verification code successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     404 {object} errors.CustomError "Not Found - User not found"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/send-verification-code [post]
func (c *AuthController) SendVerificationCode(ctx *gin.Context) {
	phoneNumber := ctx.GetString("phone_number")
	if phoneNumber == "" {
		logger.Error(nil, "Phone number not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}
	var verifyPhoneRequest dto.VerifyPhoneRequest
	verifyPhoneRequest.PhoneNumber = phoneNumber

	err := c.authUseCase.SendVerificationCode(ctx, &verifyPhoneRequest)
	if err != nil {
		switch err {
		case errors.ErrInvalidPhoneNumber:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "verify code sent successfully"})
}

// @Summary     Active user
// @Description Active user
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.CodeRequest true "Verify code request"
// @Success     200 {object} dto.TokenResponse "Verify code successfully"
// @Failure     400 {object} errors.CustomError "Bad Request - Invalid request body"
// @Failure     404 {object} errors.CustomError "Not Found - Verification code not found"
// @Failure     500 {object} errors.CustomError "Internal Server Error"
// @Router      /auth/verify-phone [post]
func (c *AuthController) VerifyPhone(ctx *gin.Context) {
	phoneNumber := ctx.GetString("phone_number")
	if phoneNumber == "" {
		logger.Error(nil, "Phone number not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	var request dto.CodeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}

	var verifyCodeRequest dto.VerifyCodeRequest
	verifyCodeRequest.PhoneNumber = phoneNumber
	verifyCodeRequest.Code = request.Code

	response, err := c.authUseCase.VerifyPhone(ctx, &verifyCodeRequest)
	if err != nil {
		switch err {
		case errors.ErrInvalidVerificationCode:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrVerificationCodeNotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}
