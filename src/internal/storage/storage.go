package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mg MongoInstance

// TODO: Create a login for server
const dbName = "birdai"
const mongoURI = "mongodb://localhost:27017"

// TODO: Create structs for different documents:
// 	Users, Admins, Birds, Posts, Sounds, Pictures

// Create interface with get and create a struct
// Example from Jesper

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

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

type resultDocument struct {
	ID        string             `bson:"_id"`
	Name      string             `bson:"username"`
	Auth      string             `bson:"auth_id"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

func Disconnect() {
	if mg.Client != nil {
		if err := mg.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB %s", err)
		}
	}
}

// TestGet Simple test function to get a user with the username user1
func TestGet() {
	if err := Connect(); err != nil {
		fmt.Println(err)
	}
	/*
		soundPath, _ := filepath.Abs("../BirdAI_D7057E/src/internal/storage/yippie.wav")
		fmt.Println(soundPath)
		soundData, err := os.ReadFile(soundPath)
		if err != nil {
			log.Fatal(err)
		}

		soundBson := Sound{
			ID:    primitive.NewObjectID(),
			Sound: soundData,
		}

		resID, err := mg.Db.Collection("sounds").InsertOne(context.TODO(), soundBson)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("inserted document with ID %v\n", resID.InsertedID)

		var result Sound
		fmt.Println(resID.InsertedID)
		filter := bson.M{"_id": resID.InsertedID}
		err = mg.Db.Collection("sounds").FindOne(context.Background(), filter).Decode(&result)

		output, err := os.Create("testSound.wav")
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()

		//soundReader := bytes.NewReader(result.Sound)
		//encoder := wav.NewEncoder(output, 44100, 16, 1, 1)
	*/

	/*

		imagePath, _ := filepath.Abs("../BirdAI_D7057E/src/internal/storage/testImage.png")
		fmt.Println(imagePath)
		imageData, err := os.ReadFile(imagePath)
		if err != nil {
			log.Fatal(err)
		}

		imageBson := Image{
			ID:    primitive.NewObjectID(),
			Image: imageData,
		}

		resID, err := mg.Db.Collection("images").InsertOne(context.TODO(), imageBson)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("inserted document with ID %v\n", resID.InsertedID)

		var result Image
		//objectID, err := primitive.ObjectIDFromHex(string(resID.InsertedID))
		fmt.Println(resID.InsertedID)
		filter := bson.M{"_id": resID.InsertedID}
		err = mg.Db.Collection("images").FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			log.Fatal("Error finding document", err)
		}
		//fmt.Println("Found document", result)
		fmt.Println("ID", result.ID)

		output, err := os.Create("testOutput.png")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer output.Close()

		imageReader := bytes.NewReader(result.Image)

		imageBin, _, err := image.Decode(imageReader)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = png.Encode(output, imageBin)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Kanske en bild!")

	*/

	//fmt.Println("Name", result.Image)
	/*
		var result resultDocument
		filter := bson.M{"name": "Skata"}
		err = mg.Db.Collection("birds").FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			log.Fatal("Error finding document", err)
		}
		fmt.Println("Found document", result)
		fmt.Println("ID", result.ID)
		fmt.Println("Name", result.Name)

		objectID, err := primitive.ObjectIDFromHex(result.ID)
		fmt.Println(objectID)
		filter = bson.M{"_id": objectID}
		err = mg.Db.Collection("images").FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			log.Fatal("Error finding document", err)
		}
		fmt.Println("Found document", result)
		fmt.Println("ID", result.ID)
		fmt.Println("Name", result.Name)

	*/
	Disconnect()
}
