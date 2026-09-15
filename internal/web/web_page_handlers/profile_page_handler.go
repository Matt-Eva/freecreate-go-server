package web_page_handlers

import (
	"freecreate/internal/config"
	"freecreate/internal/query_handlers"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func ProfilePageHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, profileTmpl *template.Template, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", 303)
			return
		}

		getMyCreatorsParams := query_handlers.GetMyCreatorsParams{
			UserId: userId,
		}

		myCreators, getMyCreatorsErr := query_handlers.HandleGetMyCreators(ctx, pgCore, pgCoreQueries, getMyCreatorsParams)
		if getMyCreatorsErr != nil {
			http.Error(w, getMyCreatorsErr.Message, getMyCreatorsErr.Code)
			return
		}

		type PageData struct {
			UniversalPageData
			MyCreators []query_handlers.MyCreatorsStruct
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				CsrfToken: csrf.TemplateField(r),
				LoggedIn:      true,
				LoggedInClass: "logged_in",
			},
			MyCreators: myCreators,
		}

		profileTmpl.ExecuteTemplate(w, "profile", pageData)
	}
}
