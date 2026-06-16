package web_api_handlers

import (
	"encoding/json"
	"errors"
	"freecreate/auth"
	"freecreate/gorm_models"
	"freecreate/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

func GetEditWritingHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, uErr := auth.GetUser(sessionStore, gormPGClient, w, r)
		if uErr != nil {
			logger.Log(uErr)
			http.Error(w, uErr.Error(), http.StatusUnauthorized)
			return
		}

		writingUUID := chi.URLParam(r, "writingUUID")
		if writingUUID == "" {
			err := errors.New("writingUUID cannot be empty")
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		type EditWriting struct {
			Title       string         `json:"title"`
			Description string         `json:"description"`
			WritingType string         `json:"writingType"`
			Tags        pq.StringArray `json:"tags" gorm:"type:text[]"`
			UUID        uuid.UUID      `json:"writingUUID"`
			CreatorID   uint           `json:"creatorId"`
			IsPublished bool           `json:"isPublished"`
		}

		var editWriting EditWriting

		wErr := gormPGClient.Model(pgModels.Writing{}).Where("uuid = ? AND user_id = ?", writingUUID, userId).First(&editWriting).Error
		if wErr != nil {
			logger.Log(wErr)
			http.Error(w, wErr.Error(), http.StatusInternalServerError)
			return
		}

		res, mErr := json.Marshal(editWriting)
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
