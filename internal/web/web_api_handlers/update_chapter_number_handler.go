package web_api_handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func UpdateChapterNumberHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries,) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		urlUuid := chi.URLParam(r, "uuid")
		fmt.Println(urlUuid)

		chapterUuid, parseUuidErr := uuid.Parse(urlUuid)
		if parseUuidErr != nil {
			logger.Log(parseUuidErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		type Body struct {
			ChapterNumber int `json:"chapterNumber"`
		}

		var body Body

		jsonErr := json.NewDecoder(r.Body).Decode(&body)
		if jsonErr != nil {
			logger.Log(jsonErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		params := pg_core_types.UpdateChapterNumberParams{
			UserId: userId,
			UUID: chapterUuid,
			ChapterNumber: body.ChapterNumber,
		}

		newChapterNumber, updateChapterErr := query_handlers.HandleUpdateChapterNumber(ctx, pgCore, pgCoreQueries, params)
		if updateChapterErr != nil {
			http.Error(w, updateChapterErr.Message, updateChapterErr.Code)
			return
		}

		type Response struct {
			NewChapterNumber int `json:"newChapterNumber"`
		}

		response := Response {
			NewChapterNumber: newChapterNumber,
		}

		jsonRes, marshalErr := json.Marshal(response)
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
