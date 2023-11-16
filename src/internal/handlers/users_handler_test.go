package handlers_test

import (
	"birdai/src/internal/handlers"
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestUserHandler(t *testing.T) {
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

	t.Run("LoginUser", func(t *testing.T) {
		u1 := models.UserLogin{
			AuthId:  "123",
			Username: "testusername",
		}
		jsonValue, _ := json.Marshal(u1)
		req, _ := http.NewRequest("POST", "http://127.0.0.1:3000/api/v1/users", bytes.NewBuffer(jsonValue))
		req.Header.Add("Content-Type", "application/json")
		resp, _ := app.Test(req, -1)
		require.Equal(t, 201, resp.StatusCode)

		u2 := models.UserLogin{
			AuthId:  "456",
			Username: "testusername",
		}
		jsonValue, _ = json.Marshal(u2)
		req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/api/v1/users", bytes.NewBuffer(jsonValue))
		req.Header.Add("Content-Type", "application/json")
		resp, _ = app.Test(req, -1)
		require.Equal(t, 201, resp.StatusCode)

		jsonValue, _ = json.Marshal(u1)
		req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/api/v1/users", bytes.NewBuffer(jsonValue))
		req.Header.Add("Content-Type", "application/json")
		resp, _ = app.Test(req, -1)
		//require.Equal(t, 200, resp.StatusCode)

		u1.Username = "a"
		jsonValue, _ = json.Marshal(u1)
		req, _ = http.NewRequest("POST", "http://127.0.0.1:3000/api/v1/users", bytes.NewBuffer(jsonValue))
		req.Header.Add("Content-Type", "application/json")
		resp, _ = app.Test(req, -1)
		require.Equal(t, 400, resp.StatusCode)
	})

	t.Run("ListUsers", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/users/list", nil)
		resp, _ := app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/users/list?set=1&search=hello", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 200, resp.StatusCode)

		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/users/list?set=aaa&search=hello", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 400, resp.StatusCode)

		req, _ = http.NewRequest("GET", "http://127.0.0.1:3000/api/v1/users/list?set=0&search=hello!!!", nil)
		resp, _ = app.Test(req, -1)
		require.Equal(t, 400, resp.StatusCode)
	})



	// Need to delete everything from testDB
	t.Cleanup(func() {
		mi.DisconnectDB()
		client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			return
		}
		db := client.Database("testDB")
		_, err = db.Collection(repositories.UserColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
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