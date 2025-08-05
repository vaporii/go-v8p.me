package repository

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/require"
	"github.com/vaporii/go-v8p.me/internal/dto"
	"github.com/vaporii/go-v8p.me/internal/model"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func setupDB(t *testing.T) *sql.DB {
	m, err := migrate.New(
		"file://../migrations",
		"sqlite3://test.db",
	)
	if err != nil {
		t.Fatal(err)
	}

	m.Down()
	if err := m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			t.Fatal(err)
		}
	}

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func createRepository(db *sql.DB) FileRepository {
	return NewFileRepository(db)
}

func TestCreateFile(t *testing.T) {
	db := setupDB(t)
	defer db.Close()
	repo := createRepository(db)

	exp := time.Now().Add(time.Hour * 24)
	lang := "dawdawdfdgffrdgred5tgr8j5487r4jngfgjh"

	req := dto.CreateFileDB{
		Name:                 "test file name !@#$%^&*(()))-_=+;'\";,./<>?/\\|😊😊😊テストテストテスト",
		Type:                 "text/plain",
		Encrypted:            false,
		Size:                 1000000,
		ServerPath:           "./data/nifeu9834nfgbdjghbjhbnsfoib",
		Alias:                "abcdef",
		Expires:              &exp,
		DisplayAs:            model.Markdown,
		HighlightingLanguage: &lang,
	}

	file, err := repo.CreateFile(req)
	require.NoError(t, err)
	require.NotEmpty(t, file)

	require.Equal(t, file.Name, req.Name)
	require.Equal(t, file.Type, req.Type)
	require.Equal(t, file.Encrypted, req.Encrypted)
	require.Equal(t, file.Size, req.Size)
	require.Equal(t, file.ServerPath, req.ServerPath)
	require.Equal(t, file.Alias, req.Alias)
	require.WithinDuration(t, *file.Expires, *req.Expires, time.Second)
	require.Equal(t, file.DisplayAs, req.DisplayAs)
	require.Equal(t, file.HighlightingLanguage, req.HighlightingLanguage)
	require.NotZero(t, file.CreatedAt)
	require.Equal(t, file.ID, int64(1))
}
