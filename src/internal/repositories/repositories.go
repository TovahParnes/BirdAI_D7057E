package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TODO: Create a login for Db
//const DBName = "birdai"
//const MongoURI = "mongodb://localhost:27017"

// Connect Connects to the db
//
// TODO: Check if connection needs any configurations
//
// TODO: Might be moved to other file

func Connect(dbName, mongoURI string) (IMongoInstance, error) {
	opts := options.Client()
	opts.ApplyURI(mongoURI)
	opts.SetConnectTimeout(10 * time.Second)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 11*time.Second)
	defer cancel()
	// Ping to check the connection
	err = client.Ping(ctx, nil)
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

// SetupRepositories return a RepositoryEndpoints struct that allows access to the different repositories and functions
// Needs to swap out, so it uses env variables instead of set names
func SetupRepositories() (RepositoryEndpoints, error) {
	mongoInstance, err := Connect(os.Getenv("DB_NAME"), os.Getenv("MONGO_URI"))
	if err != nil {
		return RepositoryEndpoints{}, err
	}
	AddAllCollections(mongoInstance)
	user := UserRepository{}
	user.SetCollection(mongoInstance.GetCollection(UserColl))
	post := PostRepository{}
	post.SetCollection(mongoInstance.GetCollection(PostColl))
	bird := BirdRepository{}
	bird.SetCollection(mongoInstance.GetCollection(BirdColl))
	media := MediaRepository{}
	media.SetCollection(mongoInstance.GetCollection(MediaColl))
	admin := AdminRepository{}
	admin.SetCollection(mongoInstance.GetCollection(AdminColl))
	return RepositoryEndpoints{
		User:  user,
		Post:  post,
		Bird:  bird,
		Media: media,
		Admin: admin,
	}, nil
}

// collection functions

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
func (m *MongoCollection) UpdateOne(query bson.M) models.Response {
	filter := bson.M{
		"_id": query["_id"],
	}
	delete(query, "_id")
	update := bson.M{"$set": query}
	result, err := m.Collection.UpdateOne(m.ctx, filter, update)
	response := m.FindOne(filter)
	if utils.IsTypeError(response) {
		return response
	}
	if result.ModifiedCount != 1 {
		// Needs to check if there is an error, if not the update was a "success" but there was no change needed
		if err != nil {
			return utils.ErrorToResponse(400, "Could not update object", err.Error())
		}
		return utils.ErrorToResponse(400, "Could not update object", "No change compared to current document")
	}
	return response
}

// DeleteOne returns the deleted document if it gets successfully deleted.
// If the document was a user it will instead return the updated version of the document.
func (m *MongoCollection) DeleteOne(query bson.M) models.Response {
	// Safety check so we don't remove based on other than ID
	// Should probably be moved to case specific
	deleteQuery := bson.M{
		"_id": query["_id"],
	}
	response := m.FindOne(deleteQuery)
	if utils.IsTypeError(response) {
		return response
	}
	switch response.Data.(type) {
	case *models.UserDB:
		update := bson.M{
			"_id":      query["_id"],
			"active":   false,
			"username": "Deleted User",
		}
		response := m.UpdateOne(update)
		if utils.IsTypeError(response) {
			return utils.ErrorToResponse(400, "Could not delete object", response.Data.(models.Err).Description)
		}
		return utils.Response("Deactivated user successfully ")
	default:
		one, err := m.Collection.DeleteOne(m.ctx, deleteQuery)
		if one.DeletedCount != 1 || err != nil {
			return utils.ErrorToResponse(400, "Could not delete object", "")
		}
		return utils.Response("Deleted successfully")
	}
}

// CreateOne adds the object to the db.
// Returns the id of the inserted document if successfully added.
func (m *MongoCollection) CreateOne(object models.HandlerObject) models.Response {
	object.SetCreatedAt()
	resId, err := m.Collection.InsertOne(m.ctx, object)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not create object", err.Error())
	}
	return utils.Response(resId.InsertedID.(string))
}

// TODO: Look if there is an easier way to handle the switch cases in both FindOne and FindAll

// FindOne searches for one document from the current collection that matches the query.
func (m *MongoCollection) FindOne(query bson.M) models.Response {
	collName := m.Collection.Name()
	switch collName {
	case UserColl:
		var result models.UserDB
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("User collection")
		}
		return utils.Response(&result)
	case AdminColl:
		var result models.AdminDB
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Admin collection")
		}
		return utils.Response(&result)
	case BirdColl:
		var result models.BirdDB
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Bird collection")
		}
		return utils.Response(&result)
	case PostColl:
		var result models.PostDB
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Post collection")
		}
		return utils.Response(&result)
	case MediaColl:
		var result models.MediaDB
		err := m.Collection.FindOne(m.ctx, query).Decode(&result)
		if err != nil {
			return utils.ErrorNotFoundInDatabase("Media collection")
		}
		return utils.Response(&result)
	default:
		// collection name was not found
		return utils.ErrorCollectionNotFound("FindOne")
	}
}

// FindAll returns all documents from the current Collection that matches the filter.
// The input limit will limit how many documents will be returned
// The input skip will skip the first n results
// Return will be of type interface{} and will need to be type asserted before use
func (m *MongoCollection) FindAll(filter bson.M, limit int, skip int) models.Response {
	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findCursor, err := m.Collection.Find(m.ctx, filter, findOptions)
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
		var resultStruct []models.UserDB
		for _, result := range results {
			var tempResult models.UserDB
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("User collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case AdminColl:
		var resultStruct []models.AdminDB
		for _, result := range results {
			var tempResult models.AdminDB
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Admin collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case BirdColl:
		var resultStruct []models.BirdDB
		for _, result := range results {
			var tempResult models.BirdDB
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Bird collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case PostColl:
		var resultStruct []models.PostDB
		for _, result := range results {
			var tempResult models.PostDB
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Post collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	case MediaColl:
		var resultStruct []models.MediaDB
		for _, result := range results {
			var tempResult models.MediaDB
			bsonBody, _ := bson.Marshal(result)
			err := bson.Unmarshal(bsonBody, &tempResult)
			if err != nil {
				return utils.ErrorNotFoundInDatabase("Media collection")
			}
			resultStruct = append(resultStruct, tempResult)
		}
		return utils.Response(resultStruct)
	default:
		// collection name was not found
		return utils.ErrorCollectionNotFound("FindAll")
	}
}
