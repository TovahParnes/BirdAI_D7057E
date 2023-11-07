package models

type AnalyzeResponse struct {
	AiBird    AIBird      `json:"aiBird" form:"aiBird"`
	BirdId    string      `json:"birdId" form:"birdId"`
	UserMedia MediaOutput `json:"userMedia" form:"userMedia"`
}

type AIList struct {
	Birds []AIBird `json:"birds" form:"birds"`
}

type AIBird struct {
	Name     string  `json:"name" form:"name"`
	Accuracy float32 `json:"accuracy" form:"accuracy"`
}
