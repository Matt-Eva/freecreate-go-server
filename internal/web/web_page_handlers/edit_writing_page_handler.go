package web_page_handlers

import (
	"freecreate/internal/config"
	"freecreate/internal/lib/logger"
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/valkey-io/valkey-go"
)

func EditWritingPageHandler(templates template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client, pgCore *pgxpool.Pool, pgCoreQueries config.PgCoreQueries)http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		type PageData struct {
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn: true,
				LoggedInClass: "logged_in",
				CsrfToken: csrf.TemplateField(r),
			},
		}

		err := templates.ExecuteTemplate(w, "edit_writing_page", pageData)
		if err != nil {
			logger.Log(err)
		}
	}
}