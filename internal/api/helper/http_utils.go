package helper

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func WriteResponse(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			log.Error().Err(err).Msg("failed to encode response payload")
		}
	}
}
