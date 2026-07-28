package web_page_handlers

import (
	"freecreate/auth"
	"freecreate/lib/logger"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func ProfilePageHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, profileTmpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, _, getUserErr := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if getUserErr != nil {
			logger.Log(getUserErr)
			http.Redirect(w, r, "/login", 303)
			return
		}

		type PageData struct {
			LoggedIn  bool
			CsrfToken template.HTML
		}

		pageData := PageData{
			LoggedIn:  true,
			CsrfToken: csrf.TemplateField(r),
		}

		profileTmpl.ExecuteTemplate(w, "profile", pageData)
	}
}
