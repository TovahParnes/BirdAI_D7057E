package utils_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParameterValidation(t *testing.T) {

	t.Run("Test IsValidId", func(t *testing.T) {
		response := utils.IsValidId("5f9d3b3b9d3b3b9d3b3b9d3b")
		require.False(t, utils.IsTypeError(response))
		response = utils.IsValidId("5f9d3b3b9d3b3b9d3b3b9d3")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidSet", func(t *testing.T) {
		str := "5"
		response := utils.IsValidSet(&str)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 5, response.Data.(int))

		str = "5.5"
		response = utils.IsValidSet(&str)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		str = ""
		response = utils.IsValidSet(&str)
		require.False(t, utils.IsTypeError(response))
		require.Equal(t, 0, response.Data.(int))

		str = "hello"
		response = utils.IsValidSet(&str)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidSearch", func(t *testing.T) {
		response := utils.IsValidSearch("hello")
		require.False(t, utils.IsTypeError(response))
		response = utils.IsValidSearch("hello!")
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidAdminCreation", func(t *testing.T) {
		admin := &models.AdminCreation{
			UserId: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Access: "admin",
		}
		response := utils.IsValidAdminCreation(admin)
		require.False(t, utils.IsTypeError(response))

		admin.Access = "superadmin"
		response = utils.IsValidAdminCreation(admin)
		require.False(t, utils.IsTypeError(response))

		admin.UserId = "1234"
		response = utils.IsValidAdminCreation(admin)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		admin.Access = "hello"
		response = utils.IsValidAdminCreation(admin)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidAdminInput", func(t *testing.T) {
		admin := &models.AdminInput{
			Access: "admin",
		}
		response := utils.IsValidAdminInput(admin)
		require.False(t, utils.IsTypeError(response))

		admin.Access = "superadmin"
		response = utils.IsValidAdminInput(admin)
		require.False(t, utils.IsTypeError(response))

		admin.Access = "hello"
		response = utils.IsValidAdminInput(admin)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidBirdInput", func(t *testing.T) {

		bird := &models.BirdInput{
			Name:        "bird",
			Description: "bird",
			SoundId:     "5f9d3b3b9d3b3b9d3b3b9d3b",
		}
		response := utils.IsValidBirdInput(bird)
		require.False(t, utils.IsTypeError(response))

		bird.Name = ""
		response = utils.IsValidBirdInput(bird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		bird.Name = "bird"
		bird.Description = ""
		response = utils.IsValidBirdInput(bird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		bird.Description = "bird"
		response = utils.IsValidBirdInput(bird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		bird.SoundId = "1234"
		response = utils.IsValidBirdInput(bird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		bird.SoundId = "5f9d3b3b9d3b3b9d3b3b9d3b"
		bird.Name = "bird!"
		response = utils.IsValidBirdInput(bird)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidMediaInput", func(t *testing.T) {
	})

	t.Run("Test IsValidPostCreation", func(t *testing.T) {
		post := &models.PostCreation{
			BirdId:   "5f9d3b3b9d3b3b9d3b3b9d3b",
			Location: "location",
			Comment:  "comment",
			Accuracy: 0.5,
		}
		response := utils.IsValidPostCreation(post)
		require.False(t, utils.IsTypeError(response))

		post.BirdId = "1234"
		response = utils.IsValidPostCreation(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		post.Location = ""
		response = utils.IsValidPostCreation(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		post.Location = "location"
		post.Comment = ""
		response = utils.IsValidPostCreation(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		post.Comment = "comment"
		post.Accuracy = 1.2
		response = utils.IsValidPostCreation(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		post.Accuracy = -0.2
		response = utils.IsValidPostCreation(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

	})
	t.Run("Test IsValidPostInput", func(t *testing.T) {
		post := &models.PostInput{
			Location: "location",
			Comment:  "comment",
		}
		response := utils.IsValidPostInput(post)
		require.False(t, utils.IsTypeError(response))

		post.Location = ""
		response = utils.IsValidPostInput(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		post.Location = "location"
		post.Comment = ""
		response = utils.IsValidPostInput(post)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

	})

	t.Run("Test IsValidUserInput", func(t *testing.T) {
		user := &models.UserInput{
			Username: "username",
			Active:   true,
		}
		response := utils.IsValidUserInput(user)
		require.False(t, utils.IsTypeError(response))

		user.Active = false
		response = utils.IsValidUserInput(user)
		require.False(t, utils.IsTypeError(response))

		user.Username = "username!"
		response = utils.IsValidUserInput(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		user.Username = ""
		response = utils.IsValidUserInput(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		user.Username = "us"
		response = utils.IsValidUserInput(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		user.Username = "usernameusernameusernameusernameusername"
		response = utils.IsValidUserInput(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})

	t.Run("Test IsValidUserLogin", func(t *testing.T) {
		user := &models.UserLogin{
			Username: "username",
		}
		response := utils.IsValidUserLogin(user)
		require.False(t, utils.IsTypeError(response))

		user.Username = "username!"
		response = utils.IsValidUserLogin(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		user.Username = ""
		response = utils.IsValidUserLogin(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		user.Username = "us"
		response = utils.IsValidUserLogin(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)

		user.Username = "usernameusernameusernameusernameusername"
		response = utils.IsValidUserLogin(user)
		require.True(t, utils.IsTypeError(response))
		require.Equal(t, http.StatusBadRequest, response.Data.(models.Err).StatusCode)
	})
}
