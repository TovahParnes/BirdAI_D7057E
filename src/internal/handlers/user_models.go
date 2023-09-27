package handlers

type UserReponse struct {
	Id        string `json:"_id" form:"_id"`
	AuthId    string `json:"authId" form:"authId"`
	Username  string `json:"username" form:"username"`
	CreatedAt string `json:"createdAt" form:"createdAt"`
}

type User struct {
	Username string `json:"username" form:"username"`
}

type TokenUser struct {
	Token
	User User `json:"user" form:"user"`
}
