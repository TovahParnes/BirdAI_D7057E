package utils

import (
	"birdai/src/internal/models"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsValidId(id string) models.Response {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrorParams("Given ID is not a valid id")
	}
	return Response(nil)
}

func IsValidSet(set *string) models.Response {
	if len(*set) == 0 {
		*set = "0"
	}
	setInt, err := strconv.Atoi(*set)
	if err != nil {
		return ErrorParams("Given set is not a valid set, must be int")
	}
	return Response(setInt)
}

func IsValidSearch(search string) models.Response {
	if containsSpecialCharacters(search) {
		return ErrorParams("Given search is not a valid search, must not contain special characters")
	}
	return Response(nil)
}

func IsValidAdminInput(admin *models.AdminInput) models.Response {
	response := IsValidId(admin.UserId)
	if IsTypeError(response) {
		return ErrorParams("Given user id is not a valid id")
	}
	if admin.Access != "admin" && admin.Access != "superadmin" {
		return ErrorParams("Given access is not a valid access, must be admin or superadmin")
	}
	return Response(nil)
}

func IsValidBirdInput(bird *models.BirdInput) models.Response {
	if bird.Name == "" {
		return ErrorParams("Name is empty")
	}
	if bird.Description == "" {
		return ErrorParams("Description is empty")
	}
	response := IsValidId(bird.ImageId)
	if IsTypeError(response) {
		return ErrorParams("Given image id is not a valid id")
	}
	response = IsValidId(bird.SoundId)
	if IsTypeError(response) {
		return ErrorParams("Given sound id is not a valid id")
	}
	if containsSpecialCharacters(bird.Name) {
		return ErrorParams("Name must not contain special characters")
	}
	return Response(nil)
}

func IsValidMediaInput(post *models.MediaInput) models.Response {
	if post.FileType != "image" && post.FileType != "sound" {
		return ErrorParams("Given type is not a valid type, must be image or sound")
	}
	return Response(nil)
}

func IsValidPostInput(post *models.PostInput) models.Response {
	response := IsValidId(post.BirdId)
	if IsTypeError(response) {
		return ErrorParams("Given bird id is not a valid id")
	}

	//TODO validate location?

	//response = IsValidId(post.ImageId)
	//if IsTypeError(response) {
	//	return ErrorParams("Given image id is not a valid id")
	//}
	//response = IsValidId(post.SoundId)
	//if IsTypeError(response) {
	//	return ErrorParams("Given sound id is not a valid id")
	//}
	return Response(nil)
}

func IsValidUserInput(user *models.UserInput) models.Response {
	if user.Active != true && user.Active != false {
		return ErrorParams("Given active is not a valid active, must be true or false")
	}
	return isValidName(user.Username)
}

func IsValidUserLogin(user *models.UserLogin) models.Response {
	response := isValidName(user.Username)
	if IsTypeError(response) {
		return response
	}

	return Response(nil)
}

func isValidName(name string) models.Response {
	if name == "" {
		return ErrorParams("Name is empty")
	}
	if len(name) < 6 || len(name) > 20 {
		return ErrorParams("Name must be between 6 and 20 characters")
	}
	if containsSpecialCharacters(name) {
		return ErrorParams("Name must not contain special characters")
	}
	return Response(nil)
}

func containsSpecialCharacters(str string) bool {
	f := func(r rune) bool {
		return r < 'A' || r > 'z'
	}
	if strings.IndexFunc(str, f) != -1 {
		return true
	}
	return false
}
