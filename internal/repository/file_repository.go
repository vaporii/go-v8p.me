package repository

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
	"github.com/vaporii/go-v8p.me/internal/apperror"
	"github.com/vaporii/go-v8p.me/internal/dto"
	"github.com/vaporii/go-v8p.me/internal/model"
)

type FileRepository interface {
	CreateFile(file dto.CreateFileDB) (*model.File, error)
}

type fileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) FileRepository {
	return &fileRepository{
		db: db,
	}
}

func (r *fileRepository) CreateFile(fileDto dto.CreateFileDB) (*model.File, error) {
	file := model.File{}
	var encrypted uint8 = 0
	if fileDto.Encrypted {
		encrypted = 1
	}
	err := r.db.QueryRow(`
		INSERT INTO files (
			name,
			type,
			encrypted,
			size,
			server_path,
			alias,
			expires,
			display_as,
			highlighting_language
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING
			id,
			name,
			type,
			encrypted,
			size,
			server_path,
			alias,
			expires,
			display_as,
			highlighting_language,
			created_at;
	`,
		fileDto.Name,
		fileDto.Type,
		encrypted,
		fileDto.Size,
		fileDto.ServerPath,
		fileDto.Alias,
		fileDto.Expires,
		fileDto.DisplayAs,
		fileDto.HighlightingLanguage,
	).Scan(
		&file.ID,
		&file.Name,
		&file.Type,
		&file.Encrypted,
		&file.Size,
		&file.ServerPath,
		&file.Alias,
		&file.Expires,
		&file.DisplayAs,
		&file.HighlightingLanguage,
		&file.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sqlite3.ErrConstraint) {
			return nil, &apperror.ConflictError{Entity: "File", CollisionField: "alias", FieldValue: fileDto.Alias}
		}
		return nil, err
	}

	return &file, nil
}
