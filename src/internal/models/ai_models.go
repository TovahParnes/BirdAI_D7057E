package models

type Analyze struct {
	Accuracy string  `json:"accuracy" form:"accuracy"`
	Name     string  `json:"name" form:"name"`
	BirdId   string  `json:"birdId" form:"birdId"`
	Picture  MediaDB `json:"picture" form:"picture"`
}
