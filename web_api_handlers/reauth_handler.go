package web_api_handlers

import (
	"freecreate/auth"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ReAuthHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.CheckSession(sessionStore, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
