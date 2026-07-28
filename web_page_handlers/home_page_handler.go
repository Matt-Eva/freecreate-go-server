package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func HomePageHandler(homeTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		if userId != 0 {
			loggedIn = true
		}

		type PageData struct {
			LoggedIn bool
		}

		pageData := PageData{
			LoggedIn: loggedIn,
		}
		homeTmpl.ExecuteTemplate(w, "home", pageData)
	}
}
