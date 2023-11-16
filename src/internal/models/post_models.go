package models

type PostDB struct {
	Id        string `bson:"_id"`
	UserId    string `bson:"user_id"`
	BirdId    string `bson:"bird_id"`
	CreatedAt string `bson:"created_at"`
	Location  string `bson:"location"`
	MediaId   string `bson:"media_id"`
}

type PostInput struct {
	Location string     `bson:"location" json:"location" form:"location"`
}

type PostOutput struct {
	Id        string      `bson:"_id" json:"_id" form:"_id"`
	User      UserOutput  `bson:"user" json:"user" form:"user"`
	Bird      BirdOutput  `bson:"bird" json:"bird" form:"bird"`
	CreatedAt string      `bson:"created_at" json:"createdAt" form:"createdAt"`
	Location  string      `bson:"location" json:"location" form:"location"`
	UserMedia MediaOutput `bson:"user_media" json:"userMedia" form:"userMedia"`
}

func PostDBToOutput(db *PostDB, user *UserOutput, bird *BirdOutput, media *MediaOutput) *PostOutput {
	return &PostOutput{
		Id:     	db.Id,
		User:   	*user,
		Bird:      	*bird,
		CreatedAt: 	db.CreatedAt,
		Location:  	db.Location, 
		UserMedia: 	*media,
	}
}