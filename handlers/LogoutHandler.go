package handlers

import (
	"freecreate/auth"
	"freecreate/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func LogoutHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.DestroySession(*sessionStore, w, r)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
