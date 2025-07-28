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
		// Complete hierarchy
		vehicleGroup.GET("/hierarchy", c.GetCompleteHierarchy)

		// Vehicle Types
		vehicleGroup.GET("/types", c.ListTypes)
		vehicleGroup.GET("/types/:type_id", c.GetType)

		// Brands
		vehicleGroup.GET("/types/:type_id/brands", c.ListBrands)
		vehicleGroup.GET("/types/:type_id/brands/:brand_id", c.GetBrand)

		// Models
		vehicleGroup.GET("/types/:type_id/brands/:brand_id/models", c.ListModels)
		vehicleGroup.GET("/types/:type_id/brands/:brand_id/models/:model_id", c.GetModel)

		// Generations
		vehicleGroup.GET("/types/:type_id/brands/:brand_id/models/:model_id/generations", c.ListGenerations)
		vehicleGroup.GET("/types/:type_id/brands/:brand_id/models/:model_id/generations/:generation_id", c.GetGeneration)
	}

	// User vehicle management (requires authentication)
	userVehicles := router.Group("/api/v1/user/vehicles")
	userVehicles.Use(middleware.AuthMiddleware())
	userVehicles.Use(middleware.RequireActiveUser())
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
		adminVehicles.PUT("/types/:type_id", c.UpdateVehicleType)
		adminVehicles.DELETE("/types/:type_id", c.DeleteVehicleType)

		// Brands management
		adminVehicles.POST("/types/:type_id/brands", c.CreateBrand)
		adminVehicles.PUT("/types/:type_id/brands/:brand_id", c.UpdateBrand)
		adminVehicles.DELETE("/types/:type_id/brands/:brand_id", c.DeleteBrand)

		// Models management
		adminVehicles.POST("/types/:type_id/brands/:brand_id/models", c.CreateModel)
		adminVehicles.PUT("/types/:type_id/brands/:brand_id/models/:model_id", c.UpdateModel)
		adminVehicles.DELETE("/types/:type_id/brands/:brand_id/models/:model_id", c.DeleteModel)

		// Generations management
		adminVehicles.POST("/types/:type_id/brands/:brand_id/models/:model_id/generations", c.CreateGeneration)
		adminVehicles.PUT("/types/:type_id/brands/:brand_id/models/:model_id/generations/:generation_id", c.UpdateGeneration)
		adminVehicles.DELETE("/types/:type_id/brands/:brand_id/models/:model_id/generations/:generation_id", c.DeleteGeneration)
	}
}

// Public endpoints

// @Summary     Get complete vehicle hierarchy
// @Description Get the complete vehicle hierarchy including all types, brands, models, and generations
// @Tags        Hierarchy
// @Accept      json
// @Produce     json
// @Success     200 {object} dto.CompleteVehicleHierarchyResponse
// @Router      /vehicles/hierarchy [get]
func (c *VehicleController) GetCompleteHierarchy(ctx *gin.Context) {
	hierarchy, err := c.vehicleUseCase.GetCompleteVehicleHierarchy(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, hierarchy)
}

