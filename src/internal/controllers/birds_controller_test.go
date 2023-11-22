package controllers_test

import (
	"birdai/src/internal/controllers"
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestBirdController(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)

	mi.AddCollection(repositories.BirdColl)
	birdColl := repositories.BirdRepository{}
	birdColl.SetCollection(mi.GetCollection(repositories.BirdColl))
	
	mi.AddCollection(repositories.MediaColl)
	mediaColl := repositories.MediaRepository{}
	mediaColl.SetCollection(mi.GetCollection(repositories.MediaColl))

	db := repositories.RepositoryEndpoints{
		Bird: birdColl,
		Media: mediaColl,}
	birdContr := controllers.NewController(db)

	testImage := &models.MediaDB{
		Data:     "testImage",
	}

	testSound := &models.MediaDB{
		Data:     "testSound",
	}

	t.Run("Test CreateMedia", func(t *testing.T) {
		response := mediaColl.CreateMedia(*testImage)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testImage.Id = response.Data.(string)

		response = mediaColl.CreateMedia(*testSound)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testSound.Id = response.Data.(string)
	})


	testBird1 := &models.BirdDB{
		Name: "test Bird 1",
		Description: "Cool test bird",
		ImageId: testImage.Id,
		SoundId: testSound.Id,

	}
	testBird2 := &models.BirdDB{
		Name: "test Bird 2",
		Description: "Rad test bird",
		ImageId: testImage.Id,
		SoundId: testSound.Id,
	}

	t.Run("Test CreateBirds", func(t *testing.T) {
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
		response := birdContr.CGetBirdById(testBird1.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.BirdOutput{}, response.Data)
		require.Equal(t, testBird1.Id, response.Data.(*models.BirdOutput).Id)

		response = birdContr.CGetBirdById("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, 404, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test ListBirds", func(t *testing.T) {
		response := birdContr.CListBirds(0, "")
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 2, len(response.Data.([]*models.BirdOutput)))
		require.IsType(t, &models.BirdOutput{}, response.Data.([]*models.BirdOutput)[0])

		response = birdContr.CListBirds(1, "")
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 0, len(response.Data.([]*models.BirdOutput)))
	})

	t.Run("Test UpdateBird", func(t *testing.T) {
		updateBird := &models.BirdInput{
			Name: "Updated bird",
			Description: "Cool test bird",
			ImageId: testImage.Id,
			SoundId: testSound.Id,
		}

		response := birdContr.CUpdateBird(testBird1.Id, updateBird)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.BirdOutput{}, response.Data)
		require.Equal(t, updateBird.Name, response.Data.(*models.BirdOutput).Name)

		response = birdContr.CUpdateBird("IncorrectID", updateBird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)

		response = birdContr.CUpdateBird(testBird1.Id, updateBird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
		require.Equal(t, "No change compared to current document", response.Data.(models.Err).Description)

		updateBird = &models.BirdInput{
			Name: "test Bird 2",
			Description: "Rad test bird",
			ImageId: "invalid id",
			SoundId: testSound.Id,
		}
		response = birdContr.CUpdateBird(testBird2.Id, updateBird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
		require.Equal(t, "Image with given id does not exist", response.Data.(models.Err).Description)
	
		updateBird = &models.BirdInput{
			Name: "test Bird 2",
			Description: "Rad test bird",
			ImageId: testImage.Id,
			SoundId: "invalid id",
		}
		response = birdContr.CUpdateBird(testBird2.Id, updateBird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
		require.Equal(t, "Sound with given id does not exist", response.Data.(models.Err).Description)
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
		_, err = db.Collection(repositories.MediaColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
	})
}