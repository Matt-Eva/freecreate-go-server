package web_page_handlers

import (
	"freecreate/internal/config"
	pg_core_queries "freecreate/internal/db/pg_core/queries"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func MyCreatorsPageHandler(templates *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// _, userId, _ := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		// if userId == 0 {
		// 	http.Redirect(w, r, "/login", 303)
		// 	return
		// }

		userId := int64(1)

		myCreators, getMyCreatorsErr := pg_core_queries.GetMyCreators(ctx, pgCore, pgCoreQueries, userId)
		if getMyCreatorsErr != nil {
			http.Error(w, getMyCreatorsErr.Message, getMyCreatorsErr.Code)
			return
		}

		type PageData struct {
			CsrfToken     template.HTML
			LoggedIn      bool
			LoggedInClass string
			MyCreators []pg_core_queries.MyCreatorsStruct
		}

		pageData := PageData{
			CsrfToken:     csrf.TemplateField(r),
			LoggedIn:      true,
			LoggedInClass: "logged_in",
			MyCreators: myCreators,
		}

		templates.ExecuteTemplate(w, "my_creators_page", pageData)
	}
}
