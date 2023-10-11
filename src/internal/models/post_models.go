package models

type Post struct {
	Id        string `bson:"_id"`
	UserId    string `bson:"user_id"`
	BirdId    string `bson:"bird_id"`
	CreatedAt string `bson:"created_at"`
	Location  string `bson:"location"`
	ImageId   string `bson:"image_id"`
	SoundId   string `bson:"sound_id"`
}

type PostInput struct {
	UserId	 string `bson:"user_id"`
	BirdId    string `bson:"bird_id"`
	CreatedAt string `bson:"created_at"`
	Location  string `bson:"location"`
	ImageId   string `bson:"image_id"`
	SoundId   string `bson:"sound_id"`
}

type PostOutput struct {
	Id        string `bson:"_id"`
	User	 UserOutput `bson:"user"`
	Bird	 BirdOutput `bson:"bird"`
	CreatedAt string `bson:"created_at"`
	Location  string `bson:"location"`
	Image   MediaOutput `bson:"image"`
	Sound   MediaOutput `bson:"sound"`
}
