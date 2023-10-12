package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Create a login for Db
const dbName = "birdai"
const mongoURI = "mongodb://localhost:27017"

// Connect Connects to the db
//
// TODO: Check if connection needs any configurations
//
// TODO: Will be moved to other file
func Connect() (IMongoInstance, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	if err != nil {
		return nil, err
	}

	m := MongoInstance{
		Client:      client,
		Db:          db,
		Collections: map[string]IMongoCollection{},
	}
	return m, nil
}

// Collection functions

// AddAllCollections adds all collections to the MongoInstance
func AddAllCollections(m IMongoInstance) {
	m.AddCollection(UserColl)
	m.AddCollection(BirdColl)
	m.AddCollection(AdminColl)
	m.AddCollection(PostColl)
	m.AddCollection(MediaColl)
}

// UpdateOne returns the updated document as a HandlerObject.
// Input query should be of type bson.M with the id of the document to update
// and all fields with the values that should be updated
func (m *MongoCollection) UpdateOne(query bson.M) (models.Response) {
	filter := bson.M{
		"_id": query["_id"],
	}
	delete(query, "_id")
	update := bson.M{"$set": query}
	result, err := m.Collection.UpdateOne(m.ctx, filter, update)
	response := m.FindOne(filter)
	if result.ModifiedCount != 1 {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return response
}

// DeleteOne returns the deleted document if it gets successfully deleted.
// If the document was a user it will instead return the updated version of the document.
func (m *MongoCollection) DeleteOne(query bson.M) (models.Response) {
	deleteQuery := bson.M{
		"_id": query["_id"],
	}
	response := m.FindOne(deleteQuery)
	if utils.IsTypeError(response) {
		return response
	}
	switch response.Data.(type) {
	case *models.UserOutput, *models.UserInput, *models.UserDB:
		update := bson.M{
			"_id":      query["_id"],
			"active":   false,
			"username": "Deleted User",
		}
		response := m.UpdateOne(update)
		if utils.IsTypeError(response) {
			return utils.ErrorToResponse(400, "Could not delete object", response.Data.(models.Err).Description)
		}
		return response
	default:
		one, err := m.Collection.DeleteOne(m.ctx, deleteQuery)
		if one.DeletedCount != 1 || err != nil{
			return utils.ErrorToResponse(400, "Could not delete object", "")
		}
		return response
	}
}

// CreateOne adds the object to the db.
// Returns the id of the inserted document if successfully added.
func (m *MongoCollection) CreateOne(object models.HandlerObject) (models.Response) {
	object.SetCreatedAt()
	resId, err := m.Collection.InsertOne(m.ctx, object)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not create object", err.Error())
	}
	return utils.Response(resId.InsertedID.(string))
}

// TODO: Look if there is an easier way to handle the switch cases in both FindOne and FindAll

// FindOne searches for one document from the current collection that matches the query.
func (m *MongoCollection) FindOne(query bson.M) (models.Response) {
	collName := m.Collection.Name()
	// There should be a way to make this code more compact as it is repeating the same thing
	switch collName {
	case UserColl:
		var result models.UserOutput
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("User collection")
		}
		return utils.Response(&result)
	case AdminColl:
		var result models.AdminOutput
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Admin collection")
		}
		return utils.Response(&result)
	case BirdColl:
		var result models.BirdOutput
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Bird collection")
		}
		return utils.Response(&result)
	case PostColl:
		var result models.PostOutput
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Post collection")
		}
		return utils.Response(&result)
	case MediaColl:
		var result models.MediaOutput
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Media collection")
		}
		return utils.Response(&result)
	default:
		// Collection name was not found
		return utils.ErrorCollectionNotFound("FindOne")
	}
}

// FindAll returns all documents from the current collection.
// Return will be of type interface{} and will need to be type asserted before use
func (m *MongoCollection) FindAll() (models.Response) {
	findCursor, err := m.Collection.Find(m.ctx, bson.M{})
	if err != nil {
		return utils.ErrorNotFoundInDatabase("")
	}
	var results []bson.M
	if err = findCursor.All(m.ctx, &results); err != nil {
		return utils.ErrorNotFoundInDatabase("")
	}
	collName := m.Collection.Name()
	// There should be a way to make this code more compact as it is repeating the same thing
	switch collName {
	case UserColl:
		var resultStruct []models.UserOutput
		for _, result := range results {
			var tempResult models.UserOutput
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("User collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case AdminColl:
		var resultStruct []models.AdminOutput
		for _, result := range results {
			var tempResult models.AdminOutput
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Admin collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case BirdColl:
		var resultStruct []models.BirdOutput
		for _, result := range results {
			var tempResult models.BirdOutput
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Bird collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case PostColl:
		var resultStruct []models.PostOutput
		for _, result := range results {
			var tempResult models.PostOutput
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Post collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case MediaColl:
		var resultStruct []models.MediaOutput
		for _, result := range results {
			var tempResult models.MediaOutput
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Media collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	default:
		// Collection name was not found
		return utils.ErrorCollectionNotFound("FindAll")
	}
}
