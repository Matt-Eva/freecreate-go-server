package web_page_handlers

import (
	"freecreate/internal/web/web_auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func SearchPageHandler(searchTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		userId, _ := web_auth.CheckAuthentication(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		loggedInClass := "logged_out"
		if userId != 0 {
			loggedIn = true
			loggedInClass = "logged_in"
		}

		query := r.URL.Query()
		searchParams := query["search"]

		type PageData struct {
			Query         string
			LoggedIn      bool
			LoggedInClass string
		}

		var pageData PageData
		pageData.LoggedIn = loggedIn
		pageData.LoggedInClass = loggedInClass

		if len(searchParams) > 0 {
			pageData.Query = searchParams[0]
		} else {
			pageData.Query = ""
		}

		searchTmpl.ExecuteTemplate(w, "search_page", pageData)
	}
}
