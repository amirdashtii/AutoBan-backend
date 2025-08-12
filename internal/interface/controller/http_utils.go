package controller

import (
	"net/http"

	customerr "github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/gin-gonic/gin"
)

// httpStatusFromError maps custom domain errors to appropriate HTTP status codes
func httpStatusFromError(err error) int {
	// 400 Bad Request: all invalid input/body/path errors
	if customerr.Is(err, customerr.ErrBadRequest) ||
		customerr.Is(err, customerr.ErrInvalidPhoneNumber) ||
		customerr.Is(err, customerr.ErrInvalidPassword) ||
		customerr.Is(err, customerr.ErrInvalidEmail) ||
		customerr.Is(err, customerr.ErrInvalidBirthday) ||
		customerr.Is(err, customerr.ErrInvalidRole) ||
		customerr.Is(err, customerr.ErrInvalidStatus) ||
		customerr.Is(err, customerr.ErrInvalidTokenFormat) ||
		customerr.Is(err, customerr.ErrInvalidTokenClaims) ||
		customerr.Is(err, customerr.ErrInvalidPurchaseDate) ||
		customerr.Is(err, customerr.ErrInvalidUserID) ||
		customerr.Is(err, customerr.ErrInvalidVehicleTypeCreateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleTypeUpdateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleTypeID) ||
		customerr.Is(err, customerr.ErrInvalidVehicleBrandID) ||
		customerr.Is(err, customerr.ErrInvalidVehicleBrandType) ||
		customerr.Is(err, customerr.ErrInvalidVehicleBrandCreateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleBrandUpdateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleModelID) ||
		customerr.Is(err, customerr.ErrInvalidVehicleModelCreateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleModelUpdateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleGenerationID) ||
		customerr.Is(err, customerr.ErrInvalidVehicleGenerationCreateRequest) ||
		customerr.Is(err, customerr.ErrInvalidVehicleGenerationUpdateRequest) ||
		customerr.Is(err, customerr.ErrInvalidUserVehicleID) ||
		customerr.Is(err, customerr.ErrInvalidUserVehicleCreateRequest) ||
		customerr.Is(err, customerr.ErrInvalidUserVehicleUpdateRequest) ||
		customerr.Is(err, customerr.ErrInvalidDate) ||
		customerr.Is(err, customerr.ErrUserVehicleIDRequired) {
		return http.StatusBadRequest
	}

	// 401 Unauthorized
	if customerr.Is(err, customerr.ErrInvalidToken) ||
		customerr.Is(err, customerr.ErrTokenNotFound) {
		return http.StatusUnauthorized
	}

	// 403 Forbidden
	if customerr.Is(err, customerr.ErrAccessDenied) ||
		customerr.Is(err, customerr.ErrUserNotActive) ||
		customerr.Is(err, customerr.ErrUserVehicleNotOwned) ||
		customerr.Is(err, customerr.ErrOilFilterNotOwned) ||
		customerr.Is(err, customerr.ErrOilChangeNotOwned) {
		return http.StatusForbidden
	}

	// 404 Not Found
	if customerr.Is(err, customerr.ErrUserNotFound) {
		return http.StatusNotFound
	}

	// Default: 500
	return http.StatusInternalServerError
}

// respondError writes a consistent error response using the custom error structure
func respondError(ctx *gin.Context, err error) {
	status := httpStatusFromError(err)
	// Ensure payload is CustomError shape
	if _, ok := err.(*customerr.CustomError); !ok {
		// Wrap unknown errors in a generic internal server error for clients
		err = customerr.ErrInternalServerError
	}
	ctx.JSON(status, gin.H{"error": err})
} 