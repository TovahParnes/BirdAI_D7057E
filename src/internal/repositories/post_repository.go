package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PostRepository is a struct to allow the post collection in mongoDB to be accessed
type PostRepository struct {
	collection IMongoCollection
}

func (p *PostRepository) SetCollection(coll IMongoCollection) {
	p.collection = coll
}

// GetPostById returns a response containing the requested post if found
func (p *PostRepository) GetPostById(id string) models.Response {
	filter := bson.M{"_id": id}
	return p.collection.FindOne(filter)
}

// CreatePost creates a new post and returns a response with the id given to the post
func (p *PostRepository) CreatePost(post models.PostDB) models.Response {
	post.Id = primitive.NewObjectID().Hex()
	return p.collection.CreateOne(&post)
}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// UpdatePost updates the post with the specified changes and returns a response
// containing the updated post.
func (p *PostRepository) UpdatePost(post models.PostInput) models.Response {
	data, err := bson.Marshal(post)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonPost := bson.M{}
	err = bson.Unmarshal(data, &bsonPost)
	if err != nil {
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return p.collection.UpdateOne(bsonPost)
}

// ListPosts returns a response with a list for set n of size SetSize that matches the given filter
func (p *PostRepository) ListPosts(filter bson.M, set int) models.Response {
	return p.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllPosts returns a response with a list for set n of size SetSize
func (p *PostRepository) ListAllPosts(set int) models.Response {
	return p.collection.FindAll(bson.M{}, SetSize, SetSize*set)
}

// DeletePost returns the post with the given id
// Returns the post that was deleted
func (p *PostRepository) DeletePost(id string) models.Response {
	return p.collection.DeleteOne(bson.M{"_id": id})
}
