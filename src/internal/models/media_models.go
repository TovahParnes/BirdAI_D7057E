package models

type Media struct {
	Id       string `bson:"_id"`
	Data     []byte `bson:"data"`
	FileType string `bson:"file_type"`
}

type MediaInput struct {
	Data     []byte `bson:"data"`
	FileType string `bson:"file_type"`
}

type MediaOutput struct {
	Id       string `bson:"_id"`
	Data     []byte `bson:"data"`
	FileType string `bson:"file_type"`
}