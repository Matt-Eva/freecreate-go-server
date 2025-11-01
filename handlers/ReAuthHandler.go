package handlers

import (
	"errors"
	"freecreate/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ReAuthHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessionStore.Get(r, "user-session")
		if session.Values["userId"] == nil {
			err := errors.New("user not logged in")
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
