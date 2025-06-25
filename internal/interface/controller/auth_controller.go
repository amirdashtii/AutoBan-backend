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

	authGroup := router.Group("/auth")
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

// @Summary     Register a new user
// @Description Register a new user with phone number and password
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.RegisterRequest true "User registration details"
// @Success     201 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var request dto.RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	if err := c.authUseCase.Register(ctx, &request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// @Summary     User login
// @Description Login a user with phone number and password
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.LoginRequest true "User login details"
// @Success     200 {object} dto.TokenResponse
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request dto.LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	response, err := c.authUseCase.Login(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusAccepted, response)
}

// @Summary     Get user sessions
// @Description Returns all active sessions for the authenticated user
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.GetSessionsResponse
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
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

// @Summary     User logout
// @Description Logout a user by invalidating the refresh token
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.LogoutRequest true "Refresh token to invalidate"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /auth/logout [post]
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

	if err := c.authUseCase.Logout(ctx, &request, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
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
// @Success     200 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
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

// @Summary     Refresh access token
// @Description Get new access and refresh tokens using a valid refresh token
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Param       request body dto.RefreshTokenRequest true "Current refresh token"
// @Success     200 {object} dto.TokenResponse
// @Failure     400 {object} map[string]string
// @Failure     401 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /auth/refresh-token [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	var request dto.RefreshTokenRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	response, err := c.authUseCase.RefreshToken(ctx, &request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
