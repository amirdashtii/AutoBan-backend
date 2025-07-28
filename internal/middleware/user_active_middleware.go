package middleware

import (
	"net/http"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/gin-gonic/gin"
)

func RequireActiveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		status, exists := c.Get("status")
		if !exists {
			logger.Error(errors.ErrTokenNotFound, "User status not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.ErrTokenNotFound.Error()})
			return
		}

		statusInt, ok := status.(float64)
		if !ok {
			logger.Error(errors.ErrInvalidTokenClaims, "Invalid user status type in token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.ErrInvalidTokenClaims.Error()})
			return
		}
		userStatus := entity.StatusType(int(statusInt))

		if userStatus != entity.Active {
			logger.Error(errors.ErrUserNotActive, "User is not active")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errors.ErrUserNotActive.Error()})
			return
		}
		c.Next()
	}
}
