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

type ServiceVisitController struct {
	serviceVisitUseCase usecase.ServiceVisitUseCase
}

func NewServiceVisitController() *ServiceVisitController {
	serviceVisitUseCase := usecase.NewServiceVisitUseCase()
	return &ServiceVisitController{serviceVisitUseCase: serviceVisitUseCase}
}

func ServiceVisitRoutes(router *gin.Engine) {
	c := NewServiceVisitController()
	userVehicleGroup := router.Group("/api/v1/user/vehicles/:vehicle_id/service-visits")
	userVehicleGroup.Use(middleware.AuthMiddleware())
	userVehicleGroup.Use(middleware.RequireActiveUser())
	{
		userVehicleGroup.POST("", c.CreateServiceVisit)
		userVehicleGroup.GET("", c.ListServiceVisits)
		userVehicleGroup.GET("/last", c.GetLastServiceVisit)
		userVehicleGroup.GET("/:visit_id", c.GetServiceVisit)
		userVehicleGroup.PUT("/:visit_id", c.UpdateServiceVisit)
		userVehicleGroup.DELETE("/:visit_id", c.DeleteServiceVisit)
	}
}

// CreateServiceVisit godoc
// @Summary Create a new service visit
// @Description Create a new service visit record for a user vehicle with optional services
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Param service_visit body dto.CreateServiceVisitRequest true "Service visit data"
// @Success 201 {object} dto.ServiceVisitResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/service-visits [post]
func (c *ServiceVisitController) CreateServiceVisit(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicle_id")
	userID := ctx.GetString("user_id")

	var request dto.CreateServiceVisitRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		respondError(ctx, errors.ErrBadRequest)
		return
	}

	// اعتبارسنجی مالکیت و وجود
	response, err := c.serviceVisitUseCase.CreateServiceVisit(ctx, userID, vehicleID, request)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// GetServiceVisit godoc
// @Summary Get service visit by ID
// @Description Get a specific service visit record by its ID
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Param visit_id path string true "Service visit ID"
// @Success 200 {object} dto.ServiceVisitResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/service-visits/{visit_id} [get]
func (c *ServiceVisitController) GetServiceVisit(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicle_id")
	visitID := ctx.Param("visit_id")
	userID := ctx.GetString("user_id")

	response, err := c.serviceVisitUseCase.GetServiceVisit(ctx, userID, vehicleID, visitID)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// ListServiceVisits godoc
// @Summary List service visits for a user vehicle
// @Description Get all service visit records for a specific user vehicle
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Success 200 {object} dto.ListServiceVisitsResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/service-visits [get]
func (c *ServiceVisitController) ListServiceVisits(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicle_id")
	userID := ctx.GetString("user_id")

	response, err := c.serviceVisitUseCase.ListServiceVisits(ctx, userID, vehicleID)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// UpdateServiceVisit godoc
// @Summary Update service visit
// @Description Update an existing service visit record
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Param visit_id path string true "Service visit ID"
// @Param service_visit body dto.UpdateServiceVisitRequest true "Updated service visit data"
// @Success 200 {object} dto.ServiceVisitResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/service-visits/{visit_id} [put]
func (c *ServiceVisitController) UpdateServiceVisit(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicle_id")
	visitID := ctx.Param("visit_id")
	userID := ctx.GetString("user_id")

	var request dto.UpdateServiceVisitRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		respondError(ctx, errors.ErrBadRequest)
		return
	}

	response, err := c.serviceVisitUseCase.UpdateServiceVisit(ctx, userID, vehicleID, visitID, request)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// DeleteServiceVisit godoc
// @Summary Delete service visit
// @Description Delete a service visit record
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Param visit_id path string true "Service visit ID"
// @Success 204 "No Content"
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/service-visits/{visit_id} [delete]
func (c *ServiceVisitController) DeleteServiceVisit(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicle_id")
	visitID := ctx.Param("visit_id")
	userID := ctx.GetString("user_id")

	err := c.serviceVisitUseCase.DeleteServiceVisit(ctx, userID, vehicleID, visitID)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.Status(http.StatusNoContent)
}

// GetLastServiceVisit godoc
// @Summary Get last service visit for a user vehicle
// @Description Get the most recent service visit record for a specific user vehicle
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param vehicle_id path string true "Vehicle ID"
// @Success 200 {object} dto.ServiceVisitResponse
// @Failure 400 {object} errors.CustomError
// @Failure 401 {object} errors.CustomError
// @Failure 403 {object} errors.CustomError
// @Failure 404 {object} errors.CustomError
// @Failure 500 {object} errors.CustomError
// @Router /user/vehicles/{vehicle_id}/service-visits/last [get]
func (c *ServiceVisitController) GetLastServiceVisit(ctx *gin.Context) {
	vehicleID := ctx.Param("vehicle_id")
	userID := ctx.GetString("user_id")

	response, err := c.serviceVisitUseCase.GetLastServiceVisit(ctx, userID, vehicleID)
	if err != nil {
		respondError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
