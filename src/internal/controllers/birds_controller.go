package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

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
	response := c.db.Bird.UpdateBird(id, *bird)
	if utils.IsTypeError(response) {
		return response
	}
	birdDB := response.Data.(*models.BirdDB)
	birdResponse := c.CBirdDBToOutput(birdDB)
	return birdResponse
}

func (c *Controller) CBirdDBToOutput(bird *models.BirdDB) models.Response {
	if bird.Description == "" {
		return utils.ErrorParams("Description cannot be empty")
	}
	birdOutput := models.BirdDBToOutput(bird)
	return utils.Response(birdOutput)
}
