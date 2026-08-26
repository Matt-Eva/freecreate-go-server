package web_page_handlers

import (
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func LoginPageHandler(sessionStore *sessions.CookieStore, valkeyClient valkey.Client, loginTmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		if userId != 0 {
			http.Redirect(w, r, "/profile", 303)
			return
		}

		type PageData struct {
			CsrfToken     template.HTML
			LoggedIn      bool
			LoggedInClass string
		}

		pageData := PageData{
			CsrfToken:     csrf.TemplateField(r),
			LoggedIn:      false,
			LoggedInClass: "logged_out",
		}

		loginTmpl.ExecuteTemplate(w, "login_page", pageData)
	}
}
