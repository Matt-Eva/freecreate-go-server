package handlers

import (
	"encoding/json"
	"errors"
	"freecreate/auth"
	"freecreate/logger"
	"freecreate/pgModels"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func CreateCreatorHandler(sessionStore *sessions.CookieStore, gormPGClient *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, aErr := auth.CheckSession(sessionStore, w, r)
		if aErr != nil {
			logger.Log(aErr)
			http.Error(w, aErr.Error(), http.StatusUnauthorized)
			return
		}

		type Body struct {
			CreatorName string `json:"creatorName"`
		}

		var body Body

		jErr := json.NewDecoder(r.Body).Decode(&body)
		if jErr != nil {
			logger.Log(jErr)
			http.Error(w, jErr.Error(), http.StatusUnprocessableEntity)
			return
		}

		var user pgModels.User

		uErr := gormPGClient.Where("session_uuid = ?", userId).First(&user).Error
		if uErr != nil && uErr == gorm.ErrRecordNotFound {
			logger.Log(uErr)

			aErr := auth.DestroySession(sessionStore, w, r)
			if aErr != nil {
				logger.Log(aErr)
				http.Error(w, aErr.Error(), http.StatusInternalServerError)
				return
			}
			
			http.Error(w, "Looks like your session expired! Logging you out - please log back in and try again.", http.StatusNotFound)
			return
		} else if uErr != nil {
			logger.Log(uErr)
			http.Error(w, uErr.Error(), http.StatusInternalServerError)
			return
		}

		creatorUUID := uuid.New()

		newCreator := pgModels.Creator{
			UserID: user.ID,
			UUID:   creatorUUID,
			Name:   body.CreatorName,
		}

		result := gormPGClient.Create(&newCreator)

		if result.Error != nil {
			var pgErr *pgconn.PgError
			errors.As(result.Error, &pgErr)

			if pgErr.Code == "23505" {
				http.Error(w, "you cannot create two creators with the same name", http.StatusConflict)
				return
			} else {
				logger.Log(result.Error)
				http.Error(w, result.Error.Error(), http.StatusInternalServerError)
				return
			}
		}

		type Response struct {
			Name string    `json:"name"`
			ID   uint      `json:"id"`
			UUID uuid.UUID `json:"uuid"`
		}

		response := Response{
			Name: newCreator.Name,
			ID:   newCreator.ID,
			UUID: newCreator.UUID,
		}

		res, eErr := json.Marshal(response)
		if eErr != nil {
			logger.Log(eErr)
			http.Error(w, eErr.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}
}
