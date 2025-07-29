package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("FRONTEND_URL"),
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
