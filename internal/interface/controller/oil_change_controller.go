package controller

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/middleware"
	"github.com/amirdashtii/AutoBan/internal/usecase"
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
	oilChangeGroup := router.Group("/api/v1/user/vehicles/:vehicle_id/oil-changes")
	oilChangeGroup.Use(middleware.AuthMiddleware())
	oilChangeGroup.Use(middleware.RequireActiveUser())
	{
		oilChangeGroup.GET("", c.ListOilChanges)
		oilChangeGroup.GET("/last", c.GetLastOilChange)
		oilChangeGroup.GET("/:oil_change_id", c.GetOilChange)
	}

}

// ListOilChanges godoc
// @Summary List all oil changes
// @Description Get a list of all oil changes for a user vehicle
// @Tags Oil Changes
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Param vehicle_id path int true "Vehicle ID"
// @Success 200 {object} dto.ListOilChangesResponse
// @Failure 400 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/oil-changes [get]
func (c *OilChangeController) ListOilChanges(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")

	response, err := c.oilChangeUseCase.ListOilChanges(ctx, userID, vehicleID)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetLastOilChange godoc
// @Summary Last oil change
// @Description Get the last oil change for a user vehicle
// @Tags Oil Changes
// @Produce json
// @Security    BearerAuth
// @Param vehicle_id path int true "Vehicle ID"
// @Success 200 {object} dto.OilChangeResponse
// @Failure 400 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/oil-changes/last [get]
func (c *OilChangeController) GetLastOilChange(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")

	response, err := c.oilChangeUseCase.GetLastOilChange(ctx, userID, vehicleID)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// GetOilChange godoc
// @Summary Get oil change information
// @Description Get information about a specific oil change
// @Tags Oil Changes
// @Accept json
// @Produce json
// @Security    BearerAuth
// @Param vehicle_id path int true "Vehicle ID"
// @Param oil_change_id path int true "Oil change ID"
// @Success 200 {object} dto.OilChangeResponse
// @Failure 400 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/oil-changes/{oil_change_id} [get]
func (c *OilChangeController) GetOilChange(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")
	oilChangeID := ctx.Param("oil_change_id")

	response, err := c.oilChangeUseCase.GetOilChange(ctx, userID, vehicleID, oilChangeID)
	if err != nil {
		respondError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
