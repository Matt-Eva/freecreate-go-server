package web_page_handlers

import (
	"freecreate/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func SignupPageHandler(signupTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, userId, _ := web_auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		if userId != 0 {
			http.Redirect(w, r, "/profile", 303)
			return
		}

		type PageData struct {
			CsrfToken template.HTML
			LoggedInClass string
		}

		pageData := PageData{
			CsrfToken: csrf.TemplateField(r),
			LoggedInClass: "logged_out",
		}

		signupTmpl.ExecuteTemplate(w, "signup_page", pageData)
	}
}

func renderSignupPage(signupTmpl *template.Template, w http.ResponseWriter, r *http.Request) {

}
