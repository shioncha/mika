package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/shioncha/mika/backend/database"
	"github.com/shioncha/mika/backend/router"
)

func main() {
	godotenv.Load(".env")

	client := database.SetupClient()
	defer database.CloseClient(client)

	router := router.SetupRouter(client)
	router.Run(":8080")
}
