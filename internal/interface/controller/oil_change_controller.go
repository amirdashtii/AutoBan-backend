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

type OilChangeController struct {
	oilChangeUseCase usecase.OilChangeUseCase
}

func NewOilChangeController() *OilChangeController {
	oilChangeUseCase := usecase.NewOilChangeUseCase()
	return &OilChangeController{oilChangeUseCase: oilChangeUseCase}
}

func OilChangeRoutes(router *gin.Engine) {
	c := NewOilChangeController()

	// Oil change management (requires authentication)
	oilChangeGroup := router.Group("/api/v1/oil-changes")
	oilChangeGroup.Use(middleware.AuthMiddleware())
	{
		oilChangeGroup.POST("", c.CreateOilChange)
		oilChangeGroup.GET("/:id", c.GetOilChange)
		oilChangeGroup.PUT("/:id", c.UpdateOilChange)
		oilChangeGroup.DELETE("/:id", c.DeleteOilChange)
		oilChangeGroup.GET("/list/:user_vehicle_id", c.ListOilChanges)
		oilChangeGroup.GET("/last/:user_vehicle_id", c.GetLastOilChange)
	}

}

// CreateOilChange godoc
// @Summary Create a new oil change record for a user vehicle
// @Description Create a new oil change record for a user vehicle
// @Tags Oil Changes
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Param request body dto.CreateOilChangeRequest true "Oil change information"
// @Success 201 {object} dto.OilChangeResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-changes [post]
func (c *OilChangeController) CreateOilChange(ctx *gin.Context) {
	var request dto.CreateOilChangeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody.Error()})
		return
	}

	response, err := c.oilChangeUseCase.CreateOilChange(ctx, request)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilChangeCreateRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilChangeCreateRequest})
		case errors.ErrInvalidDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidDate})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToCreateOilChange})
		}
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// GetOilChange godoc
// @Summary Get oil change information
// @Description Get information about a specific oil change
// @Tags Oil Changes
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Param id path int true "Oil change ID"
// @Success 200 {object} dto.OilChangeResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-changes/{id} [get]
func (c *OilChangeController) GetOilChange(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := c.oilChangeUseCase.GetOilChange(ctx, id)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilChangeID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilChangeID})
		case errors.ErrFailedToGetOilChange:
			ctx.JSON(http.StatusNotFound, gin.H{"error": errors.ErrFailedToGetOilChange})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToGetOilChange})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ListOilChanges godoc
// @Summary List all oil changes
// @Description Get a list of all oil changes for a user vehicle
// @Tags Oil Changes
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Param user_vehicle_id path int true "User vehicle ID"
// @Success 200 {object} dto.ListOilChangesResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-changes/list/{user_vehicle_id} [get]
func (c *OilChangeController) ListOilChanges(ctx *gin.Context) {
	userVehicleID := ctx.Param("user_vehicle_id")
	if userVehicleID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserVehicleIDRequired})
		return
	}

	response, err := c.oilChangeUseCase.ListOilChanges(ctx, userVehicleID)
	if err != nil {
		switch err {
		case errors.ErrFailedToListOilChanges:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToListOilChanges})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateOilChange godoc
// @Summary Update oil change
// @Description Update the information of an oil change
// @Tags Oil Changes
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Param id path int true "Oil change ID"
// @Param request body dto.UpdateOilChangeRequest true "Update oil change information"
// @Success 200 {object} dto.OilChangeResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-changes/{id} [put]
func (c *OilChangeController) UpdateOilChange(ctx *gin.Context) {
	id := ctx.Param("id")

	var request dto.UpdateOilChangeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	response, err := c.oilChangeUseCase.UpdateOilChange(ctx, id, request)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilChangeUpdateRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilChangeUpdateRequest})
		case errors.ErrInvalidOilChangeID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilChangeID})
		case errors.ErrFailedToGetOilChange:
			ctx.JSON(http.StatusNotFound, gin.H{"error": errors.ErrFailedToGetOilChange})
		case errors.ErrInvalidDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidDate})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToUpdateOilChange})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// DeleteOilChange godoc
// @Summary Delete oil change
// @Description Delete an oil change record
// @Tags Oil Changes
// @Produce json
// @Security    BearerAuth
// @Param id path int true "Oil change ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-changes/{id} [delete]
func (c *OilChangeController) DeleteOilChange(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.oilChangeUseCase.DeleteOilChange(ctx, id)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilChangeID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilChangeID})
		case errors.ErrFailedToDeleteOilChange:
			ctx.JSON(http.StatusNotFound, gin.H{"error": errors.ErrFailedToDeleteOilChange})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToDeleteOilChange})
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetLastOilChange godoc
// @Summary Last oil change
// @Description Get the last oil change for a user vehicle
// @Tags Oil Changes
// @Produce json
// @Security    BearerAuth
// @Param user_vehicle_id path int true "User vehicle ID"
// @Success 200 {object} dto.OilChangeResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-changes/last/{user_vehicle_id} [get]
func (c *OilChangeController) GetLastOilChange(ctx *gin.Context) {
	userVehicleID := ctx.Param("user_vehicle_id")
	if userVehicleID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserVehicleIDRequired})
		return
	}

	response, err := c.oilChangeUseCase.GetLastOilChange(ctx, userVehicleID)
	if err != nil {
		switch err {
		case errors.ErrFailedToGetOilChange:
			ctx.JSON(http.StatusNotFound, gin.H{"error": errors.ErrFailedToGetOilChange})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrFailedToGetOilChange})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}
