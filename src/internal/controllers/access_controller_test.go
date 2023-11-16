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

	mi.AddCollection(repositories.BirdColl)
	birdColl := repositories.BirdRepository{}
	birdColl.SetCollection(mi.GetCollection(repositories.BirdColl))

	mi.AddCollection(repositories.MediaColl)
	mediaColl := repositories.MediaRepository{}
	mediaColl.SetCollection(mi.GetCollection(repositories.MediaColl))

	mi.AddCollection(repositories.PostColl)
	postColl := repositories.PostRepository{}
	postColl.SetCollection(mi.GetCollection(repositories.PostColl))

	db := repositories.RepositoryEndpoints{
		User: userColl,
		Admin: adminColl,
		Bird: birdColl,
		Media: mediaColl,
		Post: postColl,
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

	testAdmin1 := &models.AdminDB{
		UserId: testUser1.Id,
		Access: "admin",
	}

	testAdmin2 := &models.AdminDB{
		UserId: testUser2.Id,
		Access: "superadmin",
	}

	t.Run("Test CreateAdmin", func(t *testing.T) {
		adminInput := &models.AdminInput{
			UserId: testUser1.Id,
			Access: testAdmin1.Access,
		}
		response := contr.CCreateAdmin(adminInput)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		testAdmin1.Id = response.Data.(*models.AdminOutput).Id

		adminInput = &models.AdminInput{
			UserId: testUser2.Id,
			Access: testAdmin2.Access,
		}
		response = contr.CCreateAdmin(adminInput)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.AdminOutput{}, response.Data.(*models.AdminOutput))
		testAdmin2.Id = response.Data.(*models.AdminOutput).Id
	})

	testImage := &models.MediaDB{
		Data:     "testImage",
	}

	testSound := &models.MediaDB{
		Data:     "testSound",
	}

	testMediaInput := &models.MediaInput{
		Data:     "testSound",
	}

	t.Run("Test CreateMedia", func(t *testing.T) {
		response := mediaColl.CreateMedia(*testImage)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testImage.Id = response.Data.(string)

		response = mediaColl.CreateMedia(*testSound)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testSound.Id = response.Data.(string)
	})


	testBird1 := &models.BirdDB{
		Name: "test Bird 1",
		Description: "Cool test bird",
		ImageId: testImage.Id,
		SoundId: testSound.Id,

	}

	t.Run("Test CreateBirds", func(t *testing.T) {
		response := birdColl.CreateBird(*testBird1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testBird1.Id = response.Data.(string)
	})

	testPost1 := &models.PostDB{
		BirdId: testBird1.Id,
		Location: "place 1",
	}

	testPost2 := &models.PostDB{
		BirdId: testBird1.Id,
		Location: "place 2",
	}

	t.Run("Test CreatePost", func(t *testing.T) {
		postCreation := &models.PostCreation{
			BirdId: testBird1.Id,
			Location: testPost1.Location,
			Media: *testMediaInput,
		}

		response := contr.CCreatePost(testUser1.Id, postCreation)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.PostOutput{}, response.Data.(*models.PostOutput))
		testPost1.Id = response.Data.(*models.PostOutput).Id
		testPost1.MediaId = response.Data.(*models.PostOutput).UserMedia.Id

		postCreation = &models.PostCreation{
			BirdId: testPost2.BirdId,
			Location: testPost2.Location,
			Media: *testMediaInput,
		}

		response = contr.CCreatePost(testUser3.Id, postCreation)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.PostOutput{}, response.Data.(*models.PostOutput))
		testPost2.Id = response.Data.(*models.PostOutput).Id
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
		response := contr.CIsPostsUser(testUser1.Id, testPost1.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is posts user", response.Data.(string))

		response = contr.CIsPostsUser(testUser3.Id, testPost2.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is posts user", response.Data.(string))

		response = contr.CIsPostsUser(testUser3.Id, testPost1.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
		
		response = contr.CIsPostsUser("IncorrectID", testPost1.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsCurrentUser", func(t *testing.T) {
		response := contr.CIsCurrentUser(testAdmin1.UserId, testUser1.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is current user", response.Data.(string))

		response = contr.CIsCurrentUser(testAdmin1.UserId, testUser2.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsPostsUserOrAdmin", func(t *testing.T) {
		//Admin
		response := contr.CIsPostsUserOrAdmin(testUser1.Id, testPost2.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is admin", response.Data.(string))

		//Posts user
		response = contr.CIsPostsUserOrAdmin(testUser1.Id, testPost1.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is posts user", response.Data.(string))

		//Not admin or posts user
		response = contr.CIsPostsUserOrAdmin(testUser3.Id, testPost1.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test CIsCurrentUserOrAdmin", func(t *testing.T) {
		//Current User
		response := contr.CIsCurrentUserOrAdmin(testUser3.Id, testUser3.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is current user", response.Data.(string))

		//Admin
		response = contr.CIsCurrentUserOrAdmin(testUser1.Id, testUser3.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is admin", response.Data.(string))

		//Not admin or current user
		response = contr.CIsCurrentUserOrAdmin(testUser3.Id, testUser2.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test CIsCurrentUserOrSuperAdmin", func(t *testing.T) {
		//Current User
		response := contr.CIsCurrentUserOrSuperAdmin(testUser3.Id, testUser3.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is current user", response.Data.(string))

		//Superadmin
		response = contr.CIsCurrentUserOrSuperAdmin(testUser2.Id, testUser3.Id)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, "Is superadmin", response.Data.(string))

		//Not superadmin or current user
		response = contr.CIsCurrentUserOrSuperAdmin(testUser1.Id, testUser3.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusForbidden, response.Data.(models.Err).StatusCode)
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
		_, err = db.Collection(repositories.BirdColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
		_, err = db.Collection(repositories.MediaColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
		_, err = db.Collection(repositories.PostColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
	})

}