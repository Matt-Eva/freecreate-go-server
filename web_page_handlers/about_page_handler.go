package web_page_handlers

import (
	"freecreate/auth"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/valkey-io/valkey-go"
)

func AboutPageHandler(aboutTmpl *template.Template, sessionStore *sessions.CookieStore, valkeyClient valkey.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_, userId, _ := auth.GetUser(ctx, sessionStore, valkeyClient, w, r)
		loggedIn := false
		if userId != 0 {
			loggedIn = true
		}

		type PageData struct {
			Title    string
			LoggedIn bool
		}

		pageData := PageData{
			Title:    "about",
			LoggedIn: loggedIn,
		}

		aboutTmpl.ExecuteTemplate(w, "about_page", pageData)
	}
}

func write103Header(w http.ResponseWriter) {
	w.Header().Add("Link", "</static/globals.css>; rel=preload; as=style")
	w.Header().Add("Link", "</static/header_component.css>; rel=preload; as=style")
	w.WriteHeader(103)
}
