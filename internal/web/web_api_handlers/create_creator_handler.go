package web_api_handlers

import (
	"freecreate/internal/config"
	"freecreate/internal/web/web_auth"
	"net/http"

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
	}
}
