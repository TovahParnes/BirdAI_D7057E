package handlers

import (
	"birdai/src/internal/repositories"

	jwtware "github.com/gofiber/contrib/jwt"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(app *fiber.App, db repositories.IMongoInstance) {
	app.Use(cors.New())
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
	usersRoute.Post("/", handler.LoginUser)
	usersRoute.Patch("/:id", JWTProtected(), handler.UpdateUser)
	usersRoute.Delete("/:id", JWTProtected(), handler.DeleteUser)

	birdsRoute := app.Group("/birds")
	birdsRoute.Get("/list", handler.ListBirds)
	birdsRoute.Get("/:id", handler.GetBirdById)
	birdsRoute.Patch("/:id", handler.UpdateBird)

	postsRoute := app.Group("/posts")
	postsRoute.Get("/list", handler.ListPosts)
	usersRoute.Get("/:id/posts/list", handler.ListUsersPosts)
	postsRoute.Get("/:id", handler.GetPostById)
	postsRoute.Post("/", handler.CreatePost)
	postsRoute.Patch("/:id", handler.UpdatePost)
	postsRoute.Delete("/:id", handler.DeletePost)

	adminRoute := app.Group("/admins")
	adminRoute.Get("/list", handler.ListAdmins)
	adminRoute.Get("/me", handler.GetAdminMe)
	adminRoute.Get("/:id", handler.GetAdminById)
	adminRoute.Post("/", handler.CreateAdmin)
	adminRoute.Patch("/:id", handler.UpdateAdmin)
	adminRoute.Delete("/:id", handler.DeleteAdmin)

}

func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}

	return jwtware.New(config)
}
