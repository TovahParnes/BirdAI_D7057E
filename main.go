package main

import (
	"birdai/src"
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

//	@title			BirdAI API
//	@version		2.0
//	@description	A server for BirdAI API, for managing users, admin, birds, posts and more.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Tovah Parnes
//	@contact.email	??

//	@license.name	MIT License
//	@license.url	https://opensource.org/license/mit/

//	@host		localhost:4000
//	@BasePath	/
//	@schemes	http
func main() {
	// get env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	// setup global context
	ctx := context.Background()

	app, err := src.Setup(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.Fatal(app.Listen(":4000"))
}
