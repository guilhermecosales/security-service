package main

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guilhermecosales/security-service/internal/database"
	"github.com/guilhermecosales/security-service/internal/handlers"
	"github.com/guilhermecosales/security-service/internal/repository"
	"github.com/guilhermecosales/security-service/internal/server"
	"github.com/guilhermecosales/security-service/internal/service"
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

	userRepo := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo)

	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	handlers.NewHealthHandler(r)
	handlers.NewUserHandler(r, userService)

	srv := server.NewServer(envConfig, r)

	log.Info().Msgf("Starting '%s' in '%s' mode on port :%s",
		envConfig.ApplicationName, envConfig.Environment, envConfig.ApplicationPort)
	log.Fatal().Err(srv.ListenAndServe())
}
