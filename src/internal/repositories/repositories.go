package repositories

import (
	"birdai/src/internal/models"
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
func (m *MongoCollection) UpdateOne(query bson.M) (models.HandlerObject, error) {
	filter := bson.M{
		"_id": query["_id"],
	}
	delete(query, "_id")
	update := bson.M{"$set": query}
	result, err := m.Collection.UpdateOne(m.ctx, filter, update)
	updated, err := m.FindOne(filter)
	if result.ModifiedCount != 1 {
		return nil, err
	}
	return updated, err
}

// DeleteOne returns the deleted document if it gets successfully deleted.
// If the document was a user it will instead return the updated version of the document.
func (m *MongoCollection) DeleteOne(query bson.M) (models.HandlerObject, error) {
	deleteQuery := bson.M{
		"_id": query["_id"],
	}
	data, err := m.FindOne(deleteQuery)
	if err != nil {
		return nil, err
	}
	switch data.(type) {
	case *models.User:
		update := bson.M{
			"_id":      query["_id"],
			"active":   false,
			"username": "Deleted User",
		}
		updated, err := m.UpdateOne(update)
		if err != nil {
			return nil, err
		}
		return updated, nil
	default:
		one, err := m.Collection.DeleteOne(m.ctx, deleteQuery)
		if one.DeletedCount != 1 {
			return nil, err
		}
		return data, nil
	}
}

// CreateOne adds the object to the db.
// Returns the id of the inserted document if successfully added.
func (m *MongoCollection) CreateOne(object models.HandlerObject) (string, error) {
	object.SetCreatedAt()
	resId, err := m.Collection.InsertOne(m.ctx, object)
	if err != nil {
		return "", err
	}
	return resId.InsertedID.(string), nil
}

// TODO: Look if there is an easier way to handle the switch cases in both FindOne and FindAll

// FindOne searches for one document from the current collection that matches the query.
func (m *MongoCollection) FindOne(query bson.M) (models.HandlerObject, error) {
	collName := m.Collection.Name()
	// There should be a way to make this code more compact as it is repeating the same thing
	switch collName {
	case UserColl:
		var result models.User
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, err
	case AdminColl:
		var result models.Admin
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, err
	case BirdColl:
		var result models.Bird
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, err
	case PostColl:
		var result models.Post
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, err
	case MediaColl:
		var result models.Media
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return nil, err
		}
		return &result, err
	default:
		// Collection name was not found
		return nil, nil
	}
}

// FindAll returns all documents from the current collection.
// Return will be of type interface{} and will need to be type asserted before use
func (m *MongoCollection) FindAll() (interface{}, error) {
	findCursor, err := m.Collection.Find(m.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var results []bson.M
	if err = findCursor.All(m.ctx, &results); err != nil {
		return nil, err
	}
	collName := m.Collection.Name()
	// There should be a way to make this code more compact as it is repeating the same thing
	switch collName {
	case UserColl:
		var resultStruct []models.User
		for _, result := range results {
			var tempResult models.User
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return nil, err
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return resultStruct, nil
	case BirdColl:
		var resultStruct []models.Bird
		for _, result := range results {
			var tempResult models.Bird
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return nil, err
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return resultStruct, nil
	case PostColl:
		var resultStruct []models.Post
		for _, result := range results {
			var tempResult models.Post
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return nil, err
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return resultStruct, nil
	case MediaColl:
		var resultStruct []models.Media
		for _, result := range results {
			var tempResult models.Media
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return nil, err
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return resultStruct, nil
	default:
		// Collection name was not found
		return nil, nil
	}
}
