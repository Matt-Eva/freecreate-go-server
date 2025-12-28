package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(userId)

		type Body struct {
			CreatorId uint `json:"creatorId"`
			Title string `json:"title"`
			Description string `json:"description"`
			Tags []string `json:"tags"`
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
			Tags: body.Tags,
			Description: body.Description,
			UUID: uuid.New(),
			IsPublished: false,
			LastPublished: time.Now(),
		}

		fmt.Println(newWriting)
		

		

	}
}
