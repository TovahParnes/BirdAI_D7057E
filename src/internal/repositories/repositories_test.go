package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"testing"
)

// TestConnection tests connecting to the database and getting all collections.
//
// Needs a db with name "testDB" and all collections named under repositories_helpers.go
func TestConnection(t *testing.T) {
	t.Run("Test connect", func(t *testing.T) {
		mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
		require.Nil(t, err)
		mi.DisconnectDB()
	})
}

// TestRepository functions
func TestRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	repositories.AddAllCollections(mi)

	// Users collection will be used for testing
	userColl := mi.GetCollection(repositories.UserColl)

	// Birds collection will be used to test delete
	birdColl := mi.GetCollection(repositories.BirdColl)

	testUser1 := &models.UserDB{
		Id:       primitive.NewObjectID().Hex(),
		Username: "test User 1",
		AuthId:   "123123123",
		Active:   true,
	}
	testUser2 := &models.UserDB{
		Id:       primitive.NewObjectID().Hex(),
		Username: "test User 2",
		AuthId:   "123123",
		Active:   true,
	}

	testBird := &models.BirdDB{
		Id:   primitive.NewObjectID().Hex(),
		Name: "Skata",
	}

	t.Run("Test CreateOne", func(t *testing.T) {
		response := userColl.CreateOne(testUser1)
		require.Equal(t, testUser1.Id, response.Data.(string))
		response = userColl.CreateOne(testUser2)
		require.Equal(t, testUser2.Id, response.Data.(string))

		// Add a bird to be able to test delete later
		response = birdColl.CreateOne(testBird)
		require.Equal(t, testBird.Id, response.Data.(string))
	})

	t.Run("Test FindOne", func(t *testing.T) {
		filter := bson.M{"_id": testUser1.Id}
		response := userColl.FindOne(filter)
		require.Equal(t, testUser1.AuthId, response.Data.(*models.UserDB).AuthId)
		testUser1 = response.Data.(*models.UserDB)

		filter = bson.M{"_id": testUser2.Id}
		response = userColl.FindOne(filter)
		require.Equal(t, testUser2.AuthId, response.Data.(*models.UserDB).AuthId)
		testUser2 = response.Data.(*models.UserDB)

		filter = bson.M{"_id": "IncorrectID"}
		response = userColl.FindOne(filter)
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test FindAll", func(t *testing.T) {
		response := userColl.FindAll(bson.M{}, 0, 0)
		require.Equal(t, 2, len(response.Data.([]models.UserDB)))
		response = userColl.FindAll(bson.M{"_id": testUser2.Id}, 0, 0)
		require.Equal(t, 1, len(response.Data.([]models.UserDB)))
		response = userColl.FindAll(bson.M{"_id": "IncorrectId"}, 0, 0)
		require.Equal(t, 0, len(response.Data.([]models.UserDB)))
	})

	t.Run("Test UpdateOne", func(t *testing.T) {
		update := bson.M{
			"_id":      testUser1.Id,
			"username": "Test User 1 is the best!",
		}
		response := userColl.UpdateOne(update)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, update["username"], response.Data.(*models.UserDB).Username)

		response = userColl.UpdateOne(update)
		require.True(t, utils.IsTypeError(response))
	})
	// Delete both
	t.Run("Test DeleteOne", func(t *testing.T) {
		response := userColl.DeleteOne(bson.M{"_id": testUser1.Id})
		require.False(t, utils.IsTypeError(response))
		testUser1.Username = "Deleted User"
		testUser1.Active = false
		require.Equal(t, testUser1, response.Data.(*models.UserDB))

		response = birdColl.DeleteOne(bson.M{"_id": testBird.Id})
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, testBird, response.Data.(*models.BirdDB))
	})
	mi.DisconnectDB()

	// Need to delete everything from testDB
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	db := client.Database("testDB")
	_, err = db.Collection(repositories.BirdColl).DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return
	}
	_, err = db.Collection(repositories.UserColl).DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return
	}
}
