package repositories

import (
	"birdai/src/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mg models.MongoInstance

// TODO: Create a login for Db
const dbName = "birdai"
const mongoURI = "mongodb://localhost:27017"

// Connect TODO: Check if connection needs any configurations
// Connect Connects to the db
func Connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))

	if err != nil {
		return err
	}

	db := client.Database(dbName)

	if err != nil {
		return err
	}

	mg = models.MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

// Disconnect Disconnects from the db
func Disconnect() {
	if mg.Client != nil {
		if err := mg.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB %s", err)
		}
	}
}

// TestConnect Simple function that connects and disconnects from the db
func TestConnect() {

	if err := Connect(); err != nil {
		fmt.Println(err)
	}
	Disconnect()
}
