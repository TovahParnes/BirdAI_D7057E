package models

type AdminDB struct {
	Id     string `bson:"_id" json:"_id" form:"_id"`
	UserId string `bson:"user_id" json:"userId" form:"userId"`
	Access string `bson:"access" json:"access" form:"access"`
}

type AdminCreation struct {
	UserId string `bson:"user_id" json:"userId" form:"userId"`
	Access string `bson:"access" json:"access" form:"access"`
}
type AdminInput struct {
	Access string `bson:"access" json:"access" form:"access"`
}

type AdminOutput struct {
	Id     string     `bson:"_id" json:"_id" form:"_id"`
	User   UserOutput `bson:"user" json:"user" form:"user"`
	Access string     `bson:"access" json:"access" form:"access"`
}

func AdminDBToOutput(db *AdminDB, user *UserOutput) *AdminOutput {
	return &AdminOutput{
		Id:     db.Id,
		User:   *user,
		Access: db.Access,
	}
}
