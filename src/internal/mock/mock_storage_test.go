package mock_test

import (
	"birdai/src/internal/mock"
	"birdai/src/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"

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
		newId, err := userColl.CreateOne(user)
		require.Nil(t, err)
		filter := bson.M{"_id": newId}
		newUser, err := userColl.FindOne(filter)
		user = newUser.(*models.User)
		require.NotNil(t, newUser)
		require.Nil(t, err)
		newId, err = userColl.CreateOne(user2)
		require.Nil(t, err)
		filter = bson.M{"_id": newId}
		newUser, err = userColl.FindOne(filter)
		user2 = newUser.(*models.User)
		require.NotNil(t, newUser)
		require.Nil(t, err)
	})

	t.Run("Test FindOne collection success", func(t *testing.T) {
		filter := bson.M{"_id": user.Id}
		person, err := userColl.FindOne(filter)
		require.Equal(t, user, person)
		require.Nil(t, err)

		filter = bson.M{"_id": user2.Id}
		person, err = userColl.FindOne(filter)
		require.Equal(t, user2, person)
		require.Nil(t, err)
	})

	t.Run("Test FindOne collection failure", func(t *testing.T) {
		filter := bson.M{"_id": "testtest"}
		person, err := userColl.FindOne(filter)
		require.Nil(t, person)
		require.NotNil(t, err)
	})

	t.Run("Test FindAll collection success", func(t *testing.T) {
		persons, err := userColl.FindAll()
		require.Equal(t, []models.HandlerObject{user, user2}, persons)
		require.Nil(t, err)
	})

	t.Run("Test UpdateOne collection success", func(t *testing.T) {
		person, err := userColl.UpdateOne(bson.M{
			"_id":        user.Id,
			"username":   "bird_changed",
			"auth_id":    user.AuthId,
			"created_at": user.CreatedAt,
		})
		require.Equal(t, "bird_changed", person.(*models.User).Username)
		require.Nil(t, err)
		user = person.(*models.User)
	})

	t.Run("Test DeleteOne collection success", func(t *testing.T) {
		filter := bson.M{"_id": user.Id}
		person, err := userColl.DeleteOne(filter)
		require.Equal(t, user, person)
		require.Nil(t, err)
		filter = bson.M{"_id": user.Id}
		foundPerson, _ := userColl.FindOne(filter)
		require.Nil(t, foundPerson)
	})
}
