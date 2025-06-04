package middleware

import (
	"strings"

	"github.com/amirdashtii/AutoBan/config"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationHeader = "Authorization"
	BearerSchema        = "Bearer"
)

// AuthMiddleware checks for a valid JWT token in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
		return nil
	}

	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			logger.Error(errors.ErrTokenNotFound, "Authorization header is missing")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrTokenNotFound})
			return
		}

		// Check the Authorization schema
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != BearerSchema {
			logger.Error(errors.ErrInvalidTokenFormat, "Invalid authorization header format")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrInvalidTokenFormat})
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.ErrInvalidToken
			}
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil {
			logger.Error(err, "Failed to parse token")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrInvalidToken})
			return
		}

		if !token.Valid {
			logger.Error(errors.ErrInvalidToken, "Token is not valid")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrInvalidToken})
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Error(errors.ErrInvalidTokenClaims, "Failed to extract token claims")
			c.AbortWithStatusJSON(401, gin.H{"error": errors.ErrInvalidTokenClaims})
			return
		}

		// Store user information in context
		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])
		c.Set("phone_number", claims["phone_number"])
		c.Set("status", claims["status"])

		c.Next()
	}
}
