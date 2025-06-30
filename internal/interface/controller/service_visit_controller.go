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
	return &ServiceVisitController{
		serviceVisitUseCase: usecase.NewServiceVisitUseCase(),
	}
}

func ServiceVisitRoutes(router *gin.Engine) {
	c := NewServiceVisitController()

	// Service visit management (requires authentication)
	serviceVisitGroup := router.Group("/api/v1/service-visits")
	serviceVisitGroup.Use(middleware.AuthMiddleware())
	{
		serviceVisitGroup.POST("", c.CreateServiceVisit)
		serviceVisitGroup.GET("/:id", c.GetServiceVisit)
		serviceVisitGroup.PUT("/:id", c.UpdateServiceVisit)
		serviceVisitGroup.DELETE("/:id", c.DeleteServiceVisit)
	}

	// User vehicle specific service visit routes
	userVehicleGroup := router.Group("/api/v1/user-vehicles")
	userVehicleGroup.Use(middleware.AuthMiddleware())
	{
		userVehicleGroup.GET("/:user_vehicle_id/service-visits", c.ListServiceVisits)
		userVehicleGroup.GET("/:user_vehicle_id/service-visits/last", c.GetLastServiceVisit)
	}
}

// CreateServiceVisit godoc
// @Summary Create a new service visit
// @Description Create a new service visit record for a user vehicle with optional services
// @Tags Service Visits
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param service_visit body dto.CreateServiceVisitRequest true "Service visit data"
// @Success 201 {object} dto.ServiceVisitResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service-visits [post]
func (c *ServiceVisitController) CreateServiceVisit(ctx *gin.Context) {
	var request dto.CreateServiceVisitRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.serviceVisitUseCase.CreateServiceVisit(ctx, request)
	if err != nil {
		switch err {
		case errors.ErrInvalidServiceVisitCreateRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrInvalidDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrFailedToCreateOilChange:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		case errors.ErrFailedToCreateOilFilter:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
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
// @Param id path string true "Service visit ID"
// @Success 200 {object} dto.ServiceVisitResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service-visits/{id} [get]
func (c *ServiceVisitController) GetServiceVisit(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidServiceVisitID})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.serviceVisitUseCase.GetServiceVisit(ctx, id)
	if err != nil {
		switch err {
		case errors.ErrInvalidServiceVisitID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToGetServiceVisit:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
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
// @Param user_vehicle_id path string true "User vehicle ID"
// @Success 200 {object} dto.ListServiceVisitsResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user-vehicles/{user_vehicle_id}/service-visits [get]
func (c *ServiceVisitController) ListServiceVisits(ctx *gin.Context) {
	userVehicleID := ctx.Param("user_vehicle_id")
	if userVehicleID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserVehicleIDRequired})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.serviceVisitUseCase.ListServiceVisits(ctx, userVehicleID)
	if err != nil {
		switch err {
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToListServiceVisits:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
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
// @Param id path string true "Service visit ID"
// @Param service_visit body dto.UpdateServiceVisitRequest true "Updated service visit data"
// @Success 200 {object} dto.ServiceVisitResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service-visits/{id} [put]
func (c *ServiceVisitController) UpdateServiceVisit(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidServiceVisitID})
		return
	}

	var request dto.UpdateServiceVisitRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind JSON")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.serviceVisitUseCase.UpdateServiceVisit(ctx, id, request)
	if err != nil {
		switch err {
		case errors.ErrInvalidServiceVisitID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrInvalidServiceVisitUpdateRequest:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToGetServiceVisit:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		case errors.ErrInvalidDate:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrFailedToUpdateOilChange:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		case errors.ErrFailedToUpdateOilFilter:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
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
// @Param id path string true "Service visit ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /service-visits/{id} [delete]
func (c *ServiceVisitController) DeleteServiceVisit(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidServiceVisitID})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	err := c.serviceVisitUseCase.DeleteServiceVisit(ctx, id)
	if err != nil {
		switch err {
		case errors.ErrInvalidServiceVisitID:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToDeleteServiceVisit:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
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
// @Param user_vehicle_id path string true "User vehicle ID"
// @Success 200 {object} dto.ServiceVisitResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user-vehicles/{user_vehicle_id}/service-visits/last [get]
func (c *ServiceVisitController) GetLastServiceVisit(ctx *gin.Context) {
	userVehicleID := ctx.Param("user_vehicle_id")
	if userVehicleID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserVehicleIDRequired})
		return
	}

	// Add user_id to context
	userID := ctx.GetString("user_id")
	ctx.Set("user_id", userID)

	response, err := c.serviceVisitUseCase.GetLastServiceVisit(ctx, userVehicleID)
	if err != nil {
		switch err {
		case errors.ErrUserVehicleNotOwned:
			ctx.JSON(http.StatusForbidden, gin.H{"error": err})
		case errors.ErrFailedToGetServiceVisit:
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.ErrInternalServerError})
		}
		return
	}

	ctx.JSON(http.StatusOK, response)
}