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

func TestAccessController(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.UserColl)
	userColl := repositories.UserRepository{}
	userColl.SetCollection(mi.GetCollection(repositories.UserColl))

	mi.AddCollection(repositories.AdminColl)
	adminColl := repositories.AdminRepository{}
	adminColl.SetCollection(mi.GetCollection(repositories.AdminColl))

	db := repositories.RepositoryEndpoints{
		User: userColl,
		Admin: adminColl,
	}
	contr := controllers.NewController(db)


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
		Username: "test User 3",
		AuthId:   "789",
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

		response = userColl.CreateUser(*testUser3)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testUser3.Id = response.Data.(string)
	})

	testAdmin1 := &models.AdminInput{
		UserId: testUser1.Id,
		Access: "admin",
	}

	testAdmin2 := &models.AdminInput{
		UserId: testUser2.Id,
		Access: "superadmin",
	}

	t.Run("Test CreateAdmin", func(t *testing.T) {
		response := contr.CCreateAdmin(testAdmin1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		testAdmin1.Id = response.Data.(*models.AdminOutput).Id

		response = contr.CCreateAdmin(testAdmin2)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		testAdmin2.Id = response.Data.(*models.AdminOutput).Id
	})

	t.Run("Test IsAdmin", func(t *testing.T) {
		response := contr.CIsAdmin(testUser1.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is admin", response.Data.(string))

		response = contr.CIsAdmin(testUser2.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is admin", response.Data.(string))

		response = contr.CIsAdmin(testUser3.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
		
		response = contr.CIsAdmin("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsSuperAdmin", func(t *testing.T) {
		response := contr.CIsSuperAdmin(testUser1.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)

		response = contr.CIsSuperAdmin(testUser2.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is superadmin", response.Data.(string))

		response = contr.CIsSuperAdmin(testUser3.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
		
		response = contr.CIsSuperAdmin("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsPostsUser", func(t *testing.T) {
		
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
			_, err = db.Collection(repositories.AdminColl).DeleteMany(context.TODO(), bson.M{})
			if err != nil {
				return
			}
		})
	
	}