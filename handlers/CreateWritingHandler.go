package handlers

import (
	"encoding/json"
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

		fmt.Println(body)

		

	}
}
