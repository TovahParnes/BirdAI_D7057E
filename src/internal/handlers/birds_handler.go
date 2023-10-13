package handlers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// GetBirdById is a function to get a bird by ID
//
// @Summary		Get bird by ID
// @Description	Get bird by ID
// @Tags		Birds
// @Accept		json
// @Produce		json
// @Param		id	path	string	true	"Bird ID"
// @Success		200	{object}	models.Response{data=models.BirdOutput}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		410	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/birds/{id} [get]
func (h *Handler) GetBirdById(c *fiber.Ctx) error {
	id := c.Params("id")
	response := utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CGetBirdById(id)
	return utils.ResponseToStatus(c, response)
}

// ListBirds is a function to get a set of all birds from database, with optional search parameters
//
// @Summary		List all birds of a specified set and seach parameters
// @Description	List all birds of a specified set and seach parameters
// @Tags		Birds
// @Accept		json
// @Produce		json
// @Param		set	query		int	false	"Set of birds"
// @Param		search	query	string	false	"Search parameter for birds"
// @Success		200	{object}	models.Response{data=[]models.BirdOutput}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/birds/list [get]
func (h *Handler) ListBirds(c *fiber.Ctx) error {

	//	@Failure	503	{object}	models.Response{}
	// if no connection to db was established

	//	@Failure	404	{object}	models.Response{}
	// if birds not found
	
	
	queries := c.Queries()
	set := queries["set"]
	response := utils.IsValidSet(&set)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	search := queries["search"]
	response = utils.IsValidSearch(search)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CListBirds(set, search)
	return utils.ResponseToStatus(c, response)
}


// UpdateBird is a function to update the given bird from the databse
//
// @Summary		Update given bird
// @Description	Update given bird
// @Tags		Birds
// @Accept		json
// @Produce		json
// @Security 	Bearer
// @Param		id	path	string	true	"Bird ID"
// @Param		bird	body		models.BirdInput	true	"bird"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/birds/{id} [patch]
func (h *Handler) UpdateBird(c *fiber.Ctx) error {
response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	id := c.Params("id")
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	var bird *models.BirdInput
	if err := c.BodyParser(&bird); err != nil {
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response = utils.IsValidBirdInput(bird)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	//	@Failure		403	{object}	models.Response{}
	// if user is not admin

	response = h.controller.CUpdateBird(id, bird)
	return utils.ResponseToStatus(c, response)
}
