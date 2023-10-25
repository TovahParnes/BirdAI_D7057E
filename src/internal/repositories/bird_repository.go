package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BirdRepository is a struct to allow the birds collection in mongoDB to be accessed
type BirdRepository struct {
	collection IMongoCollection
}

func (b *BirdRepository) SetCollection(coll IMongoCollection) {
	b.collection = coll
}

// GetBirdById returns a response containing the requested bird if found
func (b *BirdRepository) GetBirdById(id string) models.Response {
	filter := bson.M{"_id": id}
	return b.collection.FindOne(filter)
}

// GetBirdByName returns a response containing the requested bird if found
func (b *BirdRepository) GetBirdByName(name string) models.Response {
	filter := bson.M{"name": name}
	return b.collection.FindOne(filter)
}

// CreateBird creates a new bird and returns a response with the id given to the bird
func (b *BirdRepository) CreateBird(bird models.BirdDB) models.Response {
	bird.Id = primitive.NewObjectID().Hex()
	return b.collection.CreateOne(&bird)

}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// UpdateBird updates the bird with the specified changes and returns a response
// containing the updated bird.
func (b *BirdRepository) UpdateBird(bird models.BirdInput) models.Response {
	data, err := bson.Marshal(bird)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonBird := bson.M{}
	err = bson.Unmarshal(data, &bsonBird)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return b.collection.UpdateOne(bsonBird)
}

// ListBirds returns a response with a list for set n of size SetSize that matches the given filter
func (b *BirdRepository) ListBirds(filter bson.M, set int) models.Response {
	return b.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllBirds returns a response with a list for set n of size SetSize
func (b *BirdRepository) ListAllBirds(set int) models.Response {
	return b.collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// DeleteBird returns the bird with the given id
// Returns the bird that was deleted
func (b *BirdRepository) DeleteBird(id string) models.Response {
	return b.collection.DeleteOne(bson.M{"_id": id})
}
