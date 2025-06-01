package controller

import (
	"net/http"

	"AutoBan/internal/dto"
	"AutoBan/internal/errors"
	"AutoBan/internal/usecase"
	"AutoBan/pkg/logger"

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
	authGroup.POST("/register", c.Register)
	authGroup.POST("/login", c.Login)
	authGroup.POST("/logout", c.Logout)
	authGroup.POST("/refresh-token", c.RefreshToken)
}

// @Summary     Register a new user
// @Description Register a new user with phone number and password
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body dto.RegisterRequest true "User registration details"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]error
// @Failure     500 {object} map[string]error
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

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// @Summary     User login
// @Description Login a user with phone number and password
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body dto.LoginRequest true "User login details"
// @Success     200 {object} dto.LoginResponse
// @Failure     400 {object} map[string]error
// @Failure     500 {object} map[string]error
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

	ctx.JSON(http.StatusOK, response)
}

// @Summary     User logout
// @Description Logout a user using refresh token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body dto.LogoutRequest true "User logout details"
// @Success     200 {object} map[string]string
// @Failure     400 {object} map[string]error
// @Failure     500 {object} map[string]error
// @Router      /api/v1/auth/logout [post]
func (c *AuthController) Logout(ctx *gin.Context) {
	var request dto.LogoutRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	if err := c.authUseCase.Logout(&request); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// @Summary     Refresh access token
// @Description Get new access token using refresh token
// @Tags        auth
// @Accept      json
// @Produce     json
// @Param       request body dto.RefreshTokenRequest true "Refresh token details"
// @Success     200 {object} dto.RefreshTokenResponse
// @Failure     400 {object} map[string]error
// @Failure     500 {object} map[string]error
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
