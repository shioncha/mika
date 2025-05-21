package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/ent"
	"github.com/shioncha/mika/backend/handler"
	"github.com/shioncha/mika/backend/middleware"
)

func SetupRouter(ah *handler.AuthHandler, ph *handler.PostHandler, th *handler.TagHandler, client *ent.Client) *gin.Engine {
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
		authorized.GET("/posts", ph.GetPosts)
		authorized.POST("/posts", ph.CreatePost)
		authorized.DELETE("/posts/:id", ph.DeletePost)
		authorized.GET("/tags", th.GetTags)
		authorized.GET("/tags/:tag/posts", th.GetPostsByTag)
	}

	router.GET("/test", func(c *gin.Context) {
		u, _ := client.Users.Query().All(context.Background())
		c.JSON(200, u)
	})

	return router
}
