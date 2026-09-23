package web_api_handlers

import (
	"freecreate/internal/config"
	"freecreate/internal/web/web_auth"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func NewWritingHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0{
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
			return
		}

		
	}
}