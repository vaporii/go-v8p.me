package storage

import (
	"context"
	"io"

	"github.com/vaporii/go-v8p.me/internal/domain"
)

type StorageService interface {
	Save(
		ctx context.Context,
		r io.Reader,
	) (domain.StorageRef, error)

	Delete(
		ctx context.Context,
		ref domain.StorageRef,
	) error
}
