package web_api_handlers

import (
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func DeleteUserHandler(sessionStore *sessions.CookieStore, valkeClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		params := pg_core_types.DeleteUserParams{
			UserId: userId,
		}

		err := query_handlers.HandleDeleteUser(ctx, pgCore, pgCoreQueries, params)
		if err != nil {
			http.Error(w, err.Message, err.Code)
			return
		}

		logoutErr := web_auth.LogoutUser(ctx, sessionStore, valkeClient, w, r)
		if logoutErr != nil {
			http.Error(w, logoutErr.Message, logoutErr.Code)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
