package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository is a struct to allow the user collection in mongoDB to be accessed
type UserRepository struct {
	collection IMongoCollection
}

func (u *UserRepository) SetCollection(coll IMongoCollection) {
	u.collection = coll
}

// GetUserById returns a response containing the requested user if found
func (u *UserRepository) GetUserById(id string) models.Response {
	filter := bson.M{"_id": id}
	return u.collection.FindOne(filter)
}

// GetUserByAuthId returns a response containing the requested user if found
func (u *UserRepository) GetUserByAuthId(authId string) models.Response {
	filter := bson.M{"auth_id": authId}
	return u.collection.FindOne(filter)
}

// CreateUser creates a new user and returns a response with the id given to the user
func (u *UserRepository) CreateUser(user models.UserDB) models.Response {
	user.Id = primitive.NewObjectID().Hex()
	return u.collection.CreateOne(&user)

}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// UpdateUser updates the user with the specified changes and returns a response
// containing the updated user.
func (u *UserRepository) UpdateUser(id string, user models.UserInput) models.Response {
	data, err := bson.Marshal(user)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonUser := bson.M{}
	err = bson.Unmarshal(data, &bsonUser)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonUser["_id"] = id
	return u.collection.UpdateOne(bsonUser)
}

// ListUsers returns a response with a list for set n of size SetSize that matches the given filter
func (u *UserRepository) ListUsers(filter bson.M, set int) models.Response {
	return u.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllUsers returns a response with a list for set n of size SetSize
func (u *UserRepository) ListAllUsers(set int) models.Response {
	return u.collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// DeleteUser returns the user with the given id
// Returns the user that was deleted
func (u *UserRepository) DeleteUser(id string) models.Response {
	return u.collection.DeleteOne(bson.M{"_id": id})
}
