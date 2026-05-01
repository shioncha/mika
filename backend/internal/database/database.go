package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shioncha/mika/backend/ent"
)

func SetupClient() *ent.Client {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	username := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")

	sslmode := "require"
	isDev := os.Getenv("ENVIRONMENT") == "dev"
	if isDev {
		sslmode = "disable"
	}

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbname, password, sslmode,
	)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	if isDev {
		log.Println("Running in development mode, auto-migrating database schema...")
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return client
}

func CloseClient(client *ent.Client) {
	client.Close()
}
