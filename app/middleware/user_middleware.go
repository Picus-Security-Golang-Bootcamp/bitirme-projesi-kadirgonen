package middleware

import (
	"net/http"
	"strings"

	"HW/app/service"
	"HW/pkg/logger"

	"github.com/gin-gonic/gin"
)

type JWTAuthMiddleware struct {
	authService service.JWTAuthService
	userService service.UserService
	logger      logger.Logger
}

func NewJWTAuthMiddleware(as service.JWTAuthService, us service.UserService, logger logger.Logger) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{
		authService: as,
		userService: us,
		logger:      logger,
	}
}

func (m *JWTAuthMiddleware) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
			authorized, claims := m.authService.VerifyToken(authHeader)
			if authorized {
				c.Set("UserId", claims.UserId)
				c.Set("Email", claims.Email)
				c.Next()
				return
			}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized access",
		})
		c.Abort()
	}
}

func (m *JWTAuthMiddleware) CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetInt("UserId")
		userIsInRole := m.userService.UserHasRole(userId, role)
		if userIsInRole {
			c.Next()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "insufficient permissions",
		})
		c.Abort()
	}
}
