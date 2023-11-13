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

func TestUserController(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.UserColl)
	userColl := repositories.UserRepository{}
	userColl.SetCollection(mi.GetCollection(repositories.UserColl))

	db := repositories.RepositoryEndpoints{
		User: userColl,}
	userContr := controllers.NewController(db)


	testUser1 := &models.UserDB{
		Username: "test User 1",
		AuthId:   "123",
		Active:   true,
	}
	testUser2 := &models.UserDB{
		Username: "test User 2",
		AuthId:   "456",
		Active:   true,
	}
	testUser3 := &models.UserDB{
		Username: "Deleted User",
		AuthId:   "789",
		Active:   false,
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

		response = userColl.CreateUser(*testUser3)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testUser3.Id = response.Data.(string)
	})

	t.Run("Test GetUserById", func(t *testing.T) {
		response :=userContr.CGetUserById(testUser1.GetId())
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserOutput{}, response.Data.(*models.UserOutput))
		require.Equal(t, testUser1.Id, response.Data.(*models.UserOutput).Id)

	
		response = userContr.CGetUserById("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)

		response =userContr.CGetUserById(testUser3.GetId())
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusGone, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test GetUserByAuthId", func(t *testing.T) {
		response := userContr.CGetUserByAuthId(testUser2.AuthId)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, testUser2.Id, response.Data.(*models.UserOutput).Id)

		response = userContr.CGetUserByAuthId("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test ListUsers", func(t *testing.T) {
		response := userContr.CListUsers(0)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 3, len(response.Data.([]*models.UserOutput)))
		require.IsType(t, &models.UserOutput{}, response.Data.([]*models.UserOutput)[0])

		response = userContr.CListUsers(1)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 0, len(response.Data.([]*models.UserOutput)))
	})

	t.Run("Test DeleteUser", func(t *testing.T) {
		response := userContr.CDeleteUser(testUser1.Id)
		require.False(t, utils.IsTypeError(response))

		response = userContr.CGetUserById(testUser1.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusGone, response.Data.(models.Err).StatusCode)

		response = userColl.GetUserById(testUser1.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserDB{}, response.Data.(*models.UserDB))
		require.Equal(t, false, response.Data.(*models.UserDB).Active)
		require.Equal(t, "Deleted User", response.Data.(*models.UserDB).Username)

		response = userContr.CDeleteUser("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test UpdateUser", func(t *testing.T) {
		updateUser := models.UserInput{
			Username: "Hello world",
			Active:   true,
		}
		response := userContr.CUpdateUser(testUser2.Id, &updateUser)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserOutput{}, response.Data.(*models.UserOutput))
		require.Equal(t, updateUser.Username, response.Data.(*models.UserOutput).Username)
		require.Equal(t, testUser2.Active, response.Data.(*models.UserOutput).Active)

		response = userContr.CUpdateUser(testUser2.Id, &updateUser)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
		require.Equal(t, "No change compared to current document", response.Data.(models.Err).Description)

		response = userContr.CUpdateUser("IncorrectID", &updateUser)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)

		updateUser = models.UserInput{
			Username: "Hello world",
			Active:   true,
		}
		response = userContr.CUpdateUser(testUser3.Id, &updateUser)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserOutput{}, response.Data.(*models.UserOutput))
		require.Equal(t, updateUser.Username, response.Data.(*models.UserOutput).Username)
		require.Equal(t, updateUser.Active, response.Data.(*models.UserOutput).Active)
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