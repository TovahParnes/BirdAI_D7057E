package handlers

import (
	"birdai/src/internal/repositories"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(app *fiber.App, db repositories.IMongoInstance) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	handler := NewHandler(db)

	usersRoute := app.Group("/users")
	usersRoute.Get("/list", handler.ListUsers)
	usersRoute.Get("/me", handler.GetUserMe)
	usersRoute.Get("/:id", handler.GetUserById)
	usersRoute.Post("/", handler.Login)
	usersRoute.Patch("/:id", handler.UpdateUser)
	usersRoute.Delete("/:id", handler.DeleteUser)

	birdsRoute := app.Group("/birds")
	birdsRoute.Get("/list", handler.ListBirds)
	birdsRoute.Get("/:id", handler.GetBirdById)
	birdsRoute.Patch("/:id", handler.UpdateBird)
}
