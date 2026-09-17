package web_api_handlers

import (
	"encoding/json"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func UpdateUserInfoHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		type Body struct {
			Username       string `json:"username"`
			UserHandle     string `json:"userHandle"`
			IsAdult        bool   `json:"isAdult"`
			ReadingHistory bool   `json:"readingHistory"`
		}

		var body Body

		jsonErr := json.NewDecoder(r.Body).Decode(&body)
		if jsonErr != nil {
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusUnprocessableEntity)
			return
		}

		queryParams := query_handlers.UpdateUserInfoParams{
			Username:       body.Username,
			UserHandle:     body.UserHandle,
			IsAdult:        body.IsAdult,
			ReadingHistory: body.ReadingHistory,
			UserId:         userId,
		}

		userInfo, updateInfoErr := query_handlers.HandleUpdateUserInfo(ctx, pgCore, pgCoreQueries, queryParams)
		if updateInfoErr != nil {
			http.Error(w, updateInfoErr.Message, updateInfoErr.Code)
			return
		}

		type Response struct {
			Username       string `json:"username"`
			UserHandle     string `json:"userHandle"`
			IsAdult        bool   `json:"isAdult"`
			ReadingHistory bool   `json:"ReadingHistory"`
		}

		res := Response{
			Username:       userInfo.Username,
			UserHandle:     userInfo.UserHandle,
			IsAdult:        userInfo.IsAdult,
			ReadingHistory: userInfo.ReadingHistory,
		}

		jsonRes, marshalErr := json.Marshal(res)
		if marshalErr != nil {
			logger.Log(marshalErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_, err := w.Write(jsonRes)
		if err != nil {
			logger.Log(err)
		}
	}
}
