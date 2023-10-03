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

<<<<<<< HEAD
<<<<<<< HEAD
func TestPostRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.PostColl)
	postColl := repositories.PostRepository{}
	postColl.SetCollection(mi.GetCollection(repositories.PostColl))
=======
func TestPostEndpoints(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.PostColl)
	postColl := repositories.PostEndpoints{Collection: mi.GetCollection(repositories.PostColl)}
>>>>>>> 5bf0e17 (implemented post repository)
=======
func TestPostRepository(t *testing.T) {
	mi, err := repositories.Connect("testDB", "mongodb://localhost:27017")
	require.Nil(t, err)
	mi.AddCollection(repositories.PostColl)
	postColl := repositories.PostRepository{}
	postColl.SetCollection(mi.GetCollection(repositories.PostColl))
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)

	testPost1 := &models.PostDB{
		UserId:   primitive.NewObjectID().Hex(),
		BirdId:   primitive.NewObjectID().Hex(),
		Location: "TestLocation",
		ImageId:  primitive.NewObjectID().Hex(),
		SoundId:  primitive.NewObjectID().Hex(),
	}

	testPost2 := &models.PostDB{
		UserId:   primitive.NewObjectID().Hex(),
		BirdId:   primitive.NewObjectID().Hex(),
		Location: "TestLocation",
		ImageId:  primitive.NewObjectID().Hex(),
		SoundId:  primitive.NewObjectID().Hex(),
	}

	t.Run("Test CreatePost", func(t *testing.T) {
		response := postColl.CreatePost(*testPost1)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testPost1.Id = response.Data.(string)
		response = postColl.CreatePost(*testPost2)
		require.False(t, utils.IsTypeError(response))
		require.IsType(t, "string", response.Data.(string))
		testPost2.Id = response.Data.(string)
	})

	t.Run("Test GetPostById", func(t *testing.T) {
		response := postColl.GetPostById(testPost1.Id)
		require.Equal(t, testPost1.Id, response.Data.(*models.PostDB).Id)
		testPost1 = response.Data.(*models.PostDB)

		response = postColl.GetPostById("IncorrectID")
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test ListPosts & ListAllPosts", func(t *testing.T) {
		response := postColl.ListPosts(bson.M{}, 0)
		require.Equal(t, 2, len(response.Data.([]models.PostDB)))
		responseAll := postColl.ListAllPosts(0)
		require.Equal(t, response.Data.([]models.PostDB), responseAll.Data.([]models.PostDB))
		response = postColl.ListPosts(bson.M{"_id": testPost2.Id}, 0)
		require.Equal(t, 1, len(response.Data.([]models.PostDB)))
		response = postColl.ListPosts(bson.M{"_id": "IncorrectId"}, 0)
		require.Equal(t, 0, len(response.Data.([]models.PostDB)))
	})

	t.Run("Test UpdatePost", func(t *testing.T) {
		updatePost := models.PostInput{
			Id:       testPost1.Id,
			Location: "Location update",
		}
		response := postColl.UpdatePost(updatePost)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, updatePost.Location, response.Data.(*models.PostDB).Location)

		response = postColl.UpdatePost(updatePost)
		require.True(t, utils.IsTypeError(response))
	})

	t.Run("Test DeletePost", func(t *testing.T) {
		response := postColl.DeletePost(testPost1.Id)
		require.False(t, utils.IsTypeError(response))
		response = postColl.DeletePost(testPost1.Id)
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
		_, err = db.Collection(repositories.PostColl).DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			return
		}
	})
}
