package main

import (
	"fmt"
	"github.com/mori-dv/RSS/internal/auth"
	"github.com/mori-dv/RSS/internal/database"
	"net/http"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusForbidden, fmt.Sprintf("Authorization error: %v", err))
			return
		}
		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusNotFound, fmt.Sprintf("connot get user: %v", err))
			return
		}
		handler(w, r, user)
	}
}
