package handlers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// GetAdminById is a function to get a pst by ID
//
// @Summary		Get admin by ID
// @Description	Get admin by ID
// @Tags		Admins
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Param		id	path	string	true	"Admin ID"
// @Success		200	{object}	models.Response{data=models.AdminOutput}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		410	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/admins/{id} [get]
func (h *Handler) GetAdminById(c *fiber.Ctx) error {
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

	response = h.controller.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CGetAdminById(id)

	return utils.ResponseToStatus(c, response)
}

// GetAdminMe is a function to get the current admin from the databse
//
// @Summary		Get current admin
// @Description	Get current admin
// @Tags		Admins
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Success		200	{object}	models.Response{data=models.AdminOutput}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		410	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/admins/me [get]
func (h *Handler) GetAdminMe(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	response = h.controller.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CGetAdminById(curUserId)

	return utils.ResponseToStatus(c, response)
}

// ListAdmins is a function to get a set of all admins from database
//
// @Summary		List all admins of a specified set
// @Description	List all admins of a specified set
// @Tags		Admins
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Param		set	query		int	false	"Set of admins"
// @Param		search	query	string	false	"Search parameter for admin"
// @Success		200	{object}	models.Response{data=[]models.AdminOutput}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/admins/list [get]
func (h *Handler) ListAdmins(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	response = h.controller.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	
	queries := c.Queries()
	set := queries["set"]
	response = utils.IsValidSet(&set)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	search := queries["search"]
	response = utils.IsValidSearch(search)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CListAdmins(set, search)
	return utils.ResponseToStatus(c, response)
}

// CreateAdmin is a function to create a new admin
//
// @Summary		Create a new admin
// @Description	Create a new admin
// @Tags		Admins
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Param		admin	body	models.AdminInput	true	"admin"
// @Success		201	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/admins/ [post]
func (h *Handler) CreateAdmin(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id

	response = h.controller.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	var admin *models.AdminInput
	if err := c.BodyParser(&admin); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response = utils.IsValidAdminInput(admin)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CCreateAdmin(curUserId, admin)

	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	return utils.CreationResponseToStatus(c, response)
}

// UpdateAdmin is a function to update the given admin from the databse
//
// @Summary		Update given admin
// @Description	Update given admin
// @Tags		Admins
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Param		id	path	string	true	"admin ID"
// @Param		admin	body	models.AdminInput	true	"admin"
// @Success		200	{object}	models.Response{}
// @Failure		400	{object}	models.Response{data=models.Err}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router			/admins/{id} [patch]
func (h *Handler) UpdateAdmin(c *fiber.Ctx) error {
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	curUserId := response.Data.(models.UserDB).Id
<<<<<<< HEAD

	response = h.controller.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	
	id := c.Params("id")
<<<<<<< HEAD
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	
=======

>>>>>>> 48bbdbc (112 Create AI analyze endpoint for Frontend)
=======

	response = h.controller.CIsSuperAdmin(curUserId)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	
	id := c.Params("id")
	response = utils.IsValidId(id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}
	
>>>>>>> f463ec0 (106 validation functions (#123))
	var admin *models.AdminInput
	if err := c.BodyParser(&admin); err != nil {
		//	@Failure	400	{object}	models.Response{}
		// something with body is wrong/missing
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	response = utils.IsValidAdminInput(admin)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CUpdateAdmin(id, admin)
	return utils.ResponseToStatus(c, response)
}

// DeleteAdmin is a function to delete the given admin from the database
//
// @Summary		Delete given admin
// @Description	Delete given admin
// @Tags		Admins
// @Accept		json
// @Produce		json
// @Security	Bearer
// @Param		id	path	string	true	"Admin ID"
// @Success		200	{object}	models.Response{}
// @Failure		401	{object}	models.Response{data=models.Err}
// @Failure		403	{object}	models.Response{data=models.Err}
// @Failure		404	{object}	models.Response{data=models.Err}
// @Failure		503	{object}	models.Response{data=models.Err}
// @Router		/admins/{id} [delete]
func (h *Handler) DeleteAdmin(c *fiber.Ctx) error {
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

	response = h.controller.CIsCurrentUserOrSuperAdmin(curUserId, id)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	response = h.controller.CDeleteAdmin(id)
	return utils.ResponseToStatus(c, response)
}
