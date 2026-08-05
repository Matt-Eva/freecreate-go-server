package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func ProfilePageHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, profileTmpl *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if userId == 0 {
			http.Redirect(w, r, "/login", 303)
			return
		}

		type PageData struct {
			LoggedIn  bool
			CsrfToken template.HTML
			LoggedInClass string
		}

		pageData := PageData{
			LoggedIn:  true,
			CsrfToken: csrf.TemplateField(r),
			LoggedInClass: "logged_in",
		}

		profileTmpl.ExecuteTemplate(w, "profile", pageData)
	}
}
