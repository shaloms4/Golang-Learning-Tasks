package Infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	domain "github.com/shaloms4/Golang-Learning-Tasks/task_7/Domain"
)

type AuthMiddleware struct {
	jwtService domain.JWTService
}

func NewAuthMiddleware(jwtService domain.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

func (m *AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Authorization header is required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid authorization header format"})
			c.Abort()
			return
		}

		claims, err := m.jwtService.ValidateToken(tokenParts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("userID", claims["userID"])
		c.Set("userRole", claims["role"])
		c.Next()
	}
}

// RequireAdmin middleware checks if the user has admin role
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, domain.ErrorResponse{Message: "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}