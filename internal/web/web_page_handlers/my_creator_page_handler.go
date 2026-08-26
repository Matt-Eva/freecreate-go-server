package web_page_handlers

import (
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func MyCreatorPageHandler(templates *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", 303)
			return
		}

		creatorUuid := chi.URLParam(r, "creator_uuid")

		myCreator, getMyCreatorErr := pg_core_queries.GetMyCreator(ctx, pgCore, pgCoreQueries, creatorUuid, userId)
		if getMyCreatorErr != nil {
			http.Error(w, getMyCreatorErr.Message, getMyCreatorErr.Code)
			return
		}

		type PageData struct {
			LoggedIn      bool
			LoggedInClass string
			MyCreator     pg_core_queries.MyCreatorStruct
		}

		pageData := PageData{
			LoggedIn:      true,
			LoggedInClass: "logged_in",
			MyCreator:     myCreator,
		}

		templates.ExecuteTemplate(w, "my_creator_page", pageData)
	}
}
