package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shioncha/mika/backend/internal/handler"
	"github.com/shioncha/mika/backend/internal/middleware"
)

func SetupRouter(
	ah *handler.AuthHandler,
	ph *handler.PostHandler,
	th *handler.TagHandler,
	am *middleware.AuthRequiredMiddleware,
) *gin.Engine {
	router := gin.Default()

	/*
	 * Public routes
	 */
	router.POST("/sign-up", ah.SignUp)
	router.POST("/sign-in", ah.SignIn)
	router.POST("/sign-out")

	/*
	 * Private routes
	 */
	authorized := router.Group("/")
	authorized.Use(am.AuthRequired())
	{
		userRoutes := authorized.Group("/users")
		{
			userRoutes.GET("/me", ah.Get)
		}

		postRoutes := authorized.Group("/posts")
		{
			postRoutes.GET("", ph.GetPosts)
			postRoutes.POST("", ph.CreatePost)
			postRoutes.GET("/:id", ph.GetPost)
			postRoutes.DELETE("/:id", ph.DeletePost)
		}

		tagRoutes := authorized.Group("/tags")
		{
			tagRoutes.GET("", th.GetTags)
			tagRoutes.GET("/:tag/posts", th.GetPostsByTag)
		}
	}

	return router
}
