package routes

import (
	users "birdai/src/internal/handlers"
	"birdai/src/internal/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(app *fiber.App, db storage.IMongoInstance) {
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

	handler := users.NewHandler(db)

	usersRoute := app.Group("/users")
	usersRoute.Get("/list", handler.ListUsers)
	usersRoute.Get("/:id", handler.GetUserById)
	usersRoute.Get("/me", handler.GetUserMe)
	usersRoute.Post("/", handler.CreateUser)
	usersRoute.Patch("/:id", handler.UpdateUser)
	usersRoute.Delete("/:id", handler.DeleteUser)
}
