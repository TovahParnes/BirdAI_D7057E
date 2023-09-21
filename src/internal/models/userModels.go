package models

type User struct {
    Id string `json:"_id" xml:"_id" form:"_id"`
    Username string `json:"username" xml:"username" form:"username"`
}

type TokenUser struct {
	Token
	User User `json:"user" xml:"user" form:"user"`
}