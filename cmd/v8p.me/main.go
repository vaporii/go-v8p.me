package main

import (
	"database/sql"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vaporii/go-v8p.me/internal/config"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	conf := config.LoadConfig()

	db, err := sql.Open("sqlite3", conf.SQLitePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	logger.Debug("starting database migration...")
	m, err := migrate.New(
		"file://internal/migrations",
		"sqlite3://"+conf.SQLitePath,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Debug("no changes.")
		} else {
			log.Fatal(err)
		}
	}

	r := chi.NewRouter()

	logger.Info("starting server", "address", conf.ServerAddress)
	log.Fatal(http.ListenAndServe(conf.ServerAddress, r))
}
