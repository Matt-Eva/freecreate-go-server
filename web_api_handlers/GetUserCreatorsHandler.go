package web_api_handlers

import (
	"encoding/json"
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func GetUserCreatorHandlers(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, aErr := auth.GetUser(sessionStore, gormPGClient, w, r)
		if aErr != nil {
			logger.Log(aErr)
			http.Error(w, aErr.Error(), http.StatusUnauthorized)
			return
		}

		type ResponseCreators struct {
			Name string    `json:"name"`
			ID   uint      `json:"id"`
			UUID uuid.UUID `json:"uuid"`
		}

		var responseCreators []ResponseCreators

		cErr := gormPGClient.Model(&pgModels.Creator{}).Where("user_id = ?", userId).Find(&responseCreators).Error
		if cErr != nil {
			logger.Log(cErr)
			http.Error(w, cErr.Error(), http.StatusInternalServerError)
			return
		}

		type Response struct {
			Creators []ResponseCreators `json:"creators"`
		}

		response := Response{
			Creators: responseCreators,
		}

		res, mErr := json.Marshal(response)
		if mErr != nil {
			logger.Log(mErr)
			http.Error(w, mErr.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
