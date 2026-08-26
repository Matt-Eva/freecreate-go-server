package web_page_handlers

import (
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func DonatePageHandler(donateTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
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
			LoggedIn      bool
			LoggedInClass string
		}

		pageData := PageData{
			LoggedIn:      loggedIn,
			LoggedInClass: loggedInClass,
		}

		donateTmpl.ExecuteTemplate(w, "donate_page", pageData)
	}
}
