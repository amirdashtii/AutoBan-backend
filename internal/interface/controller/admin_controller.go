package controller

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/middleware"
	"github.com/amirdashtii/AutoBan/internal/usecase"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminUseCase usecase.AdminUseCase
}

func NewAdminController() *AdminController {
	adminUseCase := usecase.NewAdminUseCase()
	return &AdminController{adminUseCase: adminUseCase}
}

func AdminRoutes(router *gin.Engine) {
	c := NewAdminController()

	adminGroup := router.Group("/api/v1/admin/users")
	{
		adminGroup.Use(middleware.AuthMiddleware())
		adminGroup.Use(middleware.RequireAdmin())

		adminGroup.GET("", c.ListUsers)
		adminGroup.GET("/:id", c.GetUserById)
		adminGroup.PUT("/:id", c.UpdateUser)
		adminGroup.POST("/:id/role", c.ChangeUserRole)
		adminGroup.POST("/:id/status", c.ChangeUserStatus)
		adminGroup.POST("/:id/change-password", c.ChangeUserPassword)
		adminGroup.DELETE("/:id", c.DeleteUser)
	}
}

// @Summary     List all users
// @Description Get a list of all users with pagination
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.ListUsersResponse
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users [get]
func (c *AdminController) ListUsers(ctx *gin.Context) {

	users, err := c.adminUseCase.ListUsers(ctx)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary     Get user details
// @Description Get detailed information about a specific user
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Success     200 {object} dto.User
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users/{id} [get]
func (c *AdminController) GetUserById(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.adminUseCase.GetUserById(ctx, userID)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Summary     Update user
// @Description Update user information
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.UpdateUserRequest true "User update information"
// @Success     200 {object} map[string]string "User updated successfully"
// @Failure     400 {object} errors.CustomError
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     409 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users/{id} [put]
func (c *AdminController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		respondError(ctx, errors.ErrBadRequest)
		return
	}
	err := c.adminUseCase.UpdateUser(ctx, userID, request)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// @Summary     Change user role
// @Description Change the role of a user
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.ChangeUserRoleRequest true "Change user role request"
// @Success     200 {object} map[string]string "User role changed successfully"
// @Failure     400 {object} errors.CustomError
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users/{id}/role [post]
func (c *AdminController) ChangeUserRole(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.ChangeUserRoleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		respondError(ctx, errors.ErrBadRequest)
		return
	}
	err := c.adminUseCase.ChangeUserRole(ctx, userID, request)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User role changed successfully"})
}

// @Summary     Change user status
// @Description Change the status of a user
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.ChangeUserStatusRequest true "Change user status request"
// @Success     200 {object} map[string]string "User status changed successfully"
// @Failure     400 {object} errors.CustomError
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users/{id}/status [post]
func (c *AdminController) ChangeUserStatus(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.ChangeUserStatusRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		respondError(ctx, errors.ErrBadRequest)
		return
	}
	err := c.adminUseCase.ChangeUserStatus(ctx, userID, request)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User status changed successfully"})
}

// @Summary     Change user password
// @Description Change the password of a user
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.ChangeUserPasswordRequest true "Change user password request"
// @Success     200 {object} map[string]string "User password changed successfully"
// @Failure     400 {object} errors.CustomError
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users/{id}/change-password [post]
func (c *AdminController) ChangeUserPassword(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.ChangeUserPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		respondError(ctx, errors.ErrBadRequest)
		return
	}
	err := c.adminUseCase.ChangeUserPassword(ctx, userID, request)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User password changed successfully"})
}

// @Summary     Delete user
// @Description Delete a user from the system
// @Tags        Admin - Users
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Success     200 {object} map[string]string "User deleted successfully"
// @Failure     401 {object} errors.CustomError
// @Failure     403 {object} errors.CustomError
// @Failure     404 {object} errors.CustomError
// @Failure     500 {object} errors.CustomError
// @Router      /admin/users/{id} [delete]
func (c *AdminController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	err := c.adminUseCase.DeleteUser(ctx, userID)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
