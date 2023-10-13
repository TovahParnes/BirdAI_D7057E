package models

type PostDB struct {
	Id        string `bson:"_id"`
	UserId    string `bson:"user_id"`
	BirdId    string `bson:"bird_id"`
	CreatedAt string `bson:"created_at"`
	Location  string `bson:"location"`
	ImageId   string `bson:"image_id"`
	SoundId   string `bson:"sound_id"`
}

type PostInput struct {
	BirdId    string `bson:"bird_id"`
	Location  string `bson:"location"`
	ImageId   string `bson:"image_id"`
	SoundId   string `bson:"sound_id"`
}

type PostCreate struct {
	UserId    string `bson:"user_id"`
	BirdId    string `bson:"bird_id"`
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

func PostInputToPostCreate(post *PostInput, userId string) *PostCreate {
	return &PostCreate{
		UserId:    userId,
		BirdId:    post.BirdId,
		Location:  post.Location,
		ImageId:   post.ImageId,
		SoundId:   post.SoundId,
	}
}
