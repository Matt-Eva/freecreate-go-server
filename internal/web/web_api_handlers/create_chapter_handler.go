package web_api_handlers

import (
	"encoding/json"
	"freecreate/internal/config"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func CreateChapterHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries, pgContentPools config.PgContentPools, pgContentQueries config.PgContentQueries) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0{
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		type Body struct {
			Title string `json:"title"`
			ChapterNumber int `json:"chapterNumber"`
			WritingUUID uuid.UUID `json:"writingUuid"`
		}

		var body Body

		jsonErr := json.NewDecoder(r.Body).Decode(&body)
		if jsonErr != nil {
			logger.Log(jsonErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		handleCreateChapterParams := query_handlers.HandleCreateChapterParams{
			UserId: userId,
			WritingUUID: body.WritingUUID,
			Title: body.Title,
			ChapterNumber: body.ChapterNumber,
		}

		createdChapter, createChapterErr := query_handlers.HandleCreateChapter(ctx, pgCore, pgCoreQueries, pgContentPools, pgContentQueries, handleCreateChapterParams)
		if createChapterErr != nil {
			http.Error(w, createChapterErr.Message, createChapterErr.Code)
			return
		}

		type Response struct {
			UUID uuid.UUID `json:"uuid"`
			Title string `json:"title"`
			ChapterNumber int `json:"chapterNumber"`
			Published bool `json:"published"`
		}

		response := Response{
			UUID: createdChapter.UUID,
			Title: createdChapter.Title,
			ChapterNumber: createdChapter.ChapterNumber,
			Published: createdChapter.Published,
		}

		jsonRes, marshalErr := json.Marshal(response)
		if marshalErr != nil {
			logger.Log(marshalErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, err := w.Write(jsonRes)
		if err != nil {
			logger.Log(err)
		}
	}
}