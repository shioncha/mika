package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shioncha/mika/backend/database"
	"github.com/shioncha/mika/backend/handler"
	entrepogitory "github.com/shioncha/mika/backend/internal/repository/ent"
	"github.com/shioncha/mika/backend/internal/service"
	"github.com/shioncha/mika/backend/router"
)

func main() {
	godotenv.Load(".env")

	client := database.SetupClient()
	defer database.CloseClient(client)

	userRepo := entrepogitory.NewUserRepository(client)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := router.SetupRouter(authHandler, client)
	router.Run(":8080")
}
