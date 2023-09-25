package routes

import (
	"birdai/src/internal/handlers/users_handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

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
	usersRoute.Get("/list", users_handler.ListUsers)
	usersRoute.Get("/:id", users_handler.GetUserById)
	usersRoute.Get("/me", users_handler.GetUserMe)
	usersRoute.Post("/", users_handler.CreateUser)
	


	return app
}
