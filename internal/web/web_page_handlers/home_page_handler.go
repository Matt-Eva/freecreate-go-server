package web_page_handlers

import (
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func HomePageHandler(homeTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		}

		type PageData struct {
			UniversalPageData
		}

		pageData := PageData{
			UniversalPageData: UniversalPageData{
				LoggedIn:      loggedIn,
				LoggedInClass: loggedInClass,
				CsrfToken:     csrf.TemplateField(r),
			},
		}

		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
