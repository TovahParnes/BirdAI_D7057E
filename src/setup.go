package src

import (
	"birdai/src/internal/routes"
	"birdai/src/internal/storage"
	"context"
	"log"
)

func Setup(ctx context.Context) error {
	app := routes.New()
	log.Fatal(app.Listen(":3300"))
	storage.TestGet()

	return nil
}
