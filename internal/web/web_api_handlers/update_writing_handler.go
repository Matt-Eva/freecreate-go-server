package web_api_handlers

import (
	"encoding/json"
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func UpdateWritingHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		writingUuidParam := chi.URLParam(r, "writing_uuid")

		writingUuid, parseUuidErr := uuid.Parse(writingUuidParam)
		if parseUuidErr != nil {
			logger.Log(parseUuidErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		type Body struct {
			Title       string   `json:"title"`
			Subtitle    string   `json:"subtitle"`
			CreatorId   int64    `json:"creatorId"`
			WritingType string   `json:"writingType"`
			Description string   `json:"description"`
			Topics      []string `json:"topics"`
			Tags        []string `json:"tags"`
			IsAdult		bool	 `json:"isAdult"`
		}

		var body Body;

		decodeErr := json.NewDecoder(r.Body).Decode(&body)
		if decodeErr != nil {
			logger.Log(decodeErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		params := pg_core_types.UpdateWritingParams{
			UserId: userId,
			UUID:   writingUuid,
			Title: body.Title,
			Subtitle: body.Subtitle,
			Description: body.Description,
			IsAdult: body.IsAdult,
			Topics: body.Topics,
			Tags: body.Tags,
			WritingType: body.WritingType,
			CreatorId: body.CreatorId,
		}

		updatedWriting, updateWritingErr := query_handlers.HandleUpdateWriting(ctx, pgCore, pgCoreQueries, params)
		if updateWritingErr != nil {
			http.Error(w, updateWritingErr.Message, updateWritingErr.Code)
			return
		}

		response := Body {
			Title: updatedWriting.Title,
			Subtitle: updatedWriting.Subtitle,
			Description: updatedWriting.Description,
			Topics: updatedWriting.Topics,
			Tags: updatedWriting.Tags,
			IsAdult: updatedWriting.IsAdult,
			WritingType: updatedWriting.WritingType,
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
