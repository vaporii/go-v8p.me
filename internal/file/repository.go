package file

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
	"github.com/vaporii/go-v8p.me/internal/domain"
)

var (
	ErrNotFound  = errors.New("not found")
	ErrCollision = errors.New("already exists")
	ErrInternal  = errors.New("internal error")
)

type FileRepository struct {
	db *sql.DB
}

func wrapError(op string, err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("%w: %s: %v", ErrNotFound, op, err)
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlite3.ErrConstraint && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%w: %s: %v", ErrCollision, op, err)
		}
	}

	return fmt.Errorf("%w: %s: %v", ErrInternal, op, err)
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (r *FileRepository) Begin(ctx context.Context) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *FileRepository) CreateFileTx(ctx context.Context, tx *sql.Tx, newFile *domain.File) (*domain.File, error) {
	file := domain.File{Storage: domain.StorageRef{}}

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO files (
			name,
			mime_type,
			encrypted,
			size_bytes,
			alias,
			expires,
			highlighting_language,
			storage_provider,
			storage_key
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING
			id,
			name,
			mime_type,
			encrypted,
			size_bytes,
			alias,
			expires,
			highlighting_language,
			storage_provider,
			storage_key,
			created_at;
	`,
		newFile.Name,
		newFile.Type,
		newFile.Encrypted,
		newFile.Size,
		newFile.Alias,
		newFile.Expires,
		newFile.HighlightingLanguage,
		newFile.Storage.Provider,
		newFile.Storage.Key,
	).Scan(
		&file.ID,
		&file.Name,
		&file.Type,
		&file.Encrypted,
		&file.Size,
		&file.Alias,
		&file.Expires,
		&file.HighlightingLanguage,
		&file.CreatedAt,
		&file.Storage.Provider,
		&file.Storage.Key,
	)

	if err != nil {
		return nil, wrapError("insert file", err)
	}

	return &file, nil
}
