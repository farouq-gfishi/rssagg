package main

import (
	"net/http"

	"github.com/farouq-gfishi/rssagg/internal/auth"
	"github.com/farouq-gfishi/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		user, err := cfg.DB.GetUserByApikey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		handler(w, r, user)
	}
}
