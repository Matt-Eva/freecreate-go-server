package handlers

import (
	"encoding/json"
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"
	"time"

	"github.com/google/uuid"
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
		

		type Body struct {
			CreatorId uint `json:"creatorId"`
			Title string `json:"title"`
			WritingType string `json:"writingType"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
			return
		}

		if body.CreatorId == 0 || body.Title == ""{
			http.Error(w, "must have a valid creator and a title", http.StatusUnprocessableEntity)
			return
		}

		newWriting := pgModels.Writing{
			UserID: userId,
			CreatorID: body.CreatorId,
			UUID: uuid.New(),
			IsPublished: false,
			LastPublished: time.Now(),
			WritingType: body.WritingType,
		}

		if newWriting.UserID == 0 || newWriting.CreatorID == 0 || newWriting.WritingType == ""{
			http.Error(w, "must have a valid user id, creator id, and writing type", http.StatusUnprocessableEntity)
			return
		}
		

		cErr := gormPGClient.Create(&newWriting).Error
		if cErr != nil {
			logger.Log(cErr)
			http.Error(w, cErr.Error(), http.StatusInternalServerError)
			return
		}
		
		type Response struct {
			WritingUUID uuid.UUID `json:"writingUUID"`
		}

		response := Response {
			WritingUUID: newWriting.UUID,
		}

		res, mErr := json.Marshal(response)
		if mErr != nil {
			logger.Log(mErr)
			http.Error(w, mErr.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(res)
	}
}
