package middleware

import (
	"AutoBan/internal/errors"
	"AutoBan/pkg/logger"

	"github.com/gin-gonic/gin"
)

// RequireAdmin checks if the user is either admin or superAdmin
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			logger.Error(errors.ErrTokenNotFound, "Role not found in context")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrTokenNotFound})
			return
		}

		userRole, ok := role.(string)
		if !ok {
			logger.Error(errors.ErrInvalidTokenClaims, "User role in token is invalid")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrInvalidTokenClaims})
			return
		}

		// Check if the user is either admin or superAdmin
		if userRole != "admin" && userRole != "superAdmin" {
			logger.Error(errors.ErrAccessDenied, "User does not have the required role")
			c.AbortWithStatusJSON(403, gin.H{"error": errors.ErrAccessDenied})
			return
		}

		c.Next()
	}
}
