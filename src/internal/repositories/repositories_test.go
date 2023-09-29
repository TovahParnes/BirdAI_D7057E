package repositories_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/repositories"
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	t.Run("Test connect", func(t *testing.T) {
		mi, _ := repositories.Connect()
		repositories.AddAllCollections(mi)
		result, err := mi.GetCollection(repositories.UserColl).FindAll()
		fmt.Println("Allting", result)
		fmt.Println("Första", result.([]models.User)[0].Id)

		if err != nil {
			return
		}
		mi.DisconnectDB()
	})
}
