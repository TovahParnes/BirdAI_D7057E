package main

import (
	"birdai/src"
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// get env variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env: %v", err)
	}

	// setup global context
	ctx := context.Background()

	err := src.Setup(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	//for {
	//
	//	}
}
