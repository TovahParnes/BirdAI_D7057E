package models_test

import (
	"birdai/src/internal/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModels(t *testing.T) {
	t.Run("AdminModels", func(t *testing.T) {
		db := &models.AdminDB{
			Id:     "5f9d3b3b9d3b3b9d3b3b9d3b",
			UserId: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Access: "Admin",
		}
		userOutput := &models.UserOutput{
			Id:       "5f9d3b3b9d3b3b9d3b3b9d3b",
			Username: "test",
			CreatedAt: "2020-10-30T00:00:00Z",
			Active: true,
		}
		out := models.AdminDBToOutput(db, userOutput)
		require.IsType(t, &models.AdminOutput{}, out)
	})

	t.Run("BirdModels", func(t *testing.T) {
		db := &models.BirdDB{
			Id:		"5f9d3b3b9d3b3b9d3b3b9d3b",
			Name:   "test",
			Description: "test",
			ImageId: "5f9d3b3b9d3b3b9d3b3b9d3b",
			SoundId: "5f9d3b3b9d3b3b9d3b3b9d3b",
		}
		media := &models.MediaOutput{
			Id: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Data: "test",
			FileType: "sound",
		}
		out := models.BirdDBToOutput(db, media, media)
		require.IsType(t, &models.BirdOutput{}, out)
	})

	t.Run("MediaModels", func(t *testing.T) {
		db := &models.MediaDB{
			Id: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Data: "test",
			FileType: "sound",
		}
		out := models.MediaDBToOutput(db)
		require.IsType(t, &models.MediaOutput{}, out)
	})

	t.Run("PostModels", func(t *testing.T) {
		db := &models.PostDB{
			Id: "5f9d3b3b9d3b3b9d3b3b9d3b",
			UserId: "5f9d3b3b9d3b3b9d3b3b9d3b",
			BirdId: "5f9d3b3b9d3b3b9d3b3b9d3b",
		}
		userOutput := &models.UserOutput{
			Id:       "5f9d3b3b9d3b3b9d3b3b9d3b",
			Username: "test",
			CreatedAt: "2020-10-30T00:00:00Z",
			Active: true,
		}
		media := &models.MediaOutput{
			Id: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Data: "test",
			FileType: "sound",
		}
		bird := &models.BirdOutput{
			Id:		"5f9d3b3b9d3b3b9d3b3b9d3b",
			Name:   "test",
			Description: "test",
			Image: *media,
			Sound: *media,
		}
		out := models.PostDBToOutput(db, userOutput, bird, media)
		require.IsType(t, &models.PostOutput{}, out)
	})

	t.Run("UserModels", func(t *testing.T) {
		db := &models.UserDB{
			Id: "5f9d3b3b9d3b3b9d3b3b9d3b",
			Username: "test",
			AuthId: "5f9d3b3b9d3b3b9d3b3b9d3b",
			CreatedAt: "2020-10-30T00:00:00Z",
			Active: true,
		}
		out := models.UserDBToOutput(db)
		require.IsType(t, &models.UserOutput{}, out)
	})
}