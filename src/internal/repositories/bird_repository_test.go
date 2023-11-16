package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TestRepository functions
func TestBirdRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.BirdColl)
	birdColl := repositories.BirdRepository{}
	birdColl.SetCollection(mi.GetCollection(repositories.BirdColl))

	testBird1 := &models.BirdDB{
		Name:        "Bird 1",
		Description: "This is bird 1",
		ImageId:     primitive.NewObjectID().Hex(),
		SoundId:     primitive.NewObjectID().Hex(),
	}
	testBird2 := &models.BirdDB{
		Name:        "Bird 2",
		Description: "This is bird 2",
		ImageId:     primitive.NewObjectID().Hex(),
		SoundId:     primitive.NewObjectID().Hex(),
	}

	t.Run("Test CreateBird", func(t *testing.T) {
		response := birdColl.CreateBird(*testBird1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testBird1.Id = response.Data.(string)
		response = birdColl.CreateBird(*testBird2)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testBird2.Id = response.Data.(string)
	})

	t.Run("Test GetBirdById", func(t *testing.T) {
		response := birdColl.GetBirdById(testBird1.GetId())
		require.Equal(t, testBird1.Id, response.Data.(*models.BirdDB).Id)
		testBird1 = response.Data.(*models.BirdDB)

		response = birdColl.GetBirdById("IncorrectID")
		require.True(t, utils.IsTypeError(response))

	})

	t.Run("Test GetBirdByName", func(t *testing.T) {
		response := birdColl.GetBirdByName(testBird2.Name)
		require.Equal(t, testBird2.Name, response.Data.(*models.BirdDB).Name)
		testBird2 = response.Data.(*models.BirdDB)

		response = birdColl.GetBirdByName("Incorrect Name")
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test ListBirds & ListAllBirds", func(t *testing.T) {
		response := birdColl.ListBirds(bson.M{}, 0)
		require.Equal(t, 2, len(response.Data.([]models.BirdDB)))
		responseAll := birdColl.ListAllBirds(0)
		require.Equal(t, response.Data.([]models.BirdDB), responseAll.Data.([]models.BirdDB))
		response = birdColl.ListBirds(bson.M{"_id": testBird2.Id}, 0)
		require.Equal(t, 1, len(response.Data.([]models.BirdDB)))
		response = birdColl.ListBirds(bson.M{"_id": "IncorrectId"}, 0)
		require.Equal(t, 0, len(response.Data.([]models.BirdDB)))
	})

	t.Run("Test UpdateOne", func(t *testing.T) {
		updateBird := models.BirdInput{
			Description: "Test Bird 1 is the best!!",
		}
		response := birdColl.UpdateBird(testBird1.Id, updateBird)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, updateBird.Name, response.Data.(*models.BirdDB).Name)

		response = birdColl.UpdateBird(testBird1.Id, updateBird)
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test DeleteOne", func(t *testing.T) {
		response := birdColl.DeleteBird(testBird1.GetId())
		require.False(t, utils.IsTypeError(response))
		response = birdColl.DeleteBird(testBird1.GetId())
		require.True(t, utils.IsTypeError(response))

	})

	// Need to delete everything from testDB
	t.Cleanup(func() {
		mi.DisconnectDB()
		client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			return
		}
		db := client.Database("testDB")
		_, err = db.Collection(repositories.BirdColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
	})
}
