package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserEndpoints is a struct to allow the user collection in mongoDB to be accessed
type UserEndpoints struct {
	Collection IMongoCollection
}

// GetUserById returns a response containing the requested user if found
func (u UserEndpoints) GetUserById(id string) models.Response {
	filter := bson.M{"_id": id}
	return u.Collection.FindOne(filter)
}

// GetUserByAuthId returns a response containing the requested user if found
func (u UserEndpoints) GetUserByAuthId(authId string) models.Response {
	filter := bson.M{"auth_id": authId}
	return u.Collection.FindOne(filter)
}

// CreateUser creates a new user and returns a response with the id given to the user
func (u UserEndpoints) CreateUser(user models.UserDB) models.Response {
	user.Id = primitive.NewObjectID().Hex()
	return u.Collection.CreateOne(&user)

}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// UpdateUser updates the user with the specified changes and returns a response
// containing the updated user.
func (u UserEndpoints) UpdateUser(user models.UserInput) models.Response {
	data, err := bson.Marshal(user)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonUser := bson.M{}
	err = bson.Unmarshal(data, &bsonUser)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return u.Collection.UpdateOne(bsonUser)
}

// ListUsers returns a response with a list for set n of size SetSize that matches the given filter
func (u UserEndpoints) ListUsers(filter bson.M, set int) models.Response {
	return u.Collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllUsers returns a response with a list for set n of size SetSize
func (u UserEndpoints) ListAllUsers(set int) models.Response {
	return u.Collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// DeleteUser returns the user with the given id
// Returns the user that was deleted
func (u UserEndpoints) DeleteUser(id string) models.Response {
	return u.Collection.DeleteOne(bson.M{"_id": id})
}
