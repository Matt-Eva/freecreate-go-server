package handlers

import (
	"encoding/json"
	"errors"
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

func SignupHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type Body struct {
			Email      string `json:"email"`
			BirthDay   int    `json:"birthDay"`
			BirthMonth int    `json:"birthMonth"`
			BirthYear  int    `json:"birthYear"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
		}

		birthDate := time.Date(body.BirthYear, time.Month(body.BirthMonth), body.BirthDay, 0, 0, 0, 0, time.UTC)

		email := body.Email

		var currentUser pgModels.User
		var newUser pgModels.User

		// see if current user exists
		result := gormPGClient.Where("email = ?", email).First(&currentUser)
		fmt.Println(result.Error)
		fmt.Println(result.Error == gorm.ErrRecordNotFound)
		if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
			// if they don't exist, create a new user
			newUser.Email = email
			newUser.Birthday = birthDate
			newUser.SessionUUID = uuid.New()

			result := gormPGClient.Create(&newUser)
			if result.Error != nil {
				logger.Log(result.Error)
				http.Error(w, result.Error.Error(), http.StatusUnprocessableEntity)
				return
			}
			fmt.Println(newUser)
		} else {
			err := errors.New("email address already in use")
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		err := auth.CreateSession(sessionStore, w, r, newUser.SessionUUID)
		if err != nil {
			logger.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
