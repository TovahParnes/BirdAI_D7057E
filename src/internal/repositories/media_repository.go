package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 79e1b5d (implemented media repository)
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
// MediaRepository is a struct to allow the media collection in mongoDB to be accessed
type MediaRepository struct {
	collection IMongoCollection
}

func (m *MediaRepository) SetCollection(coll IMongoCollection) {
	m.collection = coll
<<<<<<< HEAD
}

// GetMediaById returns a response containing the requested media item if found
func (m *MediaRepository) GetMediaById(id string) models.Response {
	filter := bson.M{"_id": id}
	return m.collection.FindOne(filter)
}

// CreateMedia creates a new media item and returns a response with the id given to the media item
func (m *MediaRepository) CreateMedia(media models.MediaDB) models.Response {
	media.Id = primitive.NewObjectID().Hex()
	return m.collection.CreateOne(&media)
=======
// MediaEndpoints is a struct to allow the media collection in mongoDB to be accessed
type MediaEndpoints struct {
	Collection IMongoCollection
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// GetMediaById returns a response containing the requested media item if found
func (m *MediaRepository) GetMediaById(id string) models.Response {
	filter := bson.M{"_id": id}
	return m.collection.FindOne(filter)
}

// CreateMedia creates a new media item and returns a response with the id given to the media item
func (m *MediaRepository) CreateMedia(media models.MediaDB) models.Response {
	media.Id = primitive.NewObjectID().Hex()
<<<<<<< HEAD
	return m.Collection.CreateOne(&media)
>>>>>>> 79e1b5d (implemented media repository)
=======
	return m.collection.CreateOne(&media)
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// Might not be needed for media?

// UpdateMedia updates the media item with the specified changes and returns a response
// containing the updated media item.
<<<<<<< HEAD
<<<<<<< HEAD
func (m *MediaRepository) UpdateMedia(media models.MediaInput) models.Response {
	data, err := bson.Marshal(media)
	if err != nil {
=======
func (m MediaEndpoints) UpdateMedia(media models.MediaInput) models.Response {
	data, err := bson.Marshal(media)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
>>>>>>> 79e1b5d (implemented media repository)
=======
func (m *MediaRepository) UpdateMedia(media models.MediaInput) models.Response {
	data, err := bson.Marshal(media)
	if err != nil {
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonMedia := bson.M{}
	err = bson.Unmarshal(data, &bsonMedia)
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return m.collection.UpdateOne(bsonMedia)
=======
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return m.Collection.UpdateOne(bsonMedia)
>>>>>>> 79e1b5d (implemented media repository)
=======
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return m.collection.UpdateOne(bsonMedia)
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// Probably not needed?

// ListMedia returns a response with a list for set n of size SetSize that matches the given filter
<<<<<<< HEAD
<<<<<<< HEAD
func (m *MediaRepository) ListMedia(filter bson.M, set int) models.Response {
	return m.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllMedia returns a response with a list for set n of size SetSize
func (m *MediaRepository) ListAllMedia(set int) models.Response {
	return m.collection.FindAll(bson.M{}, SetSize, SetSize*set)
=======
func (m MediaEndpoints) ListMedia(filter bson.M, set int) models.Response {
	return m.Collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllMedia returns a response with a list for set n of size SetSize
func (m MediaEndpoints) ListAllMedia(set int) models.Response {
	return m.Collection.FindAll(bson.M{}, SetSize, SetSize*set)
>>>>>>> 79e1b5d (implemented media repository)
=======
func (m *MediaRepository) ListMedia(filter bson.M, set int) models.Response {
	return m.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllMedia returns a response with a list for set n of size SetSize
func (m *MediaRepository) ListAllMedia(set int) models.Response {
	return m.collection.FindAll(bson.M{}, SetSize, SetSize*set)
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// DeleteMedia returns the media item with the given id
// Returns the media item that was deleted
<<<<<<< HEAD
<<<<<<< HEAD
func (m *MediaRepository) DeleteMedia(id string) models.Response {
	return m.collection.DeleteOne(bson.M{"_id": id})
=======
func (m MediaEndpoints) DeleteMedia(id string) models.Response {
	return m.Collection.DeleteOne(bson.M{"_id": id})
>>>>>>> 79e1b5d (implemented media repository)
=======
func (m *MediaRepository) DeleteMedia(id string) models.Response {
	return m.collection.DeleteOne(bson.M{"_id": id})
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}
