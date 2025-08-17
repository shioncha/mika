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
	uh *handler.UserHandler,
	am *middleware.AuthRequiredMiddleware,
	rm *middleware.RateLimitMiddleware,
) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("FRONTEND_URL"),
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PATCH",
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
	authRoutes := router.Group("/")
	authRoutes.Use(rm.IPBasedRateLimit())
	{
		authRoutes.POST("/sign-up", ah.SignUp)
		authRoutes.POST("/sign-in", ah.SignIn)
		authRoutes.POST("/refresh-token", ah.RefreshAccessToken)
	}
	router.POST("/sign-out", ah.SignOut)

	/*
	 * Private routes
	 */
	authorized := router.Group("/")
	authorized.Use(am.AuthRequired())
	{
		userRoutes := authorized.Group("/users")
		{
			userRoutes.GET("/me", uh.Get)
		}

		postRoutes := authorized.Group("/posts")
		{
			postRoutes.GET("", ph.GetPosts)
			postRoutes.POST("", ph.CreatePost)
			postRoutes.GET("/:id", ph.GetPost)
			postRoutes.PATCH("/:id", ph.UpdatePost)
			postRoutes.DELETE("/:id", ph.DeletePost)
		}

		tagRoutes := authorized.Group("/tags")
		{
			tagRoutes.GET("", th.GetTags)
			tagRoutes.GET("/:tag/posts", th.GetPostsByTag)
		}

		{
			authorized.PATCH("/account", uh.Update)
		}
	}

	return router
}
