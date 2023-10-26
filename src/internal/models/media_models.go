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
	Data     []byte `bson:"data" json:"data" form:"data"`
	FileType string `bson:"file_type" json:"fileType" form:"fileType"`
}
