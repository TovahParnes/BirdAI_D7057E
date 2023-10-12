package models

type UserDB struct {
	Id        string `bson:"_id" json:"_id" form:"_id"`
	Username  string `bson:"username" json:"username" form:"username"`
	AuthId    string `bson:"auth_id" json:"authId" form:"authId"`
	CreatedAt string `bson:"created_at" json:"createdAt" form:"createdAt"`
	Active    bool   `bson:"active"`
}

type UserLogin struct {
	Username  string `bson:"username" json:"username" form:"username"`
	AuthId    string `bson:"auth_id" json:"authId" form:"authId"`
}

type UserInput struct {
	Username string `json:"user" bson:"username"`
	Active    bool   `bson:"active"`
}

type UserOutput struct {
	Id        string `bson:"_id" json:"_id" form:"_id"`
	Username  string `bson:"username" json:"username" form:"username"`
	CreatedAt string `bson:"created_at" json:"createdAt" form:"createdAt"`
	Active    bool   `bson:"active"`
}