package handlers

type Token struct {
	Token string `json:"token" form:"token"`
}

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
