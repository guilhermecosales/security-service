package main

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/guilhermecosales/security-service/internal/api/handlers"
	"github.com/guilhermecosales/security-service/internal/domain/service"
	"github.com/guilhermecosales/security-service/internal/infrastructure/database"
	"github.com/guilhermecosales/security-service/internal/infrastructure/repository"
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

	userRepository := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepository)

	authenticationTokenService := service.NewAuthenticationTokenService(envConfig.JWTConfig)
	authenticationService := service.NewAuthenticationService(userRepository, authenticationTokenService)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(config.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	handlers.NewHealthHandler(r)
	handlers.NewUserHandler(r, userService)
	handlers.NewAuthenticationHandler(r, authenticationService)

	srv := server.NewServer(envConfig, r)

	log.Info().Msgf("Starting '%s' in '%s' mode on port :%s",
		envConfig.ApplicationName, envConfig.Environment, envConfig.ApplicationPort)
	log.Fatal().Err(srv.ListenAndServe())
}
