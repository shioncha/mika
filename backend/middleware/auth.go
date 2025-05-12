package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/auth"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= len("Bearer ") || authHeader[:len("Bearer ")] != "Bearer " {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]

		jwtClaims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", jwtClaims["user_id"])
		c.Next()
	}
}
