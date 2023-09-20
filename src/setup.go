package src

import (
	"birdai/src/internal/authentication"
	"context"
)

func Setup(ctx context.Context) error {
	// Setup authenticator
	_, err := authentication.NewAuthentication()
	if err != nil {
		return err
	}

	return nil
}
