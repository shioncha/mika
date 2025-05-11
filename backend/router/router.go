package router

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/handler"
)

func SetupRouter(client *ent.Client) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.POST("/sign-up", func(c *gin.Context) {
		handler.SignUp(c, client)
	})

	router.POST("/sign-in", func(c *gin.Context) {
		handler.SignIn(c, client)
	})

	router.POST("sign-out")

	router.GET("/i", i)

	router.GET("/test", func(c *gin.Context) {
		u, _ := client.Users.Query().All(context.Background())
		c.JSON(200, u)
	})

	return router
}

func loadPublicKey() (*rsa.PublicKey, error) {
	b64 := os.Getenv("JWT_PUBLIC_KEY_BASE64")
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pemBytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func validateJWT(tokenString string, publicKey *rsa.PublicKey) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func i(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[len("Bearer "):]
	publickey, _ := loadPublicKey()

	jwtClaims, err := validateJWT(tokenString, publickey)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": jwtClaims,
	})
}
