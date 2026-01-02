package handlers

import (
	"encoding/json"
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func GetMyWritingHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, uErr := auth.GetUser(sessionStore, gormPGClient, w, r)
		if uErr != nil {
			logger.Log(uErr)
			http.Error(w, uErr.Error(), http.StatusUnauthorized)
			return
		}

		type Writing struct {
			Title       string `json:"title"`
			UUID        string	`json:"writingUUID"`
			IsPublished bool	`json:"isPublished"`
		}

		var writing []Writing

		wErr := gormPGClient.Model(pgModels.Writing{}).Where("user_id = ?", userId).Find(&writing).Error
		if wErr != nil && wErr != gorm.ErrRecordNotFound {
			logger.Log(uErr)
			http.Error(w, wErr.Error(), http.StatusInternalServerError)
			return
		}

		type Response struct {
			Writing []Writing `json:"writing"`
		}

		response := Response{
			Writing: writing,
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
