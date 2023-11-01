package handlers

import (
	"birdai/src/internal/repositories"

	jwtware "github.com/gofiber/contrib/jwt"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(app *fiber.App, db repositories.RepositoryEndpoints) {
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
	handler.controller.GenerateBirds()

	// Add the JWTProtected() method if JTW key is required.
	usersRoute := app.Group("/users")
	usersRoute.Get("/list", handler.ListUsers)
	usersRoute.Get("/me", JWTProtected(), handler.GetUserMe)
	usersRoute.Get("/:id", handler.GetUserById)
	usersRoute.Post("/", handler.LoginUser)
	usersRoute.Patch("/:id", JWTProtected(), handler.UpdateUser)
	usersRoute.Delete("/:id", JWTProtected(), handler.DeleteUser)

	birdsRoute := app.Group("/birds")
	birdsRoute.Get("/list", handler.ListBirds)
	birdsRoute.Get("/:id", handler.GetBirdById)
	birdsRoute.Patch("/:id", JWTProtected(), handler.UpdateBird)

	postsRoute := app.Group("/posts")
	postsRoute.Get("/list", handler.ListPosts)
	usersRoute.Get("/:id/posts/list", handler.ListUsersPosts)
	postsRoute.Get("/:id", handler.GetPostById)
	postsRoute.Post("/", JWTProtected(), handler.CreatePost)
	postsRoute.Patch("/:id", JWTProtected(), handler.UpdatePost)
	postsRoute.Delete("/:id", JWTProtected(), handler.DeletePost)

	adminRoute := app.Group("/admins")
	adminRoute.Get("/list", JWTProtected(), handler.ListAdmins)
	adminRoute.Get("/me", JWTProtected(), handler.GetAdminMe)
	adminRoute.Post("/me", JWTProtected(), handler.CreateSuperadminMe)
	adminRoute.Get("/:id", JWTProtected(), handler.GetAdminById)
	adminRoute.Post("/", JWTProtected(), handler.CreateAdmin)
	adminRoute.Patch("/:id", JWTProtected(), handler.UpdateAdmin)
	adminRoute.Delete("/:id", JWTProtected(), handler.DeleteAdmin)

	aiRoute := app.Group("/ai")
	aiRoute.Post("/inputimage", JWTProtected(), handler.ImagePrediction)

}

func JWTProtected() func(*fiber.Ctx) error {
	// Create config for JWT authentication middleware.
	config := jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}

	return jwtware.New(config)
}
