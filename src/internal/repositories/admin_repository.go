package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AdminRepository is a struct to allow the admin collection in mongoDB to be accessed
type AdminRepository struct {
	collection IMongoCollection
}

func (a *AdminRepository) SetCollection(coll IMongoCollection) {
	a.collection = coll
}

// GetAdminById returns a response containing the requested admin if found
func (a *AdminRepository) GetAdminById(id string) models.Response {
	filter := bson.M{"_id": id}
	return a.collection.FindOne(filter)
}

// GetAdminByUserId returns a response containing the requested admin if found
func (a *AdminRepository) GetAdminByUserId(userId string) models.Response {
	filter := bson.M{"user_id": userId}
	return a.collection.FindOne(filter)
}

// ListAdmins returns a response with a list for set n of size SetSize that matches the given filter
func (a *AdminRepository) ListAdmins(filter bson.M, set int) models.Response {
	return a.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllAdmins returns a response with a list for set n of size SetSize
func (a *AdminRepository) ListAllAdmins(set int) models.Response {
	return a.collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// CreateAdmin creates a new admin and returns a response with the id given to the admin
func (a *AdminRepository) CreateAdmin(admin models.AdminDB) models.Response {
	admin.Id = primitive.NewObjectID().Hex()
	return a.collection.CreateOne(&admin)
}

// UpdateAdmin updates the admin with the specified changes and returns a response
// containing the updated admin.
func (a *AdminRepository) UpdateAdmin(admin models.AdminInput) models.Response {
	data, err := bson.Marshal(admin)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonAdmin := bson.M{}
	err = bson.Unmarshal(data, &bsonAdmin)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return a.collection.UpdateOne(bsonAdmin)
}

// DeleteAdmin returns the admin with the given id
// Returns the admin that was deleted
func (a *AdminRepository) DeleteAdmin(id string) models.Response {
	return a.collection.DeleteOne(bson.M{"_id": id})
}
