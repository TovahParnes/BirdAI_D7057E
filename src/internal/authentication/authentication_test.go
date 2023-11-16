package authentication

import (
	"birdai/src/internal/controllers"
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAccessController(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.UserColl)
	userColl := repositories.UserRepository{}
	userColl.SetCollection(mi.GetCollection(repositories.UserColl))

	db := repositories.RepositoryEndpoints{
		User: userColl,
	}
	contr := controllers.NewController(db)
	auth := NewAuthentication(db.User)
	

	t.Run("Test LoginUser", func(t *testing.T) {
		user := &models.UserLogin{
			AuthId: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Username: "test",
		}
		response := auth.LoginUser(user)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserLoginOutput{}, response.Data.(*models.UserLoginOutput))
		require.True(t, response.Data.(*models.UserLoginOutput).CreatedNew)
		response = contr.CListUsers(0)
		require.Equal(t, 1, len(response.Data.([]*models.UserOutput)))

		response = auth.LoginUser(user)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserLoginOutput{}, response.Data.(*models.UserLoginOutput))
		require.False(t, response.Data.(*models.UserLoginOutput).CreatedNew)
		response = contr.CListUsers(0)
		require.Equal(t, 1, len(response.Data.([]*models.UserOutput)))

		user.Username = "test2"
		response = auth.LoginUser(user)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserLoginOutput{}, response.Data.(*models.UserLoginOutput))
		require.Equal(t, "test", response.Data.(*models.UserLoginOutput).Username)
		require.False(t, response.Data.(*models.UserLoginOutput).CreatedNew)
		response = contr.CListUsers(0)
		require.Equal(t, 1, len(response.Data.([]*models.UserOutput)))

		user.AuthId = "5f9d3b3b9d3b3b9d3b3b9d3c"
		response = auth.LoginUser(user)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserLoginOutput{}, response.Data.(*models.UserLoginOutput))
		require.True(t, response.Data.(*models.UserLoginOutput).CreatedNew)
		response = contr.CListUsers(0)
		require.Equal(t, 2, len(response.Data.([]*models.UserOutput)))

		response = auth.LoginUser(user)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.UserLoginOutput{}, response.Data.(*models.UserLoginOutput))
		require.False(t, response.Data.(*models.UserLoginOutput).CreatedNew)
		response = contr.CDeleteUser(response.Data.(*models.UserLoginOutput).Id)
		require.False(t, utils.IsTypeError(response))
		response = auth.LoginUser(user)
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test CheckExpired", func(t *testing.T) {
		
		//TODO: Need to test this
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