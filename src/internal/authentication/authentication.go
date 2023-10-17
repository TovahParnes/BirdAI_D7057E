package authentication

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
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

func (a *Authentication) LoginUser(user *models.UserLogin) models.Response {
	// Create the Claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"authId":   user.AuthId,
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return utils.ErrorParams(err.Error())
	}
	response := a.UserColl.FindOne(bson.M{"auth_id": user.AuthId})
	if response.Data.(models.Err).StatusCode != http.StatusNotFound {
		return response
	}

	userDB := models.UserDB{Username: user.Username, AuthId: user.AuthId, Active: true}
	response = a.UserColl.CreateOne(&userDB)
	if utils.IsTypeError(response) {
		return response
	}
	var UserCopy models.UserDB
	UserCopy = *response.Data.(*models.UserDB)
	UserCopy.AuthId = t
	return utils.Response(UserCopy)
}

//func (a *Authentication) Logout(c *fiber.Ctx) models.Response {
//	return models.Response{}
//}
//
//func (a *Authentication) RefreshToken(user *models.User) models.Response {
//	// Create the Claims
//	claims := jwt.MapClaims{
//		"username": user.Username,
//		"authId":   user.AuthId,
//		"exp":      time.Now().Add(timeout * 1).Unix(),
//	}
//	// Create token
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	// Generate encoded token and send it as response.
//	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
//	if err != nil {
//		return utils.ErrorParams(err.Error())
//	}
//	return utils.Response(t)
//}

func (a *Authentication) CheckExpired(c *fiber.Ctx) models.Response {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	//exp := claims["exp"].(float64)
	//fmt.Printf("time: %f", exp-float64(time.Now().Unix()))
	//if exp < float64(time.Now().Unix()) {
	//	return utils.ErrorUnauthorized("authentication has expired")
	//}
	//if exp-float64(time.Now().Unix()) < float64(timeout/2) {
	//	response := a.RefreshToken(&models.User{
	//		Username: claims["username"].(string),
	//		AuthId:   claims["authId"].(string),
	//	})
	//	if utils.IsTypeError(response) {
	//		return response
	//	}
	//}
	return utils.Response(models.UserDB{
		Username: claims["username"].(string),
		AuthId:   claims["authId"].(string),
		Active:   true,
	})
}

