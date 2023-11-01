package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
)

func (c *Controller) CGetPostById(id string) models.Response {
	response := c.db.Post.GetPostById(id)

	if utils.IsTypeError(response) {
		return response
	}

	post := response.Data.(*models.PostDB)
	postOutput := c.CPostDBToOutput(*post)
	return postOutput
}

func (c *Controller) CListPosts(set int, search string) models.Response {
	response := c.db.Post.ListPosts(bson.M{}, set)
	posts := []*models.PostOutput{}

	if utils.IsTypeError(response) {
		return response
	}

	for _, postsObject := range response.Data.([]models.PostDB) {
		postResponse := c.CPostDBToOutput(postsObject)
		if utils.IsTypeError(postResponse) {
			return postResponse
		}
		posts = append(posts, postResponse.Data.(*models.PostOutput))
	}

	return utils.Response(posts)
}

func (c *Controller) CListUsersPosts(userId string, set int) models.Response {
	filter := bson.M{"user_id": userId}
	//birdColl := c.db.GetCollection(repositories.BirdColl)
	response := c.db.Post.ListPosts(filter, set)
	output := []*models.PostOutput{}
	for _, post := range response.Data.([]models.PostDB) {
		postOutput := c.CPostDBToOutput(post)
		if utils.IsTypeError(postOutput) {
			return postOutput
		}
		output = append(output, postOutput.Data.(*models.PostOutput))
	}
	return utils.Response(output)
}

func (c *Controller) CCreatePost(userId string, postInput *models.PostInput) models.Response {
	media := &models.MediaDB{
		Data:     postInput.Media.Data,
		FileType: postInput.Media.FileType,
	}
	response := c.db.Media.CreateMedia(*media)
	if utils.IsTypeError(response) {
		return response
	}
	post := &models.PostDB{
		UserId:   userId,
		BirdId:   postInput.BirdId,
		Location: postInput.Location,
		MediaId:  response.Data.(string),
	}
	response = c.db.Post.CreatePost(*post)
	return response
}

// TODO need to change media also but now it can change location

func (c *Controller) CUpdatePost(postId string, post *models.PostInput) models.Response {
	post.Id = postId
	response := c.db.Post.UpdatePost(*post)
	if utils.IsTypeError(response) {
		return response
	}
	postDB := response.Data.(*models.PostDB)
	postOutput := c.CPostDBToOutput(*postDB)
	return postOutput
}

func (c *Controller) CDeletePost(id string) models.Response {
	deletePostResponse := c.db.Post.DeletePost(id)
	if utils.IsTypeError(deletePostResponse) {
		return deletePostResponse
	}
	deleteMediaResponse := c.db.Media.DeleteMedia(deletePostResponse.Data.(*models.PostDB).MediaId)
	return deleteMediaResponse
}

func (c *Controller) CPostDBToOutput(post models.PostDB) (models.Response) {
	user := c.CGetUserById(post.UserId)
	if utils.IsTypeError(user) {
		return user
	}
	userOutput := user.Data.(*models.UserOutput)

	bird := c.CGetBirdById(post.BirdId)
	if utils.IsTypeError(bird) {
		return bird
	}
	birdOutput := bird.Data.(*models.BirdOutput)

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
	/*
	birdOutput := models.BirdOutput{
		Id:          "651eb6aa9dd12b111952d7b2",
		Name:        "testbird",
		Description: "Cool test bird",
	}
	*/

	userImage := c.db.Media.GetMediaById(post.MediaId)
	if utils.IsTypeError(userImage) {
		return userImage
	}
	userImageOutput := models.MediaDBToOutput(userImage.Data.(*models.MediaDB))
	postOutput := models.PostDBToOutput(&post, userOutput, birdOutput, userImageOutput)
	return utils.Response(postOutput)
}
