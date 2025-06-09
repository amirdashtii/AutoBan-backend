package middleware

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/gin-gonic/gin"
)

// RequireAdmin checks if the user has admin or super admin role
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			logger.Error(errors.ErrTokenNotFound, "Role not found in context")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound.Error()})
			return
		}

		// Type assertion with safety check
		roleInt, ok := role.(float64) // JWT numbers are decoded as float64
		if !ok {
			logger.Error(errors.ErrInvalidTokenClaims, "Invalid role type in token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidTokenClaims.Error()})
			return
		}

		userRole := entity.RoleType(int(roleInt))

		// Check if the user is either admin or superAdmin
		if userRole != entity.AdminRole && userRole != entity.SuperAdminRole {
			logger.Error(errors.ErrAccessDenied, "User does not have admin privileges")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": errors.ErrAccessDenied.Error()})
			return
		}

		c.Next()
	}
}
