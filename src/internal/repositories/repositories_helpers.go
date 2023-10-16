package repositories

import (
	"birdai/src/internal/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserColl  = "users"
	AdminColl = "admins"
	BirdColl  = "birds"
	PostColl  = "posts"
	MediaColl = "media"
)

// TODO: Add custom error variables
var (
// ParseError = errors.New("")
)

// MongoInstance Instance of Mongo
type MongoInstance struct {
	Client      *mongo.Client
	Db          *mongo.Database
	Collections map[string]IMongoCollection
}

func (m MongoInstance) GetCollection(name string) IMongoCollection {
	return m.Collections[name]
}

func (m MongoInstance) AddCollection(name string) {
	m.Collections[name] = &MongoCollection{Collection: m.Db.Collection(name)}
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
	ctx        context.Context
}

// IMongoCollection TODO: Update input when known what is needed
type IMongoCollection interface {
	UpdateOne(query bson.M) (models.Response)
	DeleteOne(query bson.M) (models.Response)
	CreateOne(object models.HandlerObject) (models.Response)
	FindOne(query bson.M) (models.Response)
	FindAll() (models.Response)
}
