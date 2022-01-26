package main

import (
	"bitsports/config"
	"bitsports/ent"
	"bitsports/ent/migrate"
	"bitsports/pkg/database"
	"context"
	"log"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client, err := database.NewClient()
	if err != nil {
		log.Fatalf("failed opening mysql client: %v", err)
	}
	defer client.Close()
	createDBSchema(client)
}

func createDBSchema(client *ent.Client) {
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

