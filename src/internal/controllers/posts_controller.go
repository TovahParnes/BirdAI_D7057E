package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetPostById(id string) models.Response {
	/*
		coll := c.db.GetCollection(repositories.PostColl)
		filter := bson.M{"_id": id}
		response := coll.FindOne(filter)
		return response
	*/
	return utils.ErrorNotImplemented("CGetPostById")
}

func (c *Controller) CListPosts(set string, search string) models.Response {
	/*
		coll := c.db.GetCollection(repositories.PostColl)
		response := coll.FindAll()
		return response
	*/
	return utils.ErrorNotImplemented("CListPosts")
}

func (c *Controller) CListUsersPosts(userId string, set string, search string) models.Response {
	filter := bson.M{"_id": userId}
	postColl := c.db.GetCollection(repositories.PostColl)
	mediaColl := c.db.GetCollection(repositories.MediaColl)
	userColl := c.db.GetCollection(repositories.UserColl)
	//birdColl := c.db.GetCollection(repositories.BirdColl)
	response := postColl.FindAll(filter, 0, 0)
	output := []models.PostOutput{}
	for _, post := range response.Data.([]*models.PostDB) {
		user := userColl.FindOne(bson.M{"_id": post.UserId})
		userOutput := models.UserOutput{
			Id:        user.Data.(*models.UserDB).Id,
			Username:  user.Data.(*models.UserDB).Username,
			CreatedAt: user.Data.(*models.UserDB).CreatedAt,
			Active:    user.Data.(*models.UserDB).Active,
		}
		//bird := birdColl.FindOne(bson.M{"_id": post.BirdId})
		//image := mediaColl.FindOne(bson.M{"_id": bird.Data.(*models.BirdDB).ImageId})
		//imageOutput := models.MediaOutput{
		//	Id:       image.Data.(*models.MediaDB).Id,
		//	Data:     []byte(image.Data.(*models.MediaDB).Data),
		//	FileType: image.Data.(*models.MediaDB).FileType,
		//}
		//sound := mediaColl.FindOne(bson.M{"_id": bird.Data.(*models.BirdDB).SoundId})
		//soundOutput := models.MediaOutput{
		//	Id:       sound.Data.(*models.MediaDB).Id,
		//	Data:     []byte(sound.Data.(*models.MediaDB).Data),
		//	FileType: sound.Data.(*models.MediaDB).FileType,
		//}
		//birdOutput := models.BirdOutput{
		//	Id:          bird.Data.(*models.BirdDB).Id,
		//	Name:        bird.Data.(*models.BirdDB).Name,
		//	Description: bird.Data.(*models.BirdDB).Description,
		//	Image:       imageOutput,
		//	Sound:       soundOutput,
		//}

		//TODO fix static birds
		birdOutput := models.BirdOutput{
			Id:          "651eb6aa9dd12b111952d7b2",
			Name:        "testbird",
			Description: "Cool test bird",
		}

		userImage := mediaColl.FindOne(bson.M{"_id": post.MediaId})
		userImageOutput := models.MediaOutput{
			Id:       userImage.Data.(*models.MediaDB).Id,
			Data:     []byte(userImage.Data.(*models.MediaDB).Data),
			FileType: userImage.Data.(*models.MediaDB).FileType,
		}

		output = append(output, models.PostOutput{
			Id:        post.Id,
			User:      userOutput,
			Bird:      birdOutput,
			CreatedAt: post.CreatedAt,
			Location:  post.Location,
			UserMedia: userImageOutput,
		})
	}
	return utils.Response(output)
}

func (c *Controller) CCreatePost(userId string, postInput *models.PostInput) models.Response {
	mediaColl := c.db.GetCollection(repositories.MediaColl)
	media := &models.MediaDB{
		Data:     postInput.Media.Data,
		FileType: postInput.Media.FileType,
	}
	response := mediaColl.CreateOne(media)
	if utils.IsTypeError(response) {
		return response
	}
	post := &models.PostDB{
		UserId:   userId,
		BirdId:   postInput.BirdId,
		Location: postInput.Location,
		MediaId:  response.Data.(*models.MediaDB).Id,
	}
	postColl := c.db.GetCollection(repositories.PostColl)
	response = postColl.CreateOne(post)
	return response
}

func (c *Controller) CUpdatePost(userId string, post *models.PostInput) models.Response {
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

func (c *Controller) CDeletePost(id string) models.Response {
	/*
		coll := c.db.GetCollection(repositories.PostColl)
		filter := bson.M{"_id": id}
		response := coll.DeleteOne(filter)
		return response
	*/
	return utils.ErrorNotImplemented("CDeletePost")
}
