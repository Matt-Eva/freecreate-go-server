package handlers

import (
	"fmt"
	"freecreate/auth"
	"freecreate/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func CreateWritingHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, uErr := auth.GetUser(sessionStore, gormPGClient, w, r)
		if uErr != nil {
			logger.Log(uErr)
			http.Error(w, "failed to fetch authorized user", http.StatusInternalServerError)
			return
		}

		fmt.Println(userId)

	}
}
