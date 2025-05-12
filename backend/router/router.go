package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/handler"
	"github.com/shioncha/mika/backend/middleware"
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

	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/i", i)
	}

	router.GET("/test", func(c *gin.Context) {
		u, _ := client.Users.Query().All(context.Background())
		c.JSON(200, u)
	})

	return router
}

func i(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(200, gin.H{"user_id": userID})
}
