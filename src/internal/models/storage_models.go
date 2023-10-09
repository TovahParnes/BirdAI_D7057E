package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Structs for documents in all collections

type User struct {
	Id        string `bson:"_id" json:"_id" form:"_id"`
	Username  string `bson:"username" json:"username" form:"username"`
	AuthId    string `bson:"auth_id" json:"authId" form:"authId"`
	CreatedAt string `bson:"created_at" json:"createdAt" form:"createdAt"`
	Active    bool   `bson:"active"`
}

type Admin struct {
	Id     string `bson:"_id"`
	UserId string `bson:"user_id"`
	Access string `bson:"access"`
}

type Bird struct {
	Id          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	ImageId     string `bson:"image_id"`
	SoundId     string `bson:"sound_id"`
}

// Post TODO: Update Location type after testing
type Post struct {
	Id        string `bson:"_id"`
	UserId    string `bson:"user_id"`
	BirdId    string `bson:"bird_id"`
	CreatedAt string `bson:"created_at"`
	Location  string `bson:"location"`
	ImageId   string `bson:"image_id"`
	SoundId   string `bson:"sound_id"`
}

type Media struct {
	Id       string `bson:"_id"`
	Data     []byte `bson:"data"`
	FileType string `bson:"file_type"`
}

// GetID for all types, to make them HandlerObjects

func (u *User) GetId() string {
	return u.Id
}

func (a *Admin) GetId() string {
	return a.Id
}

func (b *Bird) GetId() string {
	return b.Id
}

func (p *Post) GetId() string {
	return p.Id
}

func (m *Media) GetId() string {
	return m.Id
}

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now().Format(time.RFC3339)
}

func (a *Admin) SetCreatedAt() {
}

func (b *Bird) SetCreatedAt() {
}

func (p *Post) SetCreatedAt() {
	p.CreatedAt = time.Now().Format(time.RFC3339)
}

func (m *Media) SetCreatedAt() {
}

type HandlerObject interface {
	GetId() string
	SetCreatedAt()
}

type MongoInstance struct {
	Client      *mongo.Client
	Db          *mongo.Database
	Collections map[string]MongoCollection
}

func (m MongoInstance) GetCollection(name string) MongoCollection {
	return m.Collections[name]
}

func (m MongoInstance) AddCollection(name string) {
	m.Collections[name] = MongoCollection{Collection: m.Db.Collection(name)}
}

func (m MongoInstance) DisconnectDB() {
	if m.Client != nil {
		if err := m.Client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting from MongoDB %s", err)
		}
	}
}

type IMongoInstance interface {
	GetCollection(name string) IMongoCollection
	AddCollection(name string)
	DisconnectDB()
}

type MongoCollection struct {
	Collection *mongo.Collection
}

// IMongoCollection TODO: Update input when known what is needed
type IMongoCollection interface {
	FindOne(id string) (Response)
	FindAll() (Response)
	UpdateOne(query bson.D) (Response)
	DeleteOne(id string) (HandlerObject, error)
	CreateOne(object HandlerObject) (Response)
}
