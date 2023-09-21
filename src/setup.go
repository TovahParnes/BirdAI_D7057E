package src

import (
	"birdai/src/internal/storage"
	"context"
)

func Setup(ctx context.Context) error {
	storage.TestGet()

	return nil
}
