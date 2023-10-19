package handlers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

// ImagePrediction is a function to analyze pictures
//
// @Summary		Analyze image
// @Description	Send in an image to get a response of which type of bird it is
// @Tags		ai
// @Accept		json
// @Produce		json
// @Param		set	body		models.MediaInput	true	"picture"
// @Success		201	{object}	models.Response{data=[]models.Err}
// @Failure		400	{object}	models.Response{data=[]models.Err}
// @Failure		401	{object}	models.Response{data=[]models.Err}
// @Failure		503	{object}	models.Response{data=[]models.Err}
// @Security 	Bearer
// @Router		/ai/inputimage [post]
func (h *Handler) ImagePrediction(c *fiber.Ctx) error {

	//response := h.auth.CheckExpired(c)
	//if utils.IsTypeError(response) {
	//	return utils.ResponseToStatus(c, response)
	//}

	var picture *models.MediaInput
	if err := c.BodyParser(&picture); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}

	//TEMPORARY for demo
	dat, err := os.ReadFile("src/internal/handlers/TEMPBIRD.txt")
	if err != nil {
		fmt.Print(err.Error())
	}

	prediction := models.Analyze{
		Accuracy: "0.8",
		Name:     "Skata",
		Picture: models.MediaDB{
			Id:       "123123",
			Data:     string(dat),
			FileType: "JPG",
		},
	}
	response := utils.Response(prediction)

	return utils.CreationResponseToStatus(c, response)
}
