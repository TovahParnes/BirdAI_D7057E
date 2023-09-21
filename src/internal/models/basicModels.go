package models

type Token struct {
    Token string `json:"token" xml:"token" form:"token"`
}

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
