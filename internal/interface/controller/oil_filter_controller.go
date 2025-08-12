package controller

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/middleware"
	"github.com/amirdashtii/AutoBan/internal/usecase"
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
	oilFilterGroup := router.Group("/api/v1/user/vehicles/:vehicle_id/oil-filters")
	oilFilterGroup.Use(middleware.AuthMiddleware())
	oilFilterGroup.Use(middleware.RequireActiveUser())
	{
		oilFilterGroup.GET("", c.ListOilFilters)
		oilFilterGroup.GET("/last", c.GetLastOilFilter)
		oilFilterGroup.GET("/:oil_filter_id", c.GetOilFilter)
	}
}

// ListOilFilters godoc
// @Summary List oil filter changes for a user vehicle
// @Description Get all oil filter change records for a specific user vehicle
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Success 200 {object} dto.ListOilFiltersResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/oil-filters [get]
func (c *OilFilterController) ListOilFilters(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")

	response, err := c.oilFilterUseCase.ListOilFilters(ctx, userID, vehicleID)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetLastOilFilter godoc
// @Summary Get last oil filter change for a user vehicle
// @Description Get the most recent oil filter change record for a specific user vehicle
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Success 200 {object} dto.OilFilterResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/oil-filters/last [get]
func (c *OilFilterController) GetLastOilFilter(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")

	response, err := c.oilFilterUseCase.GetLastOilFilter(ctx, userID, vehicleID)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetOilFilter godoc
// @Summary Get oil filter change by ID
// @Description Get a specific oil filter change record by its ID
// @Tags Oil Filters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Param oil_filter_id path string true "Oil filter change ID"
// @Success 200 {object} dto.OilFilterResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/oil-filters/{oil_filter_id} [get]
func (c *OilFilterController) GetOilFilter(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")
	oilFilterID := ctx.Param("oil_filter_id")

	response, err := c.oilFilterUseCase.GetOilFilter(ctx, userID, vehicleID, oilFilterID)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
