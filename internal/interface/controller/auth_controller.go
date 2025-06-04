package controller

import (
	"net/http"
	"time"

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
		authGroup.POST("/refresh-token", c.RefreshToken)

		// Protected routes
		protected := authGroup.Use(middleware.AuthMiddleware())
		protected.POST("/logout", c.Logout)
		protected.GET("/sessions", c.GetUserSessions)
		protected.POST("/logout-all", c.LogoutAllDevices)
	}
}

// @title           AutoBan API
// @version         1.0
// @description     Intelligent Vehicle Maintenance Management System API
// @description     A comprehensive solution for managing vehicle repairs, periodic services,
// @description     maintenance costs, and monitoring vehicle conditions.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Amir Dashti
// @contact.url    https://github.com/amirdashtii
// @contact.email  AhDashti@gmail.com

// @license.name  GNU General Public License v3.0
// @license.url   https://www.gnu.org/licenses/gpl-3.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token for authentication

// @tag.name     auth-public
// @tag.description Public authentication endpoints - no token required
// @tag.x-order  1

// @tag.name     auth-protected
// @tag.description Protected authentication endpoints - valid token required
// @tag.x-order  2

// @Summary     Register a new user
// @Description Register a new user with phone number and password
// @Tags        auth-public
// @Accept      json
// @Produce     json
// @Param       request body dto.RegisterRequest true "User registration details"
// @Success     201 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	if err := c.authUseCase.Register(&request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// @Summary     User login
// @Description Login a user with phone number and password
// @Tags        auth-public
// @Accept      json
// @Produce     json
// @Param       request body dto.LoginRequest true "User login details"
// @Success     200 {object} dto.TokenResponse
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	response, err := c.authUseCase.Login(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, response)
}

// @Summary     Get user sessions
// @Description Returns all active sessions for the authenticated user
// @Tags        auth-protected
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.GetSessionsResponse
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/auth/sessions [get]
func (c *AuthController) GetUserSessions(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		logger.Error(nil, "User ID not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	sessions, err := c.authUseCase.GetUserSessions(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
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

	ctx.JSON(http.StatusOK, dto.GetSessionsResponse{
		Sessions: sessionResponses,
	})
}

// @Summary     User logout
// @Description Logout a user by invalidating the refresh token
// @Tags        auth-protected
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.LogoutRequest true "Refresh token to invalidate"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/auth/logout [post]
func (c *AuthController) Logout(ctx *gin.Context) {
	var request dto.LogoutRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	userID := ctx.GetString("user_id")
	if userID == "" {
		logger.Error(nil, "User ID not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	if err := c.authUseCase.Logout(&request, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}

// @Summary     Logout from all devices
// @Description Logs out the user from all devices by invalidating all refresh tokens
// @Tags        auth-protected
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/auth/logout-all [post]
func (c *AuthController) LogoutAllDevices(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	if userID == "" {
		logger.Error(nil, "User ID not found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound})
		return
	}

	err := c.authUseCase.LogoutAllDevices(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "logout from all devices successfully"})
}

// @Summary     Refresh access token
// @Description Get new access and refresh tokens using a valid refresh token
// @Tags        auth-public
// @Accept      json
// @Produce     json
// @Param       request body dto.RefreshTokenRequest true "Current refresh token"
// @Success     200 {object} dto.TokenResponse
// @Failure     400 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/v1/auth/refresh-token [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	response, err := c.authUseCase.RefreshToken(&request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
