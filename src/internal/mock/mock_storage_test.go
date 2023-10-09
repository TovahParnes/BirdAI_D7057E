package mock_test

import (
	"birdai/src/internal/mock"
	"birdai/src/internal/models"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	collName = "user"
)

func TestMockMongoInstance(t *testing.T) {
	mongoDB := mock.NewMockMongoInstance()
	t.Run("Test add collection", func(t *testing.T) {
		mongoDB.AddCollection(collName)
	})
	t.Run("Test get collection", func(t *testing.T) {
		collection := mongoDB.GetCollection(collName)
		require.NotNil(t, collection)
	})
	t.Run("Test Disconnect", func(t *testing.T) {
		mongoDB.DisconnectDB()
	})
}

func TestMockMongoCollection(t *testing.T) {
	mongoDB := mock.NewMockMongoInstance()
	mongoDB.AddCollection(collName)
	userColl := mongoDB.GetCollection(collName)
	user := &models.User{
		Id:        primitive.ObjectID{byte(1)}.String(),
		Username:  "bird1",
		AuthId:    "123",
		CreatedAt: "0",
	}
	user2 := &models.User{
		Id:        primitive.ObjectID{byte(2)}.String(),
		Username:  "bird2",
		AuthId:    "124",
		CreatedAt: "0",
	}
	t.Run("Test CreateOne collection success", func(t *testing.T) {
		response := userColl.CreateOne(user)
		require.Nil(t, response.Data.(*models.User))
		require.NotNil(t, response.Data.(*models.Err))

		response = userColl.CreateOne(user2)
		require.Nil(t, response.Data.(*models.User))
		require.NotNil(t, response.Data.(*models.Err))
	})

	t.Run("Test FindOne collection success", func(t *testing.T) {
		response := userColl.FindOne(user.Id)
		require.Equal(t, user, response.Data)
		//require.Nil(t, err)

		response = userColl.FindOne(user2.Id)
		require.Equal(t, user2, response.Data)
		//require.Nil(t, err)
	})

	t.Run("Test FindOne collection failure", func(t *testing.T) {
		response := userColl.FindOne("testtest")
		require.Nil(t, response.Data.(*models.User))
		require.NotNil(t, response.Data.(*models.Err))
	})

	t.Run("Test FindAll collection success", func(t *testing.T) {
		response := userColl.FindAll()
		require.Equal(t, []models.HandlerObject{user, user2}, response.Data.([]models.HandlerObject))
		require.Nil(t, response.Data.(*models.Err))
	})

	t.Run("Test UpdateOne collection success", func(t *testing.T) {
		response := userColl.UpdateOne(bson.D{
			{"_id", user.Id},
			{"username", "bird_changed"},
			{"auth_id", user.AuthId},
			{"created_at", user.CreatedAt},
		})
		require.Equal(t, "bird_changed", response.Data.(*models.User).Username)
		require.Nil(t, response.Data.(*models.Err))
	})

	t.Run("Test DeleteOne collection success", func(t *testing.T) {
		response := userColl.DeleteOne(user.Id)
		require.Equal(t, user, response.Data.(*models.User))
		require.Nil(t, response.Data.(*models.Err))
		response = userColl.FindOne(user.Id)
		require.Nil(t, response.Data.(*models.User))
	})
}
