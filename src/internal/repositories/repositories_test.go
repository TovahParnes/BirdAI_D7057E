package repositories_test

import (
	"birdai/src/internal/repositories"
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func TestConnection(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		fmt.Errorf(err.Error())
	}
	t.Run("Test connect", func(t *testing.T) {
		repositories.TestConnect()
	})
}
