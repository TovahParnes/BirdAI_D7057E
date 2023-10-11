package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
)

func (c *Controller) CGetPostById(id string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.PostColl)
	filter := bson.M{"_id": id}
	response := coll.FindOne(filter)
	return response
	*/
	return utils.ErrorNotImplemented("CGetPostById")
}

func (c *Controller) CListPosts(set string, search string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.PostColl)
	response := coll.FindAll()
	return response
	*/
	return utils.ErrorNotImplemented("CListPosts")
}

func (c *Controller) CListUsersPosts(userId string, set string, search string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.PostColl)
	response := coll.FindAll()
	return response
	*/
	return utils.ErrorNotImplemented("CListUsersPosts")
}

func (c *Controller) CCreatePost(authId string, post *models.Post) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.PostColl)
	response := coll.CreateOne(post)
	return response
	*/
	return utils.ErrorNotImplemented("CCreatePost")
}

func (c *Controller) CUpdatePost(userId string, post *models.Post) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.UserColl)
	response := coll.UpdateOne(bson.M{
		"_id": post.Id,
		"user_id": userId,
		"bird_id": post.BirdId,
		"created_at": post.CreatedAt,
		"location": post.Location,
		"image_id": post.ImageId,
		"sound_id": post.SoundId,
	})
	return response
	*/
	return utils.ErrorNotImplemented("CUpdatePost")
}

func (c *Controller) CDeletePost(id, authId string) (models.Response) {
	/*
	coll := c.db.GetCollection(repositories.PostColl)
	filter := bson.M{"_id": id, "user_id": authId}
	response := coll.DeleteOne(filter)
	return response
	*/
	return utils.ErrorNotImplemented("CDeletePost")
}