// @Summary     List all vehicle types
// @Description Get a list of all available vehicle types
// @Tags        Types
// @Accept      json
// @Produce     json
// @Order       1
// @Success     200 {object} dto.ListVehicleTypesResponse
// @Router      /vehicles/types [get]
func (c *VehicleController) ListTypes(ctx *gin.Context) {
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
// @Param       type_id path string true "Vehicle Type ID"
// @Order       2
// @Success     200 {object} dto.VehicleTypeResponse
// @Router      /vehicles/types/{type_id} [get]
func (c *VehicleController) GetType(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	vehicleType, err := c.vehicleUseCase.GetVehicleType(ctx, typeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, vehicleType)
}

// @Summary     List vehicle brands by type
// @Description Get a list of vehicle brands for a specific vehicle type
// @Tags        Brands
// @Accept      json
// @Produce     json
// @Param       type_id path string true "Vehicle Type ID"
// @Success     200 {object} dto.ListVehicleBrandsResponse
// @Router      /vehicles/types/{type_id}/brands [get]
func (c *VehicleController) ListBrands(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brands, err := c.vehicleUseCase.ListBrands(ctx, typeID)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Success     200 {object} dto.VehicleBrandResponse
// @Router      /vehicles/types/{type_id}/brands/{brand_id} [get]
func (c *VehicleController) GetBrand(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	brand, err := c.vehicleUseCase.GetBrand(ctx, typeID, brandID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, brand)
}

// @Summary     List vehicle models by brand
// @Description Get a list of vehicle models for a specific vehicle brand
// @Tags        Models
// @Accept      json
// @Produce     json
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Success     200 {object} dto.ListVehicleModelsResponse
// @Router      /vehicles/types/{type_id}/brands/{brand_id}/models [get]
func (c *VehicleController) ListModels(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	models, err := c.vehicleUseCase.ListModels(ctx, typeID, brandID)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Success     200 {object} dto.VehicleModelResponse
// @Router      /vehicles/types/{type_id}/brands/{brand_id}/models/{model_id} [get]
func (c *VehicleController) GetModel(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	model, err := c.vehicleUseCase.GetModel(ctx, typeID, brandID, modelID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, model)
}

// @Summary     Get vehicle generation details
// @Description Get details of a specific vehicle generation
// @Tags        Generations
// @Accept      json
// @Produce     json
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Param       generation_id path string true "Vehicle Generation ID"
// @Success     200 {object} dto.VehicleGenerationResponse
// @Router      /vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations/{generation_id} [get]
func (c *VehicleController) GetGeneration(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	generationID := ctx.Param("generation_id")
	generation, err := c.vehicleUseCase.GetGeneration(ctx, typeID, brandID, modelID, generationID)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Success     200 {object} dto.ListVehicleGenerationsResponse
// @Router      /vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations [get]
func (c *VehicleController) ListGenerations(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	generations, err := c.vehicleUseCase.ListGenerations(ctx, typeID, brandID, modelID)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       vehicleType body dto.UpdateVehicleTypeRequest true "Vehicle Type"
// @Success     200 {object} dto.VehicleTypeResponse
// @Router      /admin/vehicles/types/{type_id} [put]
func (c *VehicleController) UpdateVehicleType(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	var vehicleType dto.UpdateVehicleTypeRequest
	if err := ctx.ShouldBindJSON(&vehicleType); err != nil {
		logger.Error(err, "Failed to bind vehicle type request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedVehicleType, err := c.vehicleUseCase.UpdateVehicleType(ctx, typeID, vehicleType)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Success     204
// @Router      /admin/vehicles/types/{type_id} [delete]
func (c *VehicleController) DeleteVehicleType(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	err := c.vehicleUseCase.DeleteVehicleType(ctx, typeID)
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
// @Router      /admin/vehicles/types/{type_id}/brands [post]
func (c *VehicleController) CreateBrand(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	var request dto.CreateVehicleBrandRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind vehicle brand request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdBrand, err := c.vehicleUseCase.CreateBrand(ctx, typeID, request)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       brand body dto.UpdateVehicleBrandRequest true "Vehicle Brand"
// @Success     200 {object} dto.VehicleBrandResponse
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id} [put]
func (c *VehicleController) UpdateBrand(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	var brand dto.UpdateVehicleBrandRequest
	if err := ctx.ShouldBindJSON(&brand); err != nil {
		logger.Error(err, "Failed to bind vehicle brand request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedBrand, err := c.vehicleUseCase.UpdateBrand(ctx, typeID, brandID, brand)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Success     204
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id} [delete]
func (c *VehicleController) DeleteBrand(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	err := c.vehicleUseCase.DeleteBrand(ctx, typeID, brandID)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model body dto.CreateVehicleModelRequest true "Vehicle Model"
// @Success     201 {object} dto.VehicleModelResponse
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id}/models [post]
func (c *VehicleController) CreateModel(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	var request dto.CreateVehicleModelRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind vehicle model request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdModel, err := c.vehicleUseCase.CreateModel(ctx, typeID, brandID, request)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Param       model body dto.UpdateVehicleModelRequest true "Vehicle Model"
// @Success     200 {object} dto.VehicleModelResponse
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id} [put]
func (c *VehicleController) UpdateModel(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	var model dto.UpdateVehicleModelRequest
	if err := ctx.ShouldBindJSON(&model); err != nil {
		logger.Error(err, "Failed to bind vehicle model request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedModel, err := c.vehicleUseCase.UpdateModel(ctx, typeID, brandID, modelID, model)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Success     204
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id} [delete]
func (c *VehicleController) DeleteModel(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	err := c.vehicleUseCase.DeleteModel(ctx, typeID, brandID, modelID)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Param       generation body dto.CreateVehicleGenerationRequest true "Vehicle Generation"
// @Success     201 {object} dto.VehicleGenerationResponse
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations [post]
func (c *VehicleController) CreateGeneration(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	var request dto.CreateVehicleGenerationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		logger.Error(err, "Failed to bind vehicle generation request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	createdModel, err := c.vehicleUseCase.CreateGeneration(ctx, typeID, brandID, modelID, request)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Param       generation_id path string true "Vehicle Generation ID"
// @Param       generation body dto.CreateVehicleGenerationRequest true "Vehicle Generation"
// @Success     200 {object} dto.VehicleGenerationResponse
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations/{generation_id} [put]
func (c *VehicleController) UpdateGeneration(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	generationID := ctx.Param("generation_id")
	var generation dto.UpdateVehicleGenerationRequest
	if err := ctx.ShouldBindJSON(&generation); err != nil {
		logger.Error(err, "Failed to bind vehicle generation request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrInvalidRequestBody})
		return
	}
	updatedGeneration, err := c.vehicleUseCase.UpdateGeneration(ctx, typeID, brandID, modelID, generationID, generation)
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
// @Param       type_id path string true "Vehicle Type ID"
// @Param       brand_id path string true "Vehicle Brand ID"
// @Param       model_id path string true "Vehicle Model ID"
// @Param       generation_id path string true "Vehicle Generation ID"
// @Success     204
// @Router      /admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations/{generation_id} [delete]
func (c *VehicleController) DeleteGeneration(ctx *gin.Context) {
	typeID := ctx.Param("type_id")
	brandID := ctx.Param("brand_id")
	modelID := ctx.Param("model_id")
	generationID := ctx.Param("generation_id")
	err := c.vehicleUseCase.DeleteGeneration(ctx, typeID, brandID, modelID, generationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.Status(http.StatusNoContent)
}
