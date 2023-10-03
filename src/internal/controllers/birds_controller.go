package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
)

func (c *Controller) CGetBirdById(id string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.BirdColl)
	filter := bson.M{"_id": id}
	response := coll.FindOne(filter)
	return response
	*/
	return utils.ErrorNotImplemented("CGetBirdById")
}

func (c *Controller) CListBirds(set string, search string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.BirdColl)
	response := coll.FindAll()
	return response
	*/
	return utils.ErrorNotImplemented("CGetBirdById")
}

func (c *Controller) CUpdateBird(id string, bird *models.BirdInput) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.BirdColl)
	response := coll.UpdateOne(bson.M{
		"name": bird.Name,
		"description": bird.Description,
		"image_id": bird.ImageId,
		"sound_id": bird.SoundId,
	})
	return response
	*/
	return utils.ErrorNotImplemented("CGetBirdById")
}
