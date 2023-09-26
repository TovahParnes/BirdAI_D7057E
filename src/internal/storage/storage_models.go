package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// Structs for documents in all collections

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	AuthID    string             `bson:"auth_id"`
	CreatedAt primitive.DateTime `bson:"created_at"`
}

type Admin struct {
	ID     primitive.ObjectID
	UserID string
	Access string
}

type Bird struct {
	ID          primitive.ObjectID
	Name        string
	Description string
	ImageID     string
	SoundID     string
}

// Post TODO: Update Location type after testing
// Probably string
type Post struct {
	ID        primitive.ObjectID
	UserID    string
	BirdID    string
	CreatedAt primitive.DateTime
	Location  string
	ImageID   string
	SoundID   string
}

// Sound TODO: Update Sound type after testing
type Sound struct {
	ID    primitive.ObjectID `bson:"_id"`
	Sound []byte             `bson:"sound"`
}

// Image TODO: Update Image type after testing
type Image struct {
	ID    primitive.ObjectID `bson:"_id"`
	Image []byte             `bson:"image"`
}

// GetID for all types, to make them HandlerObjects

func (u User) GetID() primitive.ObjectID {
	return u.ID
}

func (a Admin) GetID() primitive.ObjectID {
	return a.ID
}

func (b Bird) GetID() primitive.ObjectID {
	return b.ID
}

func (p Post) GetID() primitive.ObjectID {
	return p.ID
}

func (s Sound) GetID() primitive.ObjectID {
	return s.ID
}

func (i Image) GetID() primitive.ObjectID {
	return i.ID
}

type HandlerObject interface {
	GetID() primitive.ObjectID
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
	GetCollection(name string) MongoCollection
	AddCollection(name string)
	DisconnectDB()
}

type MongoCollection struct {
	Collection *mongo.Collection
}

// IMongoCollection TODO: Add choice between ID and Name for find
type IMongoCollection interface {
	FindOne(id string) (HandlerObject, error)
	FindAll() ([]HandlerObject, error)
	UpdateOne(query bson.D) (HandlerObject, error) // Might not be bson.D
	DeleteOne(id string) (HandlerObject, error)
	CreateOne(object HandlerObject) error
}
