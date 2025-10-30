package handlers

import (
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ReAuthHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessionStore.Get(r, "user-session")
		if session.Values["userId"] == nil {

		}
	}
}
