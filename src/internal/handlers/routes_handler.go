package handlers

import (
	"birdai/src/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(app *fiber.App, db models.IMongoInstance) {
	app.Use(cors.New())
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
	usersRoute.Post("/", handler.CreateUser)
	usersRoute.Patch("/:id", handler.UpdateUser)
	usersRoute.Delete("/:id", handler.DeleteUser)
}
