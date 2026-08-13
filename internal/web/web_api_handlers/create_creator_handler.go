package web_api_handlers

import (
	"encoding/json"
	"fmt"
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/lib/api_error"
	"freecreate/internal/lib/logger"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func CreateCreatorHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, getUserErr := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if getUserErr != nil || userId == 0 {
			http.Redirect(w, r, "/login", 303)
			return
		}

		type Body struct {
			Name string `json:"name"`
		}

		var body Body

		jsonErr := json.NewDecoder(r.Body).Decode(&body)
		if jsonErr != nil {
			logger.Log(jsonErr)
			http.Error(w, api_error.InteralServerErrorMessage, http.StatusInternalServerError)
			return
		}

		creatorName := body.Name

		createdCreator, createCreatorErr := pg_core_queries.CreateCreator(ctx, pgCore, pgCoreQueries, creatorName, userId)
		if createCreatorErr != nil {
			http.Error(w, createCreatorErr.Message, createCreatorErr.Code)
			return
		}

		type Response struct {
			UUID uuid.UUID `json:"uuid"`
			Name string `json:"name"`
		}

		res := Response {
			UUID: createdCreator.UUID,
			Name: createdCreator.Name,
		}
		
		jsonRes, err := json.Marshal(res)
		if err != nil {
			logger.Log(err)
			http.Error(w, api_error.InteralServerErrorMessage, 500)
			return
		}

		fmt.Println("creator successfully created! Returning response")
		fmt.Println(jsonRes)

		// w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jsonRes)
	}
}
