package models

type Token struct {
	Token string `json:"token" form:"token"`
}

type Response struct {
	Timestamp string `json:"timestamp"`
	Data interface{} `json:"data"`
}

type Err struct {
	Success bool `json:"success"`
	StatusCode int `json:"status"`
	StatusName string `json:"name"`
	Message string `json:"message"` //the displayed error text for the user
	Description string `json:"description"` //the error description for the developer, aka err.Error()
}