package web_page_handlers

import (
	"freecreate/internal/config"
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

func MyWritingsPageHandler(templates *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		params := query_handlers.GetMyWritingsParams{
			UserId: userId,
		}

		myWritings, myWritingsErr := query_handlers.HandleGetMyWritings(ctx, pgCore, pgCoreQueries, params)
		if myWritingsErr != nil {
			http.Error(w, myWritingsErr.Message, myWritingsErr.Code)
			return
		}

		type PageData struct {
			UniversalPageData
			MyWritings []query_handlers.CreatorWritingsGroup
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      true,
				LoggedInClass: "logged_in",
				CsrfToken:     csrf.TemplateField(r),
			},
			MyWritings: myWritings,
		}

		err := templates.ExecuteTemplate(w, "my_writings_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}
