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

func TestPostController(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)

	mi.AddCollection(repositories.UserColl)
	userColl := repositories.UserRepository{}
	userColl.SetCollection(mi.GetCollection(repositories.UserColl))

	mi.AddCollection(repositories.PostColl)
	postColl := repositories.PostRepository{}
	postColl.SetCollection(mi.GetCollection(repositories.PostColl))

	mi.AddCollection(repositories.BirdColl)
	birdColl := repositories.BirdRepository{}
	birdColl.SetCollection(mi.GetCollection(repositories.BirdColl))
	
	mi.AddCollection(repositories.MediaColl)
	mediaColl := repositories.MediaRepository{}
	mediaColl.SetCollection(mi.GetCollection(repositories.MediaColl))

	db := repositories.RepositoryEndpoints{
		User: userColl,
		Post: postColl,
		Bird: birdColl,
		Media: mediaColl,
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
	testBird2 := &models.BirdDB{
		Name: "test Bird 2",
		Description: "Rad test bird",
		ImageId: testImage.Id,
		SoundId: testSound.Id,
	}

	t.Run("Test CreateBirds", func(t *testing.T) {
		response := birdColl.CreateBird(*testBird1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testBird1.Id = response.Data.(string)

		response = birdColl.CreateBird(*testBird2)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testBird2.Id = response.Data.(string)
	})

	testPost1 := &models.PostDB{
		BirdId: testBird1.Id,
		Location: "place 1",
		Accuracy: 0.5,
	}

	testPost2 := &models.PostDB{
		BirdId: testBird2.Id,
		Location: "place 2",
		Accuracy: 0.9,
	}

	t.Run("Test CreatePost", func(t *testing.T) {
		postCreation := &models.PostCreation{
			BirdId: testPost1.BirdId,
			Location: testPost1.Location,
			Accuracy: testPost1.Accuracy,
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
			Accuracy: testPost2.Accuracy,
			Media: *testMediaInput,
		}
		response = contr.CCreatePost(testUser2.Id, postCreation)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.PostOutput{}, response.Data.(*models.PostOutput))
		testPost2.Id = response.Data.(*models.PostOutput).Id

		postCreation.BirdId = "IncorrectID"

		response = contr.CCreatePost(testUser1.Id, postCreation)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test GetPostById", func(t *testing.T) {
		response := contr.CGetPostById(testPost1.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.PostOutput{}, response.Data.(*models.PostOutput))
		require.Equal(t, testPost1.Id, response.Data.(*models.PostOutput).Id)

		response = contr.CGetPostById(testPost2.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.PostOutput{}, response.Data.(*models.PostOutput))
		require.Equal(t, testPost2.Id, response.Data.(*models.PostOutput).Id)

		response = contr.CGetPostById("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test ListPosts", func(t *testing.T) {
		response := contr.CListPosts(0, "")
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 2, len(response.Data.([]*models.PostOutput)))
		require.IsType(t, &models.PostOutput{}, response.Data.([]*models.PostOutput)[0])

		response = contr.CListPosts(1, "")
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 0, len(response.Data.([]*models.PostOutput)))
	})

	t.Run("Test ListUsersPosts", func(t *testing.T) {
		response := contr.CListUsersPosts(testUser1.Id, 0)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 1, len(response.Data.([]*models.PostOutput)))
		require.IsType(t, &models.PostOutput{}, response.Data.([]*models.PostOutput)[0])

		response = contr.CListPosts(1, "")
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 0, len(response.Data.([]*models.PostOutput)))
	})


	t.Run("Test UpdatePost", func(t *testing.T) {
		updatePost := &models.PostInput{
			Location: "updated place",
			Comment: "updated comment",
		}

		response := contr.CUpdatePost(testPost1.Id, updatePost)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, &models.PostOutput{}, response.Data)
		require.Equal(t, updatePost.Location, response.Data.(*models.PostOutput).Location)
	})

	t.Run("Test DeletePost", func(t *testing.T) {
		response := contr.CDeletePost(testPost1.Id)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data)

		response = contr.CGetPostById(testPost1.Id)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)

		response = contr.CDeletePost("IncorrectID")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusNotFound, response.Data.(models.Err).StatusCode)
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
		_, err = db.Collection(repositories.MediaColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
		_, err = db.Collection(repositories.BirdColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
		_, err = db.Collection(repositories.PostColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
	})

}