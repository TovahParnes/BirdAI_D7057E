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
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 5bf0e17 (implemented post repository)
	Id       string `bson:"_id"`
	UserId   string `bson:"user_id"`
	BirdId   string `bson:"bird_id"`
	Location string `bson:"location"`
	ImageId  string `bson:"image_id"`
	SoundId  string `bson:"sound_id"`
<<<<<<< HEAD
}

type PostCreate struct {
	UserId   string `bson:"user_id"`
	BirdId   string `bson:"bird_id"`
	Location string `bson:"location"`
	ImageId  string `bson:"image_id"`
	SoundId  string `bson:"sound_id"`
=======
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
>>>>>>> f463ec0 (106 validation functions (#123))
=======
}

type PostCreate struct {
	UserId   string `bson:"user_id"`
	BirdId   string `bson:"bird_id"`
	Location string `bson:"location"`
	ImageId  string `bson:"image_id"`
	SoundId  string `bson:"sound_id"`
>>>>>>> 5bf0e17 (implemented post repository)
}

type PostOutput struct {
	Id        string      `bson:"_id"`
	User      UserOutput  `bson:"user"`
	Bird      BirdOutput  `bson:"bird"`
	CreatedAt string      `bson:"created_at"`
	Location  string      `bson:"location"`
	Image     MediaOutput `bson:"image"`
	Sound     MediaOutput `bson:"sound"`
<<<<<<< HEAD
}

func PostInputToPostCreate(post *PostInput, userId string) *PostCreate {
	return &PostCreate{
		UserId:   userId,
		BirdId:   post.BirdId,
		Location: post.Location,
		ImageId:  post.ImageId,
		SoundId:  post.SoundId,
	}
=======
>>>>>>> 5bf0e17 (implemented post repository)
}

func PostInputToPostCreate(post *PostInput, userId string) *PostCreate {
	return &PostCreate{
		UserId:   userId,
		BirdId:   post.BirdId,
		Location: post.Location,
		ImageId:  post.ImageId,
		SoundId:  post.SoundId,
	}
}
