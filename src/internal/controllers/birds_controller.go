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
	response := c.db.Bird.ListBirds(bson.M{}, set)
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

func (c *Controller) CUpdateBird(id string, bird *models.BirdInput) models.Response {
	bird.Id = id
	response := c.db.Bird.UpdateBird(*bird)
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

func (c *Controller) GenerateBirds() models.Response {
	currentBirds := c.db.Bird.ListAllBirds(0)
	if utils.IsTypeError(currentBirds) && currentBirds.Data.(models.Err).StatusCode != http.StatusNotFound {
		return currentBirds
	}
	if len(currentBirds.Data.([]models.BirdDB)) != 0 {
		for _, bird := range currentBirds.Data.([]models.BirdDB) {
			c.db.Bird.DeleteBird(bird.Id)
		}
	}

	response := c.db.Media.CreateMedia(models.MediaDB{
		Data:     "testImage",
		FileType: "image/png",
	})
	if utils.IsTypeError(response) {
		return response
	}
	imageId := response.Data.(string)
	response = c.db.Media.CreateMedia(models.MediaDB{
		Data:     "testSound",
		FileType: "audio/mpeg",
	})
	if utils.IsTypeError(response) {
		return response
	}
	soundId := response.Data.(string)
	response = c.db.Bird.CreateBird(models.BirdDB{
		Name:        "Skata",
		Description: "Cool test bird",
		ImageId:     imageId,
		SoundId:     soundId,
	})
	return response
}
