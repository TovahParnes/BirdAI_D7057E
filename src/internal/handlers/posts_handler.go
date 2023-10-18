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
// @Success		200	{object}	models.Response{data=[]models.PostOutput}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		410	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router		/posts/{id} [get]
func (h *Handler) GetPostById(c *fiber.Ctx) error {
	id := c.Params("id")
	response := h.controller.CGetPostById(id)

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
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router		/posts/list [get]
func (h *Handler) ListPosts(c *fiber.Ctx) error {
	//authId := c.GetReqHeaders()["Authid"]
	queries := c.Queries()
	set := queries["set"]
	search := queries["search"]

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	response := h.controller.CListPosts(set, search)
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
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router		/users/{id}/posts/list [get]
func (h *Handler) ListUsersPosts(c *fiber.Ctx) error {
	//authId := c.GetReqHeaders()["Authid"]
	userId := c.Params("id")
	queries := c.Queries()
	set := queries["set"]
	search := queries["search"]

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if post not found

	response := h.controller.CListUsersPosts(userId, set, search)
	return utils.ResponseToStatus(c, response)
}

// CreatePost is a function to create a new post
//
// @Summary		Create a new post
// @Description	Create a new post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		set	body		models.PostInput	true	"post"
// @Success		201	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router		/posts/ [post]
func (h *Handler) CreatePost(c *fiber.Ctx) error {
	authId := c.GetReqHeaders()["Authid"]
	var post *models.PostInput
	if err := c.BodyParser(&post); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}

	response := h.controller.CCreatePost(authId, post)

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
// @Param		id	path	string	true	"post ID"
// @Param		post	body		models.PostInput	true	"post"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		403	{object}	models.Response{data=[]models.Err}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/posts/{id} [patch]
func (h *Handler) UpdatePost(c *fiber.Ctx) error {
	//authId := c.GetReqHeaders()["Authid"]
	id := c.Params("id")

	var post *models.PostInput
	if err := c.BodyParser(&post); err != nil {
		//	@Failure	400	{object}	models.Response{}
		// something with body is wrong/missing
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or post is not the same as the one being updated

	response := h.controller.CUpdatePost(id, post)
	return utils.ResponseToStatus(c, response)
}

// DeletePost is a function to delete the given post from the database
//
// @Summary		Delete given post
// @Description	Delete given post
// @Tags		Posts
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"Post ID"
// @Success		200	{object}	models.Response{}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		403	{object}	models.Response{data=[]models.Err}
// @Failure		404	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Router			/posts/{id} [delete]
func (h *Handler) DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	authId := c.GetReqHeaders()["Authid"]

	//	@Failure	401	{object}	models.Response{}
	// Authenticate(jwt.token)

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin or user is not the same as the one being updated

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if user not found

	response := h.controller.CDeletePost(id, authId)
	return utils.ResponseToStatus(c, response)
}
