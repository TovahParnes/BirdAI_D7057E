package controllers

import (
	"birdai/src/internal/models"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

// RequestAnalyze Connects to AI and retrieves the name and accuracy of the picture
func (c *Controller) RequestAnalyze(mediaData string) models.AIList {
	if os.Getenv("USE_AI") == "true" {
		postBody, _ := json.Marshal(map[string]string{
			"media": mediaData,
		})
		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post("http://localhost:3500/evaluate_image", "application/json", responseBody)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var aiBird models.AIList
		err = json.Unmarshal(body, &aiBird)
		if err != nil {
			log.Fatalln(err)
		}
		return aiBird
	}
	return models.AIList{Birds: []models.AIBird{
		{
			Name:     "EURASIAN MAGPIE",
			Accuracy: 1,
		},
	}}
}

func (c *Controller) AiListToResponse(aiList models.AIList) []models.AnalyzeResponse {
	response := []models.AnalyzeResponse{}
	for _, ai := range aiList.Birds {
		bird := c.CGetBirdByName(ai.Name)
		response = append(response, models.AnalyzeResponse{
			AiBird: models.AIBird{
				Name:     ai.Name,
				Accuracy: ai.Accuracy,
			},
			BirdId:    bird.Data.(*models.BirdOutput).Id,
			UserMedia: bird.Data.(*models.BirdOutput).Image,
		})
	}
	return response
}
