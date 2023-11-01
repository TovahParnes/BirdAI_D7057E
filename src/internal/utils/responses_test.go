package utils_test

import (
	"birdai/src/internal/models"
	"birdai/src/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsType(t *testing.T) {
	var user1 models.UserDB
	testnormal := utils.Response(user1)
	result := utils.IsType(testnormal, models.UserDB{})
	assert.True(t, result)

	var user *models.UserDB
	testpointer := utils.Response(user)
	result = utils.IsType(testpointer, models.UserDB{})
	assert.True(t, result)

	testpointer = utils.Response(user)
	result = utils.IsType(testpointer, &models.UserDB{})
	assert.True(t, result)
}