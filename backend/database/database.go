package database

import (
	"context"
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

	connectionString := "host=" + host + " port=" + port + " user=" + username + " dbname=" + dbname + " password=" + password + " sslmode=disable"

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client
}

func CloseClient(client *ent.Client) {
	client.Close()
}
