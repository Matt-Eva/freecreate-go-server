package handlers

import (
	"freecreate/auth"
	"freecreate/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func CreateWritingHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, aErr := auth.CheckSession(sessionStore, w, r)
		if aErr != nil {
			logger.Log(aErr)
			http.Error(w, aErr.Error(), http.StatusUnauthorized)
			return
		}

		//gormPGClient.Raw(`INSERT INTO `)

	}
}
