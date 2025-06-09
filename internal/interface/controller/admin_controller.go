package controller

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/middleware"
	"github.com/amirdashtii/AutoBan/internal/usecase"
	"github.com/gin-gonic/gin"
)

// @tag.name     admin
// @tag.description Protected admin endpoints - requires admin privileges
// @tag.x-order  4

type AdminController struct {
	adminUseCase usecase.AdminUseCase
}

func NewAdminController() *AdminController {
	adminUseCase := usecase.NewAdminUseCase()
	return &AdminController{adminUseCase: adminUseCase}
}

func AdminRoutes(router *gin.Engine) {
	c := NewAdminController()

	adminGroup := router.Group("/api/v1/users")
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
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.ListUsersResponse
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users [get]
func (c *AdminController) ListUsers(ctx *gin.Context) {

	users, err := c.adminUseCase.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @Summary     Get user details
// @Description Get detailed information about a specific user
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Success     200 {object} dto.User
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     404 {object} map[string]string "User not found"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users/{id} [get]
func (c *AdminController) GetUserById(ctx *gin.Context) {
	userID := ctx.Param("id")
	user, err := c.adminUseCase.GetUserById(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// @Summary     Update user
// @Description Update user information
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.UpdateUserRequest true "User update information"
// @Success     200 {object} map[string]string "User updated successfully"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     404 {object} map[string]string "User not found"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users/{id} [put]
func (c *AdminController) UpdateUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := c.adminUseCase.UpdateUser(ctx, userID, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// @Summary     Change user role
// @Description Change the role of a user
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.ChangeUserRoleRequest true "Change user role request"
// @Success     200 {object} map[string]string "User role changed successfully"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     404 {object} map[string]string "User not found"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users/{id}/role [post]
func (c *AdminController) ChangeUserRole(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.ChangeUserRoleRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := c.adminUseCase.ChangeUserRole(ctx, userID, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User role changed successfully"})
}

// @Summary     Change user status
// @Description Change the status of a user
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.ChangeUserStatusRequest true "Change user status request"
// @Success     200 {object} map[string]string "User status changed successfully"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     404 {object} map[string]string "User not found"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users/{id}/status [post]
func (c *AdminController) ChangeUserStatus(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.ChangeUserStatusRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := c.adminUseCase.ChangeUserStatus(ctx, userID, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User status changed successfully"})
}

// @Summary     Change user password
// @Description Change the password of a user
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Param       request body dto.ChangeUserPasswordRequest true "Change user password request"
// @Success     200 {object} map[string]string "User password changed successfully"
// @Failure     400 {object} map[string]string "Bad Request"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     404 {object} map[string]string "User not found"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users/{id}/change-password [post]
func (c *AdminController) ChangeUserPassword(ctx *gin.Context) {
	userID := ctx.Param("id")
	var request dto.ChangeUserPasswordRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := c.adminUseCase.ChangeUserPassword(ctx, userID, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User password changed successfully"})
}

// @Summary     Delete user
// @Description Delete a user from the system
// @Tags        admin
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       id path string true "User ID"
// @Success     200 {object} map[string]string "User deleted successfully"
// @Failure     401 {object} map[string]string "Unauthorized"
// @Failure     403 {object} map[string]string "Forbidden - Admin access required"
// @Failure     404 {object} map[string]string "User not found"
// @Failure     500 {object} map[string]string "Internal Server Error"
// @Router      /api/v1/users/{id} [delete]
func (c *AdminController) DeleteUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	err := c.adminUseCase.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
