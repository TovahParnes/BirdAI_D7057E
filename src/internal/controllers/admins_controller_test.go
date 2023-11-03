package controllers_test

import (
	"birdai/src/internal/controllers"
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAdminController(t *testing.T) {
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

		response = contr.CCreateAdmin(testAdmin1)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusConflict, response.Data.(models.Err).StatusCode)

		response = contr.CCreateAdmin(&models.AdminInput{
			UserId: "IncorrectID",
			Access: "admin",
		})
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test GetAdminById", func(t *testing.T) {
		response := contr.CGetAdminById(testAdmin1.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		require.Equal(t, testAdmin1.Id, response.Data.(*models.AdminOutput).Id)

		response = contr.CGetAdminById(testAdmin2.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		require.Equal(t, testAdmin2.Id, response.Data.(*models.AdminOutput).Id)

		response = contr.CGetAdminById("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test GetAdminByUserId", func(t *testing.T) {
		response := contr.CGetAdminByUserId(testUser1.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		require.Equal(t, testAdmin1.Id, response.Data.(*models.AdminOutput).Id)
		fmt.Println(response.Data.(*models.AdminOutput))

		response = contr.CGetAdminByUserId(testUser2.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		require.Equal(t, testAdmin2.Id, response.Data.(*models.AdminOutput).Id)
		fmt.Println(response.Data.(*models.AdminOutput))

		response = contr.CGetAdminByUserId("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test ListAdmins", func(t *testing.T) {
		response := contr.CListAdmins(0, "")
		fmt.Println(response.Data)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 2, len(response.Data.([]*models.AdminOutput)))
		require.IsType(t, &models.AdminOutput{}, response.Data.([]*models.AdminOutput)[0])

		response = contr.CListAdmins(1, "")
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 0, len(response.Data.([]*models.AdminOutput)))
	})

	t.Run("Test UpdateAdmin", func(t *testing.T) {
		//incorrect user id
		response := contr.CUpdateAdmin(testAdmin1.Id, &models.AdminInput{
			UserId: "IncorrectID",
			Access: "admin",
		})
		fmt.Println(response.Data)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)

		//User id already exists
		response = contr.CUpdateAdmin(testAdmin1.Id, &models.AdminInput{
			UserId: testUser2.Id,
			Access: "admin",
		})
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusConflict, response.Data.(models.Err).StatusCode)

		//Can't remove last superadmin
		response = contr.CUpdateAdmin(testAdmin1.Id, &models.AdminInput{
			UserId: testUser1.Id,
			Access: "admin",
		})
		require.True(t, utils.IsTypeError(response))
		fmt.Println(response.Data)
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		//No changes to document
		response = contr.CUpdateAdmin(testAdmin1.Id, &models.AdminInput{
			UserId: testUser1.Id,
			Access: "admin",
		})
		require.True(t, utils.IsTypeError(response))
		fmt.Println(response.Data)
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		//Update sucessful
		response = contr.CUpdateAdmin(testAdmin1.Id, &models.AdminInput{
			UserId: testUser1.Id,
			Access: "superadmin",
		})
		fmt.Println(response.Data)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		require.Equal(t, testUser1.Id, response.Data.(*models.AdminOutput).User.Id)
		require.Equal(t, "superadmin", response.Data.(*models.AdminOutput).Access)
	})

	t.Run("Test DeleteAdmin", func(t *testing.T) {
		response := contr.CDeleteAdmin(testAdmin1.Id)
		require.False(t, utils.IsTypeError(response))

		response = contr.CDeleteAdmin(testAdmin2.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusConflict, response.Data.(models.Err).StatusCode)	
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