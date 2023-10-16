package utils

import "birdai/src/internal/models"

func IsAdmin(authId string) models.Response {
	return ErrorNotImplemented("IsAdmin")
}

func IsSuperAdmin(authId string) models.Response {
	return ErrorNotImplemented("IsSuperAdmin")
}

func IsPostsUser(authId string, userId string) models.Response {
	return ErrorNotImplemented("IsPostsUser")
}

func IsCurrentUser(authId string, userId string) models.Response {
	return ErrorNotImplemented("IsCurrentUser")
}
