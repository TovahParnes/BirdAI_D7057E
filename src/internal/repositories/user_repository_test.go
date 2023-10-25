package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

// TestRepository functions
func TestUserRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.UserColl)
	userColl := repositories.UserEndpoints{Collection: mi.GetCollection(repositories.UserColl)}

	testUser1 := &models.UserDB{
		Username: "test User 1",
		AuthId:   "123123123",
		Active:   true,
	}
	testUser2 := &models.UserDB{
		Username: "test User 2",
		AuthId:   "123123",
		Active:   true,
	}

	t.Run("Test CreateUser", func(t *testing.T) {
		response := userColl.CreateUser(*testUser1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testUser1.Id = response.Data.(string)
		response = userColl.CreateUser(*testUser2)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testUser2.Id = response.Data.(string)
	})

	t.Run("Test GetUserById", func(t *testing.T) {
		response := userColl.GetUserById(testUser1.GetId())
		require.Equal(t, testUser1.Id, response.Data.(*models.UserDB).Id)
		testUser1 = response.Data.(*models.UserDB)

		response = userColl.GetUserById("IncorrectID")
		require.True(t, utils.IsTypeError(response))

	})

	t.Run("Test GetUserByAuthId", func(t *testing.T) {
		response := userColl.GetUserByAuthId(testUser2.AuthId)
		require.Equal(t, testUser2.AuthId, response.Data.(*models.UserDB).AuthId)
		testUser2 = response.Data.(*models.UserDB)

		response = userColl.GetUserByAuthId("IncorrectID")
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test ListUsers & ListAllUsers", func(t *testing.T) {
		response := userColl.ListUsers(bson.M{}, 0)
		require.Equal(t, 2, len(response.Data.([]models.UserDB)))
		responseAll := userColl.ListAllUsers(0)
		require.Equal(t, response.Data.([]models.UserDB), responseAll.Data.([]models.UserDB))
		response = userColl.ListUsers(bson.M{"_id": testUser2.Id}, 0)
		require.Equal(t, 1, len(response.Data.([]models.UserDB)))
		response = userColl.ListUsers(bson.M{"_id": "IncorrectId"}, 0)
		require.Equal(t, 0, len(response.Data.([]models.UserDB)))
	})

	t.Run("Test UpdateOne", func(t *testing.T) {
		updateUser := models.UserInput{
			Username: "Test User 1 is the best!!",
		}
		response := userColl.UpdateUser(updateUser)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, updateUser.Username, response.Data.(*models.UserDB).Username)

		response = userColl.UpdateUser(updateUser)
		require.True(t, utils.IsTypeError(response))
	})
	// Delete both
	t.Run("Test DeleteOne", func(t *testing.T) {
		response := userColl.DeleteUser(testUser1.GetId())
		require.False(t, utils.IsTypeError(response))
		testUser1.Username = "Deleted User"
		testUser1.Active = false
		require.Equal(t, testUser1, response.Data.(*models.UserDB))

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
	})
}
