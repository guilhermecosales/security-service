package main

import (
	"github.com/guilhermecosales/security-service/internal/database"
	"github.com/guilhermecosales/security-service/internal/repository/user"
	"github.com/guilhermecosales/security-service/internal/server"
	"github.com/guilhermecosales/security-service/pkg/config"
	"github.com/rs/zerolog/log"
)

func main() {
	envConfig, err := config.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load environment variables")
	}

	conn, err := database.NewDatabase(envConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	defer conn.Close()

	m, err := database.NewMigration(conn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create migration")
	}
	defer m.Close()

	_ = user.NewUserRepository(conn)

	srv := server.New(envConfig)
	log.Info().Msgf("Starting '%s' in '%s' mode on port :%s",
		envConfig.ApplicationName, envConfig.Environment, envConfig.ApplicationPort)
	log.Fatal().Err(srv.ListenAndServe())
}
