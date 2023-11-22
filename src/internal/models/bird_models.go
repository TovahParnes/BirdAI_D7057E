package models

type BirdDB struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	ImageId     string `bson:"image_id"`
	SoundId     string `bson:"sound_id"`
	Species		bool `bson:"species"`
}

type BirdInput struct {
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
	Species		bool `bson:"species"`
}

func BirdInputToDB(input *BirdInput) *BirdDB {
	return &BirdDB{
		Name:        input.Name,
		Description: input.Description,
		ImageId:     input.ImageId,
		SoundId:     input.SoundId,
	}
}

func BirdDBToOutput(db *BirdDB, image *MediaOutput, sound *MediaOutput) *BirdOutput {
	return &BirdOutput{
		Id:          db.Id,
		Name:        db.Name,
		Description: db.Description,
		//Image:       *image,
		//Sound:       *sound,
		Species:	 db.Species,
	}
}
