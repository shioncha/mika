package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shioncha/mika/backend/internal/auth"
	"github.com/shioncha/mika/backend/internal/database"
	"github.com/shioncha/mika/backend/internal/handler"
	"github.com/shioncha/mika/backend/internal/middleware"
	ent "github.com/shioncha/mika/backend/internal/repository/ent"
	redis "github.com/shioncha/mika/backend/internal/repository/redis"
	"github.com/shioncha/mika/backend/internal/router"
	"github.com/shioncha/mika/backend/internal/service"
)

type App struct {
	router  *gin.Engine
	cleanup func()
}

func newApp() (*App, error) {
	privateKey, publicKey, err := auth.LoadKeys()
	if err != nil {
		return nil, err
	}

	redisClient := database.NewRedisClient()
	sessionRepo := redis.NewSessionRepository(redisClient)

	client := database.SetupClient()

	userRepo := ent.NewUserRepository(client)
	authService := service.NewAuthService(userRepo, sessionRepo, publicKey, privateKey)
	authHandler := handler.NewAuthHandler(authService)

	postRepo := ent.NewPostRepository(client)
	postService := service.NewPostService(client, postRepo)
	postHandler := handler.NewPostHandler(postService)

	tagRepo := ent.NewTagRepository(client)
	tagService := service.NewTagService(tagRepo)
	tagHandler := handler.NewTagHandler(tagService)

	authMiddleware := middleware.NewAuthRequiredMiddleware(publicKey)

	router := router.SetupRouter(authHandler, postHandler, tagHandler, authMiddleware)

	cleanupFunc := func() {
		database.CloseClient(client)
		database.CloseRedisClient(redisClient)
	}

	return &App{
		router:  router,
		cleanup: cleanupFunc,
	}, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	app, err := newApp()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	defer app.cleanup()

	if err := app.router.Run(":" + os.Getenv("BACKEND_PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
