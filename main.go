package main

import (
	"birdai/src"
	"context"
	"fmt"
	"log"
	"os"

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

// @host		127.0.0.1:3000
// @BasePath	/
// @schemes	http
func main() {
	// get env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	
	// get JWT secret variable
	if err := godotenv.Load("secret/.env"); err != nil {
		log.Fatalf("Failed to load secret env: %v", err)
	}

	// setup global context
	ctx := context.Background()

	app, err := src.Setup(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Port used: " + os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))

}
