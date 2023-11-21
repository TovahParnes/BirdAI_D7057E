package authentication

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Authentication struct {
	UserColl repositories.UserRepository
}

func NewAuthentication(userCollection repositories.UserRepository) Authentication {
	return Authentication{
		UserColl: userCollection,
	}
}

func (a *Authentication) LoginUser(user *models.UserLogin) models.Response {
	response := a.UserColl.GetUserByAuthId(user.AuthId)
	createdNow := false
	// Check if data is type error
	if utils.IsType(response, models.Err{}) {
		if response.Data.(models.Err).StatusCode == http.StatusNotFound {
			createdNow = true
			userDB := models.UserLoginToDB(user)
			response = a.UserColl.CreateUser(*userDB)
			if utils.IsTypeError(response) {
				return response
			}
		} else {
			return response
		}
	}
	// Check if user is active
	response = a.UserColl.GetUserByAuthId(user.AuthId)
	if !response.Data.(*models.UserDB).Active {
		return utils.ErrorDeleted(response.Data.(*models.UserDB).Username)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"_id":      response.Data.(*models.UserDB).Id,
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

	return utils.Response(models.UserDBToLoginOutput(response.Data.(*models.UserDB), t, createdNow))
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
		Id:       claims["_id"].(string),
		Username: claims["username"].(string),
		AuthId:   claims["authId"].(string),
		Active:   true,
	})
}
