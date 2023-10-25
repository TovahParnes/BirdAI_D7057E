package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MediaEndpoints is a struct to allow the media collection in mongoDB to be accessed
type MediaEndpoints struct {
	Collection IMongoCollection
}

// GetMediaById returns a response containing the requested media item if found
func (m MediaEndpoints) GetMediaById(id string) models.Response {
	filter := bson.M{"_id": id}
	return m.Collection.FindOne(filter)
}

// CreateMedia creates a new media item and returns a response with the id given to the media item
func (m MediaEndpoints) CreateMedia(media models.MediaDB) models.Response {
	media.Id = primitive.NewObjectID().Hex()
	return m.Collection.CreateOne(&media)
}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// Might not be needed for media?

// UpdateMedia updates the media item with the specified changes and returns a response
// containing the updated media item.
func (m MediaEndpoints) UpdateMedia(media models.MediaInput) models.Response {
	data, err := bson.Marshal(media)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonMedia := bson.M{}
	err = bson.Unmarshal(data, &bsonMedia)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return m.Collection.UpdateOne(bsonMedia)
}

// Probably not needed?

// ListMedia returns a response with a list for set n of size SetSize that matches the given filter
func (m MediaEndpoints) ListMedia(filter bson.M, set int) models.Response {
	return m.Collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllMedia returns a response with a list for set n of size SetSize
func (m MediaEndpoints) ListAllMedia(set int) models.Response {
	return m.Collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// DeleteMedia returns the media item with the given id
// Returns the media item that was deleted
func (m MediaEndpoints) DeleteMedia(id string) models.Response {
	return m.Collection.DeleteOne(bson.M{"_id": id})
}
