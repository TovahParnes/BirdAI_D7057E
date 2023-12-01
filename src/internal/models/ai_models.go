package models

type AnalyzeResponse struct {
	AiBird      AIBird `json:"aiBird" form:"aiBird"`
	BirdId      string `json:"birdId" form:"birdId"`
	Description string `json:"description" form:"description"`
}

type AIList struct {
	Birds []AIBird `json:"birds" form:"birds"`
}

type AIBird struct {
	Name     string  `json:"name" form:"name"`
	Accuracy float32 `json:"accuracy" form:"accuracy"`
}
