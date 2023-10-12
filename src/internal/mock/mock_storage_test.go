package mock_test

import (
	"birdai/src/internal/mock"
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
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
	user := &models.UserDB{
		Id:        primitive.ObjectID{byte(1)}.String(),
		Username:  "bird1",
		AuthId:    "123",
		CreatedAt: "0",
	}
	user2 := &models.UserDB{
		Id:        primitive.ObjectID{byte(2)}.String(),
		Username:  "bird2",
		AuthId:    "124",
		CreatedAt: "0",
	}
	t.Run("Test CreateOne collection success", func(t *testing.T) {
		response := userColl.CreateOne(user)
		require.False(t, utils.IsTypeError(response))

		filter := bson.M{"_id": response.Data.(*models.UserOutput).Id}
		response = userColl.FindOne(filter)
		require.NotNil(t, response.Data.(*models.UserOutput))

		response = userColl.CreateOne(user2)
		require.False(t, utils.IsTypeError(response))

		filter = bson.M{"_id": response.Data.(*models.UserOutput).Id}
		response = userColl.FindOne(filter)
	})

	t.Run("Test FindOne collection success", func(t *testing.T) {
		filter := bson.M{"_id": user.Id}
		response := userColl.FindOne(filter)
		require.Equal(t, user, response.Data)
		require.False(t, utils.IsTypeError(response))

	
		filter = bson.M{"_id": user2.Id}
		response = userColl.FindOne(filter)
		require.Equal(t, user2, response.Data)
		require.False(t, utils.IsTypeError(response))
	})

	t.Run("Test FindOne collection failure", func(t *testing.T) {
		filter := bson.M{"_id": "testtest"}
		response := userColl.FindOne(filter)
		require.True(t, utils.IsTypeError(response))
		require.False(t, utils.IsType(response, models.UserOutput{}))
	})

	t.Run("Test FindAll collection success", func(t *testing.T) {
		response := userColl.FindAll()
		require.Equal(t, []models.HandlerObject{user, user2}, response.Data.([]models.HandlerObject))
		require.Nil(t, response.Data.(*models.Err))
	})

	t.Run("Test UpdateOne collection success", func(t *testing.T) {
		response := userColl.UpdateOne(bson.M{
			"_id": user.Id,
			"username": "updated_name",
			"auth_id": user.AuthId,
			"created_at": user.CreatedAt,
		})
		require.Equal(t, "updated_name", response.Data.(*models.UserOutput).Username)
		require.False(t, utils.IsTypeError(response))
	})

	t.Run("Test DeleteOne collection success", func(t *testing.T) {
		filter := bson.M{"_id": user.Id}
		response := userColl.DeleteOne(filter)
		require.Equal(t, user, response.Data)
		require.False(t, utils.IsTypeError(response))
		response = userColl.FindOne(filter)
		require.True(t, utils.IsTypeError(response))
	})
}
