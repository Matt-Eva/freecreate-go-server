package handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func GetUserCreatorHandlers(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		userId, aErr := auth.CheckSession(sessionStore, w, r)
		if aErr != nil {
			logger.Log(aErr)
			http.Error(w, aErr.Error(), http.StatusUnauthorized)
			return
		}

		var user pgModels.User;

		uErr := gormPGClient.Where("session_uuid = ?", userId).First(&user).Error
		if uErr != nil {
			logger.Log(uErr)
			http.Error(w, uErr.Error(), http.StatusInternalServerError)
			return
		}

		var creators []pgModels.Creator;

		cErr := gormPGClient.Where("user_id = ?", user.ID).Find(&creators).Error
		if cErr != nil {
			logger.Log(cErr)
			http.Error(w, cErr.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(creators)

		type Response struct {
			Creators []pgModels.Creator `json:"creators"`
		}

		response := Response {
			Creators: creators,
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