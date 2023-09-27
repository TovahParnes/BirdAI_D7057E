package mock_test

import (
	"birdai/src/internal/mock"
	"birdai/src/internal/storage"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
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
	user := storage.User{
		ID:        primitive.ObjectID{byte(1)},
		Username:  "bird1",
		AuthID:    "123",
		CreatedAt: 0,
	}
	user2 := storage.User{
		ID:        primitive.ObjectID{byte(2)},
		Username:  "bird2",
		AuthID:    "124",
		CreatedAt: 0,
	}
	t.Run("Test CreateOne collection success", func(t *testing.T) {
		err := userColl.CreateOne(user)
		require.Nil(t, err)
		err = userColl.CreateOne(user2)
		require.Nil(t, err)
	})

	t.Run("Test FindOne collection success", func(t *testing.T) {
		person, err := userColl.FindOne(user.ID.Hex())
		require.Equal(t, user, person)
		require.Nil(t, err)

		person, err = userColl.FindOne(user2.ID.Hex())
		require.Equal(t, user2, person)
		require.Nil(t, err)
	})

	t.Run("Test FindOne collection failure", func(t *testing.T) {
		person, err := userColl.FindOne("testtest")
		require.Nil(t, person)
		require.NotNil(t, err)
	})

	t.Run("Test FindAll collection success", func(t *testing.T) {
		persons, err := userColl.FindAll()
		require.Equal(t, []storage.HandlerObject{user, user2}, persons)
		require.Nil(t, err)
	})

	t.Run("Test UpdateOne collection success", func(t *testing.T) {
		person, err := userColl.UpdateOne(bson.D{
			{"_id", user.ID},
			{"username", "bird_changed"},
			{"auth_id", user.AuthID},
			{"created_at", user.CreatedAt},
		})
		require.Equal(t, "bird_changed", person.(storage.User).Username)
		require.Nil(t, err)
	})

	t.Run("Test DeleteOne collection success", func(t *testing.T) {
		person, err := userColl.DeleteOne(user.ID.Hex())
		require.Equal(t, user, person)
		require.Nil(t, err)
		foundPerson, _ := userColl.FindOne(user.ID.Hex())
		require.Nil(t, foundPerson)
	})
}
