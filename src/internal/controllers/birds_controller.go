package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetBirdById(id string) models.Response {
	response := c.db.Bird.GetBirdById(id)
	if utils.IsTypeError(response) {
		return response
	}
	bird := response.Data.(*models.BirdDB)
	birdResponse := c.CBirdDBToOutput(bird)
	return birdResponse
}

func (c *Controller) CGetBirdByName(name string) models.Response {
	response := c.db.Bird.GetBirdByName(name)
	if utils.IsTypeError(response) {
		return response
	}
	bird := response.Data.(*models.BirdDB)
	birdResponse := c.CBirdDBToOutput(bird)
	return birdResponse
}

func (c *Controller) CListBirds(set int, search string) models.Response {
	filter := bson.M{"name": bson.M{"$regex": "(?i)"+search}}
	response := c.db.Bird.ListBirds(filter, set)
	if utils.IsTypeError(response) {
		return response
	}

	output := []*models.BirdOutput{}
	for _, bird := range response.Data.([]models.BirdDB) {
		birdResponse := c.CBirdDBToOutput(&bird)
		if utils.IsTypeError(birdResponse) {
			return birdResponse
		}

		output = append(output, birdResponse.Data.(*models.BirdOutput))
	}

	return utils.Response(output)
}

func (c *Controller) CUpdateBird(id string, bird *models.BirdInput) (models.Response) {
	media := c.db.Media.GetMediaById(bird.ImageId)
	if utils.IsTypeError(media) {
		if media.Data.(models.Err).StatusCode == http.StatusNotFound{
			return utils.ErrorNotFoundInDatabase("Image with given id does not exist")
		} else {
		return media
		}
	}

	media = c.db.Media.GetMediaById(bird.SoundId)
	if utils.IsTypeError(media) {
		if media.Data.(models.Err).StatusCode == http.StatusNotFound{
			return utils.ErrorNotFoundInDatabase("Sound with given id does not exist")
		} else {
		return media
		}
	}
	
	response := c.db.Bird.UpdateBird(id, *bird)
	if utils.IsTypeError(response) {
		return response
	}
	birdDB := response.Data.(*models.BirdDB)
	birdResponse := c.CBirdDBToOutput(birdDB)
	return birdResponse
}

func (c *Controller) CBirdDBToOutput(bird *models.BirdDB) models.Response {
	imageResponse := c.db.Media.GetMediaById(bird.ImageId)
	if utils.IsTypeError(imageResponse) {
		return imageResponse
	}
	imageOutput := models.MediaDBToOutput(imageResponse.Data.(*models.MediaDB))
	soundResponse := c.db.Media.GetMediaById(bird.SoundId)
	if utils.IsTypeError(soundResponse) {
		return soundResponse
	}
	soundOutput := models.MediaDBToOutput(soundResponse.Data.(*models.MediaDB))
	birdOutput := models.BirdDBToOutput(bird, imageOutput, soundOutput)
	return utils.Response(birdOutput)
}