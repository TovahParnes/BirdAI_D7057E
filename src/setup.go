package src

import (
	"birdai/src/internal/authentication"
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"net/http"
)

func Setup(ctx context.Context) (*fiber.App, error) {
	// Setup authenticator
	auth, err := authentication.NewAuthentication()
	if err != nil {
		return nil, err
	}

	app := fiber.New()
	app.Use(cors.New())

	state, err := generateRandomState()
	if err != nil {
		return nil, err
	}
	store := session.New()
	// Save the state inside the session.
	app.Get("/login", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}
		// Set key/value
		sess.Set("state", state)
		if err := sess.Save(); err != nil {
			return c.SendString(err.Error())
		}
		return c.Redirect(auth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	})

	// Save the state inside the session.
	app.Get("/callback", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}
		if c.Query("state") != sess.Get("state") {
			return c.SendString("Invalid state parameter.")
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(c.Context(), c.Query("code"))
		if err != nil {
			return c.SendString("Failed to exchange an authorization code for a token.")
		}

		idToken, err := auth.VerifyIDToken(c.Context(), token)
		if err != nil {
			return c.SendString("Failed to verify ID Token.")
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			return c.SendString(err.Error())
		}

		sess.Set("access_token", token.AccessToken)
		sess.Set("profile", profile)
		if err := sess.Save(); err != nil {
			return c.SendString(err.Error())
		}

		// Redirect to logged in page.
		return c.Redirect("/user", http.StatusTemporaryRedirect)
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		profile := sess.Get("profile")
		return c.JSON(profile)
	})

	return app, nil
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
