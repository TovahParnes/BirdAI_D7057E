package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

// TestRepository functions
func TestMediaRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.MediaColl)
	mediaColl := repositories.MediaEndpoints{Collection: mi.GetCollection(repositories.MediaColl)}

	testMedia := &models.MediaDB{
		Data:     "123",
		FileType: "testFile",
	}

	t.Run("Test CreateMedia", func(t *testing.T) {
		response := mediaColl.CreateMedia(*testMedia)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testMedia.Id = response.Data.(string)
	})

	t.Run("Test GetMediaById", func(t *testing.T) {
		response := mediaColl.GetMediaById(testMedia.GetId())
		require.Equal(t, testMedia.Id, response.Data.(*models.MediaDB).Id)

		response = mediaColl.GetMediaById("IncorrectID")
		require.True(t, utils.IsTypeError(response))

	})

	t.Run("Test DeleteOne", func(t *testing.T) {
		response := mediaColl.DeleteMedia(testMedia.GetId())
		require.False(t, utils.IsTypeError(response))
		response = mediaColl.DeleteMedia(testMedia.GetId())
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
