package storage

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/vaporii/go-v8p.me/internal/config"
	"github.com/vaporii/go-v8p.me/internal/domain"
)

type LocalStorage struct {
	Config config.LocalStorageConfig
}

var (
	ErrStorageSaveFailed = errors.New("storage save failed")
)

func NewLocal(cfg config.LocalStorageConfig) StorageService {
	return &LocalStorage{
		Config: cfg,
	}
}

func (s *LocalStorage) Save(
	ctx context.Context,
	r io.Reader,
) (domain.StorageRef, error) {
	key := generateKey(s.Config.FileNameLength)
	path := filepath.Join(s.Config.BaseFilesPath, key)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return domain.StorageRef{}, fmt.Errorf("%w: mkdir failed: %v", ErrStorageSaveFailed, err)
	}

	f, err := os.Create(path)
	if err != nil {
		return domain.StorageRef{}, fmt.Errorf("%w: create file failed: %v", ErrStorageSaveFailed, err)
	}
	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return domain.StorageRef{}, fmt.Errorf("%w: write failed: %v", ErrStorageSaveFailed, err)
	}

	return domain.StorageRef{
		Provider: "local",
		Key:      key,
	}, nil
}

func (s *LocalStorage) Delete(ctx context.Context, ref domain.StorageRef) error {
	return nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateKey(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
