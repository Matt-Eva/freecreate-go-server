package handlers

import (
	"fmt"
	"freecreate/auth"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func ReAuthHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := auth.CheckSession(sessionStore, w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		} else {
			fmt.Println("user logged in")
			fmt.Println(session)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}
