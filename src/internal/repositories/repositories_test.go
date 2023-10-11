package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"birdai/src/internal/utils"
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	t.Run("Test connect", func(t *testing.T) {
		mi, _ := repositories.Connect()
		repositories.AddAllCollections(mi)
		response := mi.GetCollection(repositories.UserColl).FindAll()
		if utils.IsTypeError(response) {
			return
		}
		fmt.Println("Allting", response.Data.([]models.UserOutput))
		fmt.Println("Första", response.Data.([]models.UserOutput)[0].Id)

		mi.DisconnectDB()
	})
}
