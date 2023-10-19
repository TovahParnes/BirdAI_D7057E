package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BirdEndpoints is a struct to allow the birds collection in mongoDB to be accessed
type BirdEndpoints struct {
	Collection IMongoCollection
}

// GetBirdById returns a response containing the requested bird if found
func (u BirdEndpoints) GetBirdById(id string) models.Response {
	filter := bson.M{"_id": id}
	return u.Collection.FindOne(filter)
}

// GetBirdByName returns a response containing the requested bird if found
func (u BirdEndpoints) GetBirdByName(name string) models.Response {
	filter := bson.M{"name": name}
	return u.Collection.FindOne(filter)
}

// CreateBird creates a new bird and returns a response with the id given to the bird
func (u BirdEndpoints) CreateBird(bird models.BirdDB) models.Response {
	bird.Id = primitive.NewObjectID().Hex()
	return u.Collection.CreateOne(&bird)

}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// UpdateBird updates the bird with the specified changes and returns a response
// containing the updated bird.
func (u BirdEndpoints) UpdateBird(bird models.BirdInput) models.Response {
	data, err := bson.Marshal(bird)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonBird := bson.M{}
	err = bson.Unmarshal(data, &bsonBird)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return u.Collection.UpdateOne(bsonBird)
}

// ListBirds returns a response with a list for set n of size SetSize that matches the given filter
func (u BirdEndpoints) ListBirds(filter bson.M, set int) models.Response {
	return u.Collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllBirds returns a response with a list for set n of size SetSize
func (u BirdEndpoints) ListAllBirds(set int) models.Response {
	return u.Collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// DeleteBird returns the bird with the given id
// Returns the bird that was deleted
func (u BirdEndpoints) DeleteBird(id string) models.Response {
	return u.Collection.DeleteOne(bson.M{"_id": id})
}
