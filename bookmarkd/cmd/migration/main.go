package main

import (
	"context"
	"log"

	"github.com/dolldot/scrapyard/bookmarkd/config"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent/migrate"
	"github.com/dolldot/scrapyard/bookmarkd/internal/clients"
)

func main() {
	config.InitializeConfig()

	db := clients.NewEntclient()
	defer db.Close()

	ctx := context.Background()

	err := db.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("Succesfully run the migration")
}
