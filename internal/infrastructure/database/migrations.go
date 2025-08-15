package database

import (
	"database/sql"
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewMigration(conn *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance("file://./database/migrations", "postgres", driver)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration")
		return nil, err
	}

	err = m.Up()
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Info().Msg("No migration needed")
		return m, nil
	}

	if err != nil {
		log.Info().Err(err).Msg("Failed to apply migrations")
		return nil, err
	}

	log.Info().Msg("Applied migrations")

	return m, err
}
