package main

import (
	"github.com/guilhermecosales/security-service/config"
	"github.com/guilhermecosales/security-service/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	env, err := config.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load environment variables")
	}

	srv := server.New(":" + env.ApplicationPort)

	log.Info().Msgf("Starting '%s' in '%s' mode on port :%s", env.ApplicationName, env.Environment, env.ApplicationPort)
	log.Fatal().Err(srv.ListenAndServe())
}
