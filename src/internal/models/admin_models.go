package models

type AdminDB struct {
	Id     string `bson:"_id"`
	UserId string `bson:"user_id"`
	Access string `bson:"access"`
}

type AdminInput struct {
	UserId string `bson:"user_id"`
	Access string `bson:"access"`
}

type AdminOutput struct {
	Id     string `bson:"_id"`
	User   UserOutput `bson:"user"`
	Access string `bson:"access"`
}