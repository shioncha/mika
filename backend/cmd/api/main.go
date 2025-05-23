package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shioncha/mika/backend/internal/database"
	"github.com/shioncha/mika/backend/internal/handler"
	entrepogitory "github.com/shioncha/mika/backend/internal/repository/ent"
	"github.com/shioncha/mika/backend/internal/router"
	"github.com/shioncha/mika/backend/internal/service"
)

func main() {
	godotenv.Load(".env")

	client := database.SetupClient()
	defer database.CloseClient(client)

	userRepo := entrepogitory.NewUserRepository(client)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	postRepo := entrepogitory.NewPostRepository(client)
	postService := service.NewPostService(client, postRepo)
	postHandler := handler.NewPostHandler(postService)

	tagRepo := entrepogitory.NewTagRepository(client)
	tagService := service.NewTagService(tagRepo)
	tagHandler := handler.NewTagHandler(tagService)

	router := router.SetupRouter(authHandler, postHandler, tagHandler)
	router.Run(":8080")
}
