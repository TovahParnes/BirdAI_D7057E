package controllers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"net/http"

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
	filter := bson.M{}
	
	if search != "" {
		filter = bson.M{
			 "$or": [
				bson.M{"name": { "$regex": "^Da"}}
		]}

			"$or": [
				{"location": {"$regex": "(?i)"+search}},
				bson.M{"comment": bson.M{"$regex": "(?i)"+search}},
			]
			
			bson.M{
				"location": bson.M{"$regex": "(?i)"+search},
				"comment": bson.M{"$regex": "(?i)"+search},
			},
		}
	}
	response := c.db.Post.ListPosts(filter, set)
	posts := []*models.PostOutput{}

	if utils.IsTypeError(response) {
		return response
	}

	//TEMPORATY - restarting fiber changes the bird id, meaning listing posts will
	//always from error that bird does not exisits
	for _, postsObject := range response.Data.([]models.PostDB) {
		fmt.Println(postsObject.Id)
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

func (c *Controller) CListUsersFoundBirds(userId string, set int) models.Response {
	filter := bson.M{"user_id": userId}
	response := c.db.Post.ListPosts(filter, set)
	output := []*models.PostDB{}
	for _, post := range response.Data.([]models.PostDB) {
		output = append(output, &post)
	}
	return utils.Response(output)
}

func (c *Controller) CCreatePost(userId string, post *models.PostCreation) models.Response {
	bird := c.db.Bird.GetBirdById(post.BirdId)
	if utils.IsTypeError(bird) {
		if bird.Data.(models.Err).StatusCode == http.StatusNotFound {
			return utils.ErrorToResponse(http.StatusBadRequest, "Bird not found", "Bird with that id does not exist")
		}
		return bird
	}

	media := &models.MediaDB{
		Data:     post.Media.Data,
	}
	response := c.db.Media.CreateMedia(*media)
	if utils.IsTypeError(response) {
		return response
	}

	postDB := models.PostCreationToDB(userId, post, response.Data.(string))

	response = c.db.Post.CreatePost(*postDB)
	if utils.IsTypeError(response) {
		return response
	}
	return c.CGetPostById(response.Data.(string))
}

// TODO need to change media also but now it can change location

func (c *Controller) CUpdatePost(postId string, post *models.PostInput) models.Response {
	response := c.db.Post.UpdatePost(postId, *post)
	if utils.IsTypeError(response) {
		return response
	}
	postDB := response.Data.(*models.PostDB)
	postOutput := c.CPostDBToOutput(*postDB)
	return postOutput
}

func (c *Controller) CDeletePost(id string) models.Response {
	post := c.db.Post.GetPostById(id)
	if utils.IsTypeError(post) {
		return post
	}

	deletePostResponse := c.db.Post.DeletePost(id)
	if utils.IsTypeError(deletePostResponse) {
		return deletePostResponse
	}

	deleteMediaResponse := c.db.Media.DeleteMedia(post.Data.(*models.PostDB).MediaId)
	return deleteMediaResponse
}

func (c *Controller) CPostDBToOutput(post models.PostDB) models.Response {
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

	userImage := c.db.Media.GetMediaById(post.MediaId)
	if utils.IsTypeError(userImage) {
		return userImage
	}
	userImageOutput := models.MediaDBToOutput(userImage.Data.(*models.MediaDB))
	postOutput := models.PostDBToOutput(&post, userOutput, birdOutput, userImageOutput)
	return utils.Response(postOutput)
}
