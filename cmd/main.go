package main

import (
	"github.com/guilhermecosales/security-service/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	srv := server.New(":8080")

	log.Info().Msgf("Starting server at %s", srv.Addr)
	log.Fatal().Err(srv.ListenAndServe())
}
