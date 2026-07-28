package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LoginPageHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, loginTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if userId != 0{
			http.Redirect(w, r, "/profile", 303)
			return
		}

		type PageData struct {
			CsrfToken template.HTML
			LoggedIn bool
		}

		pageData := PageData{
			CsrfToken: csrf.TemplateField(r),
			LoggedIn: false,
		}

		loginTmpl.ExecuteTemplate(w, "login_page", pageData)
	}
}
