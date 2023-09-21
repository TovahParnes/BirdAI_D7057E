package src

import (
	"birdai/src/internal/routes"
	"birdai/src/internal/storage"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)


func Setup(ctx context.Context) error {
	fiberApp := fiber.New()
	app := routes.New(fiberApp)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3300/swagger/oauth2-redirect.html",
	}))

	log.Fatal(app.Listen(":3300"))
	storage.TestGet()

	return nil
}
