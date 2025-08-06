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

// @tag.name     Users
// @tag.description Protected user endpoints - valid token required
// @tag.x-order  3

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController() *UserController {
	userUseCase := usecase.NewUserUseCase()
	return &UserController{userUseCase: userUseCase}
}

func UserRoutes(router *gin.Engine) {
	c := NewUserController()

	userGroup := router.Group("/api/v1/users")
	{
		protected := userGroup.Use(middleware.AuthMiddleware())
		{
			protected.GET("/me", c.GetProfile)
			protected.PUT("/me", c.UpdateProfile)
			protected.PUT("/me/change-password", c.ChangePassword)
			protected.DELETE("/me", c.DeleteUser)
		}
	}
}

// @Summary     Get user profile
// @Description Get the profile information of the authenticated user
// @Tags        Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.GetProfileResponse
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /users/me [get]
func (c *UserController) GetProfile(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	user, err := c.userUseCase.GetProfile(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary     Update user profile
// @Description Update the profile information of the authenticated user
// @Tags        Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.UpdateProfileRequest true "Profile update information"
// @Success     200 {object} dto.UpdateProfileResponse
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /users/me [put]
func (c *UserController) UpdateProfile(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	var request dto.UpdateProfileRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}
	user, err := c.userUseCase.UpdateProfile(ctx, userID, request)
	if err != nil {
		switch err {
		case errors.ErrEmailAlreadyExists:
			ctx.JSON(http.StatusConflict, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary     Update user password
// @Description Update the password of the authenticated user
// @Tags        Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       request body dto.UpdatePasswordRequest true "Password update information"
// @Success     200 {object} map[string]string "Password updated successfully"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /users/me/change-password [put]
func (c *UserController) ChangePassword(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	var request dto.UpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrBadRequest})
		return
	}
	err := c.userUseCase.ChangePassword(ctx, userID, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// @Summary     Delete user
// @Description Delete the authenticated user
// @Tags        Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} map[string]string "User deleted successfully"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /users/me [delete]
func (c *UserController) DeleteUser(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	err := c.userUseCase.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
