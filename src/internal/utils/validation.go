package utils

import (
	"birdai/src/internal/models"
	"unicode"

	valid "github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func IsValidId(id string) models.Response {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorParams("Given ID is not a valid id")
	}
	return Response(nil)
}

func IsValidSet(set string) models.Response {
	valid := valid.IsIn(set)
	if !valid {
		return ErrorParams("Given set is not a valid set, must be int")
	}
	return Response(nil)
}

func IsValidSearch(search string) models.Response {
	if containsSpecialCharacters(search) {
		return ErrorParams("Given search is not a valid search, must not contain special characters")
	}
	return Response(nil)
}

func IsValidAdminInput(admin models.AdminInput) models.Response {
	response := IsValidId(admin.UserId)
	if IsTypeError(response) {
		return ErrorParams("Given user id is not a valid id")
	}
	if admin.Access != "admin" && admin.Access != "superadmin" {
		return ErrorParams("Given access is not a valid access, must be admin or superadmin")
	}
	return Response(nil)
}

func IsValidBirdInput(bird models.BirdInput) models.Response {
	if bird.Name == ""{
		return ErrorParams("Name is empty")
	}
	response := IsValidId(bird.ImageId)
	if IsTypeError(response) {
		return ErrorParams("Given image id is not a valid id")
	}
	response = IsValidId(bird.SoundId)
	if IsTypeError(response) {
		return response
	}
	return Response(nil)
}

func IsValidMediaInput(post models.MediaInput) models.Response {
	if post.FileType != "image" && post.FileType != "sound" {
		return ErrorParams("Given type is not a valid type, must be image or sound")
	}
	return Response(nil)
}

func IsValidPostInput(post models.PostInput) models.Response {
	response := IsValidId(post.UserId)
	if IsTypeError(response) {
		return ErrorParams("Given user id is not a valid id")
	}
	response = IsValidId(post.BirdId)
	if IsTypeError(response) {
		return ErrorParams("Given bird id is not a valid id")
	}

	//TODO validate location?

	response = IsValidId(post.ImageId)
	if IsTypeError(response) {
		return ErrorParams("Given image id is not a valid id")
	}
	response = IsValidId(post.SoundId)
	if IsTypeError(response) {
		return ErrorParams("Given sound id is not a valid id")
	}
	return Response(nil)
}

func IsValidUserInput(user models.UserInput) models.Response {
	if user.Username == ""{
		return ErrorParams("Username is empty")
	}
	if len(user.Username) <6 || len(user.Username) > 20 {
		return ErrorParams("Username must be between 6 and 20 characters")
	}
	if containsSpecialCharacters(user.Username) {
		return ErrorParams("Username must not contain special characters")
	}

	return Response(nil)
}



func containsSpecialCharacters(str string) bool {
	for _, letter := range str {
		if unicode.IsSymbol(letter) {
			return true
		}
	}
	return false
}

