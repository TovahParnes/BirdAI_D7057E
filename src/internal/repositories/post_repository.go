package repositories

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"fmt"
>>>>>>> 5bf0e17 (implemented post repository)
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
// PostRepository is a struct to allow the post collection in mongoDB to be accessed
type PostRepository struct {
	collection IMongoCollection
}

func (p *PostRepository) SetCollection(coll IMongoCollection) {
	p.collection = coll
<<<<<<< HEAD
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
=======
// PostEndpoints is a struct to allow the post collection in mongoDB to be accessed
type PostEndpoints struct {
	Collection IMongoCollection
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// GetPostById returns a response containing the requested post if found
func (p *PostRepository) GetPostById(id string) models.Response {
	filter := bson.M{"_id": id}
	return p.collection.FindOne(filter)
}

// CreatePost creates a new post and returns a response with the id given to the post
func (p *PostRepository) CreatePost(post models.PostDB) models.Response {
	post.Id = primitive.NewObjectID().Hex()
<<<<<<< HEAD
	return p.Collection.CreateOne(&post)
>>>>>>> 5bf0e17 (implemented post repository)
=======
	return p.collection.CreateOne(&post)
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// TODO: Fix ToBson and FromBson on structs for easier handling of bson to struct and back

// UpdatePost updates the post with the specified changes and returns a response
// containing the updated post.
<<<<<<< HEAD
<<<<<<< HEAD
func (p *PostRepository) UpdatePost(post models.PostInput) models.Response {
	data, err := bson.Marshal(post)
	if err != nil {
=======
func (p PostEndpoints) UpdatePost(post models.PostInput) models.Response {
	data, err := bson.Marshal(post)
	if err != nil {
		fmt.Println(data)
		fmt.Println(err)
>>>>>>> 5bf0e17 (implemented post repository)
=======
func (p *PostRepository) UpdatePost(post models.PostInput) models.Response {
	data, err := bson.Marshal(post)
	if err != nil {
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	bsonPost := bson.M{}
	err = bson.Unmarshal(data, &bsonPost)
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
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
=======
		fmt.Println(data)
		fmt.Println(err)
=======
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
		return utils.ErrorToResponse(400, "Could not update object", err.Error())
	}
	return p.collection.UpdateOne(bsonPost)
}

// ListPosts returns a response with a list for set n of size SetSize that matches the given filter
func (p *PostRepository) ListPosts(filter bson.M, set int) models.Response {
	return p.collection.FindAll(filter, SetSize, SetSize*set)
}

// ListAllPosts returns a response with a list for set n of size SetSize
<<<<<<< HEAD
func (p PostEndpoints) ListAllPosts(set int) models.Response {
	return p.Collection.FindAll(bson.M{}, SetSize, SetSize*set)
>>>>>>> 5bf0e17 (implemented post repository)
=======
func (p *PostRepository) ListAllPosts(set int) models.Response {
	return p.collection.FindAll(bson.M{}, SetSize, SetSize*set)
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}

// DeletePost returns the post with the given id
// Returns the post that was deleted
<<<<<<< HEAD
<<<<<<< HEAD
func (p *PostRepository) DeletePost(id string) models.Response {
	return p.collection.DeleteOne(bson.M{"_id": id})
=======
func (p PostEndpoints) DeletePost(id string) models.Response {
	return p.Collection.DeleteOne(bson.M{"_id": id})
>>>>>>> 5bf0e17 (implemented post repository)
=======
func (p *PostRepository) DeletePost(id string) models.Response {
	return p.collection.DeleteOne(bson.M{"_id": id})
>>>>>>> 9295f66 (implemented admin repository and refactored repositories)
}
