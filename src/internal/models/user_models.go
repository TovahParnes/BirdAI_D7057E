package models

type UserReponse struct {
	Id        string `json:"_id" form:"_id"`
	AuthId    string `json:"authId" form:"authId"`
	Username  string `json:"username" form:"username"`
	CreatedAt string `json:"createdAt" form:"createdAt"`
}

type InputUser struct {
	Username string `json:"user" bson:"username"`
}

type TokenUser struct {
	Token
	InputUser InputUser `json:"user" form:"user"`
}
