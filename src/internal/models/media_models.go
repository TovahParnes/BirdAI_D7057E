package models

import "time"

type MediaDB struct {
	Id   string `bson:"_id" json:"_id" form:"_id"`
	Data string `bson:"data" json:"data" form:"data"`
}

type MediaInput struct {
	Data      string        `bson:"data" json:"data" form:"data"`
	StartTime time.Duration `bson:"start_time" json:"startTime" form:"startTime"`
}

type MediaOutput struct {
	Id   string `bson:"_id" json:"_id" form:"_id"`
	Data string `bson:"data" json:"data" form:"data"`
}

func MediaDBToOutput(db *MediaDB) *MediaOutput {
	return &MediaOutput{
		Id:   db.Id,
		Data: db.Data,
	}
}
