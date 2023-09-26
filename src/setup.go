package src

import (
	"birdai/src/internal/docs"
	"birdai/src/internal/routes"
	"context"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func Setup(ctx context.Context) (*fiber.App, error) {
	fiberApp := fiber.New()
	app := routes.New(fiberApp)

	//these two lines needs to be there for swagger to fuction
	docs.SwaggerInfo.Host = "localhost:4000"
	app.Get("/swagger/*", swagger.HandlerDefault) // default


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
		OAuth2RedirectUrl: "http://localhost:4000/swagger/oauth2-redirect.html",
	}))

	return app, nil
}
