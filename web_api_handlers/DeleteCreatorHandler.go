package web_api_handlers

import (
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func DeleteCreatorHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, aErr := auth.CheckSession(sessionStore, w, r)
		if aErr != nil {
			logger.Log(aErr)
			http.Error(w, aErr.Error(), http.StatusUnauthorized)
			return
		}

		creatorId := chi.URLParam(r, "creatorId")

		type User struct {
			ID uint
		}

		var user User

		qErr := gormPGClient.Model(&pgModels.User{}).Where("session_uuid = ?", userId).First(&user).Error
		if qErr != nil {
			logger.Log(qErr)
			http.Error(w, qErr.Error(), http.StatusInternalServerError)
			return
		}

		var creator pgModels.Creator

		dErr := gormPGClient.Where("user_id = ? AND id = ?", user.ID, creatorId).Delete(&creator).Error
		if dErr != nil {
			logger.Log(dErr)
			http.Error(w, dErr.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
