package handlers

import (
	"birdai/src/internal/repositories"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func New(app *fiber.App, db repositories.IMongoInstance) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
	}))
	//app.Use(jwtware.New(jwtware.Config{
	//	SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	//}))
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "UTC",
	}))

	handler := NewHandler(db)

	// Add the JWTProtected() method if JTW key is required.
	usersRoute := app.Group("/users")
	usersRoute.Get("/list", JWTProtected(), handler.ListUsers)
	usersRoute.Get("/me", JWTProtected(), handler.GetUserMe)
	usersRoute.Get("/:id", JWTProtected(), handler.GetUserById)
	usersRoute.Post("/", JWTProtected(), handler.Login)
	usersRoute.Patch("/:id", JWTProtected(), handler.UpdateUser)
	usersRoute.Delete("/:id", JWTProtected(), handler.DeleteUser)
}

func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}

	return jwtware.New(config)
}
