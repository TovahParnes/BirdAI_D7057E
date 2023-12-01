package models

type BirdDB struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	SoundId     string `bson:"sound_id"`
	Species     bool   `bson:"species"`
}

type BirdInput struct {
	Name        string `bson:"name"`
	Description string `bson:"description"`
	SoundId     string `bson:"sound_id"`
}

type BirdOutput struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Sound       string `bson:"sound"`
	Species     bool   `bson:"species"`
}

func BirdInputToDB(input *BirdInput) *BirdDB {
	return &BirdDB{
		Name:        input.Name,
		Description: input.Description,
		SoundId:     input.SoundId,
	}
}

func BirdDBToOutput(db *BirdDB) *BirdOutput {
	return &BirdOutput{
		Id:          db.Id,
		Name:        db.Name,
		Description: db.Description,
		Sound:       db.SoundId,
		Species:     db.Species,
	}
}
