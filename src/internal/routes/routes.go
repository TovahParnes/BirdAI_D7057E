package routes

import (
	"birdai/src/internal/handlers/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// New create an instance of Book app routes
func New(app *fiber.App) *fiber.App {
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))
	/*
	swaggerCfg := swagger.Config{
		BasePath: "/docs", //swagger ui base path
		FilePath: "./docs/swagger.json",
	}

	app.Use(swagger.New(swaggerCfg))
	*/

	usersRoute := app.Group("/users")
	usersRoute.Get("/set=:set<int>", users.GetAllUsers)
	usersRoute.Get("/:id", users.GetUserById)
	usersRoute.Post("/me", users.GetUserMe)


	return app
}