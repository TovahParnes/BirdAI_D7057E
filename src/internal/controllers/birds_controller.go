package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetBirdById(id string) (models.Response) {
	response := c.db.Bird.GetBirdById(id)
	if utils.IsTypeError(response) {
		return response
	}
	bird := response.Data.(*models.BirdDB)
	birdResponse := c.BirdDBToOutput(bird)
	if utils.IsTypeError(birdResponse) {
		return birdResponse
	}
	return utils.Response(birdResponse.Data.(models.BirdOutput))
}

func (c *Controller) CListBirds(set int, search string) (models.Response) {
	response := c.db.Bird.ListBirds(bson.M{}, set)
	if utils.IsTypeError(response) {
		return response
	}

	output := []models.BirdOutput{}
	for _, bird := range response.Data.([]*models.BirdDB) {
		birdResponse := c.BirdDBToOutput(bird)
		if utils.IsTypeError(birdResponse) {
			return birdResponse
		}
		output = append(output, birdResponse.Data.(models.BirdOutput))
	}

	return utils.Response(output)
}

func (c *Controller) CUpdateBird(id string, bird *models.BirdInput) (models.Response) {
	bird.Id = id
	response := c.db.Bird.UpdateBird(*bird)
	if utils.IsTypeError(response) {
		return response
	}
	birdDB := response.Data.(*models.BirdDB)
	birdResponse := c.BirdDBToOutput(birdDB)
	return birdResponse
}


func (c *Controller) BirdDBToOutput(bird *models.BirdDB) (models.Response) {
	imageResponse := c.db.Media.GetMediaById(bird.ImageId)
	if utils.IsTypeError(imageResponse) {
		return imageResponse
	}
	imageOutput := models.MediaDBToOutput(imageResponse.Data.(models.MediaDB))
	soundResponse := c.db.Media.GetMediaById(bird.SoundId)
	if utils.IsTypeError(soundResponse) {
		return soundResponse
	}
	soundOutput := models.MediaDBToOutput(soundResponse.Data.(models.MediaDB))
	birdOutput := models.BirdDBToOutput(bird, imageOutput, soundOutput)
	return utils.Response(birdOutput)

}
