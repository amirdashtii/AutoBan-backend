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

type VehicleController struct {
	vehicleUseCase usecase.VehicleUseCase
}

func NewVehicleController() *VehicleController {
	vehicleUseCase := usecase.NewVehicleUseCase()
	return &VehicleController{vehicleUseCase: vehicleUseCase}
}

func VehicleRoutes(router *gin.Engine) {
	c := NewVehicleController()

	// Public routes for vehicle catalog
	vehicleGroup := router.Group("/api/v1/vehicles")
	{
		// Vehicle Types
		vehicleGroup.GET("/types", c.ListVehicleTypes)
		vehicleGroup.GET("/types/:id", c.GetVehicleType)

		// Brands
		vehicleGroup.GET("/brands", c.ListBrands)
		vehicleGroup.GET("/brands/:id", c.GetBrand)
		vehicleGroup.GET("/types/:id/brands", c.ListBrandsByType)

		// Models
		vehicleGroup.GET("/models", c.ListModels)
		vehicleGroup.GET("/models/:id", c.GetModel)
		vehicleGroup.GET("/brands/:id/models", c.ListModelsByBrand)

		// Generations
		vehicleGroup.GET("/generations", c.ListGenerations)
		vehicleGroup.GET("/generations/:id", c.GetGeneration)
		vehicleGroup.GET("/models/:id/generations", c.ListGenerationsByModel)
	}

	// User vehicle management (requires authentication)
	userVehicles := router.Group("/api/v1/user/vehicles")
	userVehicles.Use(middleware.AuthMiddleware())
	{
		userVehicles.POST("", c.AddUserVehicle)
		userVehicles.GET("", c.ListUserVehicles)
		userVehicles.GET("/:vehicle_id", c.GetUserVehicle)
		userVehicles.PUT("/:vehicle_id", c.UpdateUserVehicle)
		userVehicles.DELETE("/:vehicle_id", c.DeleteUserVehicle)
	}

	// Admin routes for managing vehicle catalog
	adminVehicles := router.Group("/api/v1/admin/vehicles")
	adminVehicles.Use(middleware.AuthMiddleware(), middleware.RequireAdmin())
	{
		// Vehicle Types management
		adminVehicles.POST("/types", c.CreateVehicleType)
		adminVehicles.PUT("/types/:id", c.UpdateVehicleType)
		adminVehicles.DELETE("/types/:id", c.DeleteVehicleType)

		// Brands management
		adminVehicles.POST("/brands", c.CreateBrand)
		adminVehicles.PUT("/brands/:id", c.UpdateBrand)
		adminVehicles.DELETE("/brands/:id", c.DeleteBrand)

		// Models management
		adminVehicles.POST("/models", c.CreateModel)
		adminVehicles.PUT("/models/:id", c.UpdateModel)
		adminVehicles.DELETE("/models/:id", c.DeleteModel)

		// Generations management
		adminVehicles.POST("/generations", c.CreateGeneration)
		adminVehicles.PUT("/generations/:id", c.UpdateGeneration)
		adminVehicles.DELETE("/generations/:id", c.DeleteGeneration)
	}
}

// Public endpoints

// @Summary     List all vehicle types
// @Description Get a list of all available vehicle types
// @Tags        Types
// @Accept      json
// @Produce     json
// @Order       1
// @Success     200 {object} dto.ListVehicleTypesResponse
// @Router      /vehicles/types [get]
func (c *VehicleController) ListVehicleTypes(ctx *gin.Context) {
	types, err := c.vehicleUseCase.ListVehicleTypes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, types)
}

// @Summary     Get vehicle type details
// @Description Get details of a specific vehicle type
// @Tags        Types
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Type ID"
// @Order       2
// @Success     200 {object} dto.VehicleTypeResponse
// @Router      /vehicles/types/{id} [get]
func (c *VehicleController) GetVehicleType(ctx *gin.Context) {
	id := ctx.Param("id")
	vehicleType, err := c.vehicleUseCase.GetVehicleType(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, vehicleType)
}

// @Summary     List all vehicle brands
// @Description Get a list of all available vehicle brands
// @Tags        Brands
// @Accept      json
// @Produce     json
// @Order       1
// @Router      /vehicles/brands [get]
func (c *VehicleController) ListBrands(ctx *gin.Context) {
	brands, err := c.vehicleUseCase.ListBrands(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, brands)
}

