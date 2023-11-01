package models

type MediaDB struct {
	Id       string `bson:"_id" json:"_id" form:"_id"`
	Data     string `bson:"data" json:"data" form:"data"`
	FileType string `bson:"file_type" json:"fileType" form:"fileType"`
}

type MediaInput struct {
	Data     string `bson:"data" json:"data" form:"data"`
	FileType string `bson:"file_type" json:"fileType" form:"fileType"`
}

type MediaOutput struct {
	Id       string `bson:"_id" json:"_id" form:"_id"`
	Data     string `bson:"data" json:"data" form:"data"`
	FileType string `bson:"file_type" json:"fileType" form:"fileType"`
}

func MediaDBToOutput(db *MediaDB) *MediaOutput {
	return &MediaOutput{
		Id:       db.Id,
		Data:     db.Data,
		FileType: db.FileType,
	}
}
