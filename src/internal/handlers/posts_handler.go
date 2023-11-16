package handlers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// GetPostById is a function to get a pst by ID
//
// @Summary		Get post by ID
// @Description	Get post by ID
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"Post ID"
// @Success		200	{object}	models.Response{data=models.PostOutput}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		410	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/posts/{id} [get]
func (h *Handler) GetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	response := utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CGetPostById(id)

	return utils.ResponseToStatus(c, response)
}

// ListPosts is a function to get a set of all posts from database
//
// @Summary		List all posts of a specified set
// @Description	List all posts of a specified set
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		set	query		int	false	"Set of posts"
// @Param		search	query	string	false	"Search parameter for post"
// @Success		200	{object}	models.Response{data=[]models.PostOutput}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/posts/list [get]
func (h *Handler) ListPosts(c *fiber.Ctx) error {
	queries := c.Queries()
	set := queries["set"]
	response := utils.IsValidSet(&set)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	setInt := response.Data.(int)

	search := queries["search"]
	response = utils.IsValidSearch(search)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	response = h.controller.CListPosts(setInt, search)
	return utils.ResponseToStatus(c, response)
}

// ListUsersPosts is a function to get a set of all posts from a given user from database
//
// @Summary		List all posts of a specified set
// @Description	List all posts of a specified set
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"User ID"
// @Param		set	query		int	false	"Set of posts"
// @Param		search	query	string	false	"Search parameter for post"
// @Success		200	{object}	models.Response{data=[]models.PostOutput}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/users/{id}/posts/list [get]
func (h *Handler) ListUsersPosts(c *fiber.Ctx) error {
	userId := c.Params("id")
	response := utils.IsValidId(userId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	queries := c.Queries()
	set := queries["set"]
	response = utils.IsValidSet(&set)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	setInt := response.Data.(int)

	search := queries["search"]
	response = utils.IsValidSearch(search)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CListUsersPosts(userId, setInt)
	return utils.ResponseToStatus(c, response)
}

// CreatePost is a function to create a new post
//
// @Summary		Create a new post
// @Description	Create a new post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Param		post	body	models.PostCreation	true	"post"
// @Success		201	{object}	models.Response{data=models.PostOutput}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/posts [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	var post *models.PostCreation
	if err := c.BodyParser(&post); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response = utils.IsValidPostCreation(post)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CCreatePost(curUserId, post)

	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	return utils.CreationResponseToStatus(c, response)
}

// UpdatePost is a function to update the given post from the databse
//
// @Summary		Update given post
// @Description	Update given post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Param		id	path	string	true	"post ID"
// @Param		post	body		models.PostInput	true	"post"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/posts/{id} [patch]
func (h *Handler) UpdatePost(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	id := c.Params("id")
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CIsPostsUserOrAdmin(curUserId, id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	var post *models.PostInput
	if err := c.BodyParser(&post); err != nil {
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response = utils.IsValidPostInput(post)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or post is not the same as the one being updated

	response = h.controller.CUpdatePost(id, post)
	return utils.ResponseToStatus(c, response)
}

// DeletePost is a function to delete the given post from the database
//
// @Summary		Delete given post
// @Description	Delete given post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Param		id	path	string	true	"Post ID"
// @Success		200	{object}	models.Response{data=string}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/api/v1/posts/{id} [delete]
func (h *Handler) DeletePost(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	id := c.Params("id")
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CIsPostsUserOrAdmin(curUserId, id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CDeletePost(id)
	return utils.ResponseToStatus(c, response)
}
