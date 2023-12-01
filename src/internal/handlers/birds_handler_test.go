package handlers_test

import (
	"birdai/src/internal/handlers"
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestBirdHandler(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.UserColl)
	userColl := repositories.UserRepository{}
	userColl.SetCollection(mi.GetCollection(repositories.UserColl))

	mi.AddCollection(repositories.AdminColl)
	adminColl := repositories.AdminRepository{}
	adminColl.SetCollection(mi.GetCollection(repositories.AdminColl))

	mi.AddCollection(repositories.BirdColl)
	birdColl := repositories.BirdRepository{}
	birdColl.SetCollection(mi.GetCollection(repositories.BirdColl))

	mi.AddCollection(repositories.MediaColl)
	mediaColl := repositories.MediaRepository{}
	mediaColl.SetCollection(mi.GetCollection(repositories.MediaColl))

	mi.AddCollection(repositories.PostColl)
	postColl := repositories.PostRepository{}
	postColl.SetCollection(mi.GetCollection(repositories.PostColl))

	db := repositories.RepositoryEndpoints{
		User: userColl,
		Admin: adminColl,
		Bird: birdColl,
		Media: mediaColl,
		Post: postColl,
	}
	app := fiber.New()
	handlers.New(app, db)

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
		SoundId: testSound.Id,

	}
	testBird2 := &models.BirdDB{
		Name: "test Bird 2",
		Description: "Rad test bird",
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
	
	t.Run("Test ListBirds", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/birds/list", nil)
		resp, _ := app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/birds/list?set=0&search=hello", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/birds/list?set=1&search=hello", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/birds/list?set=aaa&search=hello", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 400, resp.StatusCode)
		
		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/birds/list?set=0&search=hello!!!", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("Test GetBirdById", func(t *testing.T) {
		url := "http://127.0.0.1:3000/api/v1/birds/" + testBird1.Id
		req, _ := http.NewRequest("GET", url, nil)
		resp, _ := app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		url = "http://127.0.0.1:3000/api/v1/birds/" + testBird2.Id
		req, _ = http.NewRequest("GET", url, nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		//not a valid id
		url = "http://127.0.0.1:3000/api/v1/birds/" + "invalidid"
		req, _ = http.NewRequest("GET", url, nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 400, resp.StatusCode)

		//not found
		url = "http://127.0.0.1:3000/api/v1/birds/" + "60b9b6b9e7b3d3b3f0a3b3b3"
		req, _ = http.NewRequest("GET", url, nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 404, resp.StatusCode)
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