package models

type BirdDB struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	ImageId     string `bson:"image_id"`
	SoundId     string `bson:"sound_id"`
}

type BirdInput struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	ImageId     string `bson:"image_id"`
	SoundId     string `bson:"sound_id"`
}

type BirdOutput struct {
	Id          string      `bson:"_id"`
	Name        string      `bson:"name"`
	Description string      `bson:"description"`
	Image       MediaOutput `bson:"image"`
	Sound       MediaOutput `bson:"sound"`
}
