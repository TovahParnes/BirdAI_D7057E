package models

type Token struct {
	Token string `json:"token" form:"token"`
}

type ResponseHTTP struct {
	Success bool `json:"success"` //temporary, will be removed
	StatusCode int `json:"status"`
	StatusName string `json:"name"`
	Timestamp string `json:"timestamp"`
	Data interface{} `json:"data"`
	Message string `json:"message"` //the displayed error text for the user
	Description string `json:"description"` //the error description for the developer, aka err.Error()
}
