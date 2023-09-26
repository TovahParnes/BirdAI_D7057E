package src

import (
	"birdai/src/internal/docs"
	"birdai/src/internal/routes"
	"birdai/src/internal/storage"
	"context"
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(ctx context.Context) (*fiber.App, error) {
	fiberApp := fiber.New()
	app := routes.New(fiberApp)

	//Need to use the inported docs package, useless line but needed
	docs.SwaggerInfo.Host = "localhost:3000"

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	log.Fatal(app.Listen(":3000"))
	storage.TestGet()

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:3000/swagger/oauth2-redirect.html",
	}))

	return app, nil
}
