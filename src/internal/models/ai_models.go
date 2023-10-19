package models

type Analyze struct {
	Accuracy string  `json:"accuracy" form:"accuracy"`
	Name     string  `json:"name" form:"name"`
	Picture  MediaDB `json:"picture" form:"picture"`
}
