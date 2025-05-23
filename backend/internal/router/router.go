package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/handler"
	"github.com/shioncha/mika/backend/internal/middleware"
)

func SetupRouter(ah *handler.AuthHandler, ph *handler.PostHandler, th *handler.TagHandler) *gin.Engine {
	router := gin.Default()

	router.POST("/sign-up", ah.SignUp)
	router.POST("/sign-in", ah.SignIn)
	router.POST("/sign-out")

	authorized := router.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/i", handler.I)

		authorized.GET("/posts", ph.GetPosts)
		authorized.POST("/posts", ph.CreatePost)
		authorized.GET("/posts/:id", ph.GetPost)
		authorized.DELETE("/posts/:id", ph.DeletePost)

		authorized.GET("/tags", th.GetTags)
		authorized.GET("/tags/:tag/posts", th.GetPostsByTag)
	}

	return router
}
