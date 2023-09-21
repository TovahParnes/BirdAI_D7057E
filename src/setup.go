package src

import (
	"birdai/src/internal/routes"
	"birdai/src/internal/storage"
	"context"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)


func Setup(ctx context.Context) error {
	fiberApp := fiber.New()
	app := routes.New(fiberApp)

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	log.Fatal(app.Listen(":3300"))
	storage.TestGet()

	return nil
}
