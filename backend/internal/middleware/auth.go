package middleware

import (
	"crypto/rsa"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/auth"
)

type AuthRequiredMiddleware struct {
	PublicKey *rsa.PublicKey
}

func NewAuthRequiredMiddleware(publicKey *rsa.PublicKey) *AuthRequiredMiddleware {
	return &AuthRequiredMiddleware{PublicKey: publicKey}
}

func (m *AuthRequiredMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= len("Bearer ") || authHeader[:len("Bearer ")] != "Bearer " {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]

		jwtClaims, err := auth.ValidateJWT(tokenString, m.PublicKey)
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user_id", jwtClaims["user_id"])
		c.Next()
	}
}
