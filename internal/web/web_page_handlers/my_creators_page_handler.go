package web_page_handlers

import (
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func MyCreatorsPageHandler(templates *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", 303)
			return
		}

		type PageData struct {
			CsrfToken     template.HTML
			LoggedIn      bool
			LoggedInClass string
		}

		pageData := PageData{
			CsrfToken:     csrf.TemplateField(r),
			LoggedIn:      true,
			LoggedInClass: "logged_in",
		}

		templates.ExecuteTemplate(w, "my_creators_page", pageData)
	}
}
