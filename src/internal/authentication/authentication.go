package authentication

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

type Authentication struct {
	UserColl repositories.IMongoCollection
}

func NewAuthentication(userCollection repositories.IMongoCollection) Authentication {
	return Authentication{
		UserColl: userCollection,
	}
}

func (l Authentication) LoginUser(user *models.User) (models.Response) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id": user.AuthId,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return utils.ErrorParams(err.Error())
	}
	response := l.UserColl.FindOne(bson.M{"auth_id": t})
	if utils.IsTypeError(response) {
		user.AuthId = t
		response = l.UserColl.CreateOne(user)
		if utils.IsTypeError(response) {
			return response
		}
		return response
	}
	return response
}

func (l Authentication) CheckUser(authId string) (models.Response) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id": authId,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return utils.ErrorParams(err.Error())
	}
	return l.UserColl.FindOne(bson.M{"auth_id": t})
}
