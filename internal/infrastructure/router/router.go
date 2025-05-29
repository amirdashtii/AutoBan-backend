package router

import (
	"AutoBan/internal/dto"
	"AutoBan/internal/errors"
	"AutoBan/internal/usecase/auth"
	"AutoBan/internal/usecase/user"
	"AutoBan/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	userUseCase user.UserUseCase
	authUseCase auth.AuthUseCase
}

// NewRouter creates a new Router instance
func NewRouter(userUseCase user.UserUseCase, authUseCase auth.AuthUseCase) *Router {
	return &Router{userUseCase: userUseCase, authUseCase: authUseCase}
}

// SetupRouter initializes the Gin router and sets up the API endpoints
func (r *Router) SetupRouter() *gin.Engine {
	router := gin.Default()

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", r.register)
		authGroup.POST("/login", r.login)
	}

	return router
}

func (r *Router) register(c *gin.Context) {
	var userDTO dto.UserRegisterDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		logger.Error(err, "Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	user, err := r.userUseCase.Register(userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// login handles user login and returns access and refresh tokens
// تابع login ورود کاربر را مدیریت کرده و اکسس و رفرش توکن‌ها را برمی‌گرداند

func (r *Router) login(c *gin.Context) {
	var userDTO dto.UserLoginDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		logger.Error(err, "Invalid request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	user, err := r.userUseCase.Login(userDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := r.authUseCase.GenerateAccessToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	refreshToken, err := r.authUseCase.GenerateRefreshToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
