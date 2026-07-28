package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func SearchPageHandler(searchTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		if userId != 0 {
			loggedIn = true
		}

		query := r.URL.Query()
		searchParams := query["search"]

		type PageData struct {
			Query    string
			LoggedIn bool
		}

		var pageData PageData
		pageData.LoggedIn = loggedIn

		if len(searchParams) > 0 {
			pageData.Query = searchParams[0]
		} else {
			pageData.Query = ""
		}

		searchTmpl.ExecuteTemplate(w, "search_page", pageData)
	}
}
