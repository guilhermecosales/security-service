package middleware

import (
	"net/http"

	"github.com/guilhermecosales/security-service/pkg/protocol"
)

func NewBasicAuthenticationMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				protocol.WriteResponse(w, http.StatusUnauthorized, map[string]interface{}{
					"message": "Missing Authorization header",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
