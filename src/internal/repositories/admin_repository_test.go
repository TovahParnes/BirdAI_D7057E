package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"context"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

// TestAdminRepository tests the basic functionality of the AdminRepository struct
func TestAdminRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.AdminColl)
	adminColl := repositories.AdminRepository{}
	adminColl.SetCollection(mi.GetCollection(repositories.AdminColl))

	testAdmin1 := &models.AdminDB{
		UserId: primitive.NewObjectID().Hex(),
		Access: "Admin",
	}
	testAdmin2 := &models.AdminDB{
		UserId: primitive.NewObjectID().Hex(),
		Access: "Admin",
	}

	t.Run("Test CreateAdmin", func(t *testing.T) {
		response := adminColl.CreateAdmin(*testAdmin1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testAdmin1.Id = response.Data.(string)
		response = adminColl.CreateAdmin(*testAdmin2)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testAdmin2.Id = response.Data.(string)
	})

	t.Run("Test GetAdminById", func(t *testing.T) {
		response := adminColl.GetAdminById(testAdmin1.Id)
		require.Equal(t, testAdmin1.Id, response.Data.(*models.AdminDB).Id)
		testAdmin1 = response.Data.(*models.AdminDB)

		response = adminColl.GetAdminById("IncorrectID")
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test GetAllAdmins", func(t *testing.T) {
		response := adminColl.ListAdmins(bson.M{}, 0)
		require.Equal(t, 2, len(response.Data.([]models.AdminDB)))
		responseAll := adminColl.ListAllAdmins(0)
		require.Equal(t, response.Data.([]models.AdminDB), responseAll.Data.([]models.AdminDB))
		response = adminColl.ListAdmins(bson.M{"_id": testAdmin2.Id}, 0)
		require.Equal(t, 1, len(response.Data.([]models.AdminDB)))
		response = adminColl.ListAdmins(bson.M{"_id": "IncorrectId"}, 0)
		require.Equal(t, 0, len(response.Data.([]models.AdminDB)))
	})

	t.Run("Test UpdateAdmin", func(t *testing.T) {
		updateAdmin := models.AdminInput{
			Id:     testAdmin1.Id,
			Access: "SuperAdmin",
		}
		response := adminColl.UpdateAdmin(updateAdmin)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, updateAdmin.Access, response.Data.(*models.AdminDB).Access)

		response = adminColl.UpdateAdmin(updateAdmin)
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test DeleteAdmin", func(t *testing.T) {
		response := adminColl.DeleteAdmin(testAdmin1.Id)
		require.False(t, utils.IsTypeError(response))
		response = adminColl.DeleteAdmin(testAdmin1.Id)
		require.True(t, utils.IsTypeError(response))

	})

	// Need to delete everything from testDB
	t.Cleanup(func() {
		mi.DisconnectDB()
		client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
		if err != nil {
			return
		}
		db := client.Database("testDB")
		_, err = db.Collection(repositories.AdminColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
	})
}
