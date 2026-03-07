package web_api_handlers

import (
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func LogoutHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionUUID, uErr := auth.CheckSession(sessionStore, w, r)
		if uErr != nil {
			logger.Log(uErr)
		} else {
			var user pgModels.User

			gErr := gormPGClient.Where("session_uuid = ?", sessionUUID).First(&user).Error
			if gErr != nil {
				logger.Log(uErr)
			} else {
				user.SessionUUID = uuid.New()

				sErr := gormPGClient.Save(&user).Error
				if sErr != nil {
					logger.Log(sErr)
				}

			}
		}

		err := auth.DestroySession(sessionStore, w, r)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
