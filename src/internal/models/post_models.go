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
	BirdId   string     `bson:"bird_id"`
	Location string     `bson:"location"`
	Media    MediaInput `bson:"media"`
}

type PostOutput struct {
	Id        string      `bson:"_id"`
	User      UserOutput  `bson:"user"`
	Bird      BirdOutput  `bson:"bird"`
	CreatedAt string      `bson:"created_at"`
	Location  string      `bson:"location"`
	UserMedia MediaOutput `bson:"user_media"`
}
