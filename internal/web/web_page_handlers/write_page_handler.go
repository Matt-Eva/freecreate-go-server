package web_page_handlers

import (
	"freecreate/internal/config"
	pg_core_types "freecreate/internal/db/pg_core/types"
	"freecreate/internal/lib"
	"freecreate/internal/lib/logger"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func WritePageHandler(templates *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		getCreatorsParams := pg_core_types.GetMyCreatorsParams{
			UserId: userId,
		}

		myCreators, queryErr := query_handlers.HandleGetMyCreators(ctx, pgCore, pgCoreQueries, getCreatorsParams)
		if queryErr != nil {
			http.Error(w, queryErr.Message, queryErr.Code)
			return
		}

		type PageData struct {
			UniversalPageData
			MyCreators []pg_core_types.MyCreatorsStruct
			WritingTypes []string
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      true,
				LoggedInClass: "logged_in",
				CsrfToken:     csrf.TemplateField(r),
			},
			MyCreators: myCreators,
			WritingTypes: lib.WritingTypes,
		}

		err := templates.ExecuteTemplate(w, "write_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
