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
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func NewWritingHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}

		type Body struct {
			Title       string `json:"title"`
			CreatorId   int64  `json:"creator"`
			WritingType string `json:"writingType"`
		}

		var body Body

		decodeErr := json.NewDecoder(r.Body).Decode(&body)
		if decodeErr != nil {
			logger.Log(decodeErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		params := pg_core_types.NewWritingParams{
			Title:       body.Title,
			CreatorId:   body.CreatorId,
			UserId:      userId,
			WritingType: body.WritingType,
		}

		createdWriting, createWritingErr := query_handlers.HandleNewWriting(ctx, pgCore, pgCoreQueries, params)
		if createWritingErr != nil {
			http.Error(w, createWritingErr.Message, createWritingErr.Code)
			return
		}

		type Response struct {
			UUID        uuid.UUID `json:"uuid"`
			CreatorId   int64     `json:"creatorId"`
			Title       string    `json:"title"`
			Subtitle    string    `json:"subtitle"`
			WritingType string    `json:"writingType"`
			Topics      []string  `json:"topics"`
			Tags        []string  `json:"tags"`
			Description string    `json:"description"`
		}

		res := Response{
			UUID:        createdWriting.UUID,
			CreatorId:   createdWriting.CreatorId,
			Title:       createdWriting.Title,
			Subtitle:    createdWriting.Subtitle,
			WritingType: createdWriting.WritingType,
			Topics:      createdWriting.Topics,
			Tags:        createdWriting.Tags,
			Description: createdWriting.Description,
		}

		jsonRes, marshalErr := json.Marshal(res)
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
