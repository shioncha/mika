package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/handler"
	"github.com/shioncha/mika/backend/middleware"
)

func SetupRouter(ah *handler.AuthHandler, client *ent.Client) *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	router.POST("/sign-up", ah.SignUp)

	router.POST("/sign-in", ah.SignIn)

	router.POST("sign-out")

	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/i", handler.I)
		authorized.GET("/posts", func(c *gin.Context) {
			handler.GetPost(c, client)
		})
		authorized.POST("/posts", func(c *gin.Context) {
			handler.CreatePost(c, client)
		})
		authorized.DELETE("/posts/:id", func(c *gin.Context) {
			handler.DeletePost(c, client)
		})
		authorized.GET("/tags", func(c *gin.Context) {
			handler.GetTags(c, client)
		})
		authorized.GET("/tags/:tag/posts", func(c *gin.Context) {
			handler.GetPostsByTag(c, client)
		})
	}

	router.GET("/test", func(c *gin.Context) {
		u, _ := client.Users.Query().All(context.Background())
		c.JSON(200, u)
	})

	return router
}
