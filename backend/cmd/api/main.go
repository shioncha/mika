package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shioncha/mika/backend/internal/auth"
	"github.com/shioncha/mika/backend/internal/database"
	"github.com/shioncha/mika/backend/internal/handler"
	"github.com/shioncha/mika/backend/internal/middleware"
	entrepogitory "github.com/shioncha/mika/backend/internal/repository/ent"
	"github.com/shioncha/mika/backend/internal/router"
	"github.com/shioncha/mika/backend/internal/service"
)

func main() {
	godotenv.Load(".env")

	privateKey, publicKey, err := auth.LoadKeys()
	if err != nil {
		panic("Failed to load JWT keys: " + err.Error())
	}

	client := database.SetupClient()
	defer database.CloseClient(client)

	userRepo := entrepogitory.NewUserRepository(client)
	authService := service.NewAuthService(userRepo, publicKey, privateKey)
	authHandler := handler.NewAuthHandler(authService)

	postRepo := entrepogitory.NewPostRepository(client)
	postService := service.NewPostService(client, postRepo)
	postHandler := handler.NewPostHandler(postService)

	tagRepo := entrepogitory.NewTagRepository(client)
	tagService := service.NewTagService(tagRepo)
	tagHandler := handler.NewTagHandler(tagService)

	authMiddleware := middleware.NewAuthRequiredMiddleware(publicKey)

	router := router.SetupRouter(authHandler, postHandler, tagHandler, authMiddleware)
	router.Run(":8080")
}
