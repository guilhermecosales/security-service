package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/guilhermecosales/security-service/pkg/config"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewDatabase(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.Username,
		cfg.DatabaseConfig.Password,
		cfg.DatabaseConfig.DatabaseName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		return nil, err
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5) // TODO: study this concept

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
		return nil, err
	}

	log.Info().Msg("Database connection established")
	return db, nil
}
