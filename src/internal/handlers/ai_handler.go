package handlers

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"encoding/base64"
	"fmt"
	"github.com/go-audio/audio"
	"github.com/gofiber/fiber/v2"
	"os"
	"strings"
	"time"
)

// ImagePrediction is a function to analyze pictures
//
// @Summary		Analyze image
// @Description	Send in an image to get a response of which type of bird it is
// @Tags		AI
// @Accept		json
// @Produce		json
// @Param		set	body		models.MediaInput	true	"picture"
// @Success		200	{object}	models.Response{data=[]models.AnalyzeResponse}
// @Failure		500	{object}	models.Err
// @Failure		503	{object}	models.Err
// @Security 	Bearer
// @Router		/api/v1/ai/inputimage [post]
func (h *Handler) ImagePrediction(c *fiber.Ctx) error {

	//TODO JPEG, PNG

	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	var picture *models.MediaInput
	if err := c.BodyParser(&picture); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}

	var aiBirds models.AIList
	var dat string

	if os.Getenv("USE_AI") == "true" {
		dat = picture.Data
	} else {
		//TEMPORARY for demo
		byteDat, err := os.ReadFile("src/internal/handlers/TEMPBIRD.txt")
		if err != nil {
			fmt.Print(err.Error())
		}
		dat = string(byteDat)
	}
	aiBirds, err := h.controller.RequestAnalyze(dat, "evaluate_image")
	if err != nil {
		return err
	}

	aiResponse := h.controller.AiListToResponse(aiBirds, false)

	response = utils.Response(aiResponse)

	return utils.CreationResponseToStatus(c, response)
}

// SoundPrediction is a function to analyze pictures
//
// @Summary		Analyze sound
// @Description	Send in a sound to get a response of which type of bird it is
// @Tags		AI
// @Accept		json
// @Produce		json
// @Param		set	body		models.MediaInput	true	"picture"
// @Success		200	{object}	models.Response{data=[]models.AnalyzeResponse}
// @Failure		500	{object}	models.Err
// @Failure		503	{object}	models.Err
// @Security 	Bearer
// @Router		/api/v1/ai/inputsound [post]
func (h *Handler) SoundPrediction(c *fiber.Ctx) error {

	//TODO sound files: mp3, wav
	response := h.auth.CheckExpired(c)
	if utils.IsTypeError(response) {
		return utils.ResponseToStatus(c, response)
	}

	var sound *models.MediaInput
	if err := c.BodyParser(&sound); err != nil {
		//	@Failure	400	{object}	models.Response{}
		return utils.ResponseToStatus(c, utils.ErrorParams(err.Error()))
	}
	if sound.Data == "" {
		return models.Err{
			Description: "cannot have data be empty",
		}
	}

	var aiBirds models.AIList
	var dat string

	if os.Getenv("USE_AI") == "true" {
		dat = sound.Data
	} else {
		//TEMPORARY for demo
		byteDat, err := os.ReadFile("src/internal/handlers/TEMP_SOUND.txt")
		if err != nil {
			fmt.Print(err.Error())
		}
		dat = string(byteDat)
	}
	startBaseUrl := "data:audio/wav;base64,"
	if strings.HasPrefix(dat, startBaseUrl) {
		dat = dat[len(startBaseUrl):]
	}
	startBaseUrl = "data:audio/mpeg;base64,"
	if strings.HasPrefix(dat, startBaseUrl) {
		dat = dat[len(startBaseUrl):]
	}
	audioBytes, err := base64.StdEncoding.DecodeString(dat)
	if err != nil {
		return err
	}
	var format string
	if h.controller.IsWAV(audioBytes) {
		format = "WAV"
	} else if h.controller.IsMP3(audioBytes) {
		format = "MP3"
	} else {
		return models.Err{
			Description: "wrong format, not wav and mp3",
		}
	}

	// Decode audio data based on format
	var audioData *audio.IntBuffer
	switch format {
	case "WAV":
		audioData, err = h.controller.DecodeWAV(audioBytes)
	case "MP3":
		audioData, err = h.controller.DecodeMP3(audioBytes)
	}
	if err != nil {
		return err
	}

	// Calculate the new duration you want (e.g., 10 seconds)
	targetDuration := 20 * time.Second

	// Trim the audio data to the target duration
	trimmedData := h.controller.TrimAudioData(audioData, targetDuration, sound.StartTime*2*time.Millisecond)

	// Encode the trimmed audio data to base64
	trimmedBase64 := h.controller.EncodeAudioData(trimmedData)

	// Print or use the trimmed base64 string as needed
	//fmt.Println("Trimmed Base64:", trimmedBase64)

	dat = startBaseUrl + trimmedBase64

	aiBirds, err = h.controller.RequestAnalyze(dat, "evaluate_sound")
	if err != nil {
		return err
	}

	aiResponse := h.controller.AiListToResponse(aiBirds, true)

	response = utils.Response(aiResponse)

	return utils.CreationResponseToStatus(c, response)
}
