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

type OilFilterController struct {
	oilFilterUseCase usecase.OilFilterUseCase
}

func NewOilFilterController() *OilFilterController {
	return &OilFilterController{
		oilFilterUseCase: usecase.NewOilFilterUseCase(),
	}
}

func OilFilterRoutes(router *gin.Engine) {
	c := NewOilFilterController()

	// User vehicle specific oil filter routes
	oilFilterGroup := router.Group("/api/v1/oil-filters")
	oilFilterGroup.Use(middleware.AuthMiddleware())
	{
		oilFilterGroup.POST("", c.CreateOilFilter)
		oilFilterGroup.GET("/:id", c.GetOilFilter)
		oilFilterGroup.PUT("/:id", c.UpdateOilFilter)
		oilFilterGroup.DELETE("/:id", c.DeleteOilFilter)
		oilFilterGroup.GET("/list/:user_vehicle_id", c.ListOilFilters)
		oilFilterGroup.GET("/last/:user_vehicle_id", c.GetLastOilFilter)
	}
}

// CreateOilFilter godoc
// @Summary Create a new oil filter change
// @Description Create a new oil filter change record for a user vehicle
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param oil_filter body dto.CreateOilFilterRequest true "Oil filter change data"
// @Success 201 {object} dto.OilFilterResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-filters [post]
func (c *OilFilterController) CreateOilFilter(ctx *gin.Context) {
	var request dto.CreateOilFilterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.oilFilterUseCase.CreateOilFilter(ctx, request)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilFilterCreateRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrInvalidDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// GetOilFilter godoc
// @Summary Get oil filter change by ID
// @Description Get a specific oil filter change record by its ID
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Oil filter change ID"
// @Success 200 {object} dto.OilFilterResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-filters/{id} [get]
func (c *OilFilterController) GetOilFilter(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilFilterID})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.oilFilterUseCase.GetOilFilter(ctx, id)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilFilterID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToGetOilFilter:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// ListOilFilters godoc
// @Summary List oil filter changes for a user vehicle
// @Description Get all oil filter change records for a specific user vehicle
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_vehicle_id path string true "User vehicle ID"
// @Success 200 {object} dto.ListOilFiltersResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-filters/list/{user_vehicle_id} [get]
func (c *OilFilterController) ListOilFilters(ctx *gin.Context) {
	userVehicleID := ctx.Param("user_vehicle_id")
	if userVehicleID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserVehicleIDRequired})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.oilFilterUseCase.ListOilFilters(ctx, userVehicleID)
	if err != nil {
		switch err {
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToListOilFilters:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// UpdateOilFilter godoc
// @Summary Update oil filter change
// @Description Update an existing oil filter change record
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Oil filter change ID"
// @Param oil_filter body dto.UpdateOilFilterRequest true "Updated oil filter change data"
// @Success 200 {object} dto.OilFilterResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-filters/{id} [put]
func (c *OilFilterController) UpdateOilFilter(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilFilterID})
		return
	}

	var request dto.UpdateOilFilterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.oilFilterUseCase.UpdateOilFilter(ctx, id, request)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilFilterID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrInvalidOilFilterUpdateRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToGetOilFilter:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		case errors.ErrInvalidDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// DeleteOilFilter godoc
// @Summary Delete oil filter change
// @Description Delete an oil filter change record
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Oil filter change ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-filters/{id} [delete]
func (c *OilFilterController) DeleteOilFilter(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidOilFilterID})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	err := c.oilFilterUseCase.DeleteOilFilter(ctx, id)
	if err != nil {
		switch err {
		case errors.ErrInvalidOilFilterID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToDeleteOilFilter:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.Status(http.StatusNoContent)
}

// GetLastOilFilter godoc
// @Summary Get last oil filter change for a user vehicle
// @Description Get the most recent oil filter change record for a specific user vehicle
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_vehicle_id path string true "User vehicle ID"
// @Success 200 {object} dto.OilFilterResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /oil-filters/last/{user_vehicle_id} [get]
func (c *OilFilterController) GetLastOilFilter(ctx *gin.Context) {
	userVehicleID := ctx.Param("user_vehicle_id")
	if userVehicleID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserVehicleIDRequired})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.oilFilterUseCase.GetLastOilFilter(ctx, userVehicleID)
	if err != nil {
		switch err {
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToGetOilFilter:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}
