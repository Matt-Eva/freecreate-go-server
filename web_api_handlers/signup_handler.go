package web_api_handlers

// import (
// 	"encoding/json"
// 	"errors"
// 	"freecreate/auth"
// 	"freecreate/gorm_models"
// 	"freecreate/logger"
// 	"net/http"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/gorilla/sessions"
// 	"github.com/jackc/pgx/v5/pgconn"
// 	"gorm.io/gorm"
// )

// func SignupHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		type Body struct {
// 			Email      string `json:"email"`
// 			BirthDay   int    `json:"birthDay"`
// 			BirthMonth int    `json:"birthMonth"`
// 			BirthYear  int    `json:"birthYear"`
// 		}

// 		var body Body

// 		jErr := json.NewDecoder(r.Body).Decode(&body)
// 		if jErr != nil {
// 			logger.Log(jErr)
// 			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
// 		}

// 		birthDate := time.Date(body.BirthYear, time.Month(body.BirthMonth), body.BirthDay, 0, 0, 0, 0, time.UTC)

// 		email := body.Email

// 		newUser := pgModels.User{
// 			Email:       email,
// 			Birthday:    birthDate,
// 			SessionUUID: uuid.New(),
// 		}

// 		if newUser.Email == "" {
// 			err := errors.New("email cannot be empty")
// 			logger.Log(err)
// 			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
// 			return
// 		}

// 		result := gormPGClient.Create(&newUser)

// 		if result.Error != nil {
// 			var pgErr *pgconn.PgError
// 			errors.As(result.Error, &pgErr)
// 			if pgErr.Code == "23505" {
// 				http.Error(w, "email address already in use with another account", http.StatusConflict)
// 				return
// 			} else {
// 				logger.Log(result.Error)
// 				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 		}

// 		err := auth.LoginUser(sessionStore, w, r, newUser.SessionUUID)
// 		if err != nil {
// 			logger.Log(err)
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}
// }
