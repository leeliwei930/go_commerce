package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leeliwei930/go_commerce/ent/migrate"
	"github.com/leeliwei930/go_commerce/services"
)

func MigrateDatabase() {

	dbClient := services.GetDefaultDBClient()
	defer dbClient.Close()
	ctx := context.Background()

	// start migration
	if migrateErr := dbClient.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); migrateErr != nil {
		log.Fatalf("Unable to perform database migration: %v", migrateErr)
	}
}
