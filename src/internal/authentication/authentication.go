package authentication

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

type Authentication struct {
	UserColl repositories.IMongoCollection
}

func NewAuthentication(userCollection repositories.IMongoCollection) Authentication {
	return Authentication{
		UserColl: userCollection,
	}
}

func (l Authentication) LoginUser(user *models.User) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id": user.AuthId,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	foundUser, err := l.UserColl.FindOne(bson.M{"auth_id": t})
	if err != nil {
		user.AuthId = t
		createdUser, err := l.UserColl.CreateOne(user)
		if err != nil {
			return "", err
		}
		return createdUser, nil
	}
	return foundUser.GetId(), nil
}

func (l Authentication) CheckUser(authId string) (models.HandlerObject, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id": authId,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}
	return l.UserColl.FindOne(bson.M{"auth_id": t})
}
