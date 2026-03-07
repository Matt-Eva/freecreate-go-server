package web_api_handlers

import (
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func DeleteAccountHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := auth.CheckSession(sessionStore, w, r)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		var user pgModels.User

		result := gormPGClient.Where("session_uuid = ?", userId).Delete(&user)
		if result.Error != nil {
			logger.Log(result.Error)
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		aErr := auth.DestroySession(sessionStore, w, r)
		if aErr != nil {
			logger.Log(aErr)
			http.Error(w, aErr.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