// @Summary     List vehicle brands by type
// @Description Get a list of vehicle brands for a specific vehicle type
// @Tags        Brands
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Type ID"
// @Success     200 {object} dto.ListVehicleBrandsResponse
// @Router      /vehicles/types/{id}/brands [get]
func (c *VehicleController) ListBrandsByType(ctx *gin.Context) {
	typeID := ctx.Param("id")
	brands, err := c.vehicleUseCase.ListBrandsByType(ctx, typeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, brands)
}

// @Summary     Get vehicle brand details
// @Description Get details of a specific vehicle brand
// @Tags        Brands
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Brand ID"
// @Success     200 {object} dto.VehicleBrandResponse
// @Router      /vehicles/brands/{id} [get]
func (c *VehicleController) GetBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	brand, err := c.vehicleUseCase.GetBrand(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

// @Summary     List all vehicle models
// @Description Get a list of all available vehicle models
// @Tags        Models
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.ListVehicleModelsResponse
// @Router      /vehicles/models [get]
func (c *VehicleController) ListModels(ctx *gin.Context) {
	models, err := c.vehicleUseCase.ListModels(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, models)
}

// @Summary     List vehicle models by brand
// @Description Get a list of vehicle models for a specific vehicle brand
// @Tags        Models
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Brand ID"
// @Success     200 {object} dto.ListVehicleModelsResponse
// @Router      /vehicles/brands/{id}/models [get]
func (c *VehicleController) ListModelsByBrand(ctx *gin.Context) {
	brandID := ctx.Param("id")
	models, err := c.vehicleUseCase.ListModelsByBrand(ctx, brandID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, models)
}

// @Summary     Get vehicle model details
// @Description Get details of a specific vehicle model
// @Tags        Models
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Model ID"
// @Success     200 {object} dto.VehicleModelResponse
// @Router      /vehicles/models/{id} [get]
func (c *VehicleController) GetModel(ctx *gin.Context) {
	id := ctx.Param("id")
	model, err := c.vehicleUseCase.GetModel(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, model)
}

// @Summary     List all vehicle generations
// @Description Get a list of all available vehicle generations
// @Tags        Generations
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.ListVehicleGenerationsResponse
// @Router      /vehicles/generations [get]
func (c *VehicleController) ListGenerations(ctx *gin.Context) {
	generations, err := c.vehicleUseCase.ListGenerations(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, generations)
}

// @Summary     Get vehicle generation details
// @Description Get details of a specific vehicle generation
// @Tags        Generations
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Generation ID"
// @Success     200 {object} dto.VehicleGenerationResponse
// @Router      /vehicles/generations/{id} [get]
func (c *VehicleController) GetGeneration(ctx *gin.Context) {
	id := ctx.Param("id")
	generation, err := c.vehicleUseCase.GetGeneration(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, generation)
}

// @Summary     List vehicle generations by model
// @Description Get a list of vehicle generations for a specific vehicle model
// @Tags        Generations
// @Accept      json
// @Produce     json
// @Param       id path string true "Vehicle Model ID"
// @Success     200 {object} dto.ListVehicleGenerationsResponse
// @Router      /vehicles/models/{id}/generations [get]
func (c *VehicleController) ListGenerationsByModel(ctx *gin.Context) {
	modelID := ctx.Param("id")
	generations, err := c.vehicleUseCase.ListGenerationsByModel(ctx, modelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, generations)
}

// requires authentication

// @Summary     Add new vehicle to
// @Description Add new vehicle to user
// @Tags        User - Vehicles
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       vehicleType body dto.CreateUserVehicleRequest true "UserVehicle Type"
// @Success     201 {object} dto.UserVehicleResponse
// @Router      /user/vehicles [post]
func (c *VehicleController) AddUserVehicle(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	var userVehicle dto.CreateUserVehicleRequest
	if err := ctx.ShouldBindJSON(&userVehicle); err != nil {
		logger.Error(err, "Failed to bind user vehicle request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdUserVehicle, err := c.vehicleUseCase.AddUserVehicle(ctx, userID, &userVehicle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, createdUserVehicle)
}

// @Summary     List all user vehicles
// @Description Get a list of all user vehicle
// @Tags        User - Vehicles
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Success     200 {object} dto.ListUserVehiclesResponse
// @Router      /user/vehicles [get]
func (c *VehicleController) ListUserVehicles(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	userVehicles, err := c.vehicleUseCase.ListUserVehicles(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, userVehicles)
}

// @Summary     Get user vehicle details
// @Description Get details of a specific user vehicle
// @Tags        User - Vehicles
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       vehicle_id path string true "Vehicle ID"
// @Success     200 {object} dto.UserVehicleResponse
// @Router      /user/vehicles/{vehicle_id} [get]
func (c *VehicleController) GetUserVehicle(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")
	userVehicle, err := c.vehicleUseCase.GetUserVehicle(ctx, userID, vehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, userVehicle)
}

// @Summary     Update user vehicle details
// @Description Update the details of a specific user vehicle
// @Tags        User - Vehicles
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       vehicle_id path string true "Vehicle ID"
// @Param       userVehicle body dto.UpdateUserVehicleRequest true "User Vehicle"
// @Success     200 {object} dto.UpdateUserVehicleRequest
// @Router      /user/vehicles/{vehicle_id} [put]
func (c *VehicleController) UpdateUserVehicle(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")
	var userVehicle dto.UpdateUserVehicleRequest
	if err := ctx.ShouldBindJSON(&userVehicle); err != nil {
		logger.Error(err, "Failed to bind user vehicle request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedUserVehicle, err := c.vehicleUseCase.UpdateUserVehicle(ctx, userID, vehicleID, &userVehicle)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, updatedUserVehicle)
}

// @Summary  delete a user vehicle
// @Description Delete a user vehicle
// @Tags        User - Vehicles
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       vehicle_id path string true "Vehicle ID"
// @Success     204
// @Router      /user/vehicles/{vehicle_id} [delete]
func (c *VehicleController) DeleteUserVehicle(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	vehicleID := ctx.Param("vehicle_id")

	err := c.vehicleUseCase.DeleteUserVehicle(ctx, userID, vehicleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)

}

// Admin endpoints

// @Summary     Create a new vehicle type
// @Description Create a new vehicle type
// @Tags        Admin - Types
// @Accept      json
// @Produce     json
// @Security    BearerAuth
// @Param       vehicleType body dto.CreateVehicleTypeRequest true "Vehicle Type"
// @Success     201 {object} dto.VehicleTypeResponse
// @Router      /admin/vehicles/types [post]
func (c *VehicleController) CreateVehicleType(ctx *gin.Context) {
	var vehicleType dto.CreateVehicleTypeRequest
	if err := ctx.ShouldBindJSON(&vehicleType); err != nil {
		logger.Error(err, "Failed to bind vehicle type request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdVehicleType, err := c.vehicleUseCase.CreateVehicleType(ctx, vehicleType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, createdVehicleType)
}

// @Summary     Update a vehicle type
// @Description Update a vehicle type
// @Tags        Admin - Types
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Type ID"
// @Param       vehicleType body dto.UpdateVehicleTypeRequest true "Vehicle Type"
// @Success     200 {object} dto.VehicleTypeResponse
// @Router      /admin/vehicles/types/{id} [put]
func (c *VehicleController) UpdateVehicleType(ctx *gin.Context) {
	id := ctx.Param("id")
	var vehicleType dto.UpdateVehicleTypeRequest
	if err := ctx.ShouldBindJSON(&vehicleType); err != nil {
		logger.Error(err, "Failed to bind vehicle type request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedVehicleType, err := c.vehicleUseCase.UpdateVehicleType(ctx, id, vehicleType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, updatedVehicleType)
}

// @Summary     Delete a vehicle type
// @Description Delete a vehicle type
// @Tags        Admin - Types
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Type ID"
// @Success     204
// @Router      /admin/vehicles/types/{id} [delete]
func (c *VehicleController) DeleteVehicleType(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.vehicleUseCase.DeleteVehicleType(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary     Create a new vehicle brand
// @Description Create a new vehicle brand
// @Tags        Admin - Brands
// @Accept      json
// @Security    BearerAuth
// @Param       brand body dto.CreateVehicleBrandRequest true "Vehicle Brand"
// @Success     201 {object} dto.VehicleBrandResponse
// @Router      /admin/vehicles/brands [post]
func (c *VehicleController) CreateBrand(ctx *gin.Context) {
	var brand dto.CreateVehicleBrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		logger.Error(err, "Failed to bind vehicle brand request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdBrand, err := c.vehicleUseCase.CreateBrand(ctx, brand)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, createdBrand)
}

// @Summary     Update a vehicle brand
// @Description Update a vehicle brand
// @Tags        Admin - Brands
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Brand ID"
// @Param       brand body dto.UpdateVehicleBrandRequest true "Vehicle Brand"
// @Success     200 {object} dto.VehicleBrandResponse
// @Router      /admin/vehicles/brands/{id} [put]
func (c *VehicleController) UpdateBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	var brand dto.UpdateVehicleBrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		logger.Error(err, "Failed to bind vehicle brand request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedBrand, err := c.vehicleUseCase.UpdateBrand(ctx, id, brand)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, updatedBrand)
}

// @Summary     Delete a vehicle brand
// @Description Delete a vehicle brand
// @Tags        Admin - Brands
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Brand ID"
// @Success     204
// @Router      /admin/vehicles/brands/{id} [delete]
func (c *VehicleController) DeleteBrand(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.vehicleUseCase.DeleteBrand(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary     Create a new vehicle model
// @Description Create a new vehicle model
// @Tags        Admin - Models
// @Accept      json
// @Security    BearerAuth
// @Param       model body dto.CreateVehicleModelRequest true "Vehicle Model"
// @Success     201 {object} dto.VehicleModelResponse
// @Router      /admin/vehicles/brands/{brandId}/models [post]
func (c *VehicleController) CreateModel(ctx *gin.Context) {
	var model dto.CreateVehicleModelRequest
	if err := ctx.ShouldBindJSON(&model); err != nil {
		logger.Error(err, "Failed to bind vehicle model request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdModel, err := c.vehicleUseCase.CreateModel(ctx, model)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, createdModel)
}

// @Summary     Update a vehicle model
// @Description Update a vehicle model
// @Tags        Admin - Models
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Model ID"
// @Param       model body dto.UpdateVehicleModelRequest true "Vehicle Model"
// @Success     200 {object} dto.VehicleModelResponse
// @Router      /admin/vehicles/models/{id} [put]
func (c *VehicleController) UpdateModel(ctx *gin.Context) {
	id := ctx.Param("id")
	var model dto.UpdateVehicleModelRequest
	if err := ctx.ShouldBindJSON(&model); err != nil {
		logger.Error(err, "Failed to bind vehicle model request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedModel, err := c.vehicleUseCase.UpdateModel(ctx, id, model)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, updatedModel)
}

// @Summary     Delete a vehicle model
// @Description Delete a vehicle model
// @Tags        Admin - Models
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Model ID"
// @Success     204
// @Router      /admin/vehicles/models/{id} [delete]
func (c *VehicleController) DeleteModel(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.vehicleUseCase.DeleteModel(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary     Create a new vehicle generation
// @Description Create a new vehicle generation
// @Tags        Admin - Generations
// @Accept      json
// @Security    BearerAuth
// @Param       generation body dto.CreateVehicleGenerationRequest true "Vehicle Generation"
// @Success     201 {object} dto.VehicleGenerationResponse
// @Router      /admin/vehicles/generations [post]
func (c *VehicleController) CreateGeneration(ctx *gin.Context) {
	var generation dto.CreateVehicleGenerationRequest
	if err := ctx.ShouldBindJSON(&generation); err != nil {
		logger.Error(err, "Failed to bind vehicle generation request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdModel, err := c.vehicleUseCase.CreateGeneration(ctx, generation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusCreated, createdModel)
}

// @Summary     Update a vehicle generation
// @Description Update a vehicle generation
// @Tags        Admin - Generations
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle generation ID"
// @Param       generation body dto.CreateVehicleGenerationRequest true "Vehicle Generation"
// @Success     200 {object} dto.VehicleGenerationResponse
// @Router      /admin/vehicles/generations/{id} [put]
func (c *VehicleController) UpdateGeneration(ctx *gin.Context) {
	id := ctx.Param("id")
	var generation dto.UpdateVehicleGenerationRequest
	if err := ctx.ShouldBindJSON(&generation); err != nil {
		logger.Error(err, "Failed to bind vehicle generation request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedGeneration, err := c.vehicleUseCase.UpdateGeneration(ctx, id, generation)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, updatedGeneration)
}

// @Summary     Delete a vehicle generation
// @Description Delete a vehicle generation
// @Tags        Admin - Generations
// @Accept      json
// @Security    BearerAuth
// @Param       id path string true "Vehicle Generation ID"
// @Success     204
// @Router      /admin/vehicles/generations/{id} [delete]
func (c *VehicleController) DeleteGeneration(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.vehicleUseCase.DeleteGeneration(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)
}
