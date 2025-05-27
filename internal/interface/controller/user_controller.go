package controller

import (
	"net/http"

	"AutoBan/internal/dto"
	"AutoBan/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

// UserController struct
// ساختار UserController

type UserController struct {
	userUseCase user.UserUseCase
}

// NewUserController creates a new UserController
// تابع NewUserController یک UserController جدید ایجاد می‌کند

func NewUserController(uuc user.UserUseCase) *UserController {
	return &UserController{userUseCase: uuc}
}

// Register handles user registration
// تابع Register ثبت‌نام کاربر را مدیریت می‌کند

func (uc *UserController) Register(c *gin.Context) {
	var userDTO dto.UserRegisterDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUseCase.Register(userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Login handles user login
// تابع Login ورود کاربر را مدیریت می‌کند

func (uc *UserController) Login(c *gin.Context) {
	var userDTO dto.UserLoginDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.userUseCase.Login(userDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
