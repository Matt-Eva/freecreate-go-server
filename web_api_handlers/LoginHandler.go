package web_api_handlers

import (
	"encoding/json"
	"errors"
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

func LoginHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Body struct {
			Email string `json:"email"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
			return
		}

		email := body.Email

		var user pgModels.User

		result := gormPGClient.Where("email = ?", email).First(&user)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			err := errors.New("we could not find a user with that email address")
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		} else if result.Error != nil {
			logger.Log(result.Error)
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		user.SessionUUID = uuid.New()

		sResult := gormPGClient.Save(&user)
		if sResult.Error != nil {
			logger.Log(sResult.Error)
			http.Error(w, sResult.Error.Error(), http.StatusInternalServerError)
			return
		}

		err := auth.CreateSession(sessionStore, w, r, user.SessionUUID)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}